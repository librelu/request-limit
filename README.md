# REQIEST-LIMIT

This is a microservice that contains the `rate-limit` feature. For more details about `rate-limit` can be found in [Redis Doc INCR cmd details](https://redis.io/commands/incr)

This service uses `rate-limit checker` as middleware in external `track` API. I use `rate-limit` as middleware to avoid the same user having many requests in a short period, also using middleware can flexibly reuse in different APIs if external APIs have checking needs.

| URL | Method | Description | Response |
|-----|--------|-------------|----------|
|/api/v1/track| GET | track user requests' amount by Client IP. If current IP reach the limit, the API is response `Error`| 200: { "track" : `{{times}}`} <br> 200: `Error` <br> 400: `null` (validation failed)|
|/healthcheck| GET | response if the server is working| 200: `null`|

The default address is `127.0.0.1:8080`.

# How to setup

The server using `make` command for operation. The `make` commands contains:

| Command name | Description |
|-------------|-------------|
| `make up` | start up server by `docker-compose` file|
|`make test`| test server|
|`vegeta-attack`| use vegeta pressure testing|
|`vegeta-report`| getting the attack report|


# Testing in Heroku

The server had deployed to heroku. The hostname is `https://request-limit.herokuapp.com/`

| URL | Method |Description | Response |
|-------------|------------|-------------|--------|
| https://request-limit.herokuapp.com/api/v1/track | GET | track user requests' amount by Client IP. If current IP reach the limit, the API is response `Error`| 200: { "track" : `{{times}}`} <br> 200: `Error` <br> 400: `null` (validation failed)|
| https://request-limit.herokuapp.com/healthcheck | GET | response if the server is working| 200: `null`|

# Vegeta Report

```
Requests      [total, rate, throughput]         12000, 100.01, 0.50
Duration      [total, attack, wait]             2m0s, 2m0s, 269.498ms
Latencies     [min, mean, 50, 90, 95, 99, max]  254.887ms, 282.939ms, 272.542ms, 289.994ms, 354.898ms, 474.514ms, 2.955s
Bytes In      [total, mean]                     60531, 5.04
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           0.50%
Status Codes  [code:count]                      200:60  403:11940  
Error Set:
403 Forbidden
```

This is a Vegeta report testing in heroku. 50% success 50% failed forbidden because of reach the request limit.
