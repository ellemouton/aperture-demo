package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ellemouton/aperture-demo/contentrpc"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := cli.NewApp()
	app.Name = "appcli"
	app.Usage = "Control plane for your content delivery app"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpcserver",
			Value: "localhost:8080",
			Usage: "content app daemon address host:port",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "addbook",
			Usage:  "add a book to the content delivery system",
			Action: addBook,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "title",
				},
				cli.StringFlag{
					Name: "author",
				},
				cli.StringFlag{
					Name: "content",
				},
			},
		},
		{
			Name:   "addquote",
			Usage:  "add a quote to the content delivery system",
			Action: addQuote,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "author",
				},
				cli.StringFlag{
					Name: "content",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient(ctx *cli.Context) (contentrpc.ContentClient, func(), error) {
	rpcServer := ctx.GlobalString("rpcserver")

	conn, err := grpc.Dial(rpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to RPC server: %v", err)
	}

	cleanup := func() { _ = conn.Close() }

	sessionsClient := contentrpc.NewContentClient(conn)
	return sessionsClient, cleanup, nil
}

func addBook(ctx *cli.Context) error {
	client, cleanup, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	title := ctx.String("title")
	if title == "" {
		return fmt.Errorf("must set a title for the book")
	}

	author := ctx.String("author")
	if author == "" {
		return fmt.Errorf("must set a author for the book")
	}

	content := ctx.String("content")
	if content == "" {
		return fmt.Errorf("must set content for the book")
	}

	resp, err := client.AddArticle(context.Background(),
		&contentrpc.AddArticleRequest{
			Title:   title,
			Author:  author,
			Content: content,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Success! New article id is: %d\n", resp.Id)
	return nil
}

func addQuote(ctx *cli.Context) error {
	client, cleanup, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	author := ctx.String("author")
	if author == "" {
		return fmt.Errorf("must set a author for the quote")
	}

	content := ctx.String("content")
	if content == "" {
		return fmt.Errorf("must set content for the quote")
	}

	resp, err := client.AddQuote(context.Background(),
		&contentrpc.AddQuoteRequest{
			Author:  author,
			Content: content,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Success! New quote id is: %d\n", resp.Id)
	return nil
}
