<p align="center">
  <img src="https://github.com/ropfoo/gothulhu/blob/main/assets/favicon.ico" width="90" title="asmodeus">
</p>
<p align="center">
  <a href="https://gothulhu.fly.dev">gothulhu</a>
</p>

A simple character generator for the Call of Cthulhu 7th edition pen and paper role-playing game.

You can either run it as a CLI or as a web server.

## API Usage

```bash
# generate a random character
curl "https://gothulhu.fly.dev/api/generate"

# generate a character with specific parameters
curl "https://gothulhu.fly.dev/api/generate?name=Peter_Shaw&gender=male&con=50"
```

List of available fields:

| Field  | Required | Type             |
| ------ | -------- | ---------------- |
| gender | no       | "male", "female" |
| name   | no       | string           |
| age    | no       | int              |
| str    | no       | int              |
| dex    | no       | int              |
| con    | no       | int              |
| siz    | no       | int              |
| intl   | no       | int              |
| wis    | no       | int              |
| cha    | no       | int              |

All optional fields are generated randomly if not provided.

Sample response:

```json
{
  "name": "Peter Shaw",
  "age": 23,
  "gender": "male",
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

The application is deployed on Fly.io and accessible at [https://gothulhu.fly.dev](https://gothulhu.fly.dev)
