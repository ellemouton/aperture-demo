package content

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "github.com/ellemouton/aperture-demo/contentrpc"
	"github.com/ellemouton/aperture-demo/db"
	"github.com/gorilla/mux"
	pricespb "github.com/lightninglabs/aperture/pricesrpc"
	"google.golang.org/grpc"
)

type Server struct {
	DB *db.DB

	*pb.UnimplementedContentServer
	contentServer *grpc.Server

	pricesServer *grpc.Server
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
	// Start the Content gRPC server.
	s.contentServer = grpc.NewServer()
	pb.RegisterContentServer(s.contentServer, s)

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	log.Printf("Content Server serving at %s", "localhost:8080")
	go func() {
		if err := s.contentServer.Serve(lis); err != nil {
			fmt.Printf("error starting content server: %v\n", err)
		}
	}()

	// Start the Content gRPC server.
	s.pricesServer = grpc.NewServer()
	pricespb.RegisterPricesServer(s.pricesServer, s)

	lis2, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		return err
	}

	log.Printf("Prices Server serving at %s", "localhost:8083")
	go func() {
		if err := s.pricesServer.Serve(lis2); err != nil {
			fmt.Printf("error starting content server: %v\n", err)
		}
	}()

	// Start the http server that listens for content requests.
	r := mux.NewRouter()
	r.HandleFunc("/test", freebeeHandler).Methods("GET")
	r.HandleFunc("/article/{id}", s.articleHandler).Methods("GET")
	r.HandleFunc("/quote/{id}", s.quoteHandler).Methods("GET")

	log.Printf("Serving HTTP server on port %s", "localhost:9000")
	go func() {
		if err := http.ListenAndServe("localhost:9000", r); err != nil {
			fmt.Printf("error starting http server: %v\n", err)
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

	id, err := s.DB.AddArticle(&db.Article{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddArticleResponse{Id: int64(id)}, nil
}

func (s *Server) AddQuote(_ context.Context,
	req *pb.AddQuoteRequest) (*pb.AddQuoteResponse, error) {

	id, err := s.DB.AddQuote(&db.Quote{
		Author:  req.Author,
		Content: req.Content,
		Price:   req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddQuoteResponse{Id: int64(id)}, nil
}

func freebeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Freebee endpoint test")
}

func (s *Server) articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	article, err := s.DB.GetArticle(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("Title: %s\nAuthor: %s\nContent: %s\n",
		article.Title, article.Author, article.Content)

	fmt.Fprintln(w, resp)
}

func (s *Server) quoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quote, err := s.DB.GetQuote(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("Quote Author: %s\nContent: %s\nPrice: %d\n",
		quote.Author, quote.Content, quote.Price)

	fmt.Fprintln(w, resp)
}
