# Listing k8s resources

This is a simple code that will list the Kubernetes resources using client-go.
The resources available at the moment is, `pods`, `deployments`, `nodes`, `statefulsets`.

### Usage
You must specify the resource you want to check by adding the `-resource` parameter, like this:
```
./list-k8s-resources -resource=nodes
```
