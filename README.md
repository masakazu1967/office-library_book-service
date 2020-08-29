# 室内文庫の蔵書管理サービス

情報戦略室で購入した書籍の管理を行います。

## Features

## Requirement

- PostgreSQL

## Installation

```bash
export DB_DIALECT="postgres"
export DB_CONNECTION_STRING="host=127.0.0.1 port=5432 user=gitpod dbname=postgres sslmode=disable"
```

## Usage

```bash
book-service-linux-amd64
```

著者

- GET /writers
- GET /writers/:id
- GET /writers/:id/books
- POST /writers

出版社

- GET /publishers
- GET /publishers/:id
- GET /publishers/:id/books
- POST /publishers

書籍

- GET /books
- GET /books/:id
- GET /books/:id/writers
- POST /books

## Node

## Author

Masakazu Kobayashi

## License

"Book Service" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
