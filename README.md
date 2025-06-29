# 🚀 LogEcho

**LogEcho** é uma aplicação escrita em Go que imprime logs em JSON utilizando [Zap](https://github.com/uber-go/zap), mantendo uma rotina periódica com mensagens `"I am alive"` e expondo uma API HTTP para registro de logs customizados no stdout.

---

## 🧰 Tecnologias

- [Go 1.22+](https://golang.org)
- [Zap Logger (Uber)](https://github.com/uber-go/zap)
- Docker (multi-stage)
- HTTP API (net/http)

---

## 🛠️ Como rodar localmente

### Clonar o projeto

```bash
git clone https://github.com/ojaoferreira/go-logecho.git
cd go-logecho
```

### Rodar com Go nativo

```bash
go run ./cmd/main.go
```

---

## 🐳 Como rodar com Docker

### Build da imagem

```bash
docker build -t logecho .
```

### Executar container

```bash
docker run -it --rm --name logecho -p 8080:8080 logecho
```

---

## 🔁 Log de vida

A cada 15 segundos, a aplicação imprime:

```json
{
  "level": "info",
  "timestamp": "2025-06-29 09:50:00",
  "caller": "cmd/main.go:25",
  "msg": "I am alive"
}
```

---

## 📡 API

POST /messages

Imprime uma mensagem customizada no log JSON.

### Exemplo de requisição:

```bash
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from LogEcho!"}'
```

---

## 🔐 Boas práticas aplicadas

- Rodando como usuário não-root no Docker
- Logs estruturados no formato JSON
- Separação de responsabilidades (logger isolado)
- Multi-stage Dockerfile para builds leves

---

## ✨ Autor

Desenvolvido por: João Paulo Ferreira <jferreira.ba@gmail.com>