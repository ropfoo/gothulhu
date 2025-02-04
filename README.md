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
    "str": [71, 33.5, 14.2],
    "dex": [47, 21.5, 9.4],
    "con": [50, 23.5, 10.2],
    "siz": [61, 28.5, 12.2],
    "int": [36, 16.5, 7.2],
    "wis": [65, 30.5, 13.2],
    "cha": [53, 24.5, 10.2]
  },
  "damage_bonus": {
    "damage": "1d4",
    "build": "+1"
  },
  "skills": [
    {
      "name": "Perception",
      "score": [50, 25, 10]
    }
  ]
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
