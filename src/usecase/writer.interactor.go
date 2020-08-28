package usecase

import (
	"log"
	"local.package/domain"
)

// WriterInteractor 著者ユースケース
type WriterInteractor struct {
	WriterRepository WriterRepository
}

// Add 著者を登録する
func (interactor *WriterInteractor) Add(w domain.Writer) (writer domain.Writer, err error) {
	identifier, err := interactor.WriterRepository.Store(w)
	if err != nil {
		return
	}
	writer, err = interactor.WriterRepository.FindByID(identifier)
	return
}

// Writers 著者一覧を取得する
func (interactor *WriterInteractor) Writers() (writers domain.Writers, err error) {
	writers, err = interactor.WriterRepository.FindAll()
	return
}

// WriterByID IDの著者の情報を取得する
func (interactor *WriterInteractor) WriterByID(identifier int) (writer domain.Writer, err error) {
	log.Printf("INFO  : WriterInteractor.WriterByID(%d)", identifier)
	writer, err = interactor.WriterRepository.FindByID(identifier)
	return
}
