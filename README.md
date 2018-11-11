# Kookaburra

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/Lusitaniae/kookaburra/master/LICENSE)

Kookaburra, named after the Aussie bird, is a dummy application built in Go that demonstrates deployments to Kubernetes along with native Prometheus observability and service discovery.

Golang was a natural choice of language due to:

 - performance
 - self contained binary which facilitates deployments
 - its great standard library which covers all the basics.

Except for the standard library, I've also used the following packages:
- httprouter _just because it is performat_ while satisfying the exercise requirements
	- the standard library net/http would also be adequate here .
- prometheus to export metrics from the application itself without middleware.

## Get Started

Clone repo
`git clone https://github.com/Lusitaniae/kookaburra.git; cd kookaburra;`

Build and Deploy application to Kubernetes
`make ship`

Deploy Prometheus
`make prometheus`

## Minikube Usage

Open Prometheys dashboard

`minikube service prometheus-service`

Open application

`minikube service kookaburra`

### Issues?
`brew install gettext # if make fails on the envsubst command`

`minikube start # if you cannot connect to your local k8s cluster`
