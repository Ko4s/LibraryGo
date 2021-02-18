package mongostorage

import (
	"context"
	"errors"
	"library/pkg/adding"
	"time"

	configreader "library/pkg/config_reader"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const bookCollectionName = "Books"
const defaulProcessTime = time.Second * 4

var (
	ErrNotFound = errors.New("storage: Element not found in db")
)

func NewStorage(cr *configreader.ConfigReader) (*Storage, error) {
	//TODO
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	connectionString := cr.GetString("DB_CONNECTION_STRING")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return nil, err
	}

	dbName := cr.GetString("DB_NAME")
	db := client.Database(dbName)
	bookCollection := db.Collection(bookCollectionName)

	return &Storage{
		db:             db,
		bookCollection: bookCollection,
	}, nil
}

type Storage struct {
	//TODO
	//reference to db or db object
	//db <= typem z drivera mongo db
	db             *mongo.Database
	bookCollection *mongo.Collection
}

//Books

// AddBook will add book to teh datavase
// @param book - type maybe will change later
func (s *Storage) AddBook(book adding.Book) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), defaulProcessTime)
	defer cancel()

	mongoBook := Book{
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
	}

	result, err := s.bookCollection.InsertOne(ctx, mongoBook)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *Storage) FindBookByID(ID string) (*adding.Book, error) {
	return nil, nil
}

func (s *Storage) FindBookByText(query string) (*[]adding.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaulProcessTime)
	defer cancel()

	mongoQuery := bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}

	coursor, err := s.bookCollection.Find(ctx, mongoQuery)

	if err != nil {
		return nil, err
	}

	bookList := []adding.Book{}

	for coursor.Next(context.Background()) {
		book := Book{}
		err := coursor.Decode(&book)

		if err != nil {
			return nil, err
		}

		bookList = append(bookList, book.toAddingBook())
	}

	return &bookList, nil
}

func (s *Storage) UpdateBook(book adding.Book) (*adding.Book, error) {
	return nil, nil
}

func (s *Storage) DeleteBook(book adding.Book) (*adding.Book, error) {
	return nil, nil
}

// Books Copy

// code below noe relevant right now, we dont have book copy yet
func (s *Storage) BookExist(book adding.Book) bool {
	return false
}

func (s *Storage) AddBookCopy(ID string, bc adding.BookCopy) error {
	return nil
}
