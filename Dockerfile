FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o url-shorter ./cmd/app/main.go

CMD ["./url-shorter"]