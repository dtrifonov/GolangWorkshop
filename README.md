# golang_workshop
```
1. https://github.com/rumyantseva/go-sofia
2. https://github.com/rumyantseva/sofia
```

# golang version 1.11.1

To use modules and to create vendor folder

```
go version

export GO111MODULE=on
go mod init
go mod tidy
go mod vendor
```

# work with kubectl

```
kubectl delete deployment go-sofia
kubectl delete deployment,service go-sofia
```
