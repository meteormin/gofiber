FROM golang:1.20-alpine as build

RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
	echo "${TIME_ZONE}" > /etc/timezone

RUN apk add --no-cache make

RUN mkdir /fiber

WORKDIR /fiber

COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 go mod download
RUN CGO_ENABLED=0 go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} make build

FROM alpine:latest

LABEL maintainer="miniyu97@gmail.com"

RUN mkdir /home/gofiber

ARG SELECT_ENV
ARG GO_GROUP
ARG GO_VERSION

WORKDIR /home/gofiber

COPY --from=build /fiber/build ./build

COPY .env${SELECT_ENV} .env

EXPOSE $APP_PORT

CMD ["./build/gofiber"]