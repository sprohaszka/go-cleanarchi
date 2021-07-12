package domain

import "sort"

//-- STRUCT
type Interactor struct {
	repository Repository
	presenter Presenter
}

//-- BUILDER
func BuildInteractor(repository Repository, presenter Presenter) Interactor {
	return Interactor {
		repository: repository,
		presenter: presenter,
	}
}

//-- INTERFACE
type Repository interface {
	GetAlbums() []Album
}

type Presenter interface {
	DisplayAlbums([]Album)
}

//-- MODEL
type Album struct {
	Title string
	Year string
}

//-- METHODS
func (i Interactor) GetBestAlbumsSortedByYear() {
	albums := i.repository.GetAlbums()

	sort.SliceStable(albums, func(i, j int) bool {
		return albums[i].Year < albums[j].Year
	})

	i.presenter.DisplayAlbums(albums)
}