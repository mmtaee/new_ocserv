// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	crypto "ocserv/pkg/crypto"

	mock "github.com/stretchr/testify/mock"
)

// CustomPasswordInterface is an autogenerated mock type for the CustomPasswordInterface type
type CustomPasswordInterface struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: passwd, hashedPassword, salt
func (_m *CustomPasswordInterface) CheckPassword(passwd string, hashedPassword string, salt string) bool {
	ret := _m.Called(passwd, hashedPassword, salt)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(passwd, hashedPassword, salt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreatePassword provides a mock function with given fields: passwd, saltLength
func (_m *CustomPasswordInterface) CreatePassword(passwd string, saltLength ...int) crypto.CustomPassword {
	_va := make([]interface{}, len(saltLength))
	for _i := range saltLength {
		_va[_i] = saltLength[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, passwd)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePassword")
	}

	var r0 crypto.CustomPassword
	if rf, ok := ret.Get(0).(func(string, ...int) crypto.CustomPassword); ok {
		r0 = rf(passwd, saltLength...)
	} else {
		r0 = ret.Get(0).(crypto.CustomPassword)
	}

	return r0
}

// NewCustomPasswordInterface creates a new instance of CustomPasswordInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomPasswordInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomPasswordInterface {
	mock := &CustomPasswordInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
