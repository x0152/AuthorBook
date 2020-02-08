package models

import (
	"fmt"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
}

func (r *Resolver) Book(args struct{ ID graphql.ID }) *BookResolver {
	return dm.Books[args.ID]
}

func (r *Resolver) Author(args struct{ ID graphql.ID }) *AuthorResolver {
	return dm.Authors[args.ID]
}

func (r *Resolver) Comment(args struct{ ID graphql.ID }) *CommentResolver {
	return dm.Comments[args.ID]
}

func (r *Resolver) CreateComment(args struct {
	Bookid graphql.ID
	Text   string
	Date   string
}) *CommentResolver {

	id := graphql.ID(fmt.Sprintf("%d", len(dm.Comments)))
	newComment := &Comment{ID: id, BookID: args.Bookid, Text: args.Text, Date: args.Date}
	newCommentResolve := &CommentResolver{newComment}
	dm.Comments[id] = newCommentResolve

	log.Printf("Добавлен коментарий %v", newComment)
	return newCommentResolve
}

func (r *Resolver) Authors() []*AuthorResolver {
	authors := make([]*AuthorResolver, 0, len(dm.Books))

	for _, author := range dm.Authors {
		authors = append(authors, author)
	}

	return authors
}

func (r *Resolver) Books() []*BookResolver {
	books := make([]*BookResolver, 0, len(dm.Books))

	for _, book := range dm.Books {
		books = append(books, book)
	}

	return books
}
