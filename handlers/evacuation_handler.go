package handlers

import (
	"net/http"

	"github.com/cloudfoundry-incubator/auctioneer"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/gogo/protobuf/proto"
	"github.com/pivotal-golang/lager"
)

type EvacuationHandler struct {
	db               db.EvacuationDB
	actualLRPDB      db.ActualLRPDB
	desiredLRPDB     db.DesiredLRPDB
	actualHub        events.Hub
	auctioneerClient auctioneer.Client
	logger           lager.Logger
	exitChan         chan struct{}
}

func NewEvacuationHandler(
	logger lager.Logger,
	db db.EvacuationDB,
	actualLRPDB db.ActualLRPDB,
	desiredLRPDB db.DesiredLRPDB,
	actualHub events.Hub,
	auctioneerClient auctioneer.Client,
	exitChan chan struct{},
) *EvacuationHandler {
	return &EvacuationHandler{
		db:               db,
		actualLRPDB:      actualLRPDB,
		desiredLRPDB:     desiredLRPDB,
		actualHub:        actualHub,
		auctioneerClient: auctioneerClient,
		logger:           logger.Session("evacuation-handler"),
		exitChan:         exitChan,
	}
}

type MessageValidator interface {
	proto.Message
	Validate() error
	Unmarshal(data []byte) error
}

func (h *EvacuationHandler) RemoveEvacuatingActualLRP(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("remove-evacuating-actual-lrp")
	logger.Info("started")
	defer logger.Info("completed")

	request := &models.RemoveEvacuatingActualLRPRequest{}
	response := &models.RemoveEvacuatingActualLRPResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	beforeActualLRPGroup, err := h.actualLRPDB.ActualLRPGroupByProcessGuidAndIndex(logger, request.ActualLrpKey.ProcessGuid, request.ActualLrpKey.Index)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	err = h.db.RemoveEvacuatingActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	go h.actualHub.Emit(models.NewActualLRPRemovedEvent(beforeActualLRPGroup))
}

func (h *EvacuationHandler) EvacuateClaimedActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("evacuate-claimed-actual-lrp")
	logger.Info("started")
	defer logger.Info("completed")

	request := &models.EvacuateClaimedActualLRPRequest{}
	response := &models.EvacuationResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		logger.Error("failed-parsing-request", err)
		response.Error = models.ConvertError(err)
		response.KeepContainer = true
		return
	}

	beforeActualLRPGroup, err := h.actualLRPDB.ActualLRPGroupByProcessGuidAndIndex(logger, request.ActualLrpKey.ProcessGuid, request.ActualLrpKey.Index)
	if err == nil {
		err = h.db.RemoveEvacuatingActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey)
		if err != nil {
			if err == models.ErrNoTable {
				logger.Error("failed-actual-lrps-table-does-not-exist", err)
				h.exitChan <- struct{}{}
			}
			logger.Error("failed-removing-evacuating-actual-lrp", err)
		} else {
			go h.actualHub.Emit(models.NewActualLRPRemovedEvent(beforeActualLRPGroup))
		}
	}

	err = h.unclaimAndRequestAuction(logger, request.ActualLrpKey)
	bbsErr := models.ConvertError(err)
	if bbsErr != nil && bbsErr.Type != models.Error_ResourceNotFound {
		response.Error = bbsErr
		response.KeepContainer = true
		return
	}
}

func (h *EvacuationHandler) EvacuateCrashedActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("evacuate-crashed-actual-lrp")
	logger.Info("started")
	defer logger.Info("completed")

	request := &models.EvacuateCrashedActualLRPRequest{}
	response := &models.EvacuationResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		logger.Error("failed-parsing-request", err)
		response.Error = models.ConvertError(err)
		return
	}

	beforeActualLRPGroup, err := h.actualLRPDB.ActualLRPGroupByProcessGuidAndIndex(logger, request.ActualLrpKey.ProcessGuid, request.ActualLrpKey.Index)
	if err == nil {
		err = h.db.RemoveEvacuatingActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey)
		if err != nil {
			if err == models.ErrNoTable {
				logger.Error("failed-actual-lrps-table-does-not-exist", err)
				h.exitChan <- struct{}{}
			}
			logger.Error("failed-removing-evacuating-actual-lrp", err)
		} else {
			go h.actualHub.Emit(models.NewActualLRPRemovedEvent(beforeActualLRPGroup))
		}
	}

	_, _, _, err = h.actualLRPDB.CrashActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey, request.ErrorMessage)
	bbsErr := models.ConvertError(err)
	if bbsErr != nil && bbsErr.Type != models.Error_ResourceNotFound {
		logger.Error("failed-crashing-actual-lrp", err)
		response.Error = bbsErr
		return
	}
}

