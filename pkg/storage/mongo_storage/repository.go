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

const (
	bookCollectionName      = "Books"
	bookCoptyCollectionName = "BookCopies"
	defaulProcessTime       = time.Second * 4
)

// Custom storage errors
var (
	ErrNotFound = errors.New("storage: Element not found in db")
)

// NewStorage is a storage constructor
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
	bookCoppiesCollection := db.Collection(bookCoptyCollectionName)

	return &Storage{
		db:                    db,
		bookCollection:        bookCollection,
		bookCoppiesCollection: bookCoppiesCollection,
	}, nil
}

// Storage provide methods to work with books and books copy collections
type Storage struct {
	//TODO
	//reference to db or db object
	//db <= typem z drivera mongo db
	db                    *mongo.Database
	bookCollection        *mongo.Collection
	bookCoppiesCollection *mongo.Collection
}

//Books

// AddBook will add book to teh datavase
func (s *Storage) AddBook(book *adding.Book) (string, error) {

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

// FindBookByID returns book based on given ID
// If books doesnt exist ErrNotFound is returned
func (s *Storage) FindBookByID(ID string) (*adding.Book, error) {
	return nil, nil
}

// FindBookByText return list of books based on given query text
// Its sreaching for text in Book title, author or genre
// If something goes wrong method returns error, else return pointer to slice of books
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

func (s *Storage) UpdateBook(book *adding.Book) (*adding.Book, error) {
	return nil, nil
}

func (s *Storage) DeleteBook(book *adding.Book) (*adding.Book, error) {
	return nil, nil
}

// Books Copy

// BookExist checks if book with given title and author exist in db
func (s *Storage) BookExist(book *adding.Book) (*adding.Book, error) {

	bookFromDB := Book{}

	filter := bson.M{
		"title":  book.Title,
		"author": book.Author,
	}

	result := s.bookCollection.FindOne(context.Background(), filter)

	if err := result.Decode(&bookFromDB); err != nil {
		return nil, err
	}

	bookToReturn := bookFromDB.toAddingBook()
	return &bookToReturn, nil
}

// AddBookCopy book adds 
func (s *Storage) AddBookCopy(bc adding.BookCopy) error {
	bookID, err := primitive.ObjectIDFromHex(bc.BookID)

	if err != nil {
		return err
	}

	newBookCopy := BookCopy{
		BookID:          bookID,
		Comment:         bc.Comment,
		PublicationDate: bc.PublicationDate,
	}

	_, err = s.bookCoppiesCollection.InsertOne(context.TODO(), newBookCopy)

	if err != nil {
		return err
	}
	
	return nil
}

// DestructiveReset drops all collections
func (s *Storage) DestructiveReset() {
	s.bookCollection.Drop(context.Background())
}

// Ping pings client and check if we hace connection with db
func (s *Storage) Ping() {
	err := s.db.Client().Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
}
