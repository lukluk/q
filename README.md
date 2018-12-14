# q (quali)

## q is shell script repository
![quali](https://gallery.kissclipart.com/20180829/qxe/kissclipart-cauldron-halloween-clipart-cauldron-halloween-clip-19b25e2dd7bbaa8a.jpg "")

## how to install

```
$ get clone https://github.com/lukluk/q.git
$ cd q && cp bin/* /usr/local/bin
```

## how to use

### set repo

`q repo {repo-ssh-uri}`

example:

```
//read-write access
$ q repo git@github.com:lukluk/repo-demo.git
//read-only access
$ q repo https://github.com/lukluk/repo-demo.git
```

### create new script
`q new {SCRIPT}`

```
$ q new cx
<vim> kubectl config current-context
````

more examples:

```
$ q new podexec
<vim> kubectl exec  -it $(kubectl get pods | awk '{print $1}' | grep $1 -m1) -- bash
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

### list of scripts
```
$ q
podexec
podname
cx
```

### update collection

```
$ q pull
```

### publish your script (share to team)

```
$ q push
```
