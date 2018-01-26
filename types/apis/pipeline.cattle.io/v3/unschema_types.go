package v3

type PipelineProvider interface {
	RunPipeline(*Pipeline, string) (*Activity, error)
	RerunActivity(*Activity) error
	RunStage(*Activity, int) error
	RunStep(*Activity, int, int) error
	StopActivity(*Activity) error
	SyncActivity(*Activity) error
	GetStepLog(*Activity, int, int, map[string]interface{}) (string, error)
	OnActivityCompelte(*Activity)
	OnCreateAccount(*GitAccount) error
	OnDeleteAccount(*GitAccount) error
	Reset() error


}

type ExecuteEngine interface {
	GetName() string
	RUNPipeline(pipeline *Pipeline, triggerType string) (*Activity, error)
	StopActivity(activity *Activity)
	GetStepLog(*Step) (string, error)
	OnStart(*Activity) error
	OnComplete(*Activity) error
}

type SourceCodeProvider interface{
	
}