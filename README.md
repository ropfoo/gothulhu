# Gothulhu

Gothulhu is a simple character generator for Call of Cthulhu 7th edition.

You can either run it as a CLI or as a web server.

```bash
# run cli
go run cmd/gothulhu/main.go

# run server
go run cmd/server/main.go

# run server with hot reloading
air

# run tests
go test ./...
```

## API Usage

```bash
# generate a character
curl "http://localhost:8000/api/generate?name=Peter_Shaw&age=23"
```

List of available fields:

| Field | Required |
| ----- | -------- |
| name  | yes      |
| age   | yes      |
| str   | no       |
| dex   | no       |
| con   | no       |
| siz   | no       |
| intl  | no       |
| wis   | no       |
| cha   | no       |

All optional fields are generated randomly if not provided.

## Docker

```bash
# build docker image
docker build -t gothulhu .

# run docker image
docker run -p 8000:8000 gothulhu
```
