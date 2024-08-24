# Use the official MySQL image from Docker Hub
FROM mysql:8.0

# Set environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
ENV MYSQL_DATABASE=bookstore
ENV MYSQL_USER=bstore
ENV MYSQL_PASSWORD=${MYSQL_ADMIN_PASSWORD}

# Copy the initialization scripts into the container
COPY schema/1-create-schemas.sql /docker-entrypoint-initdb.d/
COPY schema/2-authors.sql /docker-entrypoint-initdb.d/
COPY schema/3-books.sql /docker-entrypoint-initdb.d/
COPY schema/4-customers.sql /docker-entrypoint-initdb.d/
COPY schema/5-orders.sql /docker-entrypoint-initdb.d/
COPY schema/6-order-items.sql /docker-entrypoint-initdb.d/

# Expose MySQL port
EXPOSE 3306
