# Installation

Cluster needs to have RBAC enabled. Default values file assumes your using traefik as an ingress. 
Modify depending on cluster ingress. Tested on Kubernetes Version `v1.10.0`.

```
git clone git@github.com:joshrendek/the-counter.git
cd the-counter
helm upgrade -i --namespace the-counter the-counter deployment/the-counter
```

Demo: https://the-counter.svc.bluescripts.net

# Test after installation

There is a helm test file provided that can be run like this (after installing):

```
kubectl delete po -n the-counter the-counter-credentials-test # if it already exists
helm test the-counter
```

# Local development

Assuming you have a valid `~/.kube/config` (either a real cluster or minikube):

```
go run main.go
```

You can then see the number of pods running in the default namespace:

```
curl localhost:8080
```

## Running tests

Depending on which tests you want to run you will need to run them inside the provided docker container:

``` 
make docker/test
```

# Publishing

```
make docker/build docker/push
```

# TODO
* [x] mock out k8s call to test
* [x] add helm test
* [x] Add health check
* [x] Add liveness check
* [x] set gin release mode


