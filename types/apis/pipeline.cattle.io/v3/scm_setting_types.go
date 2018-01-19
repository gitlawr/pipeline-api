package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SCMSetting struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec SCMSettingSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status SCMSettingStatus `json:"status"`
}

type SCMSettingSpec struct {
	IsAuth       bool   `json:"isAuth" yaml:"isAuth"`
	ScmType      string `json:"scmType,omitempty" yaml:"scmType,omitempty"`
	HostName     string `json:"hostName,omitempty" yaml:"hostName,omitempty"`
	Scheme       string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	HomePage     string `json:"homepage,omitempty" yaml:"homepage,omitempty"`
	ClientID     string `json:"clientID,omitempty" yaml:"clientID,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	RedirectURL  string `json:"redirectURL,omitempty" yaml:"redirectURL,omitempty"`
}

type SCMSettingStatus struct {
}