func (h *EvacuationHandler) EvacuateRunningActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("evacuate-running-actual-lrp")

	response := &models.EvacuationResponse{}
	response.KeepContainer = true
	defer writeResponse(w, response)

	request := &models.EvacuateRunningActualLRPRequest{}
	err := parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	guid := request.ActualLrpKey.ProcessGuid
	index := request.ActualLrpKey.Index
	actualLRPGroup, err := h.actualLRPDB.ActualLRPGroupByProcessGuidAndIndex(logger, guid, index)
	if err != nil {
		if err == models.ErrResourceNotFound {
			response.KeepContainer = false
			return
		}
		logger.Error("failed-fetching-lrp-group", err)
		response.Error = models.ConvertError(err)
		return
	}

	instance := actualLRPGroup.Instance
	evacuating := actualLRPGroup.Evacuating

	// If the instance is not there, clean up the corresponding evacuating LRP, if one exists.
	if instance == nil {
		err = h.db.RemoveEvacuatingActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey)
		if err != nil {
			if err == models.ErrNoTable {
				logger.Error("failed-actual-lrps-table-does-not-exist", err)
				h.exitChan <- struct{}{}
			}
			if err == models.ErrActualLRPCannotBeRemoved {
				logger.Debug("remove-evacuating-actual-lrp-failed")
				response.KeepContainer = false
				return
			}
			logger.Error("failed-removing-evacuating-actual-lrp", err)
			response.Error = models.ConvertError(err)
			return
		}

		go h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Evacuating: evacuating}))
		response.KeepContainer = false
		return
	}

	if (instance.State == models.ActualLRPStateUnclaimed && instance.PlacementError == "") ||
		(instance.State == models.ActualLRPStateClaimed && !instance.ActualLRPInstanceKey.Equal(request.ActualLrpInstanceKey)) {
		if evacuating != nil && !evacuating.ActualLRPInstanceKey.Equal(request.ActualLrpInstanceKey) {
			logger.Error("already-evacuated-by-different-cell", err)
			response.KeepContainer = false
			return
		}

		group, err := h.db.EvacuateActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey, request.ActualLrpNetInfo, request.Ttl)
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		if err == models.ErrActualLRPCannotBeEvacuated {
			logger.Error("cannot-evacuate-actual-lrp", err)
			response.KeepContainer = false
			return
		}

		response.KeepContainer = true

		if err != nil {
			logger.Error("failed-evacuating-actual-lrp", err)
			response.Error = models.ConvertError(err)
		} else {
			go h.actualHub.Emit(models.NewActualLRPCreatedEvent(group))
		}

		return
	}

	if (instance.State == models.ActualLRPStateUnclaimed && instance.PlacementError != "") ||
		(instance.State == models.ActualLRPStateRunning && !instance.ActualLRPInstanceKey.Equal(request.ActualLrpInstanceKey)) ||
		instance.State == models.ActualLRPStateCrashed {
		response.KeepContainer = false
		err = h.db.RemoveEvacuatingActualLRP(logger, &evacuating.ActualLRPKey, &evacuating.ActualLRPInstanceKey)
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		if err == nil {
			go h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Evacuating: evacuating}))
		}
		if err != nil && err != models.ErrActualLRPCannotBeRemoved {
			response.KeepContainer = true
			response.Error = models.ConvertError(err)
		}
		return
	}

	if (instance.State == models.ActualLRPStateClaimed || instance.State == models.ActualLRPStateRunning) &&
		instance.ActualLRPInstanceKey.Equal(request.ActualLrpInstanceKey) {
		group, err := h.db.EvacuateActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey, request.ActualLrpNetInfo, request.Ttl)
		if err != nil {
			if err == models.ErrNoTable {
				logger.Error("failed-actual-lrps-table-does-not-exist", err)
				h.exitChan <- struct{}{}
			}
			response.Error = models.ConvertError(err)
			return
		}

		go h.actualHub.Emit(models.NewActualLRPCreatedEvent(group))

		err = h.unclaimAndRequestAuction(logger, request.ActualLrpKey)
		if err != nil {
			response.Error = models.ConvertError(err)
			return
		}
	}
}

func (h *EvacuationHandler) EvacuateStoppedActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("evacuate-stopped-actual-lrp")

	request := &models.EvacuateStoppedActualLRPRequest{}
	response := &models.EvacuationResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		logger.Error("failed-to-parse-request", err)
		response.Error = models.ConvertError(err)
		return
	}

	guid := request.ActualLrpKey.ProcessGuid
	index := request.ActualLrpKey.Index

	group, err := h.actualLRPDB.ActualLRPGroupByProcessGuidAndIndex(logger, guid, index)
	if err != nil {
		logger.Error("failed-fetching-actual-lrp-group", err)
		response.Error = models.ConvertError(err)
		return
	}

	err = h.db.RemoveEvacuatingActualLRP(logger, request.ActualLrpKey, request.ActualLrpInstanceKey)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		logger.Error("failed-removing-evacuating-actual-lrp", err)
	} else if group.Evacuating != nil {
		go h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Evacuating: group.Evacuating}))
	}

	if group.Instance == nil || !group.Instance.ActualLRPInstanceKey.Equal(request.ActualLrpInstanceKey) {
		logger.Debug("cannot-remove-actual-lrp")
		response.Error = models.ErrActualLRPCannotBeRemoved
		return
	}

	err = h.actualLRPDB.RemoveActualLRP(logger, guid, index, request.ActualLrpInstanceKey)
	if err != nil {
		logger.Error("failed-to-remove-actual-lrp", err)
		response.Error = models.ConvertError(err)
		return
	} else {
		go h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Instance: group.Instance}))
	}
}

func (h *EvacuationHandler) unclaimAndRequestAuction(logger lager.Logger, lrpKey *models.ActualLRPKey) error {
	before, after, err := h.actualLRPDB.UnclaimActualLRP(logger, lrpKey)
	if err != nil {
		return err
	}

	go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))

	desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, lrpKey.ProcessGuid)
	if err != nil {
		logger.Error("failed-fetching-desired-lrp", err)
		return nil
	}

	schedInfo := desiredLRP.DesiredLRPSchedulingInfo()
	startRequest := auctioneer.NewLRPStartRequestFromSchedulingInfo(&schedInfo, int(lrpKey.Index))
	err = h.auctioneerClient.RequestLRPAuctions([]*auctioneer.LRPStartRequest{&startRequest})
	if err != nil {
		logger.Error("failed-requesting-auction", err)
	}

	return nil
}
