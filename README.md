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

## Qualidade de código (SonarCloud)

Configuração atual em `sonar-project.properties`:

- `sonar.organization=fiap-11soat-grupo-21`
- `sonar.projectKey=FIAP-11soat-grupo-21_hackathon-auth-microservice`

> Observação: o `projectKey` atual referencia `hackathon-auth-microservice`. Se este repositório deve ser analisado com chave específica de `user-microservice`, ajuste o valor no `sonar-project.properties`.

## SonarScanner (exemplo)

Se quiser rodar análise local/manual (com token configurado no ambiente):

```bash
sonar-scanner
```

## Próximos ajustes recomendados

- Definir e documentar variáveis de ambiente necessárias para execução local.
- Confirmar/corrigir `sonar.projectKey` para refletir este repositório.
- Adicionar pipeline CI para executar testes e análise SonarCloud automaticamente.
