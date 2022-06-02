package db

type Article struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type Quote struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Price   int64  `json:"price"`
}

type Meme struct {
	Name  string `json:"name"`
	Image string `json:"image-file"`
	Price int64  `json:"price"`
}
