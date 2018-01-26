package v3

const (
	StepTypeTask           = "task"
	StepTypeDeploy         = "deploy"
	StepTypeSCM            = "scm"
	StepTypeBuild          = "build"
	StepTypeService        = "service"
	StepTypeUpgradeService = "upgradeService"
	StepTypeUpgradeStack   = "upgradeStack"
	StepTypeUpgradeCatalog = "upgradeCatalog"
	TriggerTypeCron        = "cron"
	TriggerTypeManual      = "manual"
	TriggerTypeWebhook     = "webhook"

	ActivityStepWaiting  = "Waiting"
	ActivityStepBuilding = "Building"
	ActivityStepSuccess  = "Success"
	ActivityStepFail     = "Fail"
	ActivityStepSkip     = "Skipped"
	ActivityStepAbort    = "Abort"

	ActivityStageWaiting  = "Waiting"
	ActivityStagePending  = "Pending"
	ActivityStageBuilding = "Building"
	ActivityStageSuccess  = "Success"
	ActivityStageFail     = "Fail"
	ActivityStageDenied   = "Denied"
	ActivityStageSkip     = "Skipped"
	ActivityStageAbort    = "Abort"

	ActivityWaiting  = "Waiting"
	ActivityPending  = "Pending"
	ActivityBuilding = "Building"
	ActivitySuccess  = "Success"
	ActivityFail     = "Fail"
	ActivityDenied   = "Denied"
	ActivityAbort    = "Abort"
)
