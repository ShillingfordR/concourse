// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler"
)

type FakeBuildFactory struct {
	CreateStub        func(atc.JobConfig, atc.ResourceConfigs, []db.BuildInput) atc.Plan
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 []db.BuildInput
	}
	createReturns struct {
		result1 atc.Plan
	}
}

func (fake *FakeBuildFactory) Create(arg1 atc.JobConfig, arg2 atc.ResourceConfigs, arg3 []db.BuildInput) atc.Plan {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 []db.BuildInput
	}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeBuildFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBuildFactory) CreateArgsForCall(i int) (atc.JobConfig, atc.ResourceConfigs, []db.BuildInput) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3
}

func (fake *FakeBuildFactory) CreateReturns(result1 atc.Plan) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 atc.Plan
	}{result1}
}

var _ scheduler.BuildFactory = new(FakeBuildFactory)
