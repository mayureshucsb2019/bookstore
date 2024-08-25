package db

import (
	"sync"

	"github.com/mayureshucsb2019/bookstore/service/common"
)

var bookRepoInstance *BookRepository
var bookRepoOnce sync.Once

func NewBookRepository(db *common.DBConnection) *BookRepository {
	bookRepoOnce.Do(func() {
		bookRepoInstance = &BookRepository{
			DB: db.DB,
		}
	})
	return bookRepoInstance
}
