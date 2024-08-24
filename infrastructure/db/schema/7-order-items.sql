USE bookstore;

-- Create the OrderItems table
CREATE TABLE IF NOT EXISTS OrderItems (
    order_id INT,
    isbn VARCHAR(255),
    quantity INT NOT NULL,
    PRIMARY KEY (order_id, isbn),
    FOREIGN KEY (order_id) REFERENCES Orders(id),
    FOREIGN KEY (isbn) REFERENCES Books(isbn)
);