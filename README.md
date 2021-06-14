# Kyber

Kyber is a request repeater for HTTP-based APIs, made in Go.

# Table of content

- [Installing](#installing)
    * [Docker](#docker)
- [Getting started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Installing

> :warning: Kyber is right now in a very early version, which means that the 'minor' will be changed when some broken changes are introduced into the application. So, use it carefully.

### Docker

The application is also available through [Docker](https://hub.docker.com/r/mrandycodes/kyber).

```bash
docker run -it --rm -p 8080:8080 mrandycodes/kyber:latest -host 0.0.0.0
```
You need to pass `-host 0.0.0.0` in order to allow the container to capture the request from the host network.

## Getting Started

Just start adding as many hosts you want with the `add route endpoint`, the cloned request will be sent to each one of them.

```bash
curl --location --request POST 'http://localhost:8080/routes' \
--header 'Content-Type: application/json' \
--data-raw '{
  "route": "http://localhost:8080"
}'
```

Then you can list the available hosts with the `list routes endpoint` as follow:

```bash
curl --location --request GET 'http://localhost:8080/routes' \
--header 'Content-Type: application/json'
```

If you want to remove one route, just use the `remove route endpoint` specifying the route that you want to delete:

```bash
curl --location --request DELETE 'http://localhost:8080/routes' \
--header 'Content-Type: application/json' \
--data-raw '{
  "route": "http://localhost:8080"
}'
```

When you are ready, use the `repeater endpoint` like you will calling to one of your own api endpoints like this:

```bash
curl --location --request POST 'http://localhost:8080/api/:your-endpoint-here' \
--header 'Content-Type: application/json' \
--data-raw '{
  "foo": "data"
}'
```

Kyber will use the endpoint that you specify, the same payload, the same request method and even the same headers to make a request to each route that you'd already added to the collection of routes. 

## Contributing

Right now we are working in new releases to provide a stable release. Once we finished, we will develop a contributions guide to anyone that are interested to help.

## License

MIT License, see [LICENSE](https://github.com/mrandycodes/kyber/blob/main/LICENSE)