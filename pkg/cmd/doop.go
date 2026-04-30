// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-cli/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-cli/internal/requestflag"
	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/urfave/cli/v3"
)

var doopDownload = cli.Command{
	Name:    "download",
	Usage:   "Download a file using application/octet-stream",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleDoopDownload,
	HideHelpCommand: true,
}

func handleDoopDownload(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	response, err := client.Doop.Download(ctx, options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
