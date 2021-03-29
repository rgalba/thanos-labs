package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
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
