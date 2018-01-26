package config

import (
	"github.com/urfave/cli"
)

type config struct {
	JenkinsUser     string
	JenkinsToken    string
	JenkinsAddress  string
}

var Config config

func Parse(context *cli.Context) {
	Config.JenkinsAddress = context.String("jenkins_address")
	Config.JenkinsUser = context.String("jenkins_user")
	Config.JenkinsToken = context.String("jenkins_token")
}
