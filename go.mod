module github.com/masakazu1967/office-library_book-service

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/lib/pq v1.8.0
	github.com/masakazu1967/office-library_book-service/domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/masakazu1967/office-library_book-service/infrastructure v0.0.0-00010101000000-000000000000
	github.com/masakazu1967/office-library_book-service/interfaces/controllers v0.0.0-00010101000000-000000000000 // indirect
	github.com/masakazu1967/office-library_book-service/interfaces/database v0.0.0-00010101000000-000000000000 // indirect
	github.com/masakazu1967/office-library_book-service/usecase v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/masakazu1967/office-library_book-service/infrastructure => ./infrastructure

replace github.com/masakazu1967/office-library_book-service/interfaces/controllers => ./interfaces/controllers

replace github.com/masakazu1967/office-library_book-service/interfaces/database => ./interfaces/database

replace github.com/masakazu1967/office-library_book-service/domain => ./domain

replace github.com/masakazu1967/office-library_book-service/usecase => ./usecase
