package book

import graphql "github.com/graph-gophers/graphql-go"

type Book struct {
	ID          graphql.ID
	Name        string
	Description string
	PublishDate string
	AuthorID    graphql.ID
}

type BookResolver struct {
	Data *Book
}

func (book *BookResolver) ID() graphql.ID {
	return book.Data.ID
}

func (book *BookResolver) Name() string {
	return book.Data.Name
}

func (book *BookResolver) Description() string {
	return book.Data.Description
}

func (book *BookResolver) PublishDate() string {
	return book.Data.PublishDate
}

func (book *BookResolver) AuthorID() graphql.ID {
	return book.Data.AuthorID
}
