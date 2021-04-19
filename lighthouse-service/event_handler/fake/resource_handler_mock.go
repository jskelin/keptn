// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package event_handler_mock

import (
	keptnapimodels "github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// ResourceHandlerMock is a mock implementation of event_handler.ResourceHandler.
//
// 	func TestSomethingThatUsesResourceHandler(t *testing.T) {
//
// 		// make and configure a mocked event_handler.ResourceHandler
// 		mockedResourceHandler := &ResourceHandlerMock{
// 			GetServiceResourceFunc: func(project string, stage string, service string, resourceURI string) (*keptnapimodels.Resource, error) {
// 				panic("mock out the GetServiceResource method")
// 			},
// 		}
//
// 		// use mockedResourceHandler in code that requires event_handler.ResourceHandler
// 		// and then make assertions.
//
// 	}
type ResourceHandlerMock struct {
	// GetServiceResourceFunc mocks the GetServiceResource method.
	GetServiceResourceFunc func(project string, stage string, service string, resourceURI string) (*keptnapimodels.Resource, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetServiceResource holds details about calls to the GetServiceResource method.
		GetServiceResource []struct {
			// Project is the project argument value.
			Project string
			// Stage is the stage argument value.
			Stage string
			// Service is the service argument value.
			Service string
			// ResourceURI is the resourceURI argument value.
			ResourceURI string
		}
	}
	lockGetServiceResource sync.RWMutex
}

// GetServiceResource calls GetServiceResourceFunc.
func (mock *ResourceHandlerMock) GetServiceResource(project string, stage string, service string, resourceURI string) (*keptnapimodels.Resource, error) {
	if mock.GetServiceResourceFunc == nil {
		panic("ResourceHandlerMock.GetServiceResourceFunc: method is nil but ResourceHandler.GetServiceResource was just called")
	}
	callInfo := struct {
		Project     string
		Stage       string
		Service     string
		ResourceURI string
	}{
		Project:     project,
		Stage:       stage,
		Service:     service,
		ResourceURI: resourceURI,
	}
	mock.lockGetServiceResource.Lock()
	mock.calls.GetServiceResource = append(mock.calls.GetServiceResource, callInfo)
	mock.lockGetServiceResource.Unlock()
	return mock.GetServiceResourceFunc(project, stage, service, resourceURI)
}

// GetServiceResourceCalls gets all the calls that were made to GetServiceResource.
// Check the length with:
//     len(mockedResourceHandler.GetServiceResourceCalls())
func (mock *ResourceHandlerMock) GetServiceResourceCalls() []struct {
	Project     string
	Stage       string
	Service     string
	ResourceURI string
} {
	var calls []struct {
		Project     string
		Stage       string
		Service     string
		ResourceURI string
	}
	mock.lockGetServiceResource.RLock()
	calls = mock.calls.GetServiceResource
	mock.lockGetServiceResource.RUnlock()
	return calls
}