# Receipt Rack API

API: `http://localhost:8080/`

Endpoints: Full details listed in `api.yml`

* POST - `/receipts/process`
* GET - `/receipts/:id/points`

setup `.env` file like below

`cp .env.example .env`

```env
PORT=8080
```

Install dependencies & start server

```bash
go mod tidy
go run main.go
```
