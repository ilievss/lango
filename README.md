# lango
A web service for language translations

### Tests

Run all tests

```bash
go test ./...
```

to show test coverage

```bash
go test -cover ./...
```

### Build

```bash
go install
```

### Run

```bash
$GOPATH/bin/lango
```

### Usage

To translate a word just send a POST request:

```bash
curl -X POST -d '{"english-word": "Spain"}'
```
