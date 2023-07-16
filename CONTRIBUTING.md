## Contributing

First off, thank you for considering contributing to `unkey-go`. It's people like you that make it such a great tool.

Please read the [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.

**Note**: Please get your contributions approved by the maintainers of [unkey-go](https://github.com/WilfredAlmeida/unkey-go) or the official [unkey](https://github.com/unkeyed/unkey/) before you start working on them. This will help avoid any conflicts and also help you understand if the contribution is in line with the project goals. You can do this by opening an issue with appropriate label and describing your contribution in detail.

### You have found a bug

Please open an issue with the label `bug`. To clearly explain the bug, describe it in as much detail as possible. Please add the steps-to-reproduce as well.

### You are requesting a new feature

Please open an issue with the label `enhancement` and include as much information as you can about the need for the feature. The maintainers will convert it into a discussion if needed.

### You want to add a new feature

Please open an issue with the label `enhancement` and include as much information as you can about the need for the feature. The maintainers will convert it into a discussion if needed. Once approved, the maintainers will assign the issue to you or anyone who wishes to work on the issue.

### You want support

Please open an issue with the label `help wanted`. The maintainers and other users of the platform will respond to it.

### You want to contact the maintainers

Open an issue or discussion on GitHub. Alternatively you can ping us on [Discord](https://unkey.dev/discord) as well.

### Commits

Please keep commits modular. One big commit or commits with many files and breaking changes will be hard to understand and maintain and might lead to rejection of the contribution.

### Commit Messages

Please follow the semantic commit messages pattern similar to the one mentioned [here](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716).

### Pull Requests
Each pull request must have it own branch named `<issue>-description`. For eg. `1-add-readme`.

### Anything Else

Open an issue with the label `random` and discuss them there. You can ping us on [Discord](https://unkey.dev/discord) as well.


## Developing Locally

1. Clone the repository
2. Run `go get` to download all the dependencies
3. That's it! You are ready to go.

`unkey-go` doesn't have any external dependencies. It uses the standard library for all its operations. The only external dependency is for loading `.env` files necessary for testing secrets.

### Testing

The tests are written in dedicated files. For tests you need a `.env` file with the following secrets:
- `AUTH_TOKEN`: Bearer auth token. Needed for creating & revoking keys.
- `API_ID`: The API Id.

## Development Versions
Following are the softwares & their versions used for development:
- `go version go1.20.5 linux/amd64`
- `Ubuntu 22 on WSL2`
- IDE: `MS Word`