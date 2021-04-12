package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetricGeneratorUnit(t *testing.T) {
	t.Parallel()
	app := runApp(t, buildApp(t))
	defer docker.Stop(t, []string{app}, &docker.StopOptions{})
	verify(t)
}

func verify(t *testing.T) {
	expectedMetric := "go_build_info"
	status, body := http_helper.HttpGet(t, "http://localhost:8081/metrics", nil)

	assert.Equal(t, 200, status)
	assert.Contains(t, body, expectedMetric)
}
