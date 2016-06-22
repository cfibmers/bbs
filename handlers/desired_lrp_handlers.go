package handlers

import (
	"net/http"

	"github.com/cloudfoundry-incubator/auctioneer"
	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/rep"
	"github.com/cloudfoundry/gunk/workpool"
	"github.com/pivotal-golang/lager"
)

type DesiredLRPHandler struct {
	desiredLRPDB       db.DesiredLRPDB
	actualLRPDB        db.ActualLRPDB
	desiredHub         events.Hub
	actualHub          events.Hub
	auctioneerClient   auctioneer.Client
	repClientFactory   rep.ClientFactory
	serviceClient      bbs.ServiceClient
	updateWorkersCount int
	logger             lager.Logger
	exitChan           chan struct{}
}

func NewDesiredLRPHandler(
	logger lager.Logger,
	updateWorkersCount int,
	desiredLRPDB db.DesiredLRPDB,
	actualLRPDB db.ActualLRPDB,
	desiredHub events.Hub,
	actualHub events.Hub,
	auctioneerClient auctioneer.Client,
	repClientFactory rep.ClientFactory,
	serviceClient bbs.ServiceClient,
	exitChan chan struct{},
) *DesiredLRPHandler {
	return &DesiredLRPHandler{
		desiredLRPDB:       desiredLRPDB,
		actualLRPDB:        actualLRPDB,
		desiredHub:         desiredHub,
		actualHub:          actualHub,
		auctioneerClient:   auctioneerClient,
		repClientFactory:   repClientFactory,
		serviceClient:      serviceClient,
		updateWorkersCount: updateWorkersCount,
		logger:             logger.Session("desired-lrp-handler"),
		exitChan:           exitChan,
	}
}

func (h *DesiredLRPHandler) DesiredLRPs(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("desired-lrps")

	request := &models.DesiredLRPsRequest{}
	response := &models.DesiredLRPsResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		filter := models.DesiredLRPFilter{Domain: request.Domain}
		response.DesiredLrps, err = h.desiredLRPDB.DesiredLRPs(logger, filter)
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *DesiredLRPHandler) DesiredLRPByProcessGuid(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("desired-lrp-by-process-guid")

	request := &models.DesiredLRPByProcessGuidRequest{}
	response := &models.DesiredLRPResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		response.DesiredLrp, err = h.desiredLRPDB.DesiredLRPByProcessGuid(logger, request.ProcessGuid)
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *DesiredLRPHandler) DesiredLRPSchedulingInfos(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("desired-lrp-scheduling-infos")

	request := &models.DesiredLRPsRequest{}
	response := &models.DesiredLRPSchedulingInfosResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		filter := models.DesiredLRPFilter{Domain: request.Domain}
		response.DesiredLrpSchedulingInfos, err = h.desiredLRPDB.DesiredLRPSchedulingInfos(logger, filter)
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *DesiredLRPHandler) DesireDesiredLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("desire-lrp")

	request := &models.DesireLRPRequest{}
	response := &models.DesiredLRPLifecycleResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	err = h.desiredLRPDB.DesireLRP(logger, request.DesiredLrp)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, request.DesiredLrp.ProcessGuid)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	go h.desiredHub.Emit(models.NewDesiredLRPCreatedEvent(desiredLRP))

	schedulingInfo := request.DesiredLrp.DesiredLRPSchedulingInfo()
	h.startInstanceRange(logger, 0, schedulingInfo.Instances, &schedulingInfo)
}

func (h *DesiredLRPHandler) UpdateDesiredLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("update-desired-lrp")

	request := &models.UpdateDesiredLRPRequest{}
	response := &models.DesiredLRPLifecycleResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		logger.Error("failed-parsing-request", err)
		response.Error = models.ConvertError(err)
		return
	}

	logger = logger.WithData(lager.Data{"guid": request.ProcessGuid})

	logger.Debug("updating-desired-lrp")
	beforeDesiredLRP, err := h.desiredLRPDB.UpdateDesiredLRP(logger, request.ProcessGuid, request.Update)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		logger.Debug("failed-updating-desired-lrp")
		response.Error = models.ConvertError(err)
		return
	}
	logger.Debug("completed-updating-desired-lrp")

	desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, request.ProcessGuid)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		logger.Error("failed-fetching-desired-lrp", err)
		return
	}

	if request.Update.Instances != nil {
		logger.Debug("updating-lrp-instances")
		previousInstanceCount := beforeDesiredLRP.Instances

		requestedInstances := *request.Update.Instances - previousInstanceCount

		logger = logger.WithData(lager.Data{"instances-delta": requestedInstances})
		if requestedInstances > 0 {
			logger.Debug("increasing-the-instances")
			schedulingInfo := desiredLRP.DesiredLRPSchedulingInfo()
			h.startInstanceRange(logger, previousInstanceCount, *request.Update.Instances, &schedulingInfo)
		}

		if requestedInstances < 0 {
			logger.Debug("decreasing-the-instances")
			numExtraActualLRP := previousInstanceCount + requestedInstances
			h.stopInstancesFrom(logger, request.ProcessGuid, int(numExtraActualLRP))
		}
	}

	go h.desiredHub.Emit(models.NewDesiredLRPChangedEvent(beforeDesiredLRP, desiredLRP))
}

