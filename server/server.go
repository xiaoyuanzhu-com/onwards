package server

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Server struct {
	app *pocketbase.PocketBase
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	s.app = pocketbase.New()
	s.app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./ui/dist"), false))
		return se.Next()
	})
	if err := s.app.Start(); err != nil {
		log.Printf("failed to start pocketbase: %v", err)
	}
}
