# redis-pubsub-microservices

Projeto simples em Go para simular uma comunicação assíncrona entre microserviços usando Redis Pub/Sub.

---

# Objetivo

Este projeto demonstra um cenário básico de microserviços comunicando-se via Redis Pub/Sub:

- **publisher:** publica palavras aleatórias num canal Redis a cada 5 segundos  
- **subscriber:** escuta o canal e exibe as mensagens recebidas no console  

O Redis funciona como um broker de mensagens, intermediando a comunicação entre os serviços.

## Estrutura

- `publisher/` — Serviço que publica mensagens no canal Redis.
- `subscriber/` — Serviço que escuta mensagens do canal Redis.
- `logger-service/` Outro serviço que também escuta mensagens do canal Redis.
- `redis/` — Serviço Redis via Docker.
- `docker-compose.yml` — Orquestra os containers.
- `go.mod` e `go.sum` — Módulos e dependências Go separados para cada serviço.

---

## Requisitos

- Docker (versão recente)
- Docker Compose
- Go (para desenvolvimento local e builds manuais)

---

## Como gerar os builds separados

Cada serviço tem seu próprio módulo Go e Dockerfile.

Na raiz do projeto, rode os comandos:

```bash
docker build -t publisher -f publisher/Dockerfile ./publisher
docker build -t subscriber -f subscriber/Dockerfile ./subscriber
docker build -t logger-service -f logger-service/Dockerfile ./logger-service
```

Como subir o projeto

```bash
docker compose up
```

O docker-compose.yml vai subir:
- Redis
- Publisher
- Subscriber

Funcionamento
- O publisher publica mensagens no canal Redis a cada X segundos.
- O subscriber escuta as mensagens e imprime no terminal.

Pode-se adicionar novos subscribers (ex: logger-service) para escutar as mesmas mensagens.

Exemplo de saída no terminal

```bash
redis       | 1:M 21 May 2025 20:22:58.281 * Ready to accept connections tcp
publisher   | 2025/05/21 20:22:58 🚀 Publisher started...
subscriber  | 2025/05/21 20:22:58 📥 Subscriber listening on channel: random-words
logger      | 2025/05/21 20:22:58 📝 Logger-service started and listening on channel: random-words
publisher   | 2025/05/21 20:23:03 📢 [20:23:03] Published: docker
subscriber  | 2025/05/21 20:23:03 📝 Received: docker
logger      | 2025/05/21 20:23:03 🗂️ Logger received: docker


```

---