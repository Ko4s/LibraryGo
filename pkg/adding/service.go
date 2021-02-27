package adding

//Maybe later add adding errrors

//Service provides book adding opertions
type Service interface {
	AddBook(...Book) error //Think about returning ID or whole Book structure from db
}

//Repository provides access to books repository
type Repository interface {
	AddBook(book *Book) error
	//maybe add method GetBooks()
	BookExist(book *Book) (*Book, error)
	AddBookCopy(bc *BookCopy) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddBook(books ...Book) error {

	//sprawdzic czy ksiązka istnieje jak nie to dodać
	for _, book := range books {

		//refactos function bookExist to getBookID
		bookExist, _ := s.r.BookExist(&book)

		if bookExist == nil {
			//Later add adding of Book Copy
			continue
		}

		err := s.r.AddBook(&book)

		if err != nil {
			return err
		}
	}

	return nil
}
