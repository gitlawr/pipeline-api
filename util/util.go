package util
import(
	model "github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3"
	"regexp"
	"math/rand"
	"time"
)


func HasStepCondition(s model.Step) bool {
	return len(s.Conditions.All) > 0 || len(s.Conditions.Any) > 0
}

func HasStageCondition(s model.Stage) bool {
	return len(s.Conditions.All) > 0 || len(s.Conditions.Any) > 0
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func GetParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}


func IsStageSuccess(stage model.ActivityStage) bool {
	if stage.Status == model.ActivityStageFail || stage.Status == model.ActivityStageDenied {
		return false
	}
	successSteps := 0
	for _, step := range stage.ActivitySteps {
		if step.Status == model.ActivityStepSuccess || step.Status == model.ActivityStepSkip {
			successSteps++
		}
	}
	return successSteps == len(stage.ActivitySteps)
}


//GetServices gets run services before the step
func GetServices(activity *model.Activity, stageOrdinal int, stepOrdinal int) []model.CIService {
	services := []model.CIService{}
	for i := 0; i <= stageOrdinal; i++ {
		for j := 0; j < len(activity.Pipeline.Stages[i].Steps); j++ {
			if i == stageOrdinal && j >= stepOrdinal {
				break
			}
			step := activity.Pipeline.Stages[i].Steps[j]
			if step.IsService && step.Type == model.StepTypeTask {
				service := model.CIService{
					ContainerName: activity.Id + step.Alias,
					Name:          step.Alias,
					Image:         step.Image,
				}
				services = append(services, service)
			}
		}
	}
	return services
}
