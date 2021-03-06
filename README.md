# Microservices in Go

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ammerzon_MC523-golang-rest&metric=alert_status)](https://sonarcloud.io/dashboard?id=ammerzon_MC523-golang-rest) [![Build Status](https://www.travis-ci.com/ammerzon/MC523-golang-rest.svg?branch=main)](https://www.travis-ci.com/ammerzon/MC523-golang-rest)

Second exercise for the subject MC523 SS21 at FH OΓ Campus Hagenberg based on the blog post [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql).

## π Requirements

- `go`
- `docker`
- `docker-compose`
- `skaffold`
- `helm`
- `kubectl`

## π Get started

```bash
make run-docker
```

## β οΈ Limitations

* This is a demo application and therefore does ignore common security practices.

## βΈοΈ Kubernetes Deployment

Your `kubectl` must be configured correctly. For local deployment [minikube](https://minikube.sigs.k8s.io/docs/) or [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/) can be used.

1. Deploy the application
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
skaffold run
```

2. Forward to a local port
```bash
kubectl port-forward deployment/backend 8010:8010 -n golang-rest
```

3. Create the schema and insert test data

## β»οΈ Refactorings

### Folder structure
To blog proposed the following structure:

```shell
.
βββ app.go
βββ main.go
βββ main_test.go
βββ model.go
βββ go.sum
βββ go.mod
```

The current implementation reorganized the folder structure in the following way and added Docker support:

```shell
.
βββ Dockerfile
βββ cmd
β   βββ main.go
β   βββ main_test.go
βββ db
β   βββ schema.sql
βββ docker-compose.yaml
βββ go.mod
βββ go.sum
βββ internal
    βββ config
    βββ models
    βββ services
```

### Features

* Added a product search endpoint (`/search/product`)
* Added a price range filter to the `/product` endpoint
* Added a sort option to the `/product` endpoint
