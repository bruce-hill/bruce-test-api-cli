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

var formTest = requestflag.WithInnerFlags(cli.Command{
	Name:  "form-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "version",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "date",
			Required:  true,
			QueryPath: "date",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "datetime",
			Required:  true,
			QueryPath: "datetime",
		},
		&requestflag.Flag[requestflag.TimeValue]{
			Name:      "time",
			Required:  true,
			QueryPath: "time",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			QueryPath: "filter",
		},
		&requestflag.Flag[any]{
			Name:      "id-or-index",
			QueryPath: "idOrIndex",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			BodyPath: "blorp",
		},
		&requestflag.Flag[[]any]{
			Name:     "many-something",
			BodyPath: "many_somethings",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pet",
			BodyPath: "pets",
		},
		&requestflag.Flag[any]{
			Name:     "pls-null",
			BodyPath: "pls_null",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
		&requestflag.Flag[any]{
			Name:     "something",
			BodyPath: "something",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleFormTest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.meta",
			InnerField: "meta",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			InnerField: "status",
		},
	},
	"pet": {
		&requestflag.InnerFlag[string]{
			Name:       "pet.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pet.age",
			InnerField: "age",
		},
	},
	"preferences": {
		&requestflag.InnerFlag[bool]{
			Name:       "preferences.alerts",
			InnerField: "alerts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preferences.theme",
			InnerField: "theme",
		},
	},
})

var jsonTest = requestflag.WithInnerFlags(cli.Command{
	Name:  "json-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "version",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "date",
			Required:  true,
			QueryPath: "date",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "datetime",
			Required:  true,
			QueryPath: "datetime",
		},
		&requestflag.Flag[requestflag.TimeValue]{
			Name:      "time",
			Required:  true,
			QueryPath: "time",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			QueryPath: "filter",
		},
		&requestflag.Flag[any]{
			Name:      "id-or-index",
			QueryPath: "idOrIndex",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			BodyPath: "blorp",
		},
		&requestflag.Flag[[]any]{
			Name:     "many-something",
			BodyPath: "many_somethings",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pet",
			BodyPath: "pets",
		},
		&requestflag.Flag[any]{
			Name:     "pls-null",
			BodyPath: "pls_null",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
		&requestflag.Flag[any]{
			Name:     "something",
			BodyPath: "something",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleJsonTest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.meta",
			InnerField: "meta",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			InnerField: "status",
		},
	},
	"pet": {
		&requestflag.InnerFlag[string]{
			Name:       "pet.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pet.age",
			InnerField: "age",
		},
	},
	"preferences": {
		&requestflag.InnerFlag[bool]{
			Name:       "preferences.alerts",
			InnerField: "alerts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preferences.theme",
			InnerField: "theme",
		},
	},
})

var updateCount = cli.Command{
	Name:  "update-count",
	Usage: "Perform update-count operation",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleUpdateCount,
	HideHelpCommand: true,
}

func handleFormTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := brucetestapi.FormTestParams{
		Version: cmd.Value("version").(int64),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatRepeat,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FormTest(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "form-test", obj, format, transform)
}

func handleJsonTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := brucetestapi.JsonTestParams{
		Version: cmd.Value("version").(int64),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.JsonTest(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "json-test", obj, format, transform)
}

func handleUpdateCount(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := brucetestapi.UpdateCountParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.UpdateCount(ctx, params, options...)
}
