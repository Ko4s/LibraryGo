package mongostorage

import (
	"library/pkg/adding"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book defines how we store book in mongoDB
type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Genre  string             `bson:"genre,omitempty"`
	//dodac hasha sparwdzajacego unikatowosc book
}

func (b *Book) toAddingBook() adding.Book {
	return adding.Book{
		ID:     b.ID.Hex(),
		Title:  b.Title,
		Author: b.Author,
		Genre:  b.Genre,
	}
}
