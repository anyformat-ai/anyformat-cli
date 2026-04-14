// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/anyformat-ai/anyformat-go"
	"github.com/anyformat-ai/anyformat-go/option"
	"github.com/stainless-sdks/anyformat-cli/internal/apiquery"
	"github.com/stainless-sdks/anyformat-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var filesCreate = cli.Command{
	Name:    "create",
	Usage:   "Upload files to a workflow, creating a file collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "file",
			Required: true,
			BodyPath: "files",
		},
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Required: true,
			BodyPath: "workflow_id",
		},
	},
	Action:          handleFilesCreate,
	HideHelpCommand: true,
}

var filesList = cli.Command{
	Name:    "list",
	Usage:   "List file collections for a workflow.",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "workflow-id",
			QueryPath: "workflow_id",
		},
	},
	Action:          handleFilesList,
	HideHelpCommand: true,
}

var filesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a file collection and all its files.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "collection-id",
			Required: true,
		},
	},
	Action:          handleFilesDelete,
	HideHelpCommand: true,
}

var filesGetExtractionResults = cli.Command{
	Name:    "get-extraction-results",
	Usage:   "Get extraction results for a file collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "collection-id",
			Required: true,
		},
	},
	Action:          handleFilesGetExtractionResults,
	HideHelpCommand: true,
}

func handleFilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.FileNewParams{}

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
	_, err = client.Files.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files create", obj, format, transform)
}

func handleFilesList(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := anyformat.FileListParams{}

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
	_, err = client.Files.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files list", obj, format, transform)
}

func handleFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("collection-id") && len(unusedArgs) > 0 {
		cmd.Set("collection-id", unusedArgs[0])
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

	return client.Files.Delete(ctx, cmd.Value("collection-id").(string), options...)
}

func handleFilesGetExtractionResults(ctx context.Context, cmd *cli.Command) error {
	client := anyformat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("collection-id") && len(unusedArgs) > 0 {
		cmd.Set("collection-id", unusedArgs[0])
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
	_, err = client.Files.GetExtractionResults(ctx, cmd.Value("collection-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files get-extraction-results", obj, format, transform)
}
