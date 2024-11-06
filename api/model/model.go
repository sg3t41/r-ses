package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/sg3t41/api/config"
)

var DB *sqlx.DB

// Model : DBレコードに共通するフィールド
type Model struct {
	ID uuid.UUID `json:"id"`
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

func Insert[T any](db sqlx.Ext, baseQuery string, record *T) (uuid.UUID, error) {
	query := baseQuery + ` RETURNING id`
	var id uuid.UUID
	switch conn := db.(type) {
	// トランザクション未使用
	case *sqlx.DB:
		if record != nil {
			err := conn.QueryRowx(query, record).Scan(&id)
			if err != nil {
				return uuid.Nil, err
			}
		} else {
			err := conn.QueryRowx(query).Scan(&id)
			if err != nil {
				return uuid.Nil, err
			}
		}
		return id, nil

	// トランザクション使用
	case *sqlx.Tx:
		if record != nil {
			err := conn.QueryRowx(query, record).Scan(&id)
			if err != nil {
				return uuid.Nil, err
			}
		} else {
			err := conn.QueryRowx(query).Scan(&id)
			if err != nil {
				return uuid.Nil, err
			}
		}
		return id, nil

	default:
		return uuid.Nil, fmt.Errorf("unsupported type: %T", db)
	}
}

func Select[R any](db sqlx.Ext, query string, record *R) (*R, error) {
	return nil, nil
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
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s", table, condition)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("GetRecords: %v", err)
	}
	defer rows.Close()

	var records []Model
	for rows.Next() {
		var record Model
		if err := rows.Scan(&record.ID); err != nil {
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

func IsExist(table string, condition string, args ...interface{}) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, condition)

	var exists int
	err := DB.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("IsExist query error: %v", err)
	}

	return true, nil
}
