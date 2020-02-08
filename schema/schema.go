package schema

import (
	"../author"
	"../book"
	"../comment"
	graphql "github.com/graph-gophers/graphql-go"
)

var Schema = `
		schema {
			query: Query
		}
		# The query type, represents all of the entry points into our object graph
		type Query {
			GetBook(id: ID!): Book 
			GetAuthor(id: ID!): Author
			GetComment(id: ID!): Comment
		}

		interface Book{
			id: ID!
			name: String!
			description: String!
			publishdate: String!
			authorid: ID!
		}

		interface Author{
			id: ID!
			name: String!
			secondname: String!
		}

		interface Comment{
			id: ID!
			text: String!
			bookid: ID!
			userid: ID!
		}
	`

type Resolver struct {
	Books   map[graphql.ID]*book.Book
	Authors map[graphql.ID]*author.Author
	Comment map[graphql.ID]*comment.Comment
}

func (r *Resolver) GetBook(args struct{ ID graphql.ID }) *book.BookResolver {
	return &book.BookResolver{&book.Book{
		ID:          "",
		Name:        "mybook",
		Description: "",
		PublishDate: "",
		AuthorID:    "",
	}}
}

func (r *Resolver) GetAuthor(args struct{ ID graphql.ID }) *author.AuthorResolver {
	return &author.AuthorResolver{&author.Author{
		ID:         "",
		Name:       "",
		SecondName: "",
	}}
}

func (r *Resolver) GetComment(args struct{ ID graphql.ID }) *comment.CommentResolver {
	return &comment.CommentResolver{&comment.Comment{
		ID:     "",
		Text:   "",
		BookID: "",
		UserID: "",
	}}
}
