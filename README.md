# REQIEST-LIMIT

This is a micro server to contains the `rate-limit` feature. For more details about rate limit can be found in [Redis Doc INCR cmd details](https://redis.io/commands/incr)

| URL | Method | Description |
|-----|--------|-------------|
|/api/v1/track| GET | track user request times by IP|
|/healthcheck| GET | check service health status|

The default ip and prot is `127.0.0.1:8000`.

# How to use it

This server using `make` command to control the APP. The make commands contains:
| Command name | Description |
|-----|--------|-------------|
| `make up` | start up server by `docker-compose` file|
|`make test`| test server|