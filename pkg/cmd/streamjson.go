// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-cli/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/urfave/cli/v3"
)

var streamJsonStream = cli.Command{
	Name:            "stream",
	Usage:           "Streams JSON objects as a chunked response (NDJSON)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleStreamJsonStream,
	HideHelpCommand: true,
}

func handleStreamJsonStream(ctx context.Context, cmd *cli.Command) error {
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.StreamJson.StreamStreaming(ctx, options...)
	return ShowJSONIterator(os.Stdout, "stream-json stream", stream, format, transform)
}
