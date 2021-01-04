package todo_app

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler: 		handler,
		MaxHeaderBytes: 1 << 20,          //1 MB
		ReadTimeout:    10 * time.Second, // 10 seconds
		WriteTimeout:   10 * time.Second, // 10 seconds
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}

