Grasp - simple website analytics
================================
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jelmer/grasp/master/LICENSE)

Grasp is a friendly fork of [Fathom Lite](https://github.com/usefathom/fathom). Upstream is only providing fixes for bugs, but not adding new features - Grasp is.

![Screenshot of the Fathom dashboard](https://github.com/jelmer/grasp/raw/master/assets/src/img/fathom.jpg?v=7)

## Installation

### Production

You can install Fathom on your server by following [our simple instructions](docs/Installation%20instructions.md).

### Development

For getting a development version of Grasp up & running, go through the following steps.

1. Ensure you have [Go](https://golang.org/doc/install#install) and [NPM](https://www.npmjs.com) installed
1. Download the code: `git clone https://github.com/jelmer/grasp.git $GOPATH/src/github.com/jelmer/grasp`
1. Compile the project into an executable: `make build`
1. (Optional) Set [custom configuration values](docs/Configuration.md)
1. (Required) Register a user account: `./grasp user add --email=<email> --password=<password>`
1. Start the webserver: `./grasp server` and then visit **http://localhost:8080** to access your analytics dashboard

## Docker

### Building

Ensure you have Docker installed and run `docker build -t grasp .`.
Run the container with `docker run -d -p 8080:8080 grasp`.

### Running

To run [our pre-built Docker image](https://ghcr.io/r/jelmer/grasp), run `docker run -d -p 8080:8080 ghcr.io/jelmer/grasp:latest`

## Tracking snippet

To start tracking, create a site in your Grasp dashboard and copy the tracking snippet to the website(s) you want to track.

### Content Security Policy

If you use a [Content Security Policy (CSP)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) to specify security policies for your website, Grasp requires the following CSP directives (replace `yourgrasp.com` with the URL to your Grasp instance):

```
script-src: yourgrasp.com;
img-src: yourgrasp.com;
```

## Copyright and license

MIT licensed. Fathom and Fathom logo are trademarks of Fathom Analytics.
