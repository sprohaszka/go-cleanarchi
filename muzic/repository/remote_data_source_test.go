package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"cleanarch/livecode/muzic/domain"
)

//-- MOCK & STUB
type mockClient struct {
	mock.Mock
}

func (s mockClient) Get(url string) ([]byte, error) {
	args := s.Called(url)
	return args.Get(0).([]byte), args.Error(1)
}

//-- CONST
var seededAlbums = []byte(`{"releases": [{"title": "La danse des canards", "date": "1980"}]}`)
var expectedAlbums = []domain.Album{
	{
		Title: "La danse des canards",
		Year:  "1980",
	},
}

//-- TEST
func TestGetAlbums(t *testing.T) {
	// Given
	mockClient := mockClient{}
	mockClient.On("Get", bestArtistUrl).Return(seededAlbums, nil)
	dataSource := RemoteDataSource{
		client: mockClient,
	}

	// When
	albums := dataSource.GetAlbums()

	// Then
	assert.Equal(t, expectedAlbums, albums)
}
