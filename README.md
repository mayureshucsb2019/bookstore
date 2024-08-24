# bookstore

* write bookstore.yaml 
* brew install openapi-generator  # On MacOS

# How to autogenerate server files from bookstore_api.yaml
* RUN $ openapi-generator generate -i bookstore_api.yaml -g go-server -o ./service
* change go.mod to have module github.com/mayureshucsb2019/bookstore instead of module github.com/GIT_USER_ID/GIT_REPO_ID
* change main.go to import openapi "github.com/mayureshucsb2019/bookstore/go" instead of import github.com/GIT_USER_ID/GIT_REPO_ID
* go get github.com/gorilla/mux
* go get github.com/go-sql-driver/mysql

# Install MySQL 
* brew install mysql
* docker build -t my-mysql-image .
* docker run -d --name mysql-container -p 3306:3306 my-mysql-image
* docker exec -it mysql-container mysql -u bstore -p

# User Docker Compose to automatically build mysql docker image
* cd to bookstore/service
* $docker-compose -f infrastructure/db/docker-compose.yaml up --build
