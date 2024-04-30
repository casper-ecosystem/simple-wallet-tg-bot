FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .

RUN go mod download
RUN go build -C ./tggateway/cmd 


FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/tggateway/cmd/cmd  /app/


CMD ["./cmd"]