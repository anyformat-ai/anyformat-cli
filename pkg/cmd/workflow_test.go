// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/anyformat-ai/anyformat-cli/internal/mocktest"
	"github.com/anyformat-ai/anyformat-cli/internal/requestflag"
)

func TestWorkflowsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create",
			"--name", "Invoice or receipt",
			"--node", "{id: x, type: parse, effort: low, engine: Fast, figure_enhancement: true, mode: standard, prompt_hint: prompt_hint, visual_grounding_enabled: true}",
			"--description", "description",
			"--edge", "{source: x, target: x, branch: branch}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workflowsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create",
			"--name", "Invoice or receipt",
			"--node", "{id: x, type: parse, effort: low, engine: Fast, figure_enhancement: true, mode: standard, prompt_hint: prompt_hint, visual_grounding_enabled: true}",
			"--description", "description",
			"--edge.source", "x",
			"--edge.target", "x",
			"--edge.branch", "branch",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Invoice or receipt\n" +
			"nodes:\n" +
			"  - id: x\n" +
			"    type: parse\n" +
			"    effort: low\n" +
			"    engine: Fast\n" +
			"    figure_enhancement: true\n" +
			"    mode: standard\n" +
			"    prompt_hint: prompt_hint\n" +
			"    visual_grounding_enabled: true\n" +
			"description: description\n" +
			"edges:\n" +
			"  - source: x\n" +
			"    target: x\n" +
			"    branch: branch\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
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
