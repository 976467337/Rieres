# smart-kiki-api

Backend do Smart Kiki (app de personal trainer). Go + Gin + GORM + Postgres, com JWT para autenticação.

## Rodando localmente

```
cp .env.example .env
make docker-up      # sobe Postgres local (smartkiki_db)
make migrate-up      # aplica as migrations
make dev             # roda a API em :8080
```

## Endpoints

- `POST /api/v1/auth/register` — cria conta (`name`, `email`, `password`, `role`: `client`|`trainer`)
- `POST /api/v1/auth/login` — autentica e retorna token JWT
- `GET /api/v1/users/me` — dados do usuário logado (requer `Authorization: Bearer <token>`)
- `GET /health` — healthcheck
