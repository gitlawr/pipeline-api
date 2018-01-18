package client

const (
	CronTriggerType                 = "cronTrigger"
	CronTriggerFieldSpec            = "spec"
	CronTriggerFieldTimezone        = "timezone"
	CronTriggerFieldTriggerOnUpdate = "triggerOnUpdate"
)

type CronTrigger struct {
	Spec            string `json:"spec,omitempty"`
	Timezone        string `json:"timezone,omitempty"`
	TriggerOnUpdate *bool  `json:"triggerOnUpdate,omitempty"`
}
