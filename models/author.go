package models

import "github.com/graph-gophers/graphql-go"

type Author struct {
	ID         graphql.ID
	Name       string
	SecondName string
}

type AuthorResolver struct {
	Data *Author
}

func (author *AuthorResolver) ID() graphql.ID {
	return author.Data.ID
}

func (author *AuthorResolver) Name() string {
	return author.Data.Name
}

func (author *AuthorResolver) SecondName() string {
	return author.Data.SecondName
}

func (author *AuthorResolver) Books() []*BookResolver {

	books := make([]*BookResolver, 0, 1)

	for _, book := range dm.Books {
		if book.Data.AuthorID == author.Data.ID {
			books = append(books, book)
		}
	}

	return books
}
