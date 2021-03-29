package test

import (
	"testing"
)

func TestMetricGeneratorIntegration(t *testing.T) {
	t.Parallel()
	const workingDir = "./fixtures/scenario_01"

	composeUp(t, workingDir, nil)
	defer composeDown(t, workingDir)
	verify(t)
}
