// Code generated by mockery v2.10.0. DO NOT EDIT.

package handler

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// UserHandlerInterface is an autogenerated mock type for the UserHandlerInterface type
type UserHandlerInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: c
func (_m *UserHandlerInterface) CreateUser(c *gin.Context) {
	_m.Called(c)
}

// DeleteSpecificUser provides a mock function with given fields: c
func (_m *UserHandlerInterface) DeleteSpecificUser(c *gin.Context) {
	_m.Called(c)
}

// DeleteUsersWithIds provides a mock function with given fields: c
func (_m *UserHandlerInterface) DeleteUsersWithIds(c *gin.Context) {
	_m.Called(c)
}

// GetAllUser provides a mock function with given fields: c
func (_m *UserHandlerInterface) GetAllUser(c *gin.Context) {
	_m.Called(c)
}

// GetSpecificUser provides a mock function with given fields: c
func (_m *UserHandlerInterface) GetSpecificUser(c *gin.Context) {
	_m.Called(c)
}

// QueryUsers provides a mock function with given fields: c
func (_m *UserHandlerInterface) QueryUsers(c *gin.Context) {
	_m.Called(c)
}

// UpdateSpecificUser provides a mock function with given fields: c
func (_m *UserHandlerInterface) UpdateSpecificUser(c *gin.Context) {
	_m.Called(c)
}