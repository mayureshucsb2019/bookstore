USE bookstore;

-- Create the Customer table
CREATE TABLE IF NOT EXISTS Customer (
    email VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    dob DATE NOT NULL,
    unit_no VARCHAR(255),
    street_name VARCHAR(255),
    city VARCHAR(255),
    state VARCHAR(255),
    country VARCHAR(255),
    zipcode VARCHAR(20),
    landmark VARCHAR(255),
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    status ENUM('Active', 'Inactive') DEFAULT 'Active',
    notes TEXT
    languages JSON
);
