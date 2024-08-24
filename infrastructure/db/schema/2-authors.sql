USE bookstore;

-- Create the Authors table
CREATE TABLE IF NOT EXISTS Authors (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    last_name VARCHAR(255) NOT NULL,
    dob VARCHAR(16) NOT NULL,
    unit_no VARCHAR(255),
    street_name VARCHAR(255),
    city VARCHAR(255),
    state VARCHAR(255),
    country VARCHAR(255),
    zipcode VARCHAR(20),
    landmark VARCHAR(255),
    languages JSON
);
