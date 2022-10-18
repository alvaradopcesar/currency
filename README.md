# Currency API

API for Currency project.

![PostgreSQL](https://img.shields.io/badge/Postgres-13.5-lightblue.svg?logo=postgresql&longCache=true&style=flat) ![Go](https://img.shields.io/badge/Golang-1.18-blue.svg?logo=go&longCache=true&style=flat)

## Getting Started

This project uses the **Go** programming language (Golang) **PostgreSQL** database engine.

### Prerequisites

[PostgreSQL](https://www.postgresql.org/) is required in version 9.6 or higher, [Go](https://golang.org/) at least in version 1.18

### Running the tests

```bash
make test
```

### Running Environment Data Base Postgres
```bash
make start-compose
```

### Running the app

```bash
make init
make start
```