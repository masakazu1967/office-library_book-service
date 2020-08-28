package usecase

import "local.package/domain"

// WriterRepository 著者リポジトリインターフェイス
type WriterRepository interface {
	Store(domain.Writer) (int, error)
	FindByID(int) (domain.Writer, error)
	FindAll() (domain.Writers, error)
}
