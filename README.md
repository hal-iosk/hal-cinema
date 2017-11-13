# HAL Cinema

## サーバーの起動

```
$ go run main.go
```

## DB設定

```
DBName: halCinema
DBUser: iosk
DBPass: iosk
```

- 初めてDBを触る場合

```
$ make migrate/first
```

- 2回目以降DBをキレイにする場合

```
$ make migrate/clean
```

---

- DBの作成 `make migrate/init`
- テーブルの作成 `make migrate/up`
- DBの削除 `make migrate/drop`
