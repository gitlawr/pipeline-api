package client

const (
	BuildImageStepConfigType                = "buildImageStepConfig"
	BuildImageStepConfigFieldBuildPath      = "buildPath"
	BuildImageStepConfigFieldDockerfilePath = "dockerFilePath"
	BuildImageStepConfigFieldPush           = "push"
	BuildImageStepConfigFieldTargetImage    = "targetImage"
)

type BuildImageStepConfig struct {
	BuildPath      string `json:"buildPath,omitempty"`
	DockerfilePath string `json:"dockerFilePath,omitempty"`
	Push           *bool  `json:"push,omitempty"`
	TargetImage    string `json:"targetImage,omitempty"`
}
