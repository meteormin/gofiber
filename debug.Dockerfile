FROM golang:1.20-alpine
WORKDIR /app
EXPOSE 8000 2345

COPY . .
COPY .env${SELECT_ENV} .env

RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go mod download
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o gofiber ./cmd/gofiber/main.go

CMD [ "/go/bin/dlv", "--listen=:2345", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/gofiber" ]
