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
curl http://localhost:8080 -X POST -d '{"english-word": "Spain"}'
```

To translate a sentence:

```bash
curl http://localhost:8080 -X POST -d '{"english-sentence": "Spain is not my home country."}'
```
