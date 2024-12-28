# Gothulhu

Gothulhu is a character generator for Call of Cthulhu 7th edition.

You can either run it as a CLI or as a server.

## Dev Setup

```bash
# run cli
go run cmd/gothulhu/main.go

# run server
go run cmd/server/main.go

# run server in docker
docker build -t gothulhu .
docker run -p 8000:8000 gothulhu
```

### Testing

```bash
go test ./...
```

```bash
go test -coverprofile=coverage.out ./...
```
