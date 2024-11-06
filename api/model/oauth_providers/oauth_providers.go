package oauth_providers

import "github.com/google/uuid"

type ProviderType string

const (
	GitHub   ProviderType = "GITHUB"
	LinkedIn ProviderType = "LINKEDIN"
)

func (p ProviderType) IsValid() bool {
	switch p {
	case GitHub, LinkedIn:
		return true
	}
	return false
}

type OauthProviders struct {
	ID   uuid.UUID    `db:"id"`
	Name ProviderType `db:"name"`
}

func Find() {
	//
}
