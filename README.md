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
$ q new cx
<vim> kubectl config current-context
````

more examples:

```
$ q new podname
<vim> kubectl get pods | awk '{print $1}' | grep $1 -m1 
$ q new podexec
<vim> kubectl exec  -it $(q podname $1) -- bash
```

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
