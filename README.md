# Hackathon User Microservice

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=coverage)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=bugs)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=FIAP-11soat-grupo-21_hackathon-auth-microservice&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=FIAP-11soat-grupo-21_hackathon-auth-microservice)

Microserviço responsável pelo gerenciamento de usuários no contexto do hackathon.

## Visão geral

Este projeto foi desenvolvido em Go e segue uma organização por camadas (inspirada em arquitetura hexagonal), separando domínio, casos de uso, adapters e infraestrutura.

Estrutura principal:

- `cmd/http/main.go`: ponto de entrada da aplicação HTTP.
- `internal/core`: domínio, DTOs, fábricas e casos de uso.
- `internal/adapter`: adapters de entrada (API HTTP) e saída (banco, Cognito).
- `internal/common`: configurações e componentes de infraestrutura compartilhados.
- `infra/`: infraestrutura como código (Terraform).

## Stack

- Go (`go.mod`)
- HTTP server em Go
- Banco de dados (adapter em `internal/adapter/driven/database`)
- AWS Cognito (adapter em `internal/adapter/driven/aws/cognito`)
- SonarCloud para análise de qualidade

## Pré-requisitos

- Go instalado
- Docker e Docker Compose (opcional para execução em containers)

## Variáveis de ambiente

Para rodar o serviço localmente, crie um arquivo `.env` na raiz do projeto com as variáveis abaixo.

| Variável | Obrigatória | Exemplo | Descrição |
|---|---|---|---|
| `APP_PORT` | Sim | `8080` | Porta HTTP da aplicação. |
| `DB_HOST` | Sim | `localhost` | Host do banco de dados. |
| `DB_PORT` | Sim | `5432` | Porta do banco de dados. |
| `DB_NAME` | Sim | `user_microservice` | Nome do banco. |
| `DB_USER` | Sim | `postgres` | Usuário do banco. |
| `DB_PASSWORD` | Sim | `postgres` | Senha do banco. |
| `AWS_REGION` | Não* | `us-east-1` | Região AWS usada em integrações. |
| `COGNITO_USER_POOL_ID` | Não* | `us-east-1_xxxxxxxx` | User Pool do Cognito. |
| `COGNITO_CLIENT_ID` | Não* | `xxxxxxxxxxxxxxxxxxxxxxxxxx` | Client ID do app no Cognito. |

* Obrigatória apenas quando as integrações AWS/Cognito estiverem habilitadas no ambiente local.

Exemplo de `.env`:

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_NAME=user_microservice
DB_USER=postgres
DB_PASSWORD=postgres

AWS_REGION=us-east-1
COGNITO_USER_POOL_ID=us-east-1_xxxxxxxx
COGNITO_CLIENT_ID=xxxxxxxxxxxxxxxxxxxxxxxxxx
```

> Dica: não versione o `.env`; mantenha um `.env.example` atualizado com valores de exemplo.

## Como executar

Os comandos abaixo estão definidos no `Makefile`.

```bash
make run
```

Para modo de desenvolvimento com flags de debug:

```bash
make dev
```

Para executar com Docker:

```bash
make docker
```

## Build e testes

Build:

```bash
make build
```

Testes:

```bash
make test
```

Cobertura de testes (resumo):

```bash
make coverage
```

Cobertura em HTML:

```bash
make coverage-html
```

## Pipeline do projeto

Fluxo recomendado para desenvolvimento e validação local:

1. Configurar variáveis de ambiente (`.env`).
2. Compilar o projeto para validar build.
3. Executar testes automatizados.
4. Verificar cobertura de testes.
5. Subir a aplicação localmente.
6. Acompanhar indicadores de qualidade no SonarCloud (badges no topo do README).

Exemplo de sequência:

```bash
make build
make test
make coverage
make run
```

Para execução via container:

```bash
make docker
```

## Funções disponíveis

As funções abaixo estão disponíveis no `Makefile`:

- `make run`: executa a aplicação (`go run cmd/http/main.go`).
- `make dev`: executa em modo debug (`-gcflags=all="-N -l"`).
- `make docker`: sobe os serviços com Docker Compose (`up --build -d`).
- `make build`: gera build da aplicação.
- `make test`: executa todos os testes (`go test -v ./...`).
- `make coverage`: calcula cobertura com `coverage.out`.
- `make coverage-html`: gera relatório HTML de cobertura.
- `sonar-scanner`: executa análise SonarCloud manual (quando configurado).

### Pipeline de qualidade (SonarCloud)

Os badges no topo refletem os principais indicadores de qualidade:
- Quality Gate
- Coverage
- Bugs
- Vulnerabilities
- Code Smells
- Technical Debt

## Qualidade de código (SonarCloud)

Configuração atual em `sonar-project.properties`:

- `sonar.organization=fiap-11soat-grupo-21`
- `sonar.projectKey=FIAP-11soat-grupo-21_hackathon-auth-microservice`

> Observação: o `projectKey` atual referencia `hackathon-auth-microservice`. Se este repositório deve ser analisado com chave específica de `user-microservice`, ajuste o valor no `sonar-project.properties`.