# q (Qucumber)

## q is shell script collector
![qcumber](https://sodapopstop.com/wp-content/uploads/2016/09/mr-q-cumber.png "Mr. Qucumber")

## how to install

```
$ get clone https://github.com/lukluk/q.git
$ cd q && cp bin/* /usr/local/bin
```

## how to use
### create new script
`q new {SCRIPT}`

```
$ q new kubectl config current-context
> cx {PRESS ENTER}
> cx created!
````

```
$ q new "kubectl exec  -it $(kubectl get pods | awk '{print #1}' | grep %1 -m1) -- bash"
> podexec {PRESS ENTER}
> podexec created!
```

`#1 normalization from $1`

`%1 represent argument (param) for script`

### call script
`q {NAME}`

```
$ q cx
asia-southeast1_demo-server
```

```
$ q podexec
> not enaugh arguments , args(1)
$ q podexec util
#<util-xsd-23xx>$
```

## WORKSPACE , coming soon 

### create team

```
$ q login
> youremail@org.com
> {enter-password}
$ q new team
> {enter-team-name}
> created!
> secret-key: blablabla 
```
### update collection

```
$ q join team
> {enter-team-name}
> {enter-secret-key}
$ q update
> updated!
```

### publish your script (share to team)

```
$ q publish {script-name}
> success!
```
