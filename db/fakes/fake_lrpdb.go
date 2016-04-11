// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/auctioneer"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeLRPDB struct {
	ActualLRPGroupsStub        func(logger lager.Logger, filter models.ActualLRPFilter) ([]*models.ActualLRPGroup, error)
	actualLRPGroupsMutex       sync.RWMutex
	actualLRPGroupsArgsForCall []struct {
		logger lager.Logger
		filter models.ActualLRPFilter
	}
	actualLRPGroupsReturns struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupsByProcessGuidStub        func(logger lager.Logger, processGuid string) ([]*models.ActualLRPGroup, error)
	actualLRPGroupsByProcessGuidMutex       sync.RWMutex
	actualLRPGroupsByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	actualLRPGroupsByProcessGuidReturns struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupByProcessGuidAndIndexStub        func(logger lager.Logger, processGuid string, index int32) (*models.ActualLRPGroup, error)
	actualLRPGroupByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPGroupByProcessGuidAndIndexArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}
	actualLRPGroupByProcessGuidAndIndexReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	CreateUnclaimedActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey) (actualLRPGroup *models.ActualLRPGroup, err error)
	createUnclaimedActualLRPMutex       sync.RWMutex
	createUnclaimedActualLRPArgsForCall []struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}
	createUnclaimedActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	UnclaimActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey) (beforeActualLRPGroup *models.ActualLRPGroup, err error)
	unclaimActualLRPMutex       sync.RWMutex
	unclaimActualLRPArgsForCall []struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}
	unclaimActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	ClaimActualLRPStub        func(logger lager.Logger, processGuid string, index int32, instanceKey *models.ActualLRPInstanceKey) (beforeActualLRPGroup *models.ActualLRPGroup, err error)
	claimActualLRPMutex       sync.RWMutex
	claimActualLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
		instanceKey *models.ActualLRPInstanceKey
	}
	claimActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	StartActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) (beforeActualLRPGroup *models.ActualLRPGroup, updated bool, err error)
	startActualLRPMutex       sync.RWMutex
	startActualLRPArgsForCall []struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}
	startActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 bool
		result3 error
	}
	CrashActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, crashReason string) (beforeActualLRPGroup *models.ActualLRPGroup, shouldRestart bool, err error)
	crashActualLRPMutex       sync.RWMutex
	crashActualLRPArgsForCall []struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		crashReason string
	}
	crashActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 bool
		result3 error
	}
	FailActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, placementError string) (beforeActualLRPGroup *models.ActualLRPGroup, err error)
	failActualLRPMutex       sync.RWMutex
	failActualLRPArgsForCall []struct {
		logger         lager.Logger
		key            *models.ActualLRPKey
		placementError string
	}
	failActualLRPReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	RemoveActualLRPStub        func(logger lager.Logger, processGuid string, index int32) error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}
	removeActualLRPReturns struct {
		result1 error
	}
	DesiredLRPsStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) (*models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	DesiredLRPSchedulingInfosStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error)
	desiredLRPSchedulingInfosMutex       sync.RWMutex
	desiredLRPSchedulingInfosArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPSchedulingInfosReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	DesireLRPStub        func(logger lager.Logger, desiredLRP *models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	UpdateDesiredLRPStub        func(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (beforeDesiredLRP *models.DesiredLRP, err error)
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	RemoveDesiredLRPStub        func(logger lager.Logger, processGuid string) error
	removeDesiredLRPMutex       sync.RWMutex
	removeDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	removeDesiredLRPReturns struct {
		result1 error
	}
	ConvergeLRPsStub        func(logger lager.Logger, cellSet models.CellSet) (startRequest []*auctioneer.LRPStartRequest, keysToRetire []*models.ActualLRPKey)
	convergeLRPsMutex       sync.RWMutex
	convergeLRPsArgsForCall []struct {
		logger  lager.Logger
		cellSet models.CellSet
	}
	convergeLRPsReturns struct {
		result1 []*auctioneer.LRPStartRequest
		result2 []*models.ActualLRPKey
	}
	GatherAndPruneLRPsStub        func(logger lager.Logger, cellSet models.CellSet) (*models.ConvergenceInput, error)
	gatherAndPruneLRPsMutex       sync.RWMutex
	gatherAndPruneLRPsArgsForCall []struct {
		logger  lager.Logger
		cellSet models.CellSet
	}
	gatherAndPruneLRPsReturns struct {
		result1 *models.ConvergenceInput
		result2 error
	}
}

func (fake *FakeLRPDB) ActualLRPGroups(logger lager.Logger, filter models.ActualLRPFilter) ([]*models.ActualLRPGroup, error) {
	fake.actualLRPGroupsMutex.Lock()
	fake.actualLRPGroupsArgsForCall = append(fake.actualLRPGroupsArgsForCall, struct {
		logger lager.Logger
		filter models.ActualLRPFilter
	}{logger, filter})
	fake.actualLRPGroupsMutex.Unlock()
	if fake.ActualLRPGroupsStub != nil {
		return fake.ActualLRPGroupsStub(logger, filter)
	} else {
		return fake.actualLRPGroupsReturns.result1, fake.actualLRPGroupsReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupsCallCount() int {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return len(fake.actualLRPGroupsArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupsArgsForCall(i int) (lager.Logger, models.ActualLRPFilter) {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return fake.actualLRPGroupsArgsForCall[i].logger, fake.actualLRPGroupsArgsForCall[i].filter
}

func (fake *FakeLRPDB) ActualLRPGroupsReturns(result1 []*models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsStub = nil
	fake.actualLRPGroupsReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuid(logger lager.Logger, processGuid string) ([]*models.ActualLRPGroup, error) {
	fake.actualLRPGroupsByProcessGuidMutex.Lock()
	fake.actualLRPGroupsByProcessGuidArgsForCall = append(fake.actualLRPGroupsByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.actualLRPGroupsByProcessGuidMutex.Unlock()
	if fake.ActualLRPGroupsByProcessGuidStub != nil {
		return fake.ActualLRPGroupsByProcessGuidStub(logger, processGuid)
	} else {
		return fake.actualLRPGroupsByProcessGuidReturns.result1, fake.actualLRPGroupsByProcessGuidReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidCallCount() int {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPGroupsByProcessGuidArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return fake.actualLRPGroupsByProcessGuidArgsForCall[i].logger, fake.actualLRPGroupsByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidReturns(result1 []*models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsByProcessGuidStub = nil
	fake.actualLRPGroupsByProcessGuidReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndex(logger lager.Logger, processGuid string, index int32) (*models.ActualLRPGroup, error) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPGroupByProcessGuidAndIndexArgsForCall = append(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}{logger, processGuid, index})
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPGroupByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPGroupByProcessGuidAndIndexStub(logger, processGuid, index)
	} else {
		return fake.actualLRPGroupByProcessGuidAndIndexReturns.result1, fake.actualLRPGroupByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexCallCount() int {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexArgsForCall(i int) (lager.Logger, string, int32) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].logger, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].processGuid, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].index
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupByProcessGuidAndIndexStub = nil
	fake.actualLRPGroupByProcessGuidAndIndexReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) CreateUnclaimedActualLRP(logger lager.Logger, key *models.ActualLRPKey) (actualLRPGroup *models.ActualLRPGroup, err error) {
	fake.createUnclaimedActualLRPMutex.Lock()
	fake.createUnclaimedActualLRPArgsForCall = append(fake.createUnclaimedActualLRPArgsForCall, struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}{logger, key})
	fake.createUnclaimedActualLRPMutex.Unlock()
	if fake.CreateUnclaimedActualLRPStub != nil {
		return fake.CreateUnclaimedActualLRPStub(logger, key)
	} else {
		return fake.createUnclaimedActualLRPReturns.result1, fake.createUnclaimedActualLRPReturns.result2
	}
}

func (fake *FakeLRPDB) CreateUnclaimedActualLRPCallCount() int {
	fake.createUnclaimedActualLRPMutex.RLock()
	defer fake.createUnclaimedActualLRPMutex.RUnlock()
	return len(fake.createUnclaimedActualLRPArgsForCall)
}

func (fake *FakeLRPDB) CreateUnclaimedActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey) {
	fake.createUnclaimedActualLRPMutex.RLock()
	defer fake.createUnclaimedActualLRPMutex.RUnlock()
	return fake.createUnclaimedActualLRPArgsForCall[i].logger, fake.createUnclaimedActualLRPArgsForCall[i].key
}

func (fake *FakeLRPDB) CreateUnclaimedActualLRPReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.CreateUnclaimedActualLRPStub = nil
	fake.createUnclaimedActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) UnclaimActualLRP(logger lager.Logger, key *models.ActualLRPKey) (beforeActualLRPGroup *models.ActualLRPGroup, err error) {
	fake.unclaimActualLRPMutex.Lock()
	fake.unclaimActualLRPArgsForCall = append(fake.unclaimActualLRPArgsForCall, struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}{logger, key})
	fake.unclaimActualLRPMutex.Unlock()
	if fake.UnclaimActualLRPStub != nil {
		return fake.UnclaimActualLRPStub(logger, key)
	} else {
		return fake.unclaimActualLRPReturns.result1, fake.unclaimActualLRPReturns.result2
	}
}

func (fake *FakeLRPDB) UnclaimActualLRPCallCount() int {
	fake.unclaimActualLRPMutex.RLock()
	defer fake.unclaimActualLRPMutex.RUnlock()
	return len(fake.unclaimActualLRPArgsForCall)
}

func (fake *FakeLRPDB) UnclaimActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey) {
	fake.unclaimActualLRPMutex.RLock()
	defer fake.unclaimActualLRPMutex.RUnlock()
	return fake.unclaimActualLRPArgsForCall[i].logger, fake.unclaimActualLRPArgsForCall[i].key
}

func (fake *FakeLRPDB) UnclaimActualLRPReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.UnclaimActualLRPStub = nil
	fake.unclaimActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) ClaimActualLRP(logger lager.Logger, processGuid string, index int32, instanceKey *models.ActualLRPInstanceKey) (beforeActualLRPGroup *models.ActualLRPGroup, err error) {
	fake.claimActualLRPMutex.Lock()
	fake.claimActualLRPArgsForCall = append(fake.claimActualLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
		instanceKey *models.ActualLRPInstanceKey
	}{logger, processGuid, index, instanceKey})
	fake.claimActualLRPMutex.Unlock()
	if fake.ClaimActualLRPStub != nil {
		return fake.ClaimActualLRPStub(logger, processGuid, index, instanceKey)
	} else {
		return fake.claimActualLRPReturns.result1, fake.claimActualLRPReturns.result2
	}
}

