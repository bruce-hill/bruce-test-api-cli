// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bruce-hill/bruce-test-api-cli/internal/mocktest"
)

func TestStreamJsonStream(t *testing.T) {
	t.Skip("Mock server doesn't support application/x-ndjson responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"stream-json", "stream",
	)
}
