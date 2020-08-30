package database

import (
	"log"
	"github.com/masakazu1967/office-library_book-service/domain"
)

// WriterRepository 著者リポジトリ
type WriterRepository struct {
	SQLHandler
}

// Store 著者をリポジトリに登録する
func (repo *WriterRepository) Store(writer domain.Writer) (id int, err error) {
	result, err := repo.Execute("INSERT INTO writers (family_name, given_name) VALUES($1, $2)", writer.FamilyName, writer.GivenName)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

// FindByID IDから著者情報を取得する
func (repo *WriterRepository) FindByID(identifier int) (writer domain.Writer, err error) {
	log.Printf("INFO  : WriterRepository.FIndByID(%d)", identifier)
	row, err := repo.Query("SELECT id, family_name, given_name FROM writers WHERE id = $1", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var familyName string
	var givenName string
	row.Next()
	err = row.Scan(&id, &familyName, &givenName)
	if err != nil {
		return
	}
	writer.ID = id
	writer.FamilyName = familyName
	writer.GivenName = givenName
	return
}

// FindAll すべての著者を取得する
func (repo *WriterRepository) FindAll() (writers domain.Writers, err error) {
	rows, err := repo.Query("SELECT id, family_name, given_name FROM writers")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var familyName string
		var givenName string
		err := rows.Scan(&id, &familyName, &givenName)
		if err != nil {
			continue
		}
		writer := domain.Writer{
			ID:         id,
			FamilyName: familyName,
			GivenName:  givenName,
		}
		writers = append(writers, writer)
	}
	return
}
