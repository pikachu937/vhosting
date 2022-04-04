package vhs

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	maxHeaderBytes = 1 << 20 // 1 MB
	readTimeOut    = 15
	writeTimeOut   = 15
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg SVConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeOut * time.Second,
		WriteTimeout:   writeTimeOut * time.Second,
	}

	logrus.Printf("server host (%s) has started at port %s\n", cfg.Host, cfg.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Printf("server shut down\n")
	return s.httpServer.Shutdown(ctx)
}
