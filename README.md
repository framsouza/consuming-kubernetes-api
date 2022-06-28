# Listing k8s resources

This is a simple code that will list the Kubernetes resources using client-go.
The resources available at the moment is, `pods`, `deployments`, `nodes`, `statefulsets`.

### Usage
You must specify the resource you want to check by adding the `-resource` parameter, like this:
```
./list-k8s-resources -resource=pods
```

You will get an output similar to this:
```
nginx
redis-574cfc6865-vjkps
coredns-558bd4d5db-ggrmh
coredns-558bd4d5db-qjbn6
etcd-kind-control-plane
kindnet-wrbtb
kube-apiserver-kind-control-plane
kube-controller-manager-kind-control-plane
kube-proxy-stplz
kube-scheduler-kind-control-plane
local-path-provisioner-547f784dff-pqm4w
```
