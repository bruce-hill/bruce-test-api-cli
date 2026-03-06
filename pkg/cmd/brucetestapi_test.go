// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bruce-hill/bruce-test-api-cli/internal/mocktest"
	"github.com/bruce-hill/bruce-test-api-cli/internal/requestflag"
)

func TestDeleteTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"delete-test",
		"--api-key", "string",
	)
}

func TestDownloadTest(t *testing.T) {
	t.Skip("Mock server doesn't support application/octet-stream responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"download-test",
		"--api-key", "string",
		"--output", "/dev/null",
	)
}

func TestFormTest(t *testing.T) {
	t.Skip("prism issues because prism is not good at its job")
	mocktest.TestRunMockTestWithFlags(
		t,
		"form-test",
		"--api-key", "string",
		"--version", "2",
		"--user-id", "usr_abc123",
		"--date", "'2019-12-27'",
		"--datetime", "'2019-12-27T18:11:19.117Z'",
		"--time", "18:11:19.117Z",
		"--blorp", "example value",
		"--filter", "{meta: {level: 0}, status: status}",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--many-something", "123",
		"--many-something", "{name: Bozo, count: 5}",
		"--many-something", "{name: Fran, count: 5}",
		"--pet", "{name: Alfie, age: 0}",
		"--pet", "{name: Brando, age: 12}",
		"--pet", "{name: Charlie, age: 0}",
		"--pls-null", "null",
		"--preferences", "{alerts: true, theme: dark}",
		"--something", "{name: Albert, count: 5}",
		"--x-flag", "feature-a",
		"--x-flag", "feature-b",
		"--x-trace-id", "trace-xyz-789",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(formTest)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"form-test",
		"--version", "2",
		"--user-id", "usr_abc123",
		"--date", "'2019-12-27'",
		"--datetime", "'2019-12-27T18:11:19.117Z'",
		"--time", "18:11:19.117Z",
		"--blorp", "example value",
		"--filter.meta", "{level: 0}",
		"--filter.status", "status",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--many-something", "123",
		"--many-something", "{name: Bozo, count: 5}",
		"--many-something", "{name: Fran, count: 5}",
		"--pet.name", "Alfie",
		"--pet.age", "0",
		"--pet.name", "Brando",
		"--pet.age", "12",
		"--pet.name", "Charlie",
		"--pet.age", "0",
		"--pls-null", "null",
		"--preferences.alerts=true",
		"--preferences.theme", "dark",
		"--something", "{name: Albert, count: 5}",
		"--x-flag", "feature-a",
		"--x-flag", "feature-b",
		"--x-trace-id", "trace-xyz-789",
	)
}

func TestJsonTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"json-test",
		"--api-key", "string",
		"--version", "3",
		"--user-id", "usr_def456",
		"--date", "'2019-12-27'",
		"--datetime", "'2019-12-27T18:11:19.117Z'",
		"--time", "18:11:19.117Z",
		"--blorp", "test data",
		"--filter", "{meta: {level: 0}, status: status}",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--many-something", "123",
		"--many-something", "{name: Bozo, count: 5}",
		"--many-something", "{name: Fran, count: 5}",
		"--pet", "{name: Alfie, age: 0}",
		"--pet", "{name: Brando, age: 12}",
		"--pet", "{name: Charlie, age: 0}",
		"--pls-null", "null",
		"--preferences", "{alerts: false, theme: light}",
		"--something", "{name: Albert, count: 5}",
		"--x-flag", "experimental",
		"--x-flag", "verbose",
		"--x-trace-id", "trace-abc-123",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(jsonTest)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"json-test",
		"--version", "3",
		"--user-id", "usr_def456",
		"--date", "'2019-12-27'",
		"--datetime", "'2019-12-27T18:11:19.117Z'",
		"--time", "18:11:19.117Z",
		"--blorp", "test data",
		"--filter.meta", "{level: 0}",
		"--filter.status", "status",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--many-something", "123",
		"--many-something", "{name: Bozo, count: 5}",
		"--many-something", "{name: Fran, count: 5}",
		"--pet.name", "Alfie",
		"--pet.age", "0",
		"--pet.name", "Brando",
		"--pet.age", "12",
		"--pet.name", "Charlie",
		"--pet.age", "0",
		"--pls-null", "null",
		"--preferences.alerts=false",
		"--preferences.theme", "light",
		"--something", "{name: Albert, count: 5}",
		"--x-flag", "experimental",
		"--x-flag", "verbose",
		"--x-trace-id", "trace-abc-123",
	)
}

func TestNullableTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"nullable-test",
		"--api-key", "string",
		"--field", "null",
	)
}

func TestUpdateCount(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"update-count",
		"--api-key", "string",
		"--body", "42",
	)
}

func TestUploadTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"upload-test",
		"--api-key", "string",
		"--file", "...",
	)
}

func TestVersion(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"version",
		"--api-key", "string",
	)
}

func TestVoidTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"void-test",
		"--api-key", "string",
	)
}
