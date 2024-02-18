# Golang Boilerplate Using DDD Structure


## clone the project
```bash
git clone https://github.com/JubaerHossain/golang-ddd && cd golang-ddd
```

## install dependencies
```bash
make install
```

## run the project [development mode]
```bash
make dev
```

## build the project
```bash
make build
```

## run the project [production mode]
```bash
make run
```

## Features
- [x] Golang
- [x] DDD
- [x] Clean Architecture
- [x] Docker
- [x] Makefile
- [x] Swagger
- [x] Gorm
- [x] JWT
- [x] Viper
- [x] Logger[Zap]
- [x] Unit Test

## Project Structure
```bash
.
├── cmd
│   └── server
│       └── server.go
├── internal
│   ├── core
│   │   └── auth
│   │       └── auth.go
│   ├── cache
│   │   └── cache.go
│   ├── database
│   │   └── database.go
|   ├── errors
│   │   └── errors.go
│   ├── health
│   │   └── health.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   └── auth.middleware.go
│   ├── monitor
│   │   └── monitor.go
│   ├── server
│   │   └── server.go
│   ├── domain
│   │   ├── model
│   │   │   └── user.go
│   │   └── repository
│   │       └── user.go
│   ├── infrastructure
│   │   ├── database
│   │   │   └── database.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   ├── router
│   │   │   └── router.go
│   │   └── security
│   │       └── security.go
│   ├── interfaces
│       ├── controller
│       │   └── user.go
│       ├── middleware
│       │   └── middleware.go
│       └── response
│           └── response.go
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── swagger
    └── swagger.yaml
```

## API Endpoints
```bash
GET /api/v1/users
POST /api/v1/users
GET /api/v1/users/{id}
PUT /api/v1/users/{id}
DELETE /api/v1/users/{id}
```

## Health Check
```bash
http://localhost:8080/health
```

## Metrics
```bash
http://localhost:8080/metrics
```

## Swagger
```bash
http://localhost:8080/swagger/index.html
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
```


