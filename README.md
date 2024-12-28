# Gothulhu

Gothulhu is a character generator for Call of Cthulhu 7th edition.

You can either run it as a CLI or as a server.

```bash
# run cli
go run cmd/gothulhu/main.go

# run server
go run cmd/server/main.go
```

## Docker

```bash
# build docker image
docker build -t gothulhu .

# run docker image
docker run -p 8000:8000 gothulhu
```

### Testing

```bash
# run tests
go test ./...
```

```bash
# run tests with coverage
go test -coverprofile=coverage.out ./...
```
