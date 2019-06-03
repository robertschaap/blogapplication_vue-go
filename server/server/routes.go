package server

import (
	c "../controllers"
)

func (s *Server) routes() {
	s.R.HandleFunc("/api/comments/new", c.CreateCommentHandler).Methods("POST")
	s.R.HandleFunc("/api/posts", c.GetPostsHandler).Methods("GET")
	s.R.HandleFunc("/api/posts/{category}", c.GetPostsByCategoryHandler).Methods("GET")
	s.R.HandleFunc("/api/post/{id}", c.GetPostHandler).Methods("GET")
	s.R.HandleFunc("/api/post/new", c.CreatePostHandler).Methods("POST")
	s.R.HandleFunc("/api/users/new", c.CreateUserHandler).Methods("POST")
	s.R.HandleFunc("/api/users/login", c.GetUserLoginHandler).Methods("POST")
}
