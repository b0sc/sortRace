# Guidelines

-   Make sure your editor supports formatting with `golang.go`, and format accordingly before committing.
-   Use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
-   Use [GitHub Flow](https://docs.github.com/en/get-started/using-github/github-flow) for branching and use [Pull Request Template](./docs/pull_request_template.md) for pull requests.
-   Donot edit the `Makefile` without consulting the maintainers, if you require editing for your local setup, ignore the changes when commiting.

## Contributing Goals

-   Write code that is plugin-based (where an `algorithm.go` is considered a plugin) and easy to extend further by other developers.
-   If test files exists, `go test` should pass before creating a pull request.
-   If test files do not exist, throughly test the code before creating a pull request.
-   If you are adding a new feature, make sure to update the README.md file with the new feature.

### Formatting

If you use `gofmt` check if there is any formatting issue before committing. Alternatively, you can comment out `gofmt` and use `golang.go` extension for formatting. Your `settings.json` should look like this:

#### VS Code

```json
{
    //"go.formatTool": "gofmt",
    "[go]": {
        "editor.defaultFormatter": "golang.go",
        "editor.insertSpaces": true,
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": "explicit"
        }
    }
}
```

#### Goland

#### NeoVim

### Code of Conduct

Please refer to our [Code of Conduct](./CODE_OF_CONDUCT.md) for more information.
