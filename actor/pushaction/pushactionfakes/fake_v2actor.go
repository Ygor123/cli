// This file was generated by counterfeiter
package pushactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/v2action"
)

type FakeV2Actor struct {
	CreateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	createApplicationMutex       sync.RWMutex
	createApplicationArgsForCall []struct {
		application v2action.Application
	}
	createApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	createApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	UpdateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		application v2action.Application
	}
	updateApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2Actor) CreateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.createApplicationMutex.Lock()
	ret, specificReturn := fake.createApplicationReturnsOnCall[len(fake.createApplicationArgsForCall)]
	fake.createApplicationArgsForCall = append(fake.createApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("CreateApplication", []interface{}{application})
	fake.createApplicationMutex.Unlock()
	if fake.CreateApplicationStub != nil {
		return fake.CreateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationReturns.result1, fake.createApplicationReturns.result2, fake.createApplicationReturns.result3
}

func (fake *FakeV2Actor) CreateApplicationCallCount() int {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return len(fake.createApplicationArgsForCall)
}

func (fake *FakeV2Actor) CreateApplicationArgsForCall(i int) v2action.Application {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return fake.createApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) CreateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	fake.createApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	if fake.createApplicationReturnsOnCall == nil {
		fake.createApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].name, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UpdateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("UpdateApplication", []interface{}{application})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateApplicationReturns.result1, fake.updateApplicationReturns.result2, fake.updateApplicationReturns.result3
}

func (fake *FakeV2Actor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV2Actor) UpdateApplicationArgsForCall(i int) v2action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.updateApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) UpdateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UpdateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeV2Actor) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.V2Actor = new(FakeV2Actor)