func (fake *FakeLRPDB) ClaimActualLRPCallCount() int {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return len(fake.claimActualLRPArgsForCall)
}

func (fake *FakeLRPDB) ClaimActualLRPArgsForCall(i int) (lager.Logger, string, int32, *models.ActualLRPInstanceKey) {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return fake.claimActualLRPArgsForCall[i].logger, fake.claimActualLRPArgsForCall[i].processGuid, fake.claimActualLRPArgsForCall[i].index, fake.claimActualLRPArgsForCall[i].instanceKey
}

func (fake *FakeLRPDB) ClaimActualLRPReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.ClaimActualLRPStub = nil
	fake.claimActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) StartActualLRP(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) (beforeActualLRPGroup *models.ActualLRPGroup, updated bool, err error) {
	fake.startActualLRPMutex.Lock()
	fake.startActualLRPArgsForCall = append(fake.startActualLRPArgsForCall, struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}{logger, key, instanceKey, netInfo})
	fake.startActualLRPMutex.Unlock()
	if fake.StartActualLRPStub != nil {
		return fake.StartActualLRPStub(logger, key, instanceKey, netInfo)
	} else {
		return fake.startActualLRPReturns.result1, fake.startActualLRPReturns.result2, fake.startActualLRPReturns.result3
	}
}

