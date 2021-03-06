## Install

### Build from sources

#### Install dependencies

The [vips](https://github.com/libvips/libvips) library needs to be installed for Skrop to build.

On macOS, that can be easily done using `brew`.

```bash
brew install vips
```

On a Linux machine, use the provided [here](https://github.com/danpersa/skrop-build-docker/blob/master/install-vips.sh)

#### Build Skrop

Skrop is 'go get' compatible. If needed, create a Go workspace first:

    mkdir ~/go-workspace
    cd ~/go-workspace
    export GOPATH=$(pwd)

This can be set up in the BASH profile (`~/.bash_profile` or `~/.bashrc`)

    export GOPATH=$HOME/go-workspace

Get the Skrop sources:

    go get https://github.com/jseb/skrop


Skrop uses [modules](https://github.com/golang/go/wiki/Modules), so make sure to have go 1.11+ and just run

```
export GO111MODULE=on
go build ./cmd/skrop/
```

### Using Docker

In order to package skrop for production, you're going to need [Docker](https://docs.docker.com).
To build a Docker image, just run the build script (the arguments are optional):

```
docker/docker-build.sh
```

Now you can start Skrop in a Docker container:

```
docker/docker-run.sh
```

### Using Heroku

Press here to deploy your own demo on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/jseb/skrop/docker)
