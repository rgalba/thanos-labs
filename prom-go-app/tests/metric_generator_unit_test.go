package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"testing"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
)

func TestMetricGeneratorUnit(t *testing.T) {
	tag := "metric-generator"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	dockerFilePath := "../"
	docker.Build(t, dockerFilePath, buildOptions)

	opts := &docker.RunOptions{
		Command:      []string{},
		Detach:       true,
		OtherOptions: []string{"-p", "8080:8080"},
	}
	containerID := docker.RunAndGetID(t, tag, opts)
	defer docker.Stop(t, []string{containerID}, &docker.StopOptions{})
	
	expectedMetric := "go_build_info"
	status, body := http_helper.HttpGet(t, "http://localhost:8080/metrics", nil)
	
	assert.NotNil(t, containerID)
	assert.Equal(t, 200, status)
	assert.Contains(t, body, expectedMetric)
}
