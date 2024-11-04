package user

import (
	"fmt"

	"github.com/sg3t41/api/model"
)

type User struct {
	model.Model
}

func IsExistByGithubId(githubId int) (bool, error) {
	isExist, err := model.IsExist("user_github", "github_id = $1", githubId)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func UpdateGithubUser() {
	q := `
		UPDATE users 
		SET
			username=$1,
			email=$2, 
			avatar_url=$3, 
			profile_url=$4, 
			full_name=$5 
		WHERE 
			github_id=$6
		`
	updatedUserID, err := model.UpdateRecord(q, user.Login, user.Email, user.AvatarURL, user.URL, user.Name, user.ID)

}

// CreateUser : 新しいUserレコードを作成する関数
func Create(username, email, passwordHash string) (string, error) {
	query := "INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	id, err := model.CreateRecord(query, username, email, passwordHash)
	if err != nil {
		return "", fmt.Errorf("CreateUser: %v", err)
	}
	return id, nil
}

// UpdateUser : 既存のUserレコードを更新する関数
func Update(id int, username, email, passwordHash string) (string, error) {
	query := "UPDATE users SET username = $1, email = $2, password_hash = $3, updated_at = $4 WHERE id = $5"
	rows, err := model.UpdateRecord(query, username, email, passwordHash, id)
	if err != nil {
		return "", fmt.Errorf("UpdateUser: %v", err)
	}
	return rows, nil
}

// SoftDeleteUser : Userレコードを論理削除する関数
func SoftDelete(id int) (int64, error) {
	query := "UPDATE users SET deleted_at = $1 WHERE id = $2"
	rows, err := model.SoftDeleteRecord(query, id)
	if err != nil {
		return 0, fmt.Errorf("SoftDeleteUser: %v", err)
	}
	return rows, nil
}