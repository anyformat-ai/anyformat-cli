// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/anyformat-cli/internal/mocktest"
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

func TestWorkflowsResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "results",
			"--workflow-id", "workflow_id",
			"--as-lists", "as_lists",
			"--output-format", "jsonl",
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
			"--content-type", "content_type",
			"--file", "file",
			"--file-base64", "file_base64",
			"--filename", "filename",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content_type: content_type\n" +
			"file: file\n" +
			"file_base64: file_base64\n" +
			"filename: filename\n" +
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
			"--content-type", "content_type",
			"--file", "file",
			"--file-base64", "file_base64",
			"--filename", "filename",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content_type: content_type\n" +
			"file: file\n" +
			"file_base64: file_base64\n" +
			"filename: filename\n" +
			"text: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "upload",
			"--workflow-id", "workflow_id",
		)
	})
}
