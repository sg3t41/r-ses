package model

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/sg3t41/api/config"
)

var DB *sqlx.DB

// Model : DBレコードに共通するフィールド
type Model struct {
	ID        string       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

// Setup : データベースのセットアップ関数
func Setup() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Port,
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Name)

	DB, err = sqlx.Open(config.DatabaseSetting.Type, dsn)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	//DBへの疎通確認
	err = DB.Ping()
	if err != nil {
		log.Fatalf("db.Ping err: %v", err)
	}
}

// CloseDB : データベース接続を閉じる
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// CreateRecord : 新しいレコードを作成する関数
func CreateRecord(query string, args ...interface{}) (string, error) {
	var lastInsertId string
	err := DB.QueryRow(query+" RETURNING id", args...).Scan(&lastInsertId)
	if err != nil {
		return "", fmt.Errorf("InsertRecord error: %v", err)
	}
	return lastInsertId, nil
}

// UpdateRecord : レコードを更新する関数
func UpdateRecord(query string, args ...interface{}) (string, error) {
	var updatedId string
	err := DB.QueryRow(query+" RETURNING id", args...).Scan(&updatedId)
	if err != nil {
		return "", fmt.Errorf("UpdateRecord error: %v", err)
	}
	return updatedId, nil
}

// SoftDeleteRecord : レコードをソフトデリートする関数
func SoftDeleteRecord(query string, args ...interface{}) (int64, error) {
	nowTime := time.Now()
	args = append(args, nowTime)
	result, err := DB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func GetRecords(table string, condition string, args ...interface{}) ([]Model, error) {
	query := fmt.Sprintf("SELECT id, created_at, updated_at, deleted_at FROM %s WHERE %s", table, condition)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("GetRecords: %v", err)
	}
	defer rows.Close()

	var records []Model
	for rows.Next() {
		var record Model
		if err := rows.Scan(&record.ID, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt); err != nil {
			return nil, fmt.Errorf("GetRecords Scan: %v", err)
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetRecords Rows Err: %v", err)
	}

	return records, nil
}

func GetRecords2[T any](table string, fields []string, condition string, args ...interface{}) ([]T, error) {
	fieldsStr := strings.Join(fields, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", fieldsStr, table, condition)

	var records []T
	err := DB.Select(&records, query, args...)
	if err != nil {
		return nil, fmt.Errorf("GetRecords: %v", err)
	}

	return records, nil
}
