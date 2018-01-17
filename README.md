hello-twirp
===========

An example program for [twitchtv/twirp: A simple RPC framework with protobuf service definitions](https://github.com/twitchtv/twirp).

## Install

Run the following command.

```
go get -u github.com/hnakamur/hello-twirp
cd $GOPATH/src/github.com/hnakamur/hello-twirp
```

## Generate/update hello.{pb,twirp}.go file from hello.proto

### Install tools

Follow instructions at [Install and Update Twirp · twitchtv/twirp Wiki](https://github.com/twitchtv/twirp/wiki/Install-and-Update-Twirp).

### Run tool to generate hello.{pb,twirp}.go file from hello.proto

Run the following command.

```
go generate
```

## Running example

Run the server.

```
go run cmd/server/main.go -addr 127.0.0.1:9999
```

Run the client for a normal case.

```
$ go run cmd/client/main.go -server http://127.0.0.1:9999
Hello World
```

Run the client for an error case.

```
$ go run cmd/client/main.go -server http://127.0.0.1:9999 -subject poo
2018/01/17 09:15:35 twirp error invalid_argument: Subject must not be poo
exit status 1
```

Run curl for a normal case.

```
$ curl -sS -X POST -H 'Content-Type: application/json' -d '{"subject":"World"}' \
    http://127.0.0.1:9999/twirp/com.github.hnakamur.hello_twirp.HelloWorld/Hello
{"text":"Hello World"}
```

Run curl for an error case.

```
$ curl -sS -X POST -H 'Content-Type: application/json' -d '{"subject":"poo"}' \
    http://127.0.0.1:9999/twirp/com.github.hnakamur.hello_twirp.HelloWorld/Hello
{"code":"invalid_argument","msg":"Subject must not be poo","meta":{"argument":"Subject"}}
```
