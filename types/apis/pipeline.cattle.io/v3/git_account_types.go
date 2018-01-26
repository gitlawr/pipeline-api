package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GitAccount struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	//###Spec GitAccountSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	//###Status GitAccountStatus `json:"status"`
	GitAccountSpec
	GitAccountStatus
}

type GitAccountSpec struct {
	Id string `json:"id,omitempty"`

	//private or shared across environment
	Private       bool   `json:"private,omitempty"`
	AccountType   string `json:"accountType,omitempty"`
	RancherUserID string `json:"rancherUserId,omitempty"`
	Login         string `json:"login,omitempty"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	HTMLURL       string `json:"html_url,omitempty"`
	AccessToken   string `json:"accessToken,omitempty"`
}

type GitAccountStatus struct {
}
