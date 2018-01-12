# cmd_health_check

Exposes an http endpoint that can be used to check the exit status of an
arbitrary bash command.

Generic version of [sidekiq_health_check](https://github.com/atongen/sidekiq_health_check).

## Example

Suppose you want to expose an http endpoint to check the status of a postgres database.
In that case, you could run something like this:

```
$ export PGPASSWORD="h0tdoggi3!"
$ cmd_health_check -verbose -cmd "/usr/bin/psql -q -c 'select 1' myDb health_check_user" -port 6480
```

Now `http://localhost:6480/ping` will respond with a 200 if the db is healthy, and a 500 otherwise.

## Install

Download the [latest release](https://github.com/atongen/cmd_health_check/releases), extract it,
and put it somewhere on your PATH.

or

```sh
$ go get github.com/atongen/cmd_health_check
```

or

```sh
$ mkdir -p $GOPATH/src/github.com/atongen
$ cd $GOPATH/src/github.com/atongen
$ git clone git@github.com:atongen/cmd_health_check.git
$ cd cmd_health_check
$ go install
$ rehash
```

## Testing

[wip]

```sh
$ cd $GOPATH/src/github.com/atongen/cmd_health_check
$ go test -cover
```

## Releases

```sh
$ mkdir -p $GOPATH/src/github.com/atongen
$ cd $GOPATH/src/github.com/atongen
$ git clone git@github.com:atongen/cmd_health_check.git
$ cd cmd_health_check
$ make release
```

## Command-Line Options

```
$ cmd_health_check -h
Usage of cmd_health_check:
  -cmd string
        Health check bash command
  -port int
        Port to listen on
  -verbose
        Print verbose output
  -version
        Print version information and exit
```
