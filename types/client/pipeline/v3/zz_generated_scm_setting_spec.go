package client

const (
	SCMSettingSpecType              = "scmSettingSpec"
	SCMSettingSpecFieldClientID     = "clientID"
	SCMSettingSpecFieldClientSecret = "clientSecret"
	SCMSettingSpecFieldHomePage     = "homepage"
	SCMSettingSpecFieldHostName     = "hostName"
	SCMSettingSpecFieldIsAuth       = "isAuth"
	SCMSettingSpecFieldRedirectURL  = "redirectURL"
	SCMSettingSpecFieldScheme       = "scheme"
	SCMSettingSpecFieldScmType      = "scmType"
)

type SCMSettingSpec struct {
	ClientID     string `json:"clientID,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	HomePage     string `json:"homepage,omitempty"`
	HostName     string `json:"hostName,omitempty"`
	IsAuth       *bool  `json:"isAuth,omitempty"`
	RedirectURL  string `json:"redirectURL,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	ScmType      string `json:"scmType,omitempty"`
}
