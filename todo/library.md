## Book definition 

Book {
    _id -> mongodbID
    name string
    author string || author był idikiem -> dokumentu z authorami ewentualnie {}
    genre string || --||-- 
    BookCopies []BookCopyIDs
}

BookCopy {
    _id
    Book BookID
    isAvailable
    comment
    publition date 
}

 -> Harry Pooter 2000, 2020, 


1. Mamy booka w systemie
    1. dodac Copy
2. Nie mamy booka w systemie 
    1. Dodac Book
    1. Dodać Copy


Insert 
    name string
    author string || author był idikiem -> dokumentu z authorami ewentualnie {}
    genre string || --||-- 
    publition date 

    1. Look po Booka
        * Ok
            Stworzymy obiekt Bookopy i dodamy 
        * Nope
            Create Book
            Insert BookCpy