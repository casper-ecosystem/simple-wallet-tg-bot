FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
ENV GOPRIVATE=github.com/Simplewallethq/*
COPY . .
ARG GITHUB_ACCESS_TOKEN

RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go mod download
RUN go build -o main .


FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/main /app/


EXPOSE 8081

CMD ["./main"]