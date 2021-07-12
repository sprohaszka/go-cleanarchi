package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"cleanarch/livecode/muzic/domain"
)

//-- MOCK & STUB
type mockDataSource struct {
	mock.Mock
}

func (s mockDataSource) GetAlbums() []domain.Album {
	args := s.Called()
	return args.Get(0).([]domain.Album)
}

//-- CONST
var seededAlbum = []domain.Album{
	{
		Title: "LedZep1",
		Year: "1969",
	},
}

//-- TEST
func TestGetAlbums_WhenNoLocalAlbum(t *testing.T) {
	// Given
	mockRemoteDataSource := mockDataSource{}
	mockRemoteDataSource.On("GetAlbums").Return(seededAlbum)
	mockLocalDataSource := mockDataSource{}
	mockLocalDataSource.On("GetAlbums").Return([]domain.Album{})
	repository := BestAlbumRepository{
		remoteDataSource: mockRemoteDataSource,
		localDataSource: mockLocalDataSource,
	}

	// When
	albums := repository.GetAlbums()

	// Then
	mockRemoteDataSource.AssertExpectations(t)
	mockLocalDataSource.AssertExpectations(t)
	assert.Equal(t, seededAlbum, albums)
}

func TestGetAlbums_WhenLocalAlbumsAvailable(t *testing.T) {
	// Given
	mockRemoteDataSource := mockDataSource{}
	mockLocalDataSource := mockDataSource{}
	mockLocalDataSource.On("GetAlbums").Return(seededAlbum)
	repository := BestAlbumRepository{
		remoteDataSource: mockRemoteDataSource,
		localDataSource: mockLocalDataSource,
	}

	// When
	albums := repository.GetAlbums()

	// Then
	mockRemoteDataSource.AssertNotCalled(t, "GetAlbums")
	mockLocalDataSource.AssertExpectations(t)
	assert.Equal(t, seededAlbum, albums)
}
