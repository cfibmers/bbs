package handlers

import (
	"net/http"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type ActualLRPHandler struct {
	db       db.ActualLRPDB
	exitChan chan<- struct{}
	logger   lager.Logger
}

func NewActualLRPHandler(logger lager.Logger, db db.ActualLRPDB, exitChan chan<- struct{}) *ActualLRPHandler {
	return &ActualLRPHandler{
		db:       db,
		exitChan: exitChan,
		logger:   logger.Session("actual-lrp-handler"),
	}
}

func (h *ActualLRPHandler) ActualLRPGroups(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("actual-lrp-groups")

	request := &models.ActualLRPGroupsRequest{}
	response := &models.ActualLRPGroupsResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		filter := models.ActualLRPFilter{Domain: request.Domain, CellID: request.CellId}
		response.ActualLrpGroups, err = h.db.ActualLRPGroups(logger, filter)
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)

	writeResponse(w, response)
}

func (h *ActualLRPHandler) ActualLRPGroupsByProcessGuid(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("actual-lrp-groups-by-process-guid")

	request := &models.ActualLRPGroupsByProcessGuidRequest{}
	response := &models.ActualLRPGroupsResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		response.ActualLrpGroups, err = h.db.ActualLRPGroupsByProcessGuid(logger, request.ProcessGuid)
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)

	writeResponse(w, response)
}

func (h *ActualLRPHandler) ActualLRPGroupByProcessGuidAndIndex(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("actual-lrp-group-by-process-guid-and-index")

	request := &models.ActualLRPGroupByProcessGuidAndIndexRequest{}
	response := &models.ActualLRPGroupResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		response.ActualLrpGroup, err = h.db.ActualLRPGroupByProcessGuidAndIndex(logger, request.ProcessGuid, request.Index)
		if err == models.ErrNoTable {
			logger.Error("failed-actual-lrps-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)

	writeResponse(w, response)
}
