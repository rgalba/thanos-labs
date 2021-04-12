package test

import (
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetricGeneratorAndPrometheusIntegration(t *testing.T) {
	t.Parallel()

	const workingDir = "./fixtures/scenario_02"

	composeUp(t, workingDir, nil)
	defer composeDown(t, workingDir)
	verify(t)
	verifyPrometheus(t)
}

func verifyPrometheus(t *testing.T) {
	status, body := http_helper.HttpGet(t, "http://localhost:9090/metrics", nil)

	assert.Equal(t, 200, status)
	assert.Contains(t, body, "prometheus_http_requests_total")
}
