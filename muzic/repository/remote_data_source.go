package repository

import (
	"encoding/json"

	"cleanarch/livecode/muzic/domain"
)

//-- STRUCT
type RemoteDataSource struct {
	client RestClient
}

//-- BUILDER
func BuildRemoteDataSource(client RestClient) DataSource {
	return RemoteDataSource{client: client}
}

//-- INTERFACE
type RestClient interface {
	Get(string) ([]byte, error)
}

//-- ENTITIES
type albumEntities struct {
	Releases []albumResponse
}

type albumResponse struct {
	Title string
	Date  string
}

//-- CONST
const bestArtistUrl = "https://musicbrainz.org/ws/2/release/?artist=b665b768-0d83-4363-950c-31ed39317c15"

//-- METHODS
func (r RemoteDataSource) GetAlbums() []domain.Album {
	jsonResp, _ := r.client.Get(bestArtistUrl)

	var albumEntities albumEntities
	json.Unmarshal(jsonResp, &albumEntities)

	albums := make([]domain.Album, 0, len(albumEntities.Releases))
	for _, album := range albumEntities.Releases {
		albums = append(albums, domain.Album{
			Title: album.Title,
			Year:  album.Date,
		})
	}

	return albums
}
