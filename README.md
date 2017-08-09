# sidekiq_health_check

Exposes an http endpoint that can be used to check for the number of sidekiq processes
running on a server.

## Install

Download the [latest release](https://github.com/atongen/sidekiq_health_check/releases), extract it,
and put it somewhere on your PATH.

or

```sh
$ go get github.com/atongen/sidekiq_health_check
```

or

```sh
$ mkdir -p $GOPATH/src/github.com/atongen
$ cd $GOPATH/src/github.com/atongen
$ git clone git@github.com:atongen/sidekiq_health_check.git
$ cd sidekiq_health_check
$ go install
$ rehash
```

## Testing

[wip]

```sh
$ cd $GOPATH/src/github.com/atongen/sidekiq_health_check
$ go test -cover
```

## Releases

```sh
$ mkdir -p $GOPATH/src/github.com/atongen
$ cd $GOPATH/src/github.com/atongen
$ git clone git@github.com:atongen/sidekiq_health_check.git
$ cd sidekiq_health_check
$ make release
```

## Command-Line Options

```
λ sidekiq_health_check -h
Usage of sidekiq_health_check:
  -num int
        Number of sidekiq processes needed to be health (default 8)
  -port int
        Port to listen on
  -v    Print version information and exit
```
