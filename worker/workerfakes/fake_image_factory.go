// This file was generated by counterfeiter
package workerfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
)

type FakeImageFactory struct {
	GetImageStub        func(lager.Logger, worker.Worker, worker.VolumeClient, worker.ImageSpec, int, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, atc.ResourceTypes) (worker.Image, error)
	getImageMutex       sync.RWMutex
	getImageArgsForCall []struct {
		arg1  lager.Logger
		arg2  worker.Worker
		arg3  worker.VolumeClient
		arg4  worker.ImageSpec
		arg5  int
		arg6  <-chan os.Signal
		arg7  worker.ImageFetchingDelegate
		arg8  worker.Identifier
		arg9  worker.Metadata
		arg10 atc.ResourceTypes
	}
	getImageReturns struct {
		result1 worker.Image
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFactory) GetImage(arg1 lager.Logger, arg2 worker.Worker, arg3 worker.VolumeClient, arg4 worker.ImageSpec, arg5 int, arg6 <-chan os.Signal, arg7 worker.ImageFetchingDelegate, arg8 worker.Identifier, arg9 worker.Metadata, arg10 atc.ResourceTypes) (worker.Image, error) {
	fake.getImageMutex.Lock()
	fake.getImageArgsForCall = append(fake.getImageArgsForCall, struct {
		arg1  lager.Logger
		arg2  worker.Worker
		arg3  worker.VolumeClient
		arg4  worker.ImageSpec
		arg5  int
		arg6  <-chan os.Signal
		arg7  worker.ImageFetchingDelegate
		arg8  worker.Identifier
		arg9  worker.Metadata
		arg10 atc.ResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.recordInvocation("GetImage", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.getImageMutex.Unlock()
	if fake.GetImageStub != nil {
		return fake.GetImageStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	} else {
		return fake.getImageReturns.result1, fake.getImageReturns.result2
	}
}

func (fake *FakeImageFactory) GetImageCallCount() int {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	return len(fake.getImageArgsForCall)
}

func (fake *FakeImageFactory) GetImageArgsForCall(i int) (lager.Logger, worker.Worker, worker.VolumeClient, worker.ImageSpec, int, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, atc.ResourceTypes) {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	return fake.getImageArgsForCall[i].arg1, fake.getImageArgsForCall[i].arg2, fake.getImageArgsForCall[i].arg3, fake.getImageArgsForCall[i].arg4, fake.getImageArgsForCall[i].arg5, fake.getImageArgsForCall[i].arg6, fake.getImageArgsForCall[i].arg7, fake.getImageArgsForCall[i].arg8, fake.getImageArgsForCall[i].arg9, fake.getImageArgsForCall[i].arg10
}

func (fake *FakeImageFactory) GetImageReturns(result1 worker.Image, result2 error) {
	fake.GetImageStub = nil
	fake.getImageReturns = struct {
		result1 worker.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImageFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.ImageFactory = new(FakeImageFactory)
