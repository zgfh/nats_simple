
# install nats

```
nats:
  image: nats
  command: --user root --pass dangerous
  ports:
     - "8222:8222"
     - "6222:6222"
     - "4222:4222"
```

https://hub.docker.com/_/nats/


# run
```
# install go
export GOPATH=`pwd`
go run  src/nats_simple/main.go 
```