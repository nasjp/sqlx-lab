# sqlx lab

for learning sqlx

```sh
make run
```

## sqlx

- `*sqlx.DB.Get` は `sql.ErrNoRows` を返却する

```go
// これはnilではない
// だからnil判定ではレコードが取得できたかわからない
u := &User{}

err := db.Get(u, `SELECT * FROM users limit 1`)
// ここ sql.ErrNoRows を返却してくれるので
// nil を返却することができ、呼び出しもとで
// nil 判定をすることでレコードが取得できなかったことを
// 知ることができる
if err == sql.ErrNoRows {
  return nil, nil
}

if err != nil {
    return nil, err
}

// err == sql.ErrNoRows がなければ、
// u が nil ではないレコードを返却することになってしまう
return u, nil
```
- `*sqlx.DB.Select` は `sql.ErrNoRows` を返却しない

```go
// これはnil
var us Users // []*User

// ここ sql.ErrNoRows を返却しない
if err := db.Select(&us, `SELECT * FROM users`); err != nil {
  return nil, err
}

// errがnilなら空の配列を返却したい
if len(us) == 0 {
  return Users{}
}

return us, nil
```
