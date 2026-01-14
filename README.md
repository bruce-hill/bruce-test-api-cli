# Bruce Test API CLI

The official CLI for the Bruce Test API REST API.

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Homebrew

```sh
brew tap bruce-hill/tools
brew install bruce-test-api
```

### Installing with Go

```sh
go install 'github.com/bruce-hill/bruce-test-api-cli/cmd/bruce-test-api@latest'
```

<!-- x-release-please-end -->

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
bruce-test-api [resource] [command] [flags]
```

```sh
bruce-test-api update-count \
  --body 123
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
