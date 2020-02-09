package models

import (
	"fmt"
	"log"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/nu7hatch/gouuid"
)

type Resolver struct {
}

func (r *Resolver) User(args struct{ ID graphql.ID }) *UserResolver {
	return dm.Users[args.ID]
}

func (r *Resolver) Authors() []*AuthorResolver {
	authors := make([]*AuthorResolver, 0, len(dm.Books))

	for _, author := range dm.Authors {
		authors = append(authors, author)
	}

	return authors
}

func (r *Resolver) Author(args struct{ ID graphql.ID }) *AuthorResolver {
	return dm.Authors[args.ID]
}

func (r *Resolver) Books() []*BookResolver {
	books := make([]*BookResolver, 0, len(dm.Books))

	for _, book := range dm.Books {
		books = append(books, book)
	}

	return books
}

func (r *Resolver) Book(args struct{ ID graphql.ID }) *BookResolver {
	return dm.Books[args.ID]
}

func (r *Resolver) CreateComment(args struct {
	Bookid graphql.ID
	Text   string
}) *CommentResolver {

	id := graphql.ID(fmt.Sprintf("%d", len(dm.Comments)))
	newComment := &Comment{
		ID:     id,
		BookID: args.Bookid,
		Text:   args.Text,
		Date:   time.Now().Format("2019-12-01 15:04:05"),
	}
	newCommentResolve := &CommentResolver{newComment}
	dm.Comments[id] = newCommentResolve

	log.Printf("Добавлен коментарий %v", newComment)
	return newCommentResolve
}

func (r *Resolver) CreateUser(args struct {
	Login     string
	Password  string
	FirstName string
	LastName  string
}) (*UserResolver, error) {
	for _, user := range dm.Users {
		if user.Data.Login == args.Login {
			return nil, fmt.Errorf("User with login already exists")
		}

	}

	id := graphql.ID(fmt.Sprintf("%d", len(dm.Users)))
	newUser := &User{
		ID:        id,
		Login:     args.Login,
		Password:  args.Password,
		FirstName: args.FirstName,
		LastName:  args.LastName,
	}
	newUserResolve := &UserResolver{newUser}
	dm.Users[id] = newUserResolve

	log.Printf("Добавлен пользователь %v", newUser)
	return newUserResolve, nil
}

func (r *Resolver) Login(args struct {
	Login    string
	Password string
}) (*AuthTokenResolver, error) {
	for _, user := range dm.Users {
		if (user.Data.Login == args.Login) && (user.Data.Password == args.Password) {
			uuid, err := uuid.NewV4()
			if err != nil {
				log.Printf("Ошибка генерации токена: %v", err)
			}
			id := graphql.ID(fmt.Sprintf("%d", len(dm.AuthToken)))
			newToken := &AuthToken{
				ID:     id,
				Token:  uuid.String(),
				UserID: user.Data.ID,
			}
			newTokenResolve := &AuthTokenResolver{newToken}
			return newTokenResolve, nil
		}
	}
	return nil, fmt.Errorf("Invalid login or password credentials")
}
