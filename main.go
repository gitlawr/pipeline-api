package main

import (
	"context"
	"net/http"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	"github.com/rancher/pipeline-api/types/config"
	"github.com/rancher/pipeline-api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
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
