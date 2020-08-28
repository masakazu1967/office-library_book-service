module github.com/masakazu1967/office-library_book-service

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/lib/pq v1.8.0
	local.package/domain v0.0.0-00010101000000-000000000000 // indirect
	local.package/infrastructure v0.0.0-00010101000000-000000000000
	local.package/interfaces/controllers v0.0.0-00010101000000-000000000000 // indirect
	local.package/interfaces/database v0.0.0-00010101000000-000000000000 // indirect
	local.package/usecase v0.0.0-00010101000000-000000000000 // indirect
)

replace local.package/infrastructure => ./infrastructure

replace local.package/domain => ./domain

replace local.package/interfaces/database => ./interfaces/database

replace local.package/interfaces/controllers => ./interfaces/controllers

replace local.package/usecase => ./usecase
