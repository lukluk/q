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
### create script with humanize param

```
$ q new hello
<vim> echo <your-name> <your-email>
$ q hello
q hello your-name your-email
$ q hello lukluk luklukaha@gmail.com
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

## WORKSPACE

### list of scripts
```
$ q
podexec
podname
cx
```

### set workspace repo

`q repo {repo-ssh-uri}`

example:

```
//read-write access
$ q repo git@github.com:lukluk/repo-demo.git
//read-only access
$ q repo https://github.com/lukluk/repo-demo.git
```
### update collection

```
$ q pull
```

### publish your script (share to team)

```
$ q push
```
