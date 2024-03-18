# FRONTEND BUILDER
FROM --platform=$BUILDPLATFORM node:latest as vitebuilder
WORKDIR /app
ADD ./frontend .
RUN npm install && npm run build

# API BUILDER
FROM golang:1.21 as gobuilder
WORKDIR /app

COPY ./api .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/parops-server

# API RUNNER
FROM ubuntu as runner
COPY --from=gobuilder /app/parops-server /app/parops-server
COPY --from=vitebuilder /app/dist /app/web
EXPOSE 1323
CMD ["/app/parops-server"]