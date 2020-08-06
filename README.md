# REQIEST-LIMIT

This is a micro server to contains the `rate-limit` feature. For more details about rate limit can be found in [Redis Doc INCR cmd details](https://redis.io/commands/incr)

| URL | Method | Description |
|-----|--------|-------------|
|/api/v1/track| GET | track user request times by Client IP|
|/healthcheck| GET | check service health status|

The default ip and prot is `127.0.0.1:8000`.

# How to use it

This server using `make` command to control the APP. The make commands contains:

| Command name | Description |
|-------------|-------------|
| `make up` | start up server by `docker-compose` file|
|`make test`| test server|
|`vegeta-attack`| use vegeta pressure testing|
|`vegeta-report`| getting the attack report|


# Testing in Heroku

The server had deploied to heroku. The hostname is `https://request-limit.herokuapp.com/`

| URL | Method |Description |
|-------------|------------|-------------|
| https://request-limit.herokuapp.com/api/v1/track | GET | track user request times by Client IP|
| https://request-limit.herokuapp.com/healthcheck | GET | check service health status|

# Vegeta Report

```
Requests      [total, rate, throughput]         12000, 100.01, 50.87
Duration      [total, attack, wait]             2m0s, 2m0s, 324.8ms
Latencies     [min, mean, 50, 90, 95, 99, max]  255.751ms, 286.243ms, 272.834ms, 290.045ms, 317.939ms, 566.455ms, 3.915s
Bytes In      [total, mean]                     55068, 4.59
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           51.00%
Status Codes  [code:count]                      200:6120  403:5880  
Error Set:
403 Forbidden
```

51% success 49% faied forbidden because of reach the request limit.