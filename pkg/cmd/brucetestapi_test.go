// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bruce-hill/bruce-test-api-cli/internal/mocktest"
	"github.com/bruce-hill/bruce-test-api-cli/internal/requestflag"
)

func TestDeleteTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "delete-test",
			"--api-key", "string",
		)
	})
}

func TestDownloadTest(t *testing.T) {
	t.Skip("Mock server doesn't support application/octet-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "download-test",
			"--api-key", "string",
			"--output", "/dev/null",
		)
	})
}

func TestFormTest(t *testing.T) {
	t.Skip("prism issues because prism is not good at its job")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "form-test",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(formTest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "form-test",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blorp: example value\n" +
			"many_somethings:\n" +
			"  - 123\n" +
			"  - name: Bozo\n" +
			"    count: 5\n" +
			"  - name: Fran\n" +
			"    count: 5\n" +
			"pets:\n" +
			"  - name: Alfie\n" +
			"    age: 0\n" +
			"  - name: Brando\n" +
			"    age: 12\n" +
			"  - name: Charlie\n" +
			"    age: 0\n" +
			"pls_null: null\n" +
			"preferences:\n" +
			"  alerts: true\n" +
			"  theme: dark\n" +
			"something:\n" +
			"  name: Albert\n" +
			"  count: 5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "form-test",
			"--api-key", "string",
			"--version", "2",
			"--user-id", "usr_abc123",
			"--date", "'2019-12-27'",
			"--datetime", "'2019-12-27T18:11:19.117Z'",
			"--time", "18:11:19.117Z",
			"--filter", "{meta: {level: 0}, status: status}",
			"--id-or-index", "0",
			"--limit", "1",
			"--tag", "string",
			"--x-flag", "feature-a",
			"--x-flag", "feature-b",
			"--x-trace-id", "trace-xyz-789",
		)
	})
}

func TestJsonTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "json-test",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(jsonTest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "json-test",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blorp: test data\n" +
			"many_somethings:\n" +
			"  - 123\n" +
			"  - name: Bozo\n" +
			"    count: 5\n" +
			"  - name: Fran\n" +
			"    count: 5\n" +
			"pets:\n" +
			"  - name: Alfie\n" +
			"    age: 0\n" +
			"  - name: Brando\n" +
			"    age: 12\n" +
			"  - name: Charlie\n" +
			"    age: 0\n" +
			"pls_null: null\n" +
			"preferences:\n" +
			"  alerts: false\n" +
			"  theme: light\n" +
			"something:\n" +
			"  name: Albert\n" +
			"  count: 5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "json-test",
			"--api-key", "string",
			"--version", "3",
			"--user-id", "usr_def456",
			"--date", "'2019-12-27'",
			"--datetime", "'2019-12-27T18:11:19.117Z'",
			"--time", "18:11:19.117Z",
			"--filter", "{meta: {level: 0}, status: status}",
			"--id-or-index", "0",
			"--limit", "1",
			"--tag", "string",
			"--x-flag", "experimental",
			"--x-flag", "verbose",
			"--x-trace-id", "trace-abc-123",
		)
	})
}

func TestNullableTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "nullable-test",
			"--api-key", "string",
			"--field", "null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("field: null")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "nullable-test",
			"--api-key", "string",
		)
	})
}

func TestUpdateCount(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "update-count",
			"--api-key", "string",
			"--body", "42",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("42")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "update-count",
			"--api-key", "string",
		)
	})
}

func TestUploadTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "upload-test",
			"--api-key", "string",
			"--file", "...",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "upload-test",
			"--api-key", "string",
		)
	})
}

func TestVersion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "version",
			"--api-key", "string",
		)
	})
}

func TestVoidTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "void-test",
			"--api-key", "string",
		)
	})
}
