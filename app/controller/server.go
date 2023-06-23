package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/mohamed-ach/go-repo/blog-rest-api/controller/blog"
)

var router *mux.Router

func initHandlers() {

	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/{id}", controller.GetPost).Methods("GET")

	router.HandleFunc("/api/post/new", controller.CreatePost).Methods("POST")

	router.HandleFunc("/api/post/update", controller.UpdatePost).Methods("PUT")

	router.HandleFunc("/api/post/delete/{id}", controller.DeletePost).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
