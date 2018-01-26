package jenkins

import (
	"errors"
	"sync"
)

type jenkinsConfig map[string]string

const JenkinsServerAddress = "JenkinsServerAddress"
const JenkinsUser = "JenkinsUser"
const JenkinsToken = "JenkinsToken"
const CreateJobURI = "CreateJobURI"
const UpdateJobURI = "UpdateJobURI"
const StopJobURI = "StopjobURI"
const CancelQueueItemURI = "CancelQueueItemURI"
const ScriptURI = "ScriptURI"
const DeleteBuildURI = "DeleteBuildURI"
const GetCrumbURI = "GetCrumbURI"
const JenkinsCrumbHeader = "JenkinsCrumbHeader"
const JenkinsCrumb = "JenkinsCrumb"
const JenkinsJobBuildURI = "JenkinsJobBuildURI"
const JenkinsJobInfoURI = "JenkinsJobInfoURI"
const JenkinsSetCredURI = "JenkinsSetCredURI"
const JenkinsDeleteCredURI = "JenkinsDeleteCredURI"
const JenkinsBuildInfoURI = "JenkinsBuildInfoURI"
const JenkinsBuildLogURI = "JenkinsBuildLogURI"
const JenkinsJobBuildWithParamsURI = "JenkinsJobBuildWithParamsURI"

var ErrConfigItemNotFound = errors.New("Jenkins configuration not fount")
var jenkinsConfLock = &sync.RWMutex{}

func (j jenkinsConfig) Set(key, value string) {
	jenkinsConfLock.Lock()
	defer jenkinsConfLock.Unlock()
	j[key] = value
}

func (j jenkinsConfig) Get(key string) (string, error) {
	jenkinsConfLock.RLock()
	defer jenkinsConfLock.RUnlock()
	if value, ok := j[key]; ok {
		return value, nil
	}
	return "", ErrConfigItemNotFound
}

var JenkinsConfig = jenkinsConfig{
	CreateJobURI:                 "/createItem",
	UpdateJobURI:                 "/job/%s/config.xml",
	StopJobURI:                   "/job/%s/lastBuild/stop",
	CancelQueueItemURI:           "/queue/cancelItem?id=%d",
	DeleteBuildURI:               "/job/%s/lastBuild/doDelete",
	GetCrumbURI:                  "/crumbIssuer/api/xml?xpath=concat(//crumbRequestField,\":\",//crumb)",
	JenkinsJobBuildURI:           "/job/%s/build",
	JenkinsJobBuildWithParamsURI: "/job/%s/buildWithParameters",
	JenkinsJobInfoURI:            "/job/%s/api/json",
	JenkinsSetCredURI:            "/credentials/store/system/domain/_/createCredentials",
	JenkinsDeleteCredURI:         "/credentials/store/system/domain/_/credential/%s/doDelete",
	JenkinsBuildInfoURI:          "/job/%s/lastBuild/api/json",
	JenkinsBuildLogURI:           "/job/%s/lastBuild/timestamps/?elapsed=HH'h'mm'm'ss's'S'ms'&appendLog",
	ScriptURI:                    "/scriptText",
}

//Script to execute on specific node
const ScriptSkel = `import hudson.util.RemotingDiagnostics; 
node = "%s"
cmd = "def proc = ['bash', '-c', '%s'].execute();proc.waitFor();println proc.in.text;"
for (slave in hudson.model.Hudson.instance.slaves) {
  if(slave.name==node){
	println RemotingDiagnostics.executeGroovy(cmd, slave.getChannel());
  }
}
//on master
if(node == "master"){
	def proc = script.execute(); proc.waitFor(); println proc.in.text
}
`

const GetActiveNodesScript = `for (slave in hudson.model.Hudson.instance.slaves) {
  if (!slave.getComputer().isOffline()){
	    println slave.name;
  }
}
`

const stepFinishScript = `def result = manager.build.result
def command =  ["sh","-c","curl -s -d '' 'pipeline-server:60080/v1/events/stepfinish?id=%v&status=${result}&stageOrdinal=%v&stepOrdinal=%v'"]
manager.listener.logger.println command.execute().text`

const stepSCMFinishScript = `def result = manager.build.result
def env = manager.build.environment
def GIT_COMMIT = env.get("GIT_COMMIT")
def GIT_URL = env.get("GIT_URL")
def GIT_BRANCH = env.get("GIT_BRANCH")
def command =  ["sh","-c","curl -s -d 'GIT_URL=${GIT_URL}&GIT_BRANCH=${GIT_BRANCH}&GIT_COMMIT=${GIT_COMMIT}' 'pipeline-server:60080/v1/events/stepfinish?id=%v&status=${result}&stageOrdinal=%v&stepOrdinal=%v'"]
manager.listener.logger.println command.execute().text`

const stepStartScript = "curl -s -d '' 'pipeline-server:60080/v1/events/stepstart?id=%v&stageOrdinal=%v&stepOrdinal=%v'"
