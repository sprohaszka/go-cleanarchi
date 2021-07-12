package view

import "cleanarch/livecode/muzic/domain"

//-- STRUCT
type Controller struct {
	interactor domain.Interactor
}

//-- BUILDER
func BuildController(interactor domain.Interactor) Controller {
	return Controller{interactor: interactor}
}

//-- METHODS
func (c Controller) retrieveAlbums() {
	c.interactor.GetBestAlbumsSortedByYear()
}
