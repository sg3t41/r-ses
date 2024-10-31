package user

//
//import (
//	"database/sql"
//	"fmt"
//
//	"github.com/sg3t41/api/model"
//)
//
//type User struct {
//	model.Model
//	Username     string `json:"username"`
//	Email        string `json:"email"`
//	PasswordHash string `json:"password_hash"`
//}
//
//// CreateUser : 新しいUserレコードを作成する関数
//func Create(username, email, passwordHash string) (int64, error) {
//	query := "INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
//	id, err := model.CreateRecord(query, username, email, passwordHash)
//	if err != nil {
//		return 0, fmt.Errorf("CreateUser: %v", err)
//	}
//	return id, nil
//}
//
//// UpdateUser : 既存のUserレコードを更新する関数
//func Update(id int, username, email, passwordHash string) (int64, error) {
//	query := "UPDATE users SET username = $1, email = $2, password_hash = $3, updated_at = $4 WHERE id = $5"
//	rows, err := model.UpdateRecord(query, username, email, passwordHash, id)
//	if err != nil {
//		return 0, fmt.Errorf("UpdateUser: %v", err)
//	}
//	return rows, nil
//}
//
//// SoftDeleteUser : Userレコードを論理削除する関数
//func SoftDelete(id int) (int64, error) {
//	query := "UPDATE users SET deleted_at = $1 WHERE id = $2"
//	rows, err := model.SoftDeleteRecord(query, id)
//	if err != nil {
//		return 0, fmt.Errorf("SoftDeleteUser: %v", err)
//	}
//	return rows, nil
//}
//
//// GetUserByEmailAndPassword : メールアドレスとパスワードハッシュに基づいてユーザーを取得する関数
//func GetUserByEmailAndPassword(email, passwordHash string) (*User, error) {
//	query := "SELECT id, created_at, updated_at, deleted_at, username, email, password_hash FROM users WHERE email = $1 AND password_hash = $2"
//	row := model.DB.QueryRow(query, email, passwordHash)
//
//	var user User
//	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Username, &user.Email, &user.PasswordHash)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, fmt.Errorf("GetUserByEmailAndPassword: user not found")
//		}
//		return nil, fmt.Errorf("GetUserByEmailAndPassword: %v", err)
//	}
//
//	return &user, nil
//}
