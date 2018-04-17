# TODO

* [ ] add mock for getCurrentNamespace
* [ ] mock out k8s call to test
* [ ] Add health check
* [ ] Add liveness check
* [ ] set gin release mode

# Deploying

```
helm upgrade -i --namespace the-counter the-counter deployment/the-counter
```