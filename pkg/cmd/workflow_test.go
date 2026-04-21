// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/anyformat-ai/anyformat-cli/internal/mocktest"
)

func TestWorkflowsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create",
		)
	})
}

func TestWorkflowsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "retrieve",
			"--workflow-id", "workflow_id",
		)
	})
}

func TestWorkflowsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "list",
			"--order", "order",
			"--page", "1",
			"--page-size", "1",
			"--sort-by", "sort_by",
			"--status", "status",
		)
	})
}

func TestWorkflowsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "delete",
			"--workflow-id", "workflow_id",
		)
	})
}

func TestWorkflowsCreateFile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create-file",
			"--workflow-id", "workflow_id",
			"--file", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"files:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "create-file",
			"--workflow-id", "workflow_id",
		)
	})
}

func TestWorkflowsGetFileResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "get-file-results",
			"--workflow-id", "workflow_id",
			"--collection-id", "collection_id",
		)
	})
}

func TestWorkflowsListFiles(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "list-files",
			"--workflow-id", "workflow_id",
			"--page", "1",
			"--page-size", "1",
		)
	})
}

func TestWorkflowsListRuns(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "list-runs",
			"--workflow-id", "workflow_id",
			"--page", "1",
			"--page-size", "1",
		)
	})
}

func TestWorkflowsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "run",
			"--workflow-id", "workflow_id",
			"--file", "file",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file: file\n" +
			"text: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "run",
			"--workflow-id", "workflow_id",
		)
	})
}

func TestWorkflowsUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "upload",
			"--workflow-id", "workflow_id",
			"--file", "file",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file: file\n" +
			"text: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "upload",
			"--workflow-id", "workflow_id",
		)
	})
}
