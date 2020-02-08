package models

import "github.com/graph-gophers/graphql-go"

type Resolver struct {
	//Books   map[graphql.ID]*book.Book
	//uthor map[graphql.ID]*author.Author
	//omment map[graphql.ID]*comment.Comment
}

func (r *Resolver) Book(args struct{ ID graphql.ID }) *BookResolver {
	return &BookResolver{&Book{
		ID:          "1",
		Name:        "mybook",
		Description: "",
		PublishDate: "",
		AuthorID:    "1",
	}}
}

func (r *Resolver) Author(args struct{ ID graphql.ID }) *AuthorResolver {
	return &AuthorResolver{&Author{
		ID:         "1",
		Name:       "",
		SecondName: "",
	}}
}

func (r *Resolver) Comment(args struct{ ID graphql.ID }) *CommentResolver {
	return &CommentResolver{&Comment{
		ID:     "1",
		Text:   "",
		BookID: "1",
		UserID: "1",
	}}
}

func (r *Resolver) Authors() []*AuthorResolver {
	var authors []*AuthorResolver

	first_author := &AuthorResolver{&Author{
		ID:         "1",
		Name:       "name",
		SecondName: "",
	}}

	authors = append(authors, first_author)

	return authors
}

func (r *Resolver) Books() []*BookResolver {
	var books []*BookResolver
	first_book := &BookResolver{&Book{
		ID:          "1",
		Name:        "mybook",
		Description: "",
		PublishDate: "",
		AuthorID:    "1",
	}}

	books = append(books, first_book)

	return books
}
