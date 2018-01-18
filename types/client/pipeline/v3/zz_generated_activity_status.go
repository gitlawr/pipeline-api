package client

const (
	ActivityStatusType     = "activityStatus"
	ActivityStatusFieldTBD = "tbd"
)

type ActivityStatus struct {
	TBD string `json:"tbd,omitempty"`
}
