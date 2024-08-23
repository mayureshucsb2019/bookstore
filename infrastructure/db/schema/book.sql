USE bookstore;

-- Create the Book table
CREATE TABLE IF NOT EXISTS Book (
    isbn VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tags JSON,
    author_name VARCHAR(255) NOT NULL,
    date_of_publish DATE NOT NULL,
    publishing_house VARCHAR(255),
    number_of_pages INT,
    cost FLOAT
);
