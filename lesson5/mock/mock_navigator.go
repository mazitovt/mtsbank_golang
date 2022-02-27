// Code generated by MockGen. DO NOT EDIT.
// Source: ./lesson5/geocoding/geocoder.go

// Package mock is a generated GoMock package.
package mock

import (
	points "mtsbank_golang/lesson5/distance/points"
	geocoding "mtsbank_golang/lesson5/geocoding"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGeocoder is a mock of Geocoder interface.
type MockGeocoder struct {
	ctrl     *gomock.Controller
	recorder *MockGeocoderMockRecorder
}

// MockGeocoderMockRecorder is the mock recorder for MockGeocoder.
type MockGeocoderMockRecorder struct {
	mock *MockGeocoder
}

// NewMockGeocoder creates a new mock instance.
func NewMockGeocoder(ctrl *gomock.Controller) *MockGeocoder {
	mock := &MockGeocoder{ctrl: ctrl}
	mock.recorder = &MockGeocoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeocoder) EXPECT() *MockGeocoderMockRecorder {
	return m.recorder
}

// Geocode mock base method.
func (m *MockGeocoder) Geocode(address string) (points.PointOnSphere, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Geocode", address)
	ret0, _ := ret[0].(points.PointOnSphere)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Geocode indicates an expected call of Geocode.
func (mr *MockGeocoderMockRecorder) Geocode(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Geocode", reflect.TypeOf((*MockGeocoder)(nil).Geocode), address)
}

// ReverseGeocode mock base method.
func (m *MockGeocoder) ReverseGeocode(point points.PointOnSphere) (geocoding.GeocodeData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseGeocode", point)
	ret0, _ := ret[0].(geocoding.GeocodeData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseGeocode indicates an expected call of ReverseGeocode.
func (mr *MockGeocoderMockRecorder) ReverseGeocode(point interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseGeocode", reflect.TypeOf((*MockGeocoder)(nil).ReverseGeocode), point)
}