func (fake *FakeLRPDB) StartActualLRPCallCount() int {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return len(fake.startActualLRPArgsForCall)
}

func (fake *FakeLRPDB) StartActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return fake.startActualLRPArgsForCall[i].logger, fake.startActualLRPArgsForCall[i].key, fake.startActualLRPArgsForCall[i].instanceKey, fake.startActualLRPArgsForCall[i].netInfo
}

func (fake *FakeLRPDB) StartActualLRPReturns(result1 *models.ActualLRPGroup, result2 bool, result3 error) {
	fake.StartActualLRPStub = nil
	fake.startActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLRPDB) CrashActualLRP(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, crashReason string) (beforeActualLRPGroup *models.ActualLRPGroup, shouldRestart bool, err error) {
	fake.crashActualLRPMutex.Lock()
	fake.crashActualLRPArgsForCall = append(fake.crashActualLRPArgsForCall, struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		crashReason string
	}{logger, key, instanceKey, crashReason})
	fake.crashActualLRPMutex.Unlock()
	if fake.CrashActualLRPStub != nil {
		return fake.CrashActualLRPStub(logger, key, instanceKey, crashReason)
	} else {
		return fake.crashActualLRPReturns.result1, fake.crashActualLRPReturns.result2, fake.crashActualLRPReturns.result3
	}
}

func (fake *FakeLRPDB) CrashActualLRPCallCount() int {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return len(fake.crashActualLRPArgsForCall)
}

func (fake *FakeLRPDB) CrashActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return fake.crashActualLRPArgsForCall[i].logger, fake.crashActualLRPArgsForCall[i].key, fake.crashActualLRPArgsForCall[i].instanceKey, fake.crashActualLRPArgsForCall[i].crashReason
}

