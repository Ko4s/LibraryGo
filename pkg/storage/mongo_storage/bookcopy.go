package mongostorage

import "go.mongodb.org/mongo-driver/bson/primitive"

// BookCopy is a model of book copy :)
type BookCopy struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	BookID          primitive.ObjectID `bson:"bookID,omitempty"`
	Comment         string             `bson:"comment,omitempty"`
	PublicationDate string             `bson:"publicationDate,omitempty"`
}
