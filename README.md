# **Journey API** [WIP]

This repo houses the WIP `Go` back end for the `Journey` project.

It can receive requests from the `React` [front end](https://github.com/Ngdustry/golang-journey-ui) (written in `Typescript`) to fetch, create, update, and delete tasks with a PostgreSQL database.

---
## Setup

To run this project locally, the following should be installed:
- Go (https://golang.org/)
- Docker (https://docs.docker.com/get-docker/)

For local development, this project uses two docker containers: one for the API and another for a PostgreSQL database.

 For live reloading, the API uses `air` (https://github.com/cosmtrek/air) to track local file changes and automatically restart the API.

## Local Environment

To spin up the API docker container:
```
make docker-up
```

To stop the API docker container:
```
make docker-down
```
