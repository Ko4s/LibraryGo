package mongostorage

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Genre  string             `bson:"genre,omitempty"`
	//dodac hasha sparwdzajacego unikatowosc book
}
