package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func composeUp(t *testing.T, workingDir string, envVars map[string]string) string {
	opts := &docker.Options{
		// Directory where docker-compose.yml lives
		WorkingDir: workingDir,

		EnvVars: envVars,
	}

	docker.RunDockerCompose(t, opts, "up", "-d")
	return ""
}

func composeDown(t *testing.T, workingDir string) {
	opts := &docker.Options{
		// Directory where docker-compose.yml lives
		WorkingDir: workingDir,
	}
	defer docker.RunDockerCompose(t, opts, "down")
}

func verify(t *testing.T) {
	expectedMetric := "go_build_info"
	status, body := http_helper.HttpGet(t, "http://localhost:8081/metrics", nil)

	assert.Equal(t, 200, status)
	assert.Contains(t, body, expectedMetric)
}
