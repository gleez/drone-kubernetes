# Drone Kubernetes 

Drone plugin to create/update Kubernetes resources.

It uses the latest k8s go api, so it is intened to use on Kubernetes 1.12+. I can't guarantee it will work for previous versions.

You can directly pull the image from [gleez/drone-kubernetes](https://hub.docker.com/r/gleez/drone-kubernetes/)

####
Fork [Sh4d1/drone-kubernetes](https://github.com/Sh4d1/drone-kubernetes)

## Supported resources
Currently, this plugin supports:
* apps/v1
  * DaemonSet
  * Deployment
  * ReplicaSet
  * StatefulSet
* apps/v1beta1
  * Deployment
  * StatefulSet
* apps/v1beta2
  * DaemonSet
  * Deployment
  * ReplicaSet
  * StatefulSet
* v1
  * ConfigMap 
  * PersistentVolume 
  * PersistentVolumeClaim 
  * Pod 
  * ReplicationController 
  * Service 
* extensions/v1beta1
  * DaemonSet
  * Deployment
  * Ingress
  * ReplicaSet

## Inspiration 

It is inspired by [vallard](https://github.com/vallard) and his plugin [drone-kube](https://github.com/vallard/drone-kube).


## Usage

Here is how you can use this plugin:
```
pipeline:
- name: deploy
  image: gleez/drone-kubernetes
  settings:
    kubernetes_template: deployment.yml
    kubernetes_namespace: default
    kubernetes_incluster: false
    kubernetes_server:
    kubernetes_cert:
    kubernetes_token:
```

## Secrets or Incluster Auth

If you build and deploy in the same kubernetes cluster, no need to  define secrets. Change ```kubernetes_incluster true```


You need to define these secrets before.
```
$ drone secret add --image=gleez/drone-kubernetes -repository <your-repo> -name KUBERNETES_SERVER -value <your API server>
```
```
$ drone secret add --image=gleez/drone-kubernetes -repository <your repo> -name KUBERNETES_CERT -value <your base64 encoded cert>
```
```
$ drone secret add --image=gleez/drone-kubernetes -repository <your repo> -name KUBERNETES_TOKEN -value <your token>
```

### How to get values of `KUBERNETES_CERT` and `KUBERNETES_TOKEN`

List secrets of `default` namespace

```
$ kubectl get -n <namespace of secret> default secret
```

Show the `ca.crt` and `token` from secret

```
$ kubectl get secret -n <namespace of secret> <name of your drone secret> -o yaml | egrep 'ca.crt:|token:'
```

You can copy/paste the encoded certificate to the `KUBERNETES_CERT` value.
For the `KUBERNETES_TOKEN`, you need to decode it:

* `echo "<encoded token>" | base64 -d`
* `kubectl describe secret -n <your namespace> <drone secret name> | grep 'token:'`


TODO
