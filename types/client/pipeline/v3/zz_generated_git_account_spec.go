package client

const (
	GitAccountSpecType               = "gitAccountSpec"
	GitAccountSpecFieldAccessToken   = "accessToken"
	GitAccountSpecFieldAccountType   = "accountType"
	GitAccountSpecFieldAvatarURL     = "avatar_url"
	GitAccountSpecFieldHTMLURL       = "html_url"
	GitAccountSpecFieldLogin         = "login"
	GitAccountSpecFieldPrivate       = "private"
	GitAccountSpecFieldRancherUserID = "rancherUserId"
)

type GitAccountSpec struct {
	AccessToken   string `json:"accessToken,omitempty"`
	AccountType   string `json:"accountType,omitempty"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	HTMLURL       string `json:"html_url,omitempty"`
	Login         string `json:"login,omitempty"`
	Private       *bool  `json:"private,omitempty"`
	RancherUserID string `json:"rancherUserId,omitempty"`
}
