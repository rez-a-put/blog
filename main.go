package main

import (
	h "blog/handler"
	u "blog/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()
}

func main() {
	// posts
	r.HandleFunc("/api/posts", h.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", h.AddPost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", h.ModifyPost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", h.RemovePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{id}", h.PublishPost).Methods("PATCH")

	// tags
	r.HandleFunc("/api/tags", h.GetTags).Methods("GET")
	r.HandleFunc("/api/tags/{id}", h.GetTag).Methods("GET")
	r.HandleFunc("/api/tags", h.AddTag).Methods("POST")
	r.HandleFunc("/api/tags/{id}", h.ModifyTag).Methods("PUT")
	r.HandleFunc("/api/tags/{id}", h.RemoveTag).Methods("DELETE")

	log.Fatal(http.ListenAndServe(u.GetEnvByKey("SERVER_HOST"), r))
}
