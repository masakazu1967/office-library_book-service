package usecase

import "github.com/masakazu1967/office-library_book-service/domain"

// WriterRepository 著者リポジトリインターフェイス
type WriterRepository interface {
	Store(domain.Writer) (int, error)
	FindByID(int) (domain.Writer, error)
	FindAll() (domain.Writers, error)
}
