// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bruce-hill/bruce-test-api-cli/internal/mocktest"
)

func TestPaginationList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pagination", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page", "1",
			"--size", "1",
			"--tag", "string",
		)
	})
}
