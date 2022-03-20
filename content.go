package content

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ellemouton/aperture-demo/contentrpc"
	"github.com/ellemouton/aperture-demo/db"
	"google.golang.org/grpc"
)

type Server struct {
	DB *db.DB

	*pb.UnimplementedContentServer
	contentServer *grpc.Server
}

func NewServer() (*Server, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	return &Server{
		DB: db,
	}, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	s.contentServer = grpc.NewServer()

	pb.RegisterContentServer(s.contentServer, s)

	log.Printf("Content Server serving at %s", "localhost:8080")
	go func() {
		if err := s.contentServer.Serve(lis); err != nil {
			fmt.Printf("error starting content server: %v\n", err)
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.contentServer.Stop()

	return s.DB.Close()
}

var _ pb.ContentServer = (*Server)(nil)

func (s *Server) AddArticle(_ context.Context,
	req *pb.AddArticleRequest) (*pb.AddArticleResponse, error) {

	err := s.DB.AddArticle(&db.Article{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddArticleResponse{}, nil
}
