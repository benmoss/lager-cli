# Lager CLI

lager-cli is a CLI for viewing [Lager](http://code.cloudfoundry.org/lager) logs.

## Installing

```bash
go get -u github.com/benmoss/lager-cli
go install github.com/benmoss/lager-cli
```

## Usage
```bash
lager-cli chug [file] - Prettify lager logs
lager-cli chug-serve [file] - Serve up pretty lager logs
lager-cli chug-unify file1, file2,... - Combine lager files in temporal order
```
