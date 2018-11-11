
# Kookaburra

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/Lusitaniae/kookaburra/master/LICENSE)

Kookaburra, named after the Aussie bird, is a dummy application built in Go that demonstrates deployments to Kubernetes along with native Prometheus observability and service discovery.

Golang was a natural choice of language due to:

 - performance;
 - self contained binary which facilitates deployments;
 - its great standard library which covers all the basics.

Except for the standard library, I've also used the following:
- httprouter _just because it is performat_ while satisfying the exercise requirements;
	- the standard library net/http would also be adequate here;
- prometheus to export metrics from the application itself without a middleman.

## Get Started

Clone repo

`git clone https://github.com/Lusitaniae/kookaburra.git; cd kookaburra;`

Build and Deploy application to Kubernetes

`make ship`

Deploy Prometheus

`make prometheus`

## Minikube Usage

Open Prometheus dashboard

`minikube service prometheus-service`

Open application

`minikube service kookaburra`

### Issues?

if make fails on the envsubst command

`brew install gettext `

if you cannot connect to your local k8s cluster

`minikube start`

## Details

### Golang

The application is very simple and the Golang standard libs allows you to meet all the requirements within a very short program, under 100 lines for this exercise.  

To that a simple but performant router were added, along with prometheus libs to natively export application metrics without a need for a middleman to read and transform the statistics into the Prometheus format.

### Dockerfile
The applications uses a very basic base image based on Alpine Linux, with only a timezone database and the application files required.

Could be improved by using a multi stage Dockerfile responsible for building the Golang binary as well, facilitating people that are new to the project and might not have a working environment for Go in their laptopts, and making it more reproducible.

We can probably also figure out how to make Golang not rely on the timezone database.

### Kubernetes
Benefiting from the fact that the application is self contained, the k8s stack is quite simple.

One load balancer and one application is all we need to go live. To that we also add *one prometheus server*, to monitor all the Kubernetes components and the application itself, and *one node exporter*, to monitor all the nodes system metrics (e.g. cpu, ram, disk, net).

When using a public cloud managed k8s service we also benefit from a tight integration with the managed services from the provider, in case of the load balancer defined, we automatically get a publicly available load balancer correctly forwarding traffic to the right pods.

### Prometheus
Due to its service discovery capabilities and good integration with k8s APIs, the Prometheus instance pretty much manages itself once some default config files are set in place.

Any changes to the cluster is automatically picked up by Prometheus, and monitoring will reflect the changes in no time. Allowing a very dynamic container based infrastructure to be neatly integrated with a monitoring system.
