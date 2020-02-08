package comment

import "github.com/graph-gophers/graphql-go"

type Comment struct {
	ID     graphql.ID
	Text   string
	BookID graphql.ID
	UserID graphql.ID
}

type CommentResolver struct {
	Data *Comment
}

func (comment *CommentResolver) ID() graphql.ID {
	return comment.Data.ID
}

func (comment *CommentResolver) Text() string {
	return comment.Data.Text
}

func (comment *CommentResolver) BookId() graphql.ID {
	return comment.Data.BookID
}

func (comment *CommentResolver) UserId() graphql.ID {
	return comment.Data.UserID
}
