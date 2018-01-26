package main

import (
	"context"
	"net/http"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	"github.com/rancher/pipeline-api/types/config"
	"github.com/rancher/pipeline-api/server"
	globalconfig "github.com/rancher/pipeline-api/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "pipeline"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = checkAndRun
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "jenkins_user",
			Usage:  "User of jenkins admin",
			EnvVar: "JENKINS_USER",
		},
		cli.StringFlag{
			Name:   "jenkins_token",
			Usage:  "token of jenkins admin",
			EnvVar: "JENKINS_TOKEN",
		},
		cli.StringFlag{
			Name:   "jenkins_address",
			Usage:  "token of jenkins admin",
			EnvVar: "JENKINS_ADDRESS",
			Value:  "http://jenkins:8080",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug mode",
			EnvVar: "DEBUG",
		},
	}
	app.Run(os.Args)
}

func checkAndRun(c *cli.Context) (rtnerr error) {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}
	globalconfig.Parse(c)

	if err := run(); err != nil {
		return err
	}
	return nil
}


func run() error{
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	pipeline, err := config.NewPipelineContext(*kubeConfig)
	if err != nil {
		return err
	}

	handler, err := server.New(context.Background(), pipeline)
	if err != nil {
		return err
	}

	logrus.Info("Listening on 0.0.0.0:1234")
	return http.ListenAndServe("0.0.0.0:1234", handler)
}
