# Oxylabs Homework

**Author:** Aleksej Zaicev

## Specifications

### Server

- Accepts QUIC connections on 2 ports
    - Publisher port
    - Subscriber port
- The server notifies publishers if a subscriber has connected
- If no subscribers are connected, the server must inform the publishers
- The server sends any messages received from publishers to all connected subscribers

### Publisher

- The application must connect via QUIC to the server
- Upon establishing a connection, waits for a message from the server that there are connected subscribers
- Upon receiving a message that subscribers have connected - starts sending string messages to the server
  every second (i.e. 1 message/second)
- Upon receiving a message that all subscribers have disconnected - waits for a message that at least 
  1 subscriber has connected

### Subscriber

- The application must connect via QUIC to the server
- Upon connection must receive messages from publishers

## Project structure

The following project consists of 3 applications (server, publisher, and subscriber) that separated into
individual binary files. The overall project follows clean architecture guidelines with dependency 
injection (DI) provided by [wire]. The `cmd` directory contains the entry points to `main` package of individual 
application. The `internal` directory contains the actual components logic implementation:
- `./internal/adapters` - location where QUIC server and client logic lives
- `./internal/drivers` - contains helper functions for DI

[wire]: https://github.com/google/wire

## Configuration

Applications can be configured with a simple configuration file written in YAML and is located under root path
`./config.yaml`.

Additional X509 certificate and key pair is provided for secure TLS communication between client and server.

## Run

In order to run the project make sure your environment has Golang 1.20 installed. To build the project run:

```sh
make build
```

This will build the binary files for server, publisher, and subscriber under `./dist` directory. From
there you can now run the project:

```sh
./dist/server -f config.yaml
...
./dist/publisher -f config.yaml
...
./dist/subscriber -f config.yaml
```

If for some reason you would like to regenerate [wire] DI, run the following:

```sh
make wire
```

> NOTE: wire must be installed on your machine. If you don't have it installed run the following: 
> ```sh
> go install github.com/google/wire/cmd/wire@latest
> ```

## Additional information

The following project is not a work complete piece of software and might contain bugs and partially
available functionality. 
