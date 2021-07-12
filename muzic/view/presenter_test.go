package view

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"cleanarch/livecode/muzic/domain"
)

//-- MOCK & STUB
type mockView struct {
	albums string
}

func (m *mockView) displayAlbums(albums string) {
	m.albums = albums
}

//-- CONST
var seededAlbums = []domain.Album{
	{
		Title: "La danse des canards",
		Year:  "1980",
	},
	{
		Title: "Tata yoyo",
		Year:  "1981",
	},
}
const expectedDisplay = "La danse des canards\nTata yoyo\n"

//-- TEST
func TestDisplayAlbums(t *testing.T) {
	// Given
	mockView := mockView{}
	presenter := Presenter{view: &mockView}

	// When
	presenter.DisplayAlbums(seededAlbums)

	// Then
	assert.Equal(t, expectedDisplay, mockView.albums)
}
