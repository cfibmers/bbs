package handlers

import (
	"net/http"

	"github.com/cloudfoundry-incubator/auctioneer"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type ActualLRPLifecycleHandler struct {
	db               db.ActualLRPDB
	desiredLRPDB     db.DesiredLRPDB
	actualHub        events.Hub
	auctioneerClient auctioneer.Client
	retirer          ActualLRPRetirer
	exitChan         chan<- struct{}
	logger           lager.Logger
}

func NewActualLRPLifecycleHandler(
	logger lager.Logger,
	db db.ActualLRPDB,
	desiredLRPDB db.DesiredLRPDB,
	actualHub events.Hub,
	auctioneerClient auctioneer.Client,
	retirer ActualLRPRetirer,
	exitChan chan<- struct{},
) *ActualLRPLifecycleHandler {
	return &ActualLRPLifecycleHandler{
		db:               db,
		desiredLRPDB:     desiredLRPDB,
		actualHub:        actualHub,
		auctioneerClient: auctioneerClient,
		retirer:          retirer,
		exitChan:         exitChan,
		logger:           logger.Session("actual-lrp-handler"),
	}
}

func (h *ActualLRPLifecycleHandler) ClaimActualLRP(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("claim-actual-lrp")

	request := &models.ClaimActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}
	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	before, after, err := h.db.ClaimActualLRP(logger, request.ProcessGuid, request.Index, request.ActualLrpInstanceKey)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}

	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	if !after.Equal(before) {
		go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
	}
}

func (h *ActualLRPLifecycleHandler) StartActualLRP(w http.ResponseWriter, req *http.Request) {
	var err error

	logger := h.logger.Session("start-actual-lrp")

	request := &models.StartActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	before, after, err := h.db.StartActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey, request.ActualLrpNetInfo)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	if before == nil {
		go h.actualHub.Emit(models.NewActualLRPCreatedEvent(after))
	} else if !before.Equal(after) {
		go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
	}
}

func (h *ActualLRPLifecycleHandler) CrashActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("crash-actual-lrp")

	request := &models.CrashActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	actualLRPKey := request.ActualLrpKey
	actualLRPInstanceKey := request.ActualLrpInstanceKey

	before, after, shouldRestart, err := h.db.CrashActualLRP(logger, actualLRPKey, actualLRPInstanceKey, request.ErrorMessage)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	if shouldRestart {
		desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, actualLRPKey.ProcessGuid)
		if err != nil {
			logger.Error("failed-fetching-desired-lrp", err)
	if err == models.ErrNoTable {
		logger.Error("failed-desired-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
			response.Error = models.ConvertError(err)
			return
		}

		schedInfo := desiredLRP.DesiredLRPSchedulingInfo()
		startRequest := auctioneer.NewLRPStartRequestFromSchedulingInfo(&schedInfo, int(actualLRPKey.Index))
		err = h.auctioneerClient.RequestLRPAuctions([]*auctioneer.LRPStartRequest{&startRequest})
		if err != nil {
			logger.Error("failed-requesting-auction", err)
			response.Error = models.ConvertError(err)
			return
		}
	}

	actualLRP, _ := after.Resolve()
	go h.actualHub.Emit(models.NewActualLRPCrashedEvent(actualLRP))
	go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
}

func (h *ActualLRPLifecycleHandler) FailActualLRP(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("fail-actual-lrp")

	request := &models.FailActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	before, after, err := h.db.FailActualLRP(logger, request.ActualLrpKey, request.ErrorMessage)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
}

func (h *ActualLRPLifecycleHandler) RemoveActualLRP(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("remove-actual-lrp")

	request := &models.RemoveActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	beforeActualLRPGroup, err := h.db.ActualLRPGroupByProcessGuidAndIndex(logger, request.ProcessGuid, request.Index)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	err = h.db.RemoveActualLRP(logger, request.ProcessGuid, request.Index, request.ActualLrpInstanceKey)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	if err != nil {
		response.Error = models.ConvertError(err)
		return

	}
	go h.actualHub.Emit(models.NewActualLRPRemovedEvent(beforeActualLRPGroup))
}

func (h *ActualLRPLifecycleHandler) RetireActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("retire-actual-lrp")
	request := &models.RetireActualLRPRequest{}
	response := &models.ActualLRPLifecycleResponse{}

	var err error
	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	err = h.retirer.RetireActualLRP(logger, request.ActualLrpKey.ProcessGuid, request.ActualLrpKey.Index)
	if err == models.ErrNoTable {
		logger.Error("failed-actual-lrps-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}

	response.Error = models.ConvertError(err)
}
