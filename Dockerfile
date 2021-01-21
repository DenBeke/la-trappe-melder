FROM golang:latest AS build

WORKDIR /random_work_dir

# first download dependencies
COPY go.mod /random_work_dir
COPY go.sum /random_work_dir
RUN go mod download

# then copy source code
COPY / /random_work_dir


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /latrappemelder ./cmd/latrappemelder


FROM golang:latest


WORKDIR /

COPY --from=build /latrappemelder /bin/latrappemelder

WORKDIR /latrappemelder

RUN chmod +x /bin/latrappemelder

EXPOSE 1234

CMD ["/bin/latrappemelder"]