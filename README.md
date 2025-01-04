<p align="center">
  <img src="https://github.com/ropfoo/gothulhu/blob/main/assets/favicon.ico" width="90" title="asmodeus">
</p>
<p align="center">
  <a href="https://www.gothulhu.app">gothulhu.app</a>
</p>

# Gothulhu

Gothulhu is a simple character generator for Call of Cthulhu 7th edition.

You can either run it as a CLI or as a web server.

## API Usage

```bash
# generate a character
curl "https://www.gothulhu.app/api/generate?name=Peter_Shaw&age=23"
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

Sample response:

```json
{
  "name": "Peter Shaw",
  "age": 23,
  "hp": 11,
  "stats": {
    "str": 71,
    "dex": 47,
    "con": 50,
    "siz": 61,
    "int": 36,
    "wis": 65,
    "cha": 53
  },
  "damage_bonus": {
    "damage": "1d4",
    "stature": "+1"
  }
}
```

## Docker

```bash
# build docker image
docker build -t gothulhu .

# run docker image
docker run -p 8000:8000 gothulhu
```

## Development

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

## Deployment

The application is deployed on Fly.io and accessible at [www.gothulhu.app](https://www.gothulhu.app)
