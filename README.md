# test-libjq-go

A test project for [libjq-go](https://github.com/flant/libjq-go)


## testing

```bash
go mod tidy
go run main.go
```

output would be like this:

```
1. "bar"
2. "bar-quux"
2. "baz-baz"
3. "Foo quux"
3. "Foo baz"
4. "bar"
4. "bar"
```


## developing

```
go mod init test-libjq-go
go mod tidy
```
