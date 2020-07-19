# Gotube

Gotube is a learning project

The purpuse is make an API to upload videos, something like YouTube, allowing people to upload, comment videos. Track views and reactions.

## Needed

To run the project is necesary to have installed Golang and MySQL

- [Golang](https://golang.org)
- [MySQL](https://www.mysql.com/)

Database could be run in a [Docker](https://www.docker.com/) container:

`docker run -it -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=gotube mysql`

## Clone and Install instructions

### Clone the project:

`git clone https://github.com/wllamasr/gotube`

### Download go modules:

`go mod download`

_(this is a non mandatory command, go modules will be download before the project start)_

### Create and edit a `.env` file

`cp .env.example .env`

### Run the project

`go run main.go`


## Features:

- [X] Register
- [X] Login
- [X] Authentication
- [ ] Upload
- [ ] Comments
- [ ] Authorization
- [ ] Testing

ToDo list will be updated as the project goes on.
