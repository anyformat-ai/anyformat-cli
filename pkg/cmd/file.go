// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/anyformat-ai/anyformat-cli/internal/apiquery"
	"github.com/anyformat-ai/anyformat-cli/internal/requestflag"
	"github.com/anyformat-ai/anyformat-go"
	"github.com/urfave/cli/v3"
)

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
