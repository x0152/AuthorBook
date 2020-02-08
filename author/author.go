package author

import graphql "github.com/graph-gophers/graphql-go"

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
