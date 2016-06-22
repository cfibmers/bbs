package handlers

import (
	"errors"
	"net/http"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type DomainHandler struct {
	db       db.DomainDB
	exitChan chan<- struct{}
	logger   lager.Logger
}

var (
	ErrDomainMissing = errors.New("domain missing from request")
	ErrMaxAgeMissing = errors.New("max-age directive missing from request")
)

func NewDomainHandler(logger lager.Logger, db db.DomainDB, exitChan chan<- struct{}) *DomainHandler {
	return &DomainHandler{
		db:       db,
		exitChan: exitChan,
		logger:   logger.Session("domain-handler"),
	}
}

func (h *DomainHandler) Domains(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("domains")
	response := &models.DomainsResponse{}
	response.Domains, err = h.db.Domains(logger)
	if err == models.ErrNoTable {
		logger.Error("failed-domains-table-does-not-exist", err)
		h.exitChan <- struct{}{}
	}
	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *DomainHandler) Upsert(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("upsert")

	request := &models.UpsertDomainRequest{}
	response := &models.UpsertDomainResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		err = h.db.UpsertDomain(logger, request.Domain, request.Ttl)
		if err == models.ErrNoTable {
			logger.Error("failed-domains-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}
