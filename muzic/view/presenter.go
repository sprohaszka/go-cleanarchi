package view

import (
	"cleanarch/livecode/muzic/domain"
)

//-- STRUCT
type Presenter struct {
	view View
}

//-- BUILDER
func BuildPresenter(view View) domain.Presenter {
	return Presenter{view: view}
}

//-- INTERFACE
type View interface {
	displayAlbums(string)
}

//-- METHODS
func (p Presenter) DisplayAlbums(albums []domain.Album) {
	p.view.displayAlbums(convertAlbum(albums))
}

func convertAlbum(albums []domain.Album) string {
	var titles string
	for _, album := range albums {
		titles += album.Title
		titles += "\n"
	}

	return titles
}
