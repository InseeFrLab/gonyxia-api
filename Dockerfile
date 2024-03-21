FROM golang:1.22.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
ADD . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN ls .
RUN swag init
RUN ls docs

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /onyxia-api

FROM alpine

WORKDIR /app
EXPOSE 8080
ENV GIN_MODE release
COPY --from=0 /onyxia-api /app/onyxia-api
COPY config.yaml /app/config.yaml
CMD ["/app/onyxia-api"]
