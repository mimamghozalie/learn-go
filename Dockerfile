FROM golang:1.20.6-alpine3.18 AS builder

WORKDIR /app
# COPY cmd ./cmd
# COPY go.mod .
# COPY go.sum .
COPY . .
RUN go mod download
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# RUN go build main.go

FROM scratch
COPY --from=builder /app/main /main
ENV PORT=8080
ENV HOST=0.0.0.0
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["/main"]