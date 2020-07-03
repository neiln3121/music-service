// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/neiln3121/music-service/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Repository) Close() {
	_m.Called()
}

// GetAlbum provides a mock function with given fields: id
func (_m *Repository) GetAlbum(id string) (*models.Album, error) {
	ret := _m.Called(id)

	var r0 *models.Album
	if rf, ok := ret.Get(0).(func(string) *models.Album); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Album)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArtists provides a mock function with given fields: offset, limit
func (_m *Repository) GetArtists(offset int, limit int) ([]*models.Artist, error) {
	ret := _m.Called(offset, limit)

	var r0 []*models.Artist
	if rf, ok := ret.Get(0).(func(int, int) []*models.Artist); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Artist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}