# HAL Cinema


## サーバーの起動

```
$ go build main.go
$ ./main
```

## DB設定初心者向け
``` bash
make create
```
## DB設定上級者向け

```
DBName: halCinema
DBUser: iosk
DBPass: iosk
```

- DB作成

```
$ make migrate/init
```

- table作成

```
$ make migrate/up
```

- DB削除

```
$ make migrate/drop
```

- seed

```
$ go run cmd/seed.go
```
