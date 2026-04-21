// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/anyformat-ai/anyformat-cli/internal/apiquery"
	"github.com/anyformat-ai/anyformat-cli/internal/requestflag"
	"github.com/anyformat-ai/anyformat-go"
	"github.com/anyformat-ai/anyformat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var workflowsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new workflow.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWorkflowsCreate,
	HideHelpCommand: true,
}

var workflowsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get workflow by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
	},
	Action:          handleWorkflowsRetrieve,
	HideHelpCommand: true,
}

var workflowsList = cli.Command{
	Name:    "list",
	Usage:   "List workflows with pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "order",
			QueryPath: "order",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page_size",
		},
		&requestflag.Flag[any]{
			Name:      "sort-by",
			QueryPath: "sort_by",
		},
		&requestflag.Flag[any]{
			Name:      "status",
			QueryPath: "status",
		},
	},
	Action:          handleWorkflowsList,
	HideHelpCommand: true,
}

var workflowsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete workflow by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
	},
	Action:          handleWorkflowsDelete,
	HideHelpCommand: true,
}

var workflowsCreateFile = cli.Command{
	Name:    "create-file",
	Usage:   "Upload files to a workflow, creating a file collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "file",
			Required: true,
			BodyPath: "files",
		},
	},
	Action:          handleWorkflowsCreateFile,
	HideHelpCommand: true,
}

var workflowsGetFileResults = cli.Command{
	Name:    "get-file-results",
	Usage:   "Get processing results for a file collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "collection-id",
			Required: true,
		},
	},
	Action:          handleWorkflowsGetFileResults,
	HideHelpCommand: true,
}

var workflowsListFiles = cli.Command{
	Name:    "list-files",
	Usage:   "List file collections for a workflow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page_size",
		},
	},
	Action:          handleWorkflowsListFiles,
	HideHelpCommand: true,
}

var workflowsListRuns = cli.Command{
	Name:    "list-runs",
	Usage:   "List extraction runs for a workflow, identified by collection UUID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page_size",
		},
	},
	Action:          handleWorkflowsListRuns,
	HideHelpCommand: true,
}

var workflowsRun = cli.Command{
	Name:    "run",
	Usage:   "Execute workflow — returns collection UUID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "file",
			BodyPath: "file",
		},
		&requestflag.Flag[any]{
			Name:     "text",
			BodyPath: "text",
		},
	},
	Action:          handleWorkflowsRun,
	HideHelpCommand: true,
}

var workflowsUpload = cli.Command{
	Name:    "upload",
	Usage:   "Upload file without executing workflow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "file",
			BodyPath: "file",
		},
		&requestflag.Flag[any]{
			Name:     "text",
			BodyPath: "text",
		},
	},
	Action:          handleWorkflowsUpload,
	HideHelpCommand: true,
}

func handleWorkflowsCreate(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.New(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows create",
		Transform:      transform,
	})
}

func handleWorkflowsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.Get(ctx, cmd.Value("workflow-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows retrieve",
		Transform:      transform,
	})
}

func handleWorkflowsList(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows list",
		Transform:      transform,
	})
}

func handleWorkflowsDelete(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Workflows.Delete(ctx, cmd.Value("workflow-id").(string), options...)
}

func handleWorkflowsCreateFile(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowNewFileParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.NewFile(
		ctx,
		cmd.Value("workflow-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows create-file",
		Transform:      transform,
	})
}

func handleWorkflowsGetFileResults(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("collection-id") && len(unusedArgs) > 0 {
		cmd.Set("collection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowGetFileResultsParams{
		WorkflowID: cmd.Value("workflow-id").(string),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.GetFileResults(
		ctx,
		cmd.Value("collection-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows get-file-results",
		Transform:      transform,
	})
}

func handleWorkflowsListFiles(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowListFilesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.ListFiles(
		ctx,
		cmd.Value("workflow-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows list-files",
		Transform:      transform,
	})
}

func handleWorkflowsListRuns(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowListRunsParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.ListRuns(
		ctx,
		cmd.Value("workflow-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows list-runs",
		Transform:      transform,
	})
}

func handleWorkflowsRun(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowRunParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.Run(
		ctx,
		cmd.Value("workflow-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows run",
		Transform:      transform,
	})
}

func handleWorkflowsUpload(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-id") && len(unusedArgs) > 0 {
		cmd.Set("workflow-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.WorkflowUploadParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workflows.Upload(
		ctx,
		cmd.Value("workflow-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "workflows upload",
		Transform:      transform,
	})
}