func (fake *FakeLRPDB) CrashActualLRPReturns(result1 *models.ActualLRPGroup, result2 bool, result3 error) {
	fake.CrashActualLRPStub = nil
	fake.crashActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLRPDB) FailActualLRP(logger lager.Logger, key *models.ActualLRPKey, placementError string) (beforeActualLRPGroup *models.ActualLRPGroup, err error) {
	fake.failActualLRPMutex.Lock()
	fake.failActualLRPArgsForCall = append(fake.failActualLRPArgsForCall, struct {
		logger         lager.Logger
		key            *models.ActualLRPKey
		placementError string
	}{logger, key, placementError})
	fake.failActualLRPMutex.Unlock()
	if fake.FailActualLRPStub != nil {
		return fake.FailActualLRPStub(logger, key, placementError)
	} else {
		return fake.failActualLRPReturns.result1, fake.failActualLRPReturns.result2
	}
}

func (fake *FakeLRPDB) FailActualLRPCallCount() int {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return len(fake.failActualLRPArgsForCall)
}

func (fake *FakeLRPDB) FailActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, string) {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return fake.failActualLRPArgsForCall[i].logger, fake.failActualLRPArgsForCall[i].key, fake.failActualLRPArgsForCall[i].placementError
}

func (fake *FakeLRPDB) FailActualLRPReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.FailActualLRPStub = nil
	fake.failActualLRPReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) RemoveActualLRP(logger lager.Logger, processGuid string, index int32) error {
	fake.removeActualLRPMutex.Lock()
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}{logger, processGuid, index})
	fake.removeActualLRPMutex.Unlock()
	if fake.RemoveActualLRPStub != nil {
		return fake.RemoveActualLRPStub(logger, processGuid, index)
	} else {
		return fake.removeActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeLRPDB) RemoveActualLRPArgsForCall(i int) (lager.Logger, string, int32) {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return fake.removeActualLRPArgsForCall[i].logger, fake.removeActualLRPArgsForCall[i].processGuid, fake.removeActualLRPArgsForCall[i].index
}

func (fake *FakeLRPDB) RemoveActualLRPReturns(result1 error) {
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDB) DesiredLRPs(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(logger, filter)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeLRPDB) DesiredLRPsArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].logger, fake.desiredLRPsArgsForCall[i].filter
}

func (fake *FakeLRPDB) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuid(logger lager.Logger, processGuid string) (*models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(logger, processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].logger, fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) DesiredLRPSchedulingInfos(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	fake.desiredLRPSchedulingInfosArgsForCall = append(fake.desiredLRPSchedulingInfosArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.desiredLRPSchedulingInfosMutex.Unlock()
	if fake.DesiredLRPSchedulingInfosStub != nil {
		return fake.DesiredLRPSchedulingInfosStub(logger, filter)
	} else {
		return fake.desiredLRPSchedulingInfosReturns.result1, fake.desiredLRPSchedulingInfosReturns.result2
	}
}

func (fake *FakeLRPDB) DesiredLRPSchedulingInfosCallCount() int {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return len(fake.desiredLRPSchedulingInfosArgsForCall)
}

func (fake *FakeLRPDB) DesiredLRPSchedulingInfosArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return fake.desiredLRPSchedulingInfosArgsForCall[i].logger, fake.desiredLRPSchedulingInfosArgsForCall[i].filter
}

func (fake *FakeLRPDB) DesiredLRPSchedulingInfosReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.DesiredLRPSchedulingInfosStub = nil
	fake.desiredLRPSchedulingInfosReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) DesireLRP(logger lager.Logger, desiredLRP *models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}{logger, desiredLRP})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(logger, desiredLRP)
	} else {
		return fake.desireLRPReturns.result1
	}
}

func (fake *FakeLRPDB) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeLRPDB) DesireLRPArgsForCall(i int) (lager.Logger, *models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].logger, fake.desireLRPArgsForCall[i].desiredLRP
}

func (fake *FakeLRPDB) DesireLRPReturns(result1 error) {
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDB) UpdateDesiredLRP(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (beforeDesiredLRP *models.DesiredLRP, err error) {
	fake.updateDesiredLRPMutex.Lock()
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}{logger, processGuid, update})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(logger, processGuid, update)
	} else {
		return fake.updateDesiredLRPReturns.result1, fake.updateDesiredLRPReturns.result2
	}
}

func (fake *FakeLRPDB) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeLRPDB) UpdateDesiredLRPArgsForCall(i int) (lager.Logger, string, *models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return fake.updateDesiredLRPArgsForCall[i].logger, fake.updateDesiredLRPArgsForCall[i].processGuid, fake.updateDesiredLRPArgsForCall[i].update
}

