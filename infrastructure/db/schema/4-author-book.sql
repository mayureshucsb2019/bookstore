USE bookstore;

-- Create the AuthorBook table
CREATE TABLE IF NOT EXISTS AuthorBook (
    author_id VARCHAR(255),
    book_isbn VARCHAR(255),
    PRIMARY KEY (author_id, book_isbn),
    FOREIGN KEY (author_id) REFERENCES Authors(id) ON DELETE CASCADE,
    FOREIGN KEY (book_isbn) REFERENCES Books(isbn) ON DELETE CASCADE
);