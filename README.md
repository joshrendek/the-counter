# Installation

Cluster needs to have RBAC enabled. Default values file assumes your using traefik as an ingress. 
Modify depending on cluster ingress.

```
git clone git@github.com:joshrendek/the-counter.git
cd the-counter
helm upgrade -i --namespace the-counter the-counter deployment/the-counter
```

# Test after installation

There is a helm test file provided that can be run like this (after installing):

```
kubectl delete po -n the-counter the-counter-credentials-test # if it already exists
helm test the-counter
```

# TODO
* [x] mock out k8s call to test
* [x] add helm test
* [x] Add health check
* [x] Add liveness check
* [x] set gin release mode


