<div id="top"></div>

<div align="center" style="margin-bottom: 50px;">
  <img src="docs/assets/bbb.png" alt="Flow Icon" width="150" height="auto"/>

  <h1>BBB Voting System </h1>

   <h4>
    <a href="https://github.com/Jav4Script/pulls">Request Feature</a>
    <span> . </span>
    <a href="https://github.com/Jav4Script/issues">Report Issue</a>
  </h4>
</div>


Um sistema altamente escalável e confiável para gerenciamento de votações em tempo real, inspirado em desafios do Big Brother Brasil.

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Índice

- [Índice](#índice)
- [Pré-requisitos](#pré-requisitos)
- [Setup do Projeto](#setup-do-projeto)
- [Comandos Úteis](#comandos-úteis)
- [Arquitetura](#arquitetura)
  - [Estrutura de Pastas](#estrutura-de-pastas)
- [Dependências e Justificativas](#dependências-e-justificativas)
- [Variáveis de Ambiente](#variáveis-de-ambiente)
- [Referências e Cheatsheets](#referências-e-cheatsheets)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Pré-requisitos 
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Setup do Projeto 
1. Clone este repositório:
    ```bash
    git clone https://github.com/Jav4Script/bbb-voting-system.git
    cd bbb-voting-system
    ```

2. Copie o arquivo `.env.example` para `.env`:
    ```bash
    cp .env.example .env
    ```

3. Gere a documentação Swagger e as dependências Wire:
    ```bash
    make swag
    make wire
    ```

4. Compile a aplicação para desenvolvimento:
    ```bash
    make build-dev
    ```

5. Execute a aplicação para desenvolvimento:
    ```bash
    make run-dev
    ```

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Comandos Úteis 

| Comando            | Descrição                          |
|--------------------|------------------------------------|
| `make swag`        | Gera a documentação Swagger        |
| `make wire`        | Gera as dependências Wire          |
| `make build-dev`   | Compila o projeto para desenvolvimento |
| `make build-prod`  | Compila o projeto para produção    |
| `make run-dev`     | Executa a aplicação para desenvolvimento |
| `make run-prod`    | Executa a aplicação para produção  |
| `make stop`        | Para todos os containers em execução |
| `make clean`       | Remove arquivos temporários        |
| `make clear-redis` | Limpa todos os dados do Redis      |

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Arquitetura 

Este sistema foi projetado com os seguintes componentes:

- API REST para gerenciamento de votos.
- Redis para armazenamento temporário de resultados.
- RabbitMQ para gerenciamento de picos de tráfego.
- PostgreSQL para persistência de dados históricos.

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

### Estrutura de Pastas

```plaintext
.
├── .air.toml                   # Configuração do Air para hot reload
├── .env                        # Arquivo de configuração de variáveis de ambiente
├── .env.example                # Exemplo de arquivo de configuração de variáveis de ambiente
├── .gitignore                  # Arquivo para especificar quais arquivos/diretórios o Git deve ignorar
├── .vscode/                    # Configurações do Visual Studio Code
├── cmd/
│   └── main.go                 # Arquivo principal da aplicação
├── docker-compose.yml          # Arquivo de configuração do Docker Compose
├── Dockerfile                  # Arquivo de configuração do Docker
├── docs/
│   ├── assets/                 # Recursos estáticos para documentação
│   ├── docs.go                 # Documentação gerada pelo Swag
│   ├── swagger.json            # Arquivo JSON da documentação Swagger
│   └── swagger.yaml            # Arquivo YAML da documentação Swagger
├── go.mod                      # Arquivo de dependências do Go
├── go.sum                      # Hashes das dependências do Go
├── HISTORY.md                  # Histórico do projeto
├── internal/
│   ├── application/
│   │   └── usecases/           # Casos de uso da aplicação
│   │       ├── cast_vote_usecase.go          # Caso de uso para registrar um voto
│   │       ├── create_participant_usecase.go # Caso de uso para criar um participante
│   │       ├── get_final_results_usecase.go  # Caso de uso para obter resultados finais
│   │       ├── sync_cache_usecase.go         # Caso de uso para sincronizar o cache
│   │       └── ...                            # Outros casos de uso
│   ├── domain/
│   │   ├── constants.go        # Constantes do domínio
│   │   ├── dtos/               # Data Transfer Objects
│   │   ├── entities/           # Entidades do domínio
│   │   │   ├── participant.go  # Entidade de participante
│   │   │   └── vote.go         # Entidade de voto
│   │   ├── errors/             # Definições de erros do domínio
│   │   ├── producer/           # Interface de produtor de mensagens
│   │   ├── repositories/       # Interfaces de repositórios
│   │   └── services/           # Interfaces de serviços
│   ├── infrastructure/
│   │   ├── config/             # Configurações da infraestrutura
│   │   │   ├── database.go     # Configuração do banco de dados
│   │   │   ├── environment.go  # Carregamento de variáveis de ambiente
│   │   │   ├── rabbitmq.go     # Configuração do RabbitMQ
│   │   │   ├── redis.go        # Configuração do Redis
│   │   │   └── wire.go         # Configuração do Wire para injeção de dependências
│   │   ├── consumer/           # Consumidores de mensagens
│   │   │   └── rabbitmq_consumer.go          # Consumidor de mensagens do RabbitMQ
│   │   ├── controllers/        # Controladores da aplicação
│   │   │   ├── participant_controller.go     # Controlador de participantes
│   │   │   ├── vote_controller.go            # Controlador de votos
│   │   │   └── ...                            # Outros controladores
│   │   ├── jobs/               # Jobs da aplicação
│   │   │   └── sync_cache_job.go             # Job para sincronizar o cache
│   │   ├── mappers/            # Mapeadores de entidades
│   │   │   ├── participant_result_mapper.go  # Mapeador de resultados de participantes
│   │   │   ├── final_results_mapper.go       # Mapeador de resultados finais
│   │   │   └── vote_mapper.go                # Mapeador de votos
│   │   ├── middlewares/        # Middlewares da aplicação
│   │   ├── models/             # Modelos da aplicação
│   │   │   ├── participant_model.go          # Modelo de participante
│   │   │   ├── vote_model.go                 # Modelo de voto
│   │   │   └── ...                            # Outros modelos
│   │   ├── producer/           # Implementação do produtor de mensagens
│   │   │   └── rabbitmq_producer.go          # Implementação do RabbitMQ
│   │   ├── repositories/       # Implementações dos repositórios
│   │   │   ├── postgres/
│   │   │   │   ├── participant_repository.go # Implementação do repositório de participantes no PostgreSQL
│   │   │   │   └── vote_repository.go        # Implementação do repositório de votos no PostgreSQL
│   │   │   └── redis/
│   │   │       └── redis_repository.go       # Implementação do repositório no Redis
│   │   ├── routes.go           # Definição das rotas da aplicação
│   │   ├── server.go           # Inicialização do servidor
│   │   └── services/           # Serviços da aplicação
│   │       └── captcha_service.go            # Serviço de CAPTCHA
├── main                        # Arquivo principal da aplicação
├── makefile                    # Arquivo de automação de tarefas
├── README.md                   # Documentação do projeto
└── scripts/
    └── queue-init.sh           # Script de inicialização da fila
├── tmp/
│   ├── build-errors.log        # Log de erros de build
│   └── main                    # Binário principal gerado pelo build
```

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Dependências e Justificativas 

- Gin: Framework web para APIs REST em Go.
- Redis: Armazenamento em memória para resultados parciais.
- RabbitMQ: Gerenciamento de filas para lidar com picos de tráfego.
- PostgreSQL: Persistência de dados históricos.
- Swag: Geração de documentação Swagger.
- Wire: Injeção de dependências.

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Variáveis de Ambiente 

Aqui estão as variáveis de ambiente necessárias para configurar o projeto. Substitua os placeholders pelos valores apropriados:

```plaintext
APP_ENV=development                                  # Ambiente da aplicação (ex: development, test, production)
DATABASE_NAME=your_database_name                     # Nome do banco de dados
DATABASE_SCHEMA=your_schema                          # Esquema do banco de dados
DATABASE_PORT=5432                                   # Porta do banco de dados
DATABASE_HOST=your_database_host                     # Host do banco de dados
DATABASE_USER=your_database_user                     # Usuário do banco de dados
DATABASE_PASSWORD=your_password                      # Senha do banco de dados
CORS_ALLOWED_ORIGINS=your_cors_allowed_origins       # Origens permitidas pelo cors
RABBITMQ_USER=your_rabbitmq_user                     # Usuário do RabbitMQ
RABBITMQ_PASSWORD=your_password                      # Senha do RabbitMQ
RABBITMQ_HOST=your_rabbitmq_host                     # Host do RabbitMQ
RABBITMQ_PORT=5672                                   # Porta do RabbitMQ
RABBITMQ_VHOST=your_vhost                            # Virtual host do RabbitMQ
VOTE_QUEUE=your_vote_queue                           # Nome da fila de votos
REDIS_URL=your_redis_url                             # URL do Redis
SYNC_CACHE_INTERVAL=your_sync_cache_interval_time    # Sync cache time in minutes
```

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Referências e Cheatsheets 

- [Docker Compose Cheatsheet](https://devhints.io/docker-compose)
- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [RabbitMQ Documentation](https://www.rabbitmq.com/documentation.html)
- [Redis Documentation](https://redis.io/documentation)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Wire Documentation](https://github.com/google/wire)
- [Gorm Documentation](https://github.com/go-gorm/gorm)
  
<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)