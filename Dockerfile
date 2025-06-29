# Etapa 1: build
FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o logecho ./cmd/main.go

# Etapa 2: imagem final
FROM alpine:latest

# Cria um usuário sem privilégios
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Instala certificados HTTPS
RUN apk --no-cache add ca-certificates

# Define diretório de trabalho
WORKDIR /app

# Copia binário
COPY --from=builder /app/logecho .

# Permissões (boa prática)
RUN chown -R appuser:appgroup /app

# Usa o usuário sem privilégios
USER appuser

EXPOSE 8080

CMD ["./logecho"]