func (h *DesiredLRPHandler) RemoveDesiredLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("remove-desired-lrp")

	request := &models.RemoveDesiredLRPRequest{}
	response := &models.DesiredLRPLifecycleResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, request.ProcessGuid)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	err = h.desiredLRPDB.RemoveDesiredLRP(logger, request.ProcessGuid)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-desired-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	go h.desiredHub.Emit(models.NewDesiredLRPRemovedEvent(desiredLRP))

	h.stopInstancesFrom(logger, request.ProcessGuid, 0)
}

func (h *DesiredLRPHandler) startInstanceRange(logger lager.Logger, lower, upper int32, schedulingInfo *models.DesiredLRPSchedulingInfo) {
	logger = logger.Session("start-instance-range", lager.Data{"lower": lower, "upper": upper})
	logger.Info("starting")
	defer logger.Info("complete")

	keys := make([]*models.ActualLRPKey, upper-lower)
	i := 0
	for actualIndex := lower; actualIndex < upper; actualIndex++ {
		key := models.NewActualLRPKey(schedulingInfo.ProcessGuid, int32(actualIndex), schedulingInfo.Domain)
		keys[i] = &key
		i++
	}

	createdIndices := h.createUnclaimedActualLRPs(logger, keys)
	start := auctioneer.NewLRPStartRequestFromSchedulingInfo(schedulingInfo, createdIndices...)

	err := h.auctioneerClient.RequestLRPAuctions([]*auctioneer.LRPStartRequest{&start})
	if err != nil {
		logger.Error("failed-to-request-auction", err)
	}
}

func (h *DesiredLRPHandler) createUnclaimedActualLRPs(logger lager.Logger, keys []*models.ActualLRPKey) []int {
	count := len(keys)
	createdIndicesChan := make(chan int, count)

	works := make([]func(), count)

	for i, key := range keys {
		key := key
		works[i] = func() {
			actualLRPGroup, err := h.actualLRPDB.CreateUnclaimedActualLRP(logger, key)
			if err != nil {
				if err == models.ErrNoTable {
					logger.Error("failed-actual-lrps-table-does-not-exist", err)
					h.exitChan <- struct{}{}
				}
				logger.Info("failed-creating-actual-lrp", lager.Data{"actual_lrp_key": key, "err-message": err.Error()})
			} else {
				go h.actualHub.Emit(models.NewActualLRPCreatedEvent(actualLRPGroup))
				createdIndicesChan <- int(key.Index)
			}
		}
	}

	throttlerSize := h.updateWorkersCount
	throttler, err := workpool.NewThrottler(throttlerSize, works)
	if err != nil {
		logger.Error("failed-constructing-throttler", err, lager.Data{"max-workers": throttlerSize, "num-works": len(works)})
		return []int{}
	}

	go func() {
		throttler.Work()
		close(createdIndicesChan)
	}()

	createdIndices := make([]int, 0, count)
	for createdIndex := range createdIndicesChan {
		createdIndices = append(createdIndices, createdIndex)
	}

	return createdIndices
}

func (h *DesiredLRPHandler) stopInstancesFrom(logger lager.Logger, processGuid string, index int) {
	actualLRPGroups, err := h.actualLRPDB.ActualLRPGroupsByProcessGuid(logger, processGuid)
	if err != nil {
		logger.Error("failed-fetching-actual-lrps", err)
		return
	}

	for i := 0; i < len(actualLRPGroups); i++ {
		group := actualLRPGroups[i]

		if group.Instance != nil {
			lrp := group.Instance
			if lrp.Index >= int32(index) {
				switch lrp.State {
				case models.ActualLRPStateUnclaimed, models.ActualLRPStateCrashed:
					err = h.actualLRPDB.RemoveActualLRP(logger, lrp.ProcessGuid, lrp.Index, &lrp.ActualLRPInstanceKey)
					if err != nil {
						if err == models.ErrNoTable {
							logger.Error("failed-actual-lrps-table-does-not-exist", err)
							h.exitChan <- struct{}{}
						}
						logger.Error("failed-removing-lrp-instance", err)
					}
				default:
					cellPresence, err := h.serviceClient.CellById(logger, lrp.CellId)
					if err != nil {
						logger.Error("failed-fetching-cell-presence", err)
						continue
					}
					repClient := h.repClientFactory.CreateClient(cellPresence.RepAddress)
					logger.Debug("stopping-lrp-instance")
					err = repClient.StopLRPInstance(lrp.ActualLRPKey, lrp.ActualLRPInstanceKey)
					if err != nil {
						logger.Error("failed-stopping-lrp-instance", err)
					}
				}
			}
		}
	}
}
