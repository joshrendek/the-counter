# TODO

* [ ] add mock for getCurrentNamespace
* [ ] mock out k8s call to test
* [x] add helm test
* [x] Add health check
* [x] Add liveness check
* [x] set gin release mode

# Deploying

```
helm upgrade -i --namespace the-counter the-counter deployment/the-counter
```