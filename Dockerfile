# build stage
FROM golang:1.21.0-bullseye AS builder

RUN apt-get update && apt-get install -y git

WORKDIR /app

COPY . .

RUN go mod download 

RUN go build -o ./out/dist cmd/api/main.go

# production stage
FROM busybox

COPY --from=builder /app/out/dist /app/

COPY template ./template

WORKDIR /app

EXPOSE 3000

CMD ["./dist"]