package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"testing"
)

func buildApp(t *testing.T) string {
	tag := "metric-generator"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	dockerFilePath := "../"
	docker.Build(t, dockerFilePath, buildOptions)
	return tag
}

func runApp(t *testing.T, tag string) string {
	opts := &docker.RunOptions{
		Command:      []string{},
		Detach:       true,
		OtherOptions: []string{"-p", "8080:8080"},
	}
	containerID := docker.RunAndGetID(t, tag, opts)
	return containerID
}
