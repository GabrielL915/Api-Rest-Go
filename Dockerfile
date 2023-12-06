# Use a versão mais recente do Go compatível com a sua aplicação
FROM golang:1.20-alpine AS builder

# Copie o go.mod e o go.sum para o diretório de trabalho no container
COPY go.mod go.sum /go/src/github.com/GabrielL915/Api-Rest-Go/

# Defina o diretório de trabalho no container
WORKDIR /go/src/github.com/GabrielL915/Api-Rest-Go

# Baixe as dependências
RUN go mod download

# Copie o restante dos arquivos do projeto
COPY . /go/src/github.com/GabrielL915/Api-Rest-Go

# Construa o binário
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Api-Rest-Go .

# Use a imagem base alpine
FROM alpine

# Adicione certificados para HTTPS
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Copie o binário compilado do builder para a imagem final
COPY --from=builder /go/src/github.com/GabrielL915/Api-Rest-Go/Api-Rest-Go /usr/local/bin/

# Exponha a porta 8080
EXPOSE 8080

# Defina o comando de execução
ENTRYPOINT ["Api-Rest-Go"]
