package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var seededAlbums = []Album{
	{
		Title: "A",
		Year:  "1980",
	},
	{
		Title: "B",
		Year:  "1990",
	},
	{
		Title: "C",
		Year:  "1970",
	},
}

type stubRepository struct{}

func (s stubRepository) GetAlbums() []Album {
	return seededAlbums
}

type mockPresenter struct {
	albums []Album
}

func (m *mockPresenter) DisplayAlbums(albums []Album) {
	m.albums = make([]Album, len(albums))
	copy(m.albums, albums)
}

func TestGetBestAlbum(t *testing.T) {
	// Given
	mockPresenter := mockPresenter{}
	interactor := Interactor{
		repository: stubRepository{},
		presenter: &mockPresenter,
	}
	expectedAlbums := []Album{
		{
			Title: "C",
			Year:  "1970",
		},
		{
			Title: "A",
			Year:  "1980",
		},
		{
			Title: "B",
			Year:  "1990",
		},
	}

	// When
	interactor.GetBestAlbumsSortedByYear()

	// Then
	assert.Equal(t, expectedAlbums, mockPresenter.albums)
}
