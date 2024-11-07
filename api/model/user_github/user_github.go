package user_github

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserGithub struct {
	UserProviderID        uuid.UUID `db:"user_provider_id"`
	UserOauthGithubDataID uuid.UUID `db:"user_oauth_github_data_id"`
	GithubID              int64     `db:"github_id"`
}

func ExistsByGithubId(db sqlx.Ext, githubId int64) (bool, error) {
	q := "SELECT EXISTS(SELECT 1 FROM user_github WHERE github_id=$1)"
	var exists bool

	switch conn := db.(type) {
	case *sqlx.DB:
		if err := conn.Get(&exists, q, githubId); err != nil {
			if err == sql.ErrNoRows {
				return false, nil
			}
			return false, err
		}

	case *sqlx.Tx:
		if err := conn.Get(&exists, q, githubId); err != nil {
			if err == sql.ErrNoRows {
				return false, nil
			}
			return false, err
		}

	default:
		return false, fmt.Errorf("ExistByGithubId unsupported type: %T", db)
	}

	return exists, nil
}
