# Proxy Server

Simple proxy server.

## Setup

Place project directory into your `GOPATH` and run:

```
$ cd simple-proxy
$ go build
```

## Usage

Run:

```
$ ./simple-proxy
```

You will find out PID after application will start. Then you can get traffic usage stats sending SIGUSR2 signal:

```
$ kill -31 <PID>
```
