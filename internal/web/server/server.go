package server

import (
	"net/http"

	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/service"
	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr: ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
