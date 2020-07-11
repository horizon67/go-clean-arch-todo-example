# go-clean-arch-todo-example

## Example structure
```
.
├── README.md
├── app
│   ├── domain
│   │   └── todo.go
│   ├── infrastructure
│   │   ├── router.go
│   │   └── sqlhandler.go
│   ├── interfaces
│   │   ├── controllers
│   │   │   ├── context.go
│   │   │   ├── error.go
│   │   │   └── todo_controller.go
│   │   └── database
│   │       ├── sqlhandler.go
│   │       └── todo_repository.go
│   ├── server.go
│   └── usecase
│       ├── todo_interactor.go
│       └── todo_repository.go
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Installation

```
$ git clone https://github.com/horizon67/go-clean-arch-todo-example.git
```

## Run

```
$ cd go-clean-arch-todo-example
$ docker-compose up
```

## Usage

```
# list
curl -i --header "Content-type: application/json" -X GET localhost:8080/todos

# show
curl -i --header "Content-type: application/json" -X GET localhost:8080/todos/1

# create
curl -i --header "Content-Type: application/json" -X POST localhost:8080/todos -d '{"Content": "aa"}'

# update
curl -i --header "Content-Type: application/json" -X PUT localhost:8080/todos/1 -d '{"Content": "bb"}'

# delete
curl -i --header "Content-Type: application/json" -X DELETE localhost:8080/todos/1
```
