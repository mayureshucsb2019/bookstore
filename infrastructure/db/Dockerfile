# Use the official MySQL image from Docker Hub
FROM mysql:8.0

# Set environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
ENV MYSQL_DATABASE=bookstore
ENV MYSQL_USER=bstore
ENV MYSQL_PASSWORD=${MYSQL_ADMIN_PASSWORD}

# Copy the initialization scripts into the container
COPY schema/create-schemas.sql /docker-entrypoint-initdb.d/
COPY schema/book.sql /docker-entrypoint-initdb.d/

# Expose MySQL port
EXPOSE 3306
