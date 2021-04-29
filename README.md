# Microservices in Go

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ammerzon_MC523-golang-rest&metric=alert_status)](https://sonarcloud.io/dashboard?id=ammerzon_MC523-golang-rest) [![Build Status](https://www.travis-ci.com/ammerzon/MC523-golang-rest.svg?branch=main)](https://www.travis-ci.com/ammerzon/MC523-golang-rest)

Second exercise for the subject MC523 SS21 at FH OÖ Campus Hagenberg based on the blog post [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql).

## 📝 Requirements

- `go`
- `docker`
- `docker-compose`
- `skaffold`
- `helm`
- `kubectl`

## 🚀 Get started

```bash
make run-docker
```

## ⚠️ Limitations

* This is a demo application and therefore does ignore common security practices.

## ☸️ Kubernetes Deployment

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

## ♻️ Refactorings

### Folder structure
To blog proposed the following structure:

```shell
.
├── app.go
├── main.go
├── main_test.go
├── model.go
├── go.sum
└── go.mod
```

The current implementation reorganized the folder structure in the following way and added Docker support:

```shell
.
├── Dockerfile
├── cmd
│   ├── main.go
│   └── main_test.go
├── db
│   └── schema.sql
├── docker-compose.yaml
├── go.mod
├── go.sum
└── internal
    ├── config
    ├── models
    └── services
```

### Features

* Added a product search endpoint (`/search/product`)
* Added a price range filter to the `/product` endpoint
* Added a sort option to the `/product` endpoint
