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
	Name:    "form-test",
	Usage:   "Demonstrates a form-data endpoint with various parameter types including path,\nquery, and header parameters. Accepts multipart form data for user updates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "version",
			Usage:    "The API version to use",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "User ID in the format usr_xxxxx",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "date",
			Usage:     "Date filter in ISO 8601 format (YYYY-MM-DD)",
			Required:  true,
			QueryPath: "date",
		},
		&requestflag.Flag[any]{
			Name:      "datetime",
			Usage:     "Full datetime filter in ISO 8601 format",
			Required:  true,
			QueryPath: "datetime",
		},
		&requestflag.Flag[any]{
			Name:      "time",
			Usage:     "Time filter in ISO 8601 format (HH:MM:SS)",
			Required:  true,
			QueryPath: "time",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			Usage:    "Required field for demonstration purposes",
			Required: true,
			BodyPath: "blorp",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Complex filter object for advanced querying",
			QueryPath: "filter",
		},
		&requestflag.Flag[any]{
			Name:      "id-or-index",
			Usage:     "Flexible identifier that can be either a numeric index or string ID",
			QueryPath: "idOrIndex",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter results by one or more tags",
			QueryPath: "tags",
		},
		&requestflag.Flag[[]any]{
			Name:     "many-something",
			Usage:    "Array of Something items",
			BodyPath: "many_somethings",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pet",
			Usage:    "List of user's pets",
			BodyPath: "pets",
		},
		&requestflag.Flag[any]{
			Name:     "pls-null",
			Usage:    "A null value field for testing null handling",
			BodyPath: "pls_null",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			Usage:    "User preference settings",
			BodyPath: "preferences",
		},
		&requestflag.Flag[any]{
			Name:     "something",
			Usage:    "A flexible type that can be either a number or an object with name and optional count properties",
			BodyPath: "something",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			Usage:      "Array of feature flag names",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			Usage:      "Trace ID string for distributed tracing",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleFormTest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.meta",
			Usage:      "Metadata filter options",
			InnerField: "meta",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by status value",
			InnerField: "status",
		},
	},
	"pet": {
		&requestflag.InnerFlag[string]{
			Name:       "pet.name",
			Usage:      "Name of the pet",
			InnerField: "name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pet.age",
			Usage:      "Age of the pet in years",
			InnerField: "age",
		},
	},
	"preferences": {
		&requestflag.InnerFlag[bool]{
			Name:       "preferences.alerts",
			Usage:      "Whether to enable alert notifications",
			InnerField: "alerts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preferences.theme",
			Usage:      "UI theme preference (e.g., dark, light)",
			InnerField: "theme",
		},
	},
})

var jsonTest = requestflag.WithInnerFlags(cli.Command{
	Name:    "json-test",
	Usage:   "Demonstrates a JSON endpoint with various parameter types including path, query,\nand header parameters. Accepts JSON body for user updates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "version",
			Usage:    "The API version to use",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "User ID in the format usr_xxxxx",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "date",
			Usage:     "Date filter in ISO 8601 format (YYYY-MM-DD)",
			Required:  true,
			QueryPath: "date",
		},
		&requestflag.Flag[any]{
			Name:      "datetime",
			Usage:     "Full datetime filter in ISO 8601 format",
			Required:  true,
			QueryPath: "datetime",
		},
		&requestflag.Flag[any]{
			Name:      "time",
			Usage:     "Time filter in ISO 8601 format (HH:MM:SS)",
			Required:  true,
			QueryPath: "time",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			Usage:    "Required field for demonstration purposes",
			Required: true,
			BodyPath: "blorp",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Complex filter object for advanced querying",
			QueryPath: "filter",
		},
		&requestflag.Flag[any]{
			Name:      "id-or-index",
			Usage:     "Flexible identifier that can be either a numeric index or string ID",
			QueryPath: "idOrIndex",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter results by one or more tags",
			QueryPath: "tags",
		},
		&requestflag.Flag[[]any]{
			Name:     "many-something",
			Usage:    "Array of Something items",
			BodyPath: "many_somethings",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pet",
			Usage:    "List of user's pets",
			BodyPath: "pets",
		},
		&requestflag.Flag[any]{
			Name:     "pls-null",
			Usage:    "A null value field for testing null handling",
			BodyPath: "pls_null",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			Usage:    "User preference settings",
			BodyPath: "preferences",
		},
		&requestflag.Flag[any]{
			Name:     "something",
			Usage:    "A flexible type that can be either a number or an object with name and optional count properties",
			BodyPath: "something",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			Usage:      "Array of feature flag names",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			Usage:      "Trace ID string for distributed tracing",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleJsonTest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.meta",
			Usage:      "Metadata filter options",
			InnerField: "meta",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by status value",
			InnerField: "status",
		},
	},
	"pet": {
		&requestflag.InnerFlag[string]{
			Name:       "pet.name",
			Usage:      "Name of the pet",
			InnerField: "name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pet.age",
			Usage:      "Age of the pet in years",
			InnerField: "age",
		},
	},
	"preferences": {
		&requestflag.InnerFlag[bool]{
			Name:       "preferences.alerts",
			Usage:      "Whether to enable alert notifications",
			InnerField: "alerts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preferences.theme",
			Usage:      "UI theme preference (e.g., dark, light)",
			InnerField: "theme",
		},
	},
})

var updateCount = cli.Command{
	Name:    "update-count",
	Usage:   "Updates the current count with a new integer value. The value must be a positive\ninteger (minimum 1).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "body",
			Usage:    "A positive integer representing the new count value",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleUpdateCount,
	HideHelpCommand: true,
}

var uploadTest = cli.Command{
	Name:    "upload-test",
	Usage:   "Accepts a binary file upload using multipart/form-data. Typical use cases\ninclude uploading images, documents, or other opaque binaries.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The binary file to upload.",
			Required: true,
			BodyPath: "file",
		},
	},
	Action:          handleUploadTest,
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UpdateCount(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "update-count", obj, format, transform)
}

func handleUploadTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := brucetestapi.UploadTestParams{}

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
	_, err = client.UploadTest(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "upload-test", obj, format, transform)
}
