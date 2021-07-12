package repository

import (
	"cleanarch/livecode/muzic/domain"
)

type EmptyDataSource struct {}

func (r EmptyDataSource) GetAlbums() []domain.Album {
	return []domain.Album{}
}
