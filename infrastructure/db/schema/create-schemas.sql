-- Create the database
CREATE DATABASE IF NOT EXISTS bookstore
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_0900_ai_ci;

-- Create the admin user if it doesn't already exist
CREATE USER IF NOT EXISTS 'bstore'@'%' IDENTIFIED BY '{{MYSQL_ADMIN_PASSWORD}}';

-- Grant all privileges on the bookstore database to the admin user
GRANT ALL PRIVILEGES ON bookstore.* TO 'bstore'@'%';

-- Another user for migration and testing
CREATE USER IF NOT EXISTS 'bstore_mig'@'%' IDENTIFIED BY 'bstore_mig';

-- Grant privileges to the migration user (optional, adjust as needed)
GRANT ALL PRIVILEGES ON bookstore.* TO 'bstore_mig'@'%';

-- Flush privileges to ensure that the changes take effect
FLUSH PRIVILEGES;
