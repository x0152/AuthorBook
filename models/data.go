package models

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/graph-gophers/graphql-go"
)

type DataManager struct {
	Books     map[graphql.ID]*BookResolver
	Authors   map[graphql.ID]*AuthorResolver
	Comments  map[graphql.ID]*CommentResolver
	Users     map[graphql.ID]*UserResolver
	AuthToken map[graphql.ID]*AuthTokenResolver
}

var dm DataManager

func init() {
	dm = DataManager{
		Books:     make(map[graphql.ID]*BookResolver),
		Authors:   make(map[graphql.ID]*AuthorResolver),
		Comments:  make(map[graphql.ID]*CommentResolver),
		Users:     make(map[graphql.ID]*UserResolver),
		AuthToken: make(map[graphql.ID]*AuthTokenResolver),
	}

	dm.Books = LoadBooks()
	dm.Authors = LoadAuthors()

	// TODO load user and comments

}

func LoadBooks() map[graphql.ID]*BookResolver {
	resolver_books := make(map[graphql.ID]*BookResolver)
	file, err := ioutil.ReadFile("books.json")

	if err != nil {
		log.Fatalf("Неудалось загрузить книги из файла books.json (файл отсутствует или недостаточно прав на чтение): %v", err)
	}

	var books []Book
	err = json.Unmarshal([]byte(file), &books)

	if err != nil {
		log.Fatalf("Неудалось декодировать json файл: %v", err)
	}

	for _, book := range books {
		b := book
		resolver_books[book.ID] = &BookResolver{&b}
	}

	log.Printf("Книги успешно загружены: %v", books)

	return resolver_books
}

func LoadAuthors() map[graphql.ID]*AuthorResolver {
	resolver_authors := make(map[graphql.ID]*AuthorResolver)
	file, err := ioutil.ReadFile("authors.json")

	if err != nil {
		log.Fatalf("Неудалось загрузить авторов из файла authors.json (файл отсутствует или недостаточно прав на чтение): %v", err)
	}

	var authors []Author
	err = json.Unmarshal([]byte(file), &authors)

	if err != nil {
		log.Fatalf("Неудалось декодировать json файл: %v", err)
	}

	for _, author := range authors {
		a := author
		resolver_authors[author.ID] = &AuthorResolver{&a}
	}

	log.Printf("Авторы успешно загружены: %v", authors)

	return resolver_authors
}
