# ProjHopExp Service

This is the ProjHopExp service

Generated with

```
micro new github.com/madisonsmartin/projHopExp --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.projHopExp
- Type: srv
- Alias: projHopExp

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./projHopExp-srv
```

Build a docker image
```
make docker
```