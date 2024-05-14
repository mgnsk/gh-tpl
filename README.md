# gh-tpl

Simple tool to use Github CLI templates on any JSON.

For template functions, see https://cli.github.com/manual/gh_help_formatting.

Also includes the sprig template functions: https://masterminds.github.io/sprig.

# Usage

```sh
$ echo '[{"field": "one"}, {"field": "two"}]' | gh-tpl '{{range .}}{{.field}}{{"\n"}}{{end}}'
one
two
```

# Install

Installation from release or from source:

```
go install github.com/mgnsk/gh-tpl@latest
```
