# **Journey API** [WIP]

This directory houses the WIP `Go` back end for the `Journey` project. 

It can receive requests from the `React` [front end](https://github.com/Ngdustry/golang-journey-ui) (written in `Typescript`) to fetch, create, update, and delete tasks.

---

## Local Environment

From the project root, you can run locally:

```
go mod download
```

Installs all required packages.

```
go run cmd/main.go
```

Runs the main Go file to connect to a `PostgreSQL` database and spin up a server to listen for requests on port [8000](http://localhost:8000).<br> 