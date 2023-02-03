
# CRUD REST API USING GOLANG GIN Framework,GORM AND MYSQL

This repository is Application of REST Api with Golang,GORM (ORM),MySQL and Unit Testing(mockgen,httptest).


## Usage/Examples

```javascript
$ git clone https://github.com/int1359/RESTApi-GIN-MySQL.git
[NAME APP]

$ go build

$ ./[name-app] [ENTER]
```


## Tech

API Server technology stack is

* Server code: Golang 1.19
* REST Server: gin Framework
* Database: MySQL
* ORM: gorm 
* Unit Testing: go test and testify,mockgen

## API Reference

#### Get all Employee

```http
  GET /Employee
```

#### Get Employee of Specific ID

```http
  GET /Employee/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of Employee to fetch |

#### POST Create Employee

```http
  POST /Employee
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| -   | - | **Required**. Employee Entity in Payload |


#### POST Create Employee to  Imaginary Service

```http
  POST /PushData
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|      -   |      -    |  Push All Employee Data to https://dummy.restapiexample.com/api/v1/create (imaginary service) |


#### Update Employee of Specific ID

```http
  PUT /Employee/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of Employee to fetch and Employee Entity in Payload |


## Deployment

To Run this project

```bash
  go run main.go
```
Run api server with mysql (see docker-compose.yaml)

```bash
docker compose up -d

docker ps

CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS         PORTS                                       NAMES
222451e3591a   golang:1.19                 "sh -c 'go run main.…"   6 seconds ago   Up 4 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   restapi-gin-mysql-app-1
b52d053ae508   mysql:8.0                   "docker-entrypoint.s…"   6 seconds ago   Up 4 seconds   3306/tcp, 33060/tcp                         restapi-gin-mysql-mysql-1
  
```

Configure DB info through the create of file .env
