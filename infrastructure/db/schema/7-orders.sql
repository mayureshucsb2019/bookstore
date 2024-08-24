USE bookstore;

-- Create the Orders table
CREATE TABLE IF NOT EXISTS Orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_email VARCHAR(255),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount FLOAT,
    FOREIGN KEY (customer_email) REFERENCES Customer(email)
);