func (fake *FakeLRPDB) UpdateDesiredLRPReturns(result1 *models.DesiredLRP, result2 error) {
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDB) RemoveDesiredLRP(logger lager.Logger, processGuid string) error {
	fake.removeDesiredLRPMutex.Lock()
	fake.removeDesiredLRPArgsForCall = append(fake.removeDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.removeDesiredLRPMutex.Unlock()
	if fake.RemoveDesiredLRPStub != nil {
		return fake.RemoveDesiredLRPStub(logger, processGuid)
	} else {
		return fake.removeDesiredLRPReturns.result1
	}
}

func (fake *FakeLRPDB) RemoveDesiredLRPCallCount() int {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return len(fake.removeDesiredLRPArgsForCall)
}

func (fake *FakeLRPDB) RemoveDesiredLRPArgsForCall(i int) (lager.Logger, string) {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return fake.removeDesiredLRPArgsForCall[i].logger, fake.removeDesiredLRPArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) RemoveDesiredLRPReturns(result1 error) {
	fake.RemoveDesiredLRPStub = nil
	fake.removeDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDB) ConvergeLRPs(logger lager.Logger, cellSet models.CellSet) (startRequest []*auctioneer.LRPStartRequest, keysToRetire []*models.ActualLRPKey) {
	fake.convergeLRPsMutex.Lock()
	fake.convergeLRPsArgsForCall = append(fake.convergeLRPsArgsForCall, struct {
		logger  lager.Logger
		cellSet models.CellSet
	}{logger, cellSet})
	fake.convergeLRPsMutex.Unlock()
	if fake.ConvergeLRPsStub != nil {
		return fake.ConvergeLRPsStub(logger, cellSet)
	} else {
		return fake.convergeLRPsReturns.result1, fake.convergeLRPsReturns.result2
	}
}

func (fake *FakeLRPDB) ConvergeLRPsCallCount() int {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return len(fake.convergeLRPsArgsForCall)
}

func (fake *FakeLRPDB) ConvergeLRPsArgsForCall(i int) (lager.Logger, models.CellSet) {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return fake.convergeLRPsArgsForCall[i].logger, fake.convergeLRPsArgsForCall[i].cellSet
}

func (fake *FakeLRPDB) ConvergeLRPsReturns(result1 []*auctioneer.LRPStartRequest, result2 []*models.ActualLRPKey) {
	fake.ConvergeLRPsStub = nil
	fake.convergeLRPsReturns = struct {
		result1 []*auctioneer.LRPStartRequest
		result2 []*models.ActualLRPKey
	}{result1, result2}
}

func (fake *FakeLRPDB) GatherAndPruneLRPs(logger lager.Logger, cellSet models.CellSet) (*models.ConvergenceInput, error) {
	fake.gatherAndPruneLRPsMutex.Lock()
	fake.gatherAndPruneLRPsArgsForCall = append(fake.gatherAndPruneLRPsArgsForCall, struct {
		logger  lager.Logger
		cellSet models.CellSet
	}{logger, cellSet})
	fake.gatherAndPruneLRPsMutex.Unlock()
	if fake.GatherAndPruneLRPsStub != nil {
		return fake.GatherAndPruneLRPsStub(logger, cellSet)
	} else {
		return fake.gatherAndPruneLRPsReturns.result1, fake.gatherAndPruneLRPsReturns.result2
	}
}

func (fake *FakeLRPDB) GatherAndPruneLRPsCallCount() int {
	fake.gatherAndPruneLRPsMutex.RLock()
	defer fake.gatherAndPruneLRPsMutex.RUnlock()
	return len(fake.gatherAndPruneLRPsArgsForCall)
}

func (fake *FakeLRPDB) GatherAndPruneLRPsArgsForCall(i int) (lager.Logger, models.CellSet) {
	fake.gatherAndPruneLRPsMutex.RLock()
	defer fake.gatherAndPruneLRPsMutex.RUnlock()
	return fake.gatherAndPruneLRPsArgsForCall[i].logger, fake.gatherAndPruneLRPsArgsForCall[i].cellSet
}

func (fake *FakeLRPDB) GatherAndPruneLRPsReturns(result1 *models.ConvergenceInput, result2 error) {
	fake.GatherAndPruneLRPsStub = nil
	fake.gatherAndPruneLRPsReturns = struct {
		result1 *models.ConvergenceInput
		result2 error
	}{result1, result2}
}

var _ db.LRPDB = new(FakeLRPDB)
