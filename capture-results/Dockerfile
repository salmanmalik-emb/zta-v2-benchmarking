FROM golang:1.18-alpine3.15 as build
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

# production environment
FROM scratch
COPY --from=build /app /
ENTRYPOINT ["./capture_results"]