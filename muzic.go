package main

import (
	"fmt"
	"log"
	"net/http"

  "cleanarch/livecode/muzic/domain"
  "cleanarch/livecode/muzic/repository"
	"cleanarch/livecode/muzic/thirdparties"
	"cleanarch/livecode/muzic/view"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Start listening")

	api := inject()

	result := api.RetrieveAlbums()
	log.Println("Results retrieved!")

	fmt.Fprintf(w, "%s", result)
}

func inject() view.Api {
  client := thirdparties.BuildRestClient()

  var api view.Api

  repository := repository.BuildRepository(
    repository.BuildRemoteDataSource(client),
    repository.EmptyDataSource{},
  )

  presenter := view.BuildPresenter(&api)

  interactor := domain.BuildInteractor(repository, presenter)

  controller := view.BuildController(interactor)

  api = view.BuildApi(controller)

  return api
}