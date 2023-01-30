
FROM golang:1.19


WORKDIR /RESTApi-GIN
COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o main .


FROM mysql:8.0


ENV MYSQL_ROOT_PASSWORD=root@123
ENV MYSQL_DATABASE=Employee
EXPOSE 3306
