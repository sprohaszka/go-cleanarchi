package repository

import "cleanarch/livecode/muzic/domain"

//-- STRUCT
type BestAlbumRepository struct {
	remoteDataSource DataSource
	localDataSource DataSource
}

//-- BUILDER
func BuildRepository(remoteDataSource DataSource, localDataSource DataSource) domain.Repository {
	return BestAlbumRepository{
		remoteDataSource: remoteDataSource,
		localDataSource: localDataSource,
	}
}

//-- INTERFACE
type DataSource interface {
	GetAlbums() []domain.Album
}

//-- METHODS
func (r BestAlbumRepository) GetAlbums() []domain.Album {
	if albums := r.localDataSource.GetAlbums(); len(albums) > 0 {
		return albums;
	} else {
		return r.remoteDataSource.GetAlbums()
	}
}