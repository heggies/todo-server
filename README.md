# TODO Server
This is a REST API that manages TODO list. Project structure is inspired by [eminetto/clean-architecture-go-v2](https://github.com/eminetto/clean-architecture-go-v2).

## Project Stack & Libraries
This project is using:

- PostgreSQL for the main database.
- [Fiber](https://github.com/gofiber/fiber) as the web framework.
- [GORM](https://github.com/go-gorm/gorm) as the ORM.

## Getting Started

### Requirements
- [Docker Compose](https://docs.docker.com/compose/install)

OR if you want to run it without Docker:
- Go 1.16
- PostgreSQL
- [Air](https://github.com/cosmtrek/air)

### Installation
```
git clone https://github.com/heggies/todo-server.git
```

### Usage
1. Create a copy of `.env.example`

```
cp .env.example .env
```

2. Set your environment variables

```
ENV=development
PORT=3000
POSTGRES_HOST=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
```

3. Start the project by running:

```
docker-compose up --build
```

OR if you want to run it without Docker:

```
air
```