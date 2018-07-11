![](./envprint.png)

# envprint [![Build Status](https://travis-ci.org/dylanmei/envprint.svg?branch=master)](https://travis-ci.org/dylanmei/envprint)

Template processor for printing configuration files interpolated with environment variables.

_Equivalent to processing templates using [gliderlabs/sigil](https://github.com/gliderlabs/sigil) with POSIX style variable expansion._

## simple example

```
$ envprint -help
Usage of envprint:
  -f string
    	Use template file instead of STDIN
  -o string
    	Write out to file instead of STDOUT
  -version
    	Prints the current version

$ echo 'hello ${WORLD}' | envprint
hello 

$ echo 'hello ${WORLD:-catania}' | envprint
hello catania

$ export WORLD=palermo

$ echo 'hello ${WORLD}' | envprint
hello palermo

$ echo 'hello ${FOOBAR:?this is not a thing}' | envprint
this is not a thing

$ echo $?
1
```

## kubernetes example

Use `envprint` and `initContainers` to prep config files

```
---

apiVersion: v1
kind: Pod
metadata:
  name: hello
spec:
  containers:
  - name: hello
    image: hello-example
    args:
    - "--properties=/etc/hello/hello.properties"
    volumeMounts:
    - name: config-dir
      mountPath: /etc/hello
  initContainers:
  - name: envprint
    image: dylanmei/envprint
    args:
    - "-f=/var/run/hello.properties.template"
    - "-o=/etc/hello/hello.properties"
    env:
    - name: USERNAME
      value: kitty
    - name: PASSWORD
      valueFrom:
        secretKeyRef:
          name: "hello-secrets"
          key: password
    volumeMounts:
    - name: config-dir
      mountPath: /etc/hello
    - name: template-dir
      mountPath: /var/run
  volumes:
  - name: config-dir
    emptyDir: {}
  - name: template-dir
    configMap:
      name: "hello-properties"

---

kind: ConfigMap
apiVersion: v1
metadata:
  name: "hello-properties"
data:
  hello.properties.template:
    home=London
    username=${USERNAME}
    password=${PASSWORD}
```
