package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GitRepoCache struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec GitRepoCacheSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status GitRepoCacheStatus `json:"status"`
}

type GitRepoCacheSpec struct {
	GitAccountID string          `json:"gitAccountID,omitempty"`
	ScmType      string          `json:"scmType,omitempty"`
	Repositories []GitRepository `json:"repositories,omitempty"`
}

type GitRepository struct {
	CloneURL    string          `json:"clone_url,omitempty"`
	Permissions map[string]bool `json:"permissions,omitempty"`
}

type GitRepoCacheStatus struct {
}
