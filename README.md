# q (Qucumber)

## q is shell script collector
![qcumber](https://pngimage.net/wp-content/uploads/2018/06/larry-the-cucumber-png-7.png "Mr. Qucumber")

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

### call script
`q {NAME}`

```
$ q cx
asia-southeast1_demo-server
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
