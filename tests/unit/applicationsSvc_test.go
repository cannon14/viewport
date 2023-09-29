package unit

import (
	"test/app/utils"
	"testing"
)

func TestExtractVulnerabilityCounts(t *testing.T) {

	data := "Vulnerabilities: CRITICAL: 1, HIGH: 3, MEDIUM: 0, LOW: 0, INFORMATIONAL: 0"
	result := utils.ExtractVulnerabilityCounts(data)

	if result["critical"] != 1 {
		t.Errorf("Result = %d; Wanted 1", result["critical"])
	}

	if result["high"] != 3 {
		t.Errorf("Result = %d; Wanted 3", result["high"])
	}
}

func TestExtractVulnerabilityCountsEmptyString(t *testing.T) {

	data := ""
	result := utils.ExtractVulnerabilityCounts(data)

	if result["critical"] != 0 {
		t.Errorf("Result = %d; Wanted 0", result["critical"])
	}

	if result["high"] != 0 {
		t.Errorf("Result = %d; Wanted 0", result["high"])
	}
}
