package infrastructure

import (
	"log"
	"database/sql"
	"github.com/kelseyhightower/envconfig"
	"local.package/interfaces/database"
		_ "github.com/lib/pq"
)

// Config SQL接続設定情報
type Config struct {
	DbDialect string `required:"true" split_words:"true"`
	DbConnectionString string `required:"true" split_words:"true"`
}

// SQLHandler Postgresql専用のハンドラー
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler SQLHandler生成
func NewSQLHandler() *SQLHandler {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("FATAL : Failed to process env: %s", err.Error())
	}
	log.Printf("INFO  : Database Dialect:%s ConnectionString:%s", config.DbDialect, config.DbConnectionString)
	conn, err := sql.Open(config.DbDialect, config.DbConnectionString)
	if err != nil {
		log.Fatalf("FATAL : %s", err.Error())
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// Execute CQRSのCommandの実行
func (handler *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

// Query CQRSのQueryの実行
func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

// SQLResult SQLの実行結果
type SQLResult struct {
	Result sql.Result
}

// LastInsertId 最終インサートしたレコードのIDを取得する
func (r SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected 更新や削除で影響を受けた行数
func(r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SQLRow SQLの行
type SQLRow struct {
	Rows *sql.Rows
}

// Scan 行の列から引数に値を設定する
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next 次行が存在するか 存在するときはtrueを、しないときはfalseを返す。
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close クローズする
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
