package client

const (
	GitRepositoryType             = "gitRepository"
	GitRepositoryFieldCloneURL    = "clone_url"
	GitRepositoryFieldPermissions = "permissions"
)

type GitRepository struct {
	CloneURL    string          `json:"clone_url,omitempty"`
	Permissions map[string]bool `json:"permissions,omitempty"`
}
