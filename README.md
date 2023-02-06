# miniyus: Go fiber Template

## go lang with fiber framework

- gorm 사용
    - 현재 리포지토리에서는 postgreSQL만 사용중

## Install

```shell
git clone https://github.com/miniyus/gofiber.git

# 도커 컨테이너 및 gorm DB 드라이버는 postgreSQL을 사용 중입니다.
docker compose up -d --build 

# git pull & go mod download & docker compose up
make deploy

# 로컬 실행 시
make start
# or
make build
./build/gofiber
```

## Dot Env

```shell
APP_ENV=develop
APP_NAME=gofiber
APP_PORT=8080

TIME_ZONE=Asia/Seoul

GO_GROUP=1000
GO_USER=1000

DB_HOST=go-pgsql
DB_DATABASE=go_fiber
DB_PORT=5432
DB_USERNAME=?
DB_PASSWORD=?
DB_AUTO_MIGRATE=true

```

## Directory Structure

```shell
/project-root
|-- build: build된 파일이 저장
|-- cmd: go 실행 및 makefile을 통한 cli 명령을 실행할 수 있는 main.go 파일
|-- config: 설정
  |-- context: fiber context Locals를 통해 가져올 항목들을 미리 정의
|-- data: local data 저장소
|-- internal: api를 실질적으로 구현하는 곳 입니다.
  |-- api: api 요청 시, 수행되는 코드들
    |-- service_directory(example: users, groups...): 특정 end point 패키징
      |-- dto: DTO 정의 및 매핑 함수 정의
      |-- factory: Handler 생성을 위한 팩토리 패턴 적용
      |-- handler: 요청을 받고 응답을 해준다.
      |-- service: handler 요청의 비즈니스 로직 처리
      |-- repositroy: db, entity를 통해 데이터 CRUD 동작 수행
      |-- routes: 그룹화된 API 적용을 위한 서브 라우터
  |-- api_error: api error, error response 관련 기능
  |-- auth: 인증 관련 기능
  |-- database: database, gorm 연결
  |-- log: 로거
  |-- permission: permission 관련 기능  
  |-- resolver: factory 패턴을 이용하여, 필요한 모듈 생성 함수를 정의해 놓은 패키지
  |-- entity: db 스키마를 가진 구조체 집합
  |-- routes: 라우팅
  |-- utils: 유틸 함수들
|-- pkg: 독립적인 기능을 수핼 할 수 있는 기능들의 집합입니다.
|-- tests: .env, 파일시스템 등의 활용을 위한 패키지의 경우 경로의 깊이가 영향을 끼치기 때문에 테스트용 폴더를 따로 구분 
```

### config

- 설정 관리
- go 언어의 구조체를 활용하여 관리

```go
package main

import "github.com/miniyus/gofiber/config"

func main() {
	config.GetConfigs()
}

```

### database

```go
package main

import (
  "github.com/miniyus/gofiber/database"
  "os"
  "time"
)

func main() {
  cfg := database.Config{
    Name:        "default",
    Driver:      "postgres",
    Host:        os.Getenv("DB_HOST"),
    Dbname:      os.Getenv("DB_DATABASE"),
    Username:    os.Getenv("DB_USERNAME"),
    Password:    os.Getenv("DB_PASSWORD"),
    Port:        os.Getenv("DB_PORT"),
    TimeZone:    os.Getenv("TIME_ZONE"),
    SSLMode:     false,
    AutoMigrate: autoMigrate,
    Logger: gormLogger.Config{
      SlowThreshold:             time.Second,
      LogLevel:                  gormLogger.Silent,
      IgnoreRecordNotFoundError: true,
      Colorful:                  true,
    },
    MaxIdleConn: 10,
    MaxOpenConn: 100,
    MaxLifeTime: time.Hour,
  }
  
  database.New(cfg)
}
```

### log

```go
package main

import (
  "github.com/miniyus/gofiber/log"
  "go.uber.org/zap/zapcore"
  "os"
)

func main() {
  cfg := log.Config{
    Name:       "default",
    TimeFormat: "2006-01-02 15:04:05",
    FilePath:   "filePath",
    Filename:   "filename",
    MaxSize:    10,
    MaxBackups: 30,
    MaxAge:     30,
    Compress:   true,
    TimeKey:    "timestamp",
    TimeZone:   os.Getenv("TIME_ZONE"),
    LogLevel:   zapcore.DebugLevel,
  }
  log.New()
}

```

### app

```go
package main

import "github.com/miniyus/gofiber/app"

func main() {
	// 애플리케이션 생성
	a := app.New()
	a.Run()
}

```

### Routes

```go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
)

const ApiPrefix = "/api"

func Api(router app.Router, a app.Application) {
	router.Route(
		"test",
		func(r fiber.Router) {
			r.Get("/", func(ctx *fiber.Ctx) error {
				return ctx.JSON("test")
			})
		},
	).Name("api.test")
}

```

### internal
- 기능 구현
