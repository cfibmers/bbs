package handlers

import (
	"net/http"
	"time"

	"github.com/cloudfoundry-incubator/auctioneer"
	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/bbs/taskworkpool"
	"github.com/cloudfoundry-incubator/rep"
	"github.com/pivotal-golang/lager"
)

type TaskHandler struct {
	db                   db.TaskDB
	logger               lager.Logger
	taskCompletionClient taskworkpool.TaskCompletionClient
	auctioneerClient     auctioneer.Client
	serviceClient        bbs.ServiceClient
	repClientFactory     rep.ClientFactory
	exitChan             chan struct{}
}

func NewTaskHandler(
	logger lager.Logger,
	db db.TaskDB,
	taskCompletionClient taskworkpool.TaskCompletionClient,
	auctioneerClient auctioneer.Client,
	serviceClient bbs.ServiceClient,
	repClientFactory rep.ClientFactory,
	exitChan chan struct{},
) *TaskHandler {
	return &TaskHandler{
		db:                   db,
		logger:               logger.Session("task-handler"),
		taskCompletionClient: taskCompletionClient,
		auctioneerClient:     auctioneerClient,
		serviceClient:        serviceClient,
		repClientFactory:     repClientFactory,
		exitChan:             exitChan,
	}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("tasks")

	request := &models.TasksRequest{}
	response := &models.TasksResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		filter := models.TaskFilter{Domain: request.Domain, CellID: request.CellId}
		response.Tasks, err = h.db.Tasks(logger, filter)
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *TaskHandler) TaskByGuid(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("task-by-guid")

	request := &models.TaskByGuidRequest{}
	response := &models.TaskResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		response.Task, err = h.db.TaskByGuid(logger, request.TaskGuid)
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *TaskHandler) DesireTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("desire-task")

	request := &models.DesireTaskRequest{}
	response := &models.TaskLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	err = h.db.DesireTask(logger, request.TaskDefinition, request.TaskGuid, request.Domain)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	taskStartRequest := auctioneer.NewTaskStartRequestFromModel(request.TaskGuid, request.Domain, request.TaskDefinition)
	err = h.auctioneerClient.RequestTaskAuctions([]*auctioneer.TaskStartRequest{&taskStartRequest})
	if err != nil {
		logger.Error("failed-requesting-task-auction", err)
		// The creation succeeded, the auction request error can be dropped
	} else {
		logger.Debug("succeeded-requesting-task-auction")
	}
}

func (h *TaskHandler) StartTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("start-task")

	request := &models.StartTaskRequest{}
	response := &models.StartTaskResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		logger = logger.WithData(lager.Data{"task-guid": request.TaskGuid, "cell-id": request.CellId})
		response.ShouldStart, err = h.db.StartTask(logger, request.TaskGuid, request.CellId)
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *TaskHandler) CancelTask(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("cancel-task")

	request := &models.TaskGuidRequest{}
	response := &models.TaskLifecycleResponse{}
	defer writeResponse(w, response)

	err := parseRequest(logger, req, request)
	if err != nil {
		logger.Error("failed-parsing-request", err)
		response.Error = models.ConvertError(err)
		return
	}

	task, cellID, err := h.db.CancelTask(logger, request.TaskGuid)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	if task.CompletionCallbackUrl != "" {
		logger.Info("task-client-completing-task")
		go h.taskCompletionClient.Submit(h.db, task)
	}

	if cellID == "" {
		return
	}

	cellPresence, err := h.serviceClient.CellById(logger, cellID)
	if err != nil {
		logger.Error("failed-fetching-cell-presence", err)
		return
	}

	repClient := h.repClientFactory.CreateClient(cellPresence.RepAddress)
	repClient.CancelTask(request.TaskGuid)
	if err != nil {
		logger.Error("failed-rep-cancel-task", err)
		return
	}

	logger.Info("cell-client-succeeded-cancelling-task")
}

func (h *TaskHandler) FailTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("fail-task")

	request := &models.FailTaskRequest{}
	response := &models.TaskLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	task, err := h.db.FailTask(logger, request.TaskGuid, request.FailureReason)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	if task.CompletionCallbackUrl != "" {
		logger.Info("task-client-completing-task")
		go h.taskCompletionClient.Submit(h.db, task)
	}
}

func (h *TaskHandler) CompleteTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("complete-task")

	request := &models.CompleteTaskRequest{}
	response := &models.TaskLifecycleResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)
	if err != nil {
		response.Error = models.ConvertError(err)
		return
	}

	task, err := h.db.CompleteTask(logger, request.TaskGuid, request.CellId, request.Failed, request.FailureReason, request.Result)
	if err != nil {
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
		response.Error = models.ConvertError(err)
		return
	}

	if task.CompletionCallbackUrl != "" {
		logger.Info("task-client-completing-task")
		go h.taskCompletionClient.Submit(h.db, task)
	}
}

func (h *TaskHandler) ResolvingTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("resolving-task")

	request := &models.TaskGuidRequest{}
	response := &models.TaskLifecycleResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		err = h.db.ResolvingTask(logger, request.TaskGuid)
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("delete-task")

	request := &models.TaskGuidRequest{}
	response := &models.TaskLifecycleResponse{}

	err = parseRequest(logger, req, request)
	if err == nil {
		err = h.db.DeleteTask(logger, request.TaskGuid)
		if err == models.ErrNoTable {
			logger.Error("failed-task-table-does-not-exist", err)
			h.exitChan <- struct{}{}
		}
	}

	response.Error = models.ConvertError(err)
	writeResponse(w, response)
}

func (h *TaskHandler) ConvergeTasks(w http.ResponseWriter, req *http.Request) {
	var err error
	logger := h.logger.Session("converge-tasks")

	request := &models.ConvergeTasksRequest{}
	response := &models.ConvergeTasksResponse{}

	defer writeResponse(w, response)

	err = parseRequest(logger, req, request)

	if err != nil {
		response.Error = models.ConvertError(err)
	}

	logger.Debug("listing-cells")
	cellSet, err := h.serviceClient.Cells(logger)
	if err == models.ErrResourceNotFound {
		logger.Debug("no-cells-found")
		cellSet = models.CellSet{}
	} else if err != nil {
		logger.Debug("failed-listing-cells")
		return
	}
	logger.Debug("succeeded-listing-cells")

	tasksToAuction, tasksToComplete := h.db.ConvergeTasks(
		logger,
		cellSet,
		time.Duration(request.KickTaskDuration),
		time.Duration(request.ExpirePendingTaskDuration),
		time.Duration(request.ExpireCompletedTaskDuration),
	)

	if len(tasksToAuction) > 0 {
		logger.Debug("requesting-task-auctions", lager.Data{"num-tasks-to-auction": len(tasksToAuction)})
		if err := h.auctioneerClient.RequestTaskAuctions(tasksToAuction); err != nil {
			taskGuids := make([]string, len(tasksToAuction))
			for i, task := range tasksToAuction {
				taskGuids[i] = task.TaskGuid
			}
			logger.Error("failed-to-request-auctions-for-pending-tasks", err,
				lager.Data{"task-guids": taskGuids})
		}
		logger.Debug("done-requesting-task-auctions", lager.Data{"num-tasks-to-auction": len(tasksToAuction)})
	}

	logger.Debug("submitting-tasks-to-be-completed", lager.Data{"num-tasks-to-complete": len(tasksToComplete)})
	for _, task := range tasksToComplete {
		h.taskCompletionClient.Submit(h.db, task)
	}
	logger.Debug("done-submitting-tasks-to-be-completed", lager.Data{"num-tasks-to-complete": len(tasksToComplete)})
}
