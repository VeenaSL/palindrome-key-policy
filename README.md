How the policy works

The policy validates the labels of the kubernetes resources before creation.

The policy rejects resources that have one or more labels with palindrome key.

The policy allows user to configure if they want to allow/deny palindrome key in labels. 

The polisy seetings is as below

```json
{
  "deny_palindrome_key": true
}
```

This repository contains a working policy written in Go.

The policy looks at the `labels` of a Kubernetes resource and rejects the request
if the label key is palindrome.

The policy should allow the creation of this Pod:
```
apiVersion: v1
kind: Pod
metadata:
  name: hello-world
  labels:
    env: production
spec:
  containers:
  - name: nginx
    image: nginx
```

And policy should reject the creation of this Pod:
```
apiVersion: v1
kind: Pod
metadata:
  name: hello-world
  labels:
    env: production
    level: debug
spec:
  containers:
  - name: nginx
    image: nginx
```
