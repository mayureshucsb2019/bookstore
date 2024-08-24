USE bookstore;

-- Create the Book table
CREATE TABLE IF NOT EXISTS Books (
    isbn VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tags JSON,
    author_names JSON,
    date_of_publish DATE NOT NULL,
    publishing_house JSON,
    number_of_pages INT,
    cost FLOAT
);
