// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-cli/internal/apiquery"
	"github.com/bruce-hill/bruce-test-api-cli/internal/requestflag"
	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var paginationList = cli.Command{
	Name:    "list",
	Usage:   "Get foos",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "size",
			Usage:     "Page size",
			Default:   50,
			QueryPath: "size",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
	},
	Action:          handlePaginationList,
	HideHelpCommand: true,
}

func handlePaginationList(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := brucetestapi.PaginationListParams{}

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
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Pagination.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "pagination list", obj, format, transform)
	} else {
		iter := client.Pagination.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "pagination list", iter, format, transform)
	}
}
