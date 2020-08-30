package database

/*
 * WriterRepoistoryはSqlHandlerを埋め込んでいるが、SqlHandlerはinfrastructureにある。
 このままでは内側から外側を呼んでいるので、Clean Architectureの外側から内側を呼ぶ原則に違反する。
 依存関係逆転の原則に従って、interfacesにこのインターフェイスを定義する。
 */
 
 // SQLHandler SQLハンドラー
 type SQLHandler interface {
	 Execute(string, ...interface{}) (Result, error)
	 Query(string, ...interface{}) (Row, error)
 }

 // Result SQL実行結果
 type Result interface {
	 LastInsertId() (int64, error)
	 RowsAffected() (int64, error)
 }

 // Row 行
 type Row interface {
	 Scan(...interface{}) error
	 Next() bool
	 Close() error
 }
 