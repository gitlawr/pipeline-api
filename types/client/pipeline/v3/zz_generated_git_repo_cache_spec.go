package client

const (
	GitRepoCacheSpecType              = "gitRepoCacheSpec"
	GitRepoCacheSpecFieldGitAccountID = "gitAccountID"
	GitRepoCacheSpecFieldRepositories = "repositories"
	GitRepoCacheSpecFieldScmType      = "scmType"
)

type GitRepoCacheSpec struct {
	GitAccountID string          `json:"gitAccountID,omitempty"`
	Repositories []GitRepository `json:"repositories,omitempty"`
	ScmType      string          `json:"scmType,omitempty"`
}
