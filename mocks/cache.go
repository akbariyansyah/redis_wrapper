// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCache) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCacheMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCache)(nil).Close))
}

// DecrBy mocks base method.
func (m *MockCache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", ctx, key, value)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockCacheMockRecorder) DecrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockCache)(nil).DecrBy), ctx, key, value)
}

// Delete mocks base method.
func (m *MockCache) Delete(ctx context.Context, keys ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), varargs...)
}

// Exist mocks base method.
func (m *MockCache) Exist(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exist indicates an expected call of Exist.
func (mr *MockCacheMockRecorder) Exist(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockCache)(nil).Exist), ctx, key)
}

// Expire mocks base method.
func (m *MockCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockCacheMockRecorder) Expire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockCache)(nil).Expire), ctx, key, expiration)
}

// Forever mocks base method.
func (m *MockCache) Forever(ctx context.Context, key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Forever", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Forever indicates an expected call of Forever.
func (mr *MockCacheMockRecorder) Forever(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forever", reflect.TypeOf((*MockCache)(nil).Forever), ctx, key, value)
}

// Get mocks base method.
func (m *MockCache) Get(ctx context.Context, key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), ctx, key, value)
}

// IncrBy mocks base method.
func (m *MockCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", ctx, key, value)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockCacheMockRecorder) IncrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockCache)(nil).IncrBy), ctx, key, value)
}

// Keys mocks base method.
func (m *MockCache) Keys(ctx context.Context, prefixKey string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", ctx, prefixKey)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Keys indicates an expected call of Keys.
func (mr *MockCacheMockRecorder) Keys(ctx, prefixKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCache)(nil).Keys), ctx, prefixKey)
}

// Put mocks base method.
func (m *MockCache) Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockCacheMockRecorder) Put(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCache)(nil).Put), ctx, key, value, expiration)
}