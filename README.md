![golangci-lint](https://github.com/crab-apple/ynab-tui/actions/workflows/golangci-lint.yml/badge.svg)
![build](https://github.com/crab-apple/ynab-tui/actions/workflows/go.yml/badge.svg)

A TUI (terminal user interface) for [YNAB](https://www.ynab.com/).

This project is in a very early stage of development. It is not usable yet.

The goal is to progressively add functionality according to the
needs of the developers or potential users.
I expect that the majority of this functionality will replicate existing
functionality in YNAB web/apps (or rather, it will just delegate to it via
YNAB's API). However, I have no intention to reach feature parity with other
YNAB UIs, nor to be constrained to what they do.

### Running the tests

At the project root, run
```text
go test ./...
```

### Running the app

1. Obtain a [personal access token](https://api.ynab.com/#personal-access-tokens) for your YNAB account.
2. Write your access token into a file `<YOUR_HOME_DIRECTORY>/.ynab/access_token`.
3. At the project root, run
   ```text
   go run ./...
   ```