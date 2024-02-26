package types

type InputAuthor struct {
	Author string `json:"author" binding:"required"`
}

type InputPublisher struct {
	Publisher string `json:"publisher" binding:"required"`
}

type InputBookshelf struct {
	Bookshelf string `json:"bookshelf" binding:"required"`
}

type InputCategory struct {
	Category string `json:"category" binding:"required"`
}

type InputLanguage struct {
	Language string `json:"language" binding:"required"`
}

type InputBook struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Isbn        string `json:"isbn" binding:"required"`
	Year        string `json:"year" binding:"required"`
	Pages       int    `json:"pages" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publisher   string `json:"publisher" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Bookshelf   string `json:"bookshelf" binding:"required"`
	Language    string `json:"language" binding:"required"`
	Image       string `json:"image"`
}

type UpdateBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Isbn        string `json:"isbn"`
	Year        string `json:"year"`
	Pages       int    `json:"pages"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Category    string `json:"category"`
	Bookshelf   string `json:"bookshelf"`
	Language    string `json:"language"`
	Image       string `json:"image"`
}
