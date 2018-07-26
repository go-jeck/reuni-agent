# Reuni Agent

Reuni system is a centralized service re-configuration manager for microservices architecture.

This repository is intended for the RESTful API server for the reuni system. This API provide:
- Configration Fetching
- Configuration synchronization 
- Environment variable setter
- Service starter
- Service stopper
- Service reloader

## Development Environment

### MacOS 
Prerequisite:
- Homebrew
- Go 1.10.3 - brew install go
- PostgreSQl 10.4 - brew install postgresql


To build application: make build

To test application: make test

To run application: make run

To install application: make install

