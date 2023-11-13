package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type Server struct {
	UnimplementedShortenerServer
	Storage storage.Storage
	Server  *grpc.Server
	Cfg     *config.AppConfig
}

func NewServer(storage storage.Storage, cfg *config.AppConfig) *Server {
	return &Server{
		Storage: storage,
		Server:  grpc.NewServer(),
		Cfg:     cfg,
	}
}

func Run(s *Server) error {

	listen, err := net.Listen("tcp", ":3200")

	if err != nil {
		return err
	}

	RegisterShortenerServer(s.Server, s.UnimplementedShortenerServer)
	fmt.Println("Сервер gRPC начал работу")
	if err := s.Server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.Server.GracefulStop()
}
