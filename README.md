# bookstore

* write bookstore.yaml 
* brew install openapi-generator  # On MacOS
* openapi-generator generate -i books_api.yaml -g go-server -o ./bookstore-server
* change go.mod to have module github.com/mayureshucsb2019/bookstore instead of module github.com/GIT_USER_ID/GIT_REPO_ID
* change main.go to import openapi "github.com/mayureshucsb2019/bookstore/go" instead of import github.com/GIT_USER_ID/GIT_REPO_ID
