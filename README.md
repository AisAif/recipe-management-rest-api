# recipe-management-rest-api
__Recipe Management__ is a digital tool designed to help users organize, store, manage and expose their recipes.

## Required Tools
- [__Go__](https://go.dev/doc/install) (Latest Version)
- __MySQL Server__ (Latest Version)
- __Docker__ (For Production)

## Setup Project
### Clone Repo
```
git clone https://github.com/AisAif/recipe-management-rest-api.git
```
### Add Environment App
Copy the file named .env.example to .env then customize it.

## Run In Development
### Install Dependencies
```
go get
```
### Run Server
```
go run main.go
```
## Run In Production
### Create Docker Image
```
docker build . -t rm-app
```
### Run Server On Container
```
docker run -dit --name rm-server -p {host_port}:{container_port} rm-app
```
Make sure {container_port} is the same as the environment port.
## API Documentation
[Click Here](https://aisaif.github.io/recipe-management-rest-api/)