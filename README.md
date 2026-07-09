# limen go auth 

## run with
```bash
    go run ./cmd/rest-api/
```

## to generate migrations, needs limen cli ( already done )
```bash
    limen generate migrations --driver postgres --dsn "host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable"
```

## to apply migrations, we use pressly goose
```bash
goose up
```