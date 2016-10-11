// This file was generated by counterfeiter
package fakes

import (
	"autoscaler/db"
	"autoscaler/models"
	"sync"
)

type FakeScheduleDB struct {
	GetActiveScheduleStub        func(appId string) (*models.ActiveSchedule, error)
	getActiveScheduleMutex       sync.RWMutex
	getActiveScheduleArgsForCall []struct {
		appId string
	}
	getActiveScheduleReturns struct {
		result1 *models.ActiveSchedule
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScheduleDB) GetActiveSchedule(appId string) (*models.ActiveSchedule, error) {
	fake.getActiveScheduleMutex.Lock()
	fake.getActiveScheduleArgsForCall = append(fake.getActiveScheduleArgsForCall, struct {
		appId string
	}{appId})
	fake.recordInvocation("GetActiveSchedule", []interface{}{appId})
	fake.getActiveScheduleMutex.Unlock()
	if fake.GetActiveScheduleStub != nil {
		return fake.GetActiveScheduleStub(appId)
	} else {
		return fake.getActiveScheduleReturns.result1, fake.getActiveScheduleReturns.result2
	}
}

func (fake *FakeScheduleDB) GetActiveScheduleCallCount() int {
	fake.getActiveScheduleMutex.RLock()
	defer fake.getActiveScheduleMutex.RUnlock()
	return len(fake.getActiveScheduleArgsForCall)
}

func (fake *FakeScheduleDB) GetActiveScheduleArgsForCall(i int) string {
	fake.getActiveScheduleMutex.RLock()
	defer fake.getActiveScheduleMutex.RUnlock()
	return fake.getActiveScheduleArgsForCall[i].appId
}

func (fake *FakeScheduleDB) GetActiveScheduleReturns(result1 *models.ActiveSchedule, result2 error) {
	fake.GetActiveScheduleStub = nil
	fake.getActiveScheduleReturns = struct {
		result1 *models.ActiveSchedule
		result2 error
	}{result1, result2}
}

func (fake *FakeScheduleDB) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeScheduleDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeScheduleDB) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScheduleDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getActiveScheduleMutex.RLock()
	defer fake.getActiveScheduleMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScheduleDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.ScheduleDB = new(FakeScheduleDB)
