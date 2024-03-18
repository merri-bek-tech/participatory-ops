FROM golang:1.21 as gobuilder
WORKDIR /app

COPY ./api .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/parops-server

FROM ubuntu as runner
COPY --from=gobuilder /app/parops-server /usr/local/bin/parops-server
EXPOSE 1323
CMD ["parops-server"]