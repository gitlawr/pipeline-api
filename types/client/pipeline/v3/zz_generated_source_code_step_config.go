package client

const (
	SourceCodeStepConfigType            = "sourceCodeStepConfig"
	SourceCodeStepConfigFieldBranch     = "branch"
	SourceCodeStepConfigFieldGitUser    = "gitUser"
	SourceCodeStepConfigFieldRepository = "repository"
)

type SourceCodeStepConfig struct {
	Branch     string `json:"branch,omitempty"`
	GitUser    string `json:"gitUser,omitempty"`
	Repository string `json:"repository,omitempty"`
}
