package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const dbName = "db.json"

type DB struct {
	content *content
}

type content struct {
	Articles []*Article `json:"articles"`
	Quotes   []*Quote   `json:"quotes"`
	Memes    []*Meme    `json:"memes"`
}

func NewDB() (*DB, error) {
	// If there is no existing DB, create a new one. Otherwise, load the
	// existing one.
	file, err := os.Open(dbName)
	if errors.Is(err, os.ErrNotExist) {
		// Create the file
		file, err = os.Create(dbName)
		if err != nil {
			return nil, err
		}

		return &DB{
			content: &content{},
		}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	content := &content{}

	err = json.Unmarshal(byteValue, content)
	if err != nil {
		return nil, err
	}

	return &DB{
		content: content,
	}, nil
}

func (d *DB) Close() error {
	return d.writeContent()
}

func (d *DB) writeContent() error {
	b, err := json.MarshalIndent(d.content, " ", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dbName, b, 0644)
}

func (d *DB) AddArticle(article *Article) (int, error) {
	d.content.Articles = append(d.content.Articles, article)

	if err := d.writeContent(); err != nil {
		return 0, err
	}

	return len(d.content.Articles), nil
}

func (d *DB) GetArticle(id int) (*Article, error) {
	if len(d.content.Articles) < id || id <= 0 {
		return nil, fmt.Errorf("no article with id %d", id)
	}

	return d.content.Articles[id-1], nil
}

func (d *DB) AddQuote(quote *Quote) (int, error) {
	d.content.Quotes = append(d.content.Quotes, quote)

	if err := d.writeContent(); err != nil {
		return 0, err
	}

	return len(d.content.Quotes), nil
}

func (d *DB) GetQuote(id int) (*Quote, error) {
	if len(d.content.Quotes) < id || id <= 0 {
		return nil, fmt.Errorf("no quote with id %d", id)
	}

	return d.content.Quotes[id-1], nil
}

func (d *DB) AddMeme(quote *Meme) (int, error) {
	d.content.Memes = append(d.content.Memes, meme)

	if err := d.writeContent(); err != nil {
		return 0, err
	}

	return len(d.content.Memes), nil
}

func (d *DB) GetMeme(id int) (*Meme, error) {
	if len(d.content.Memes) < id || id <= 0 {
		return nil, fmt.Errorf("no meme with id %d", id)
	}

	return d.content.Memes[id-1], nil
}