package factory

import (
	"sync"

	author_db "github.com/mayureshucsb2019/bookstore/service/author/db"
	book_db "github.com/mayureshucsb2019/bookstore/service/book/db"
	"github.com/mayureshucsb2019/bookstore/service/common"
	customer_db "github.com/mayureshucsb2019/bookstore/service/customer/db"
)

type RepositoryFactory struct {
	dbConn *common.DBConnection
}

var repositoryFactoryInstance *RepositoryFactory
var repositoryFactoryOnce sync.Once

func GetRepositoryFactory(dbConn *common.DBConnection) *RepositoryFactory {
	repositoryFactoryOnce.Do(func() {
		repositoryFactoryInstance = &RepositoryFactory{
			dbConn: dbConn,
		}
	})
	return repositoryFactoryInstance
}

func (f *RepositoryFactory) CreateBookRepository() *book_db.BookRepository {
	return book_db.NewBookRepository(f.dbConn)
}

func (f *RepositoryFactory) CreateAuthorRepository() *author_db.AuthorRepository {
	return author_db.NewAuthorRepository(f.dbConn)
}

func (f *RepositoryFactory) CreateCustomerRepository() *customer_db.CustomerRepository {
	return customer_db.NewCustomerRepository(f.dbConn)
}
