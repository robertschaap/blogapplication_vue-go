package main

import (
	c "./controllers"
)

func (s *server) routes() {
	s.r.HandleFunc("/api/comments/new", c.CreateCommentHandler).Methods("POST")
	s.r.HandleFunc("/api/posts", c.GetPostsHandler).Methods("GET")
	s.r.HandleFunc("/api/posts/{category}", c.GetPostsByCategoryHandler).Methods("GET")
	s.r.HandleFunc("/api/post/{id}", c.GetPostHandler).Methods("GET")
	s.r.HandleFunc("/api/post/new", c.CreatePostHandler).Methods("POST")
	s.r.HandleFunc("/api/users/new", c.CreateUserHandler).Methods("POST")
	s.r.HandleFunc("/api/users/login", c.GetUserLoginHandler).Methods("GET")
}
