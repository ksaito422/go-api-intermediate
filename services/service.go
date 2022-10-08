package services

import "database/sql"

// サービス構造体を定義
type MyAppService struct {
	// フィールドにsql.DB型を含める
	db *sql.DB
}

// コンストラクタの定義
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
