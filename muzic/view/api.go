package view

//-- STRUCT
type Api struct {
	controller Controller
	result     chan string
}

//-- BUILDER
func BuildApi(controller Controller) Api {
	return Api{
		controller: controller,
		result:     make(chan string),
	}
}

//-- METHODS
func (a *Api) RetrieveAlbums() string {
	a.controller.retrieveAlbums()
	albums := <-a.result
	return albums
}

func (a *Api) displayAlbums(albums string) {
	go func() {
		a.result <- albums
	}()
}
