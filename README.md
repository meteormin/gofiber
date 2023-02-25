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

REDIS_HOST=go-redis
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DATABASE=0

CREATE_ADMIN=true
CREATE_ADMIN_USERNAME=smyoo
CREATE_ADMIN_PASSWORD=smyoo
CREATE_ADMIN_EMAIL=admin@email.com

```

## Packages

### gofiber

- create new application
- run fiber web application

```go
package main

import "github.com/miniyus/gofiber"

func main() {
	a := gofiber.New() // 지정된 기본 설정을 가지고 새로운 application 생성
}
```

### config

- 설정 관리
- go 언어의 구조체를 활용하여 관리

```go
package main

import (
	"github.com/go-redis/redis/v9"
	fCors "github.com/gofiber/fiber/v2/middleware/cors"
	fCsrf "github.com/gofiber/fiber/v2/middleware/csrf"
	fLoggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/permission"
	worker "github.com/miniyus/goworker"
)

func main() {
	cfg := config.GetConfigs() // 기본 설정 가져오기

}

type Configs struct {
	App            app.Config
	Logger         fLoggerMiddleware.Config
	CustomLogger   map[string]cLog.Config
	Database       map[string]database.Config
	Path           Path
	Auth           Auth
	Cors           fCors.Config
	Csrf           fCsrf.Config
	Permission     []permission.Config
	CreateAdmin    CreateAdminConfig
	RedisConfig    *redis.Options
	JobQueueConfig worker.DispatcherOption
	Validation     Validation
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

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
)

func main() {
	// 애플리케이션 생성
	a := app.New(app.Config{
		Env:         app.PRD,
		Port:        8000,
		Locale:      "",
		TimeZone:    "Asia/Seoul",
		FiberConfig: fiber.Config{},
	})
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

```go
// 라우트 등록
package main

import (
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/routes"
)

func main() {
	a := gofiber.New() // 지정된 기본 설정을 가지고 새로운 application 생성

	a.Route("prefix", func(r app.Router, a app.Application) {
		routes.Api(r, a)
	}, "router-group-name")

	// 다른 방식
	a.Route("prefix", routes.Api, "router-group-name")
}

```

### 기타 기능

- apierrors: 기본 error 핸들러, 에러 응답 관련 기능
- auth: jwt 토큰 기반의 인증 기능
    - 회원가입, 로그인, 로그아웃, 패스워드 변경등의 기본 API 구현
- create_admin: 초기 최고 관리자 생성을 위한 패키지
- entity: db 연동 관련되어 gorm.Model을 이용하여 생성한 entity 구조체
- groups: group 관련 api
- internal: 내부 사용 기능
    - base64
    - datetime
    - hash
    - reflect
- job_queue: 대기열 작업 큐 + DB(job history 관련) 기능 및 hooks
- jobs: 작업 관련 api
    - 현재 작업 워커 현황 체크
    - 현재 redis에 활성화되어 있는 작업 조회
- ~~permission: 권한 관련 기능 및 미들웨어 구현~~(Feature)

### pkg

- IOContainer
    - 구조체 주입 및 인터페이스 > 구현체 바인딩
    - [test코드 참조](./pkg/IOContainer/container_test.go)
- jwt
    - json web token generator
    - [test코드 참조](./tests/jwt_test/jwt_test.go)
- rs256
    - rs256 encode and decode
    - [test코드 참조](./tests/rs256_test/rs256_test.go)
- validation
    - 커스텀 유효성 검사 및 번역
    - [test코드 참조](./pkg/validation/validator_test.go)

### 외부 패키지

- [miniyus/gollection](https://github.com/miniyus/gollection)
- [miniyus/gorm-extension](https://github.com/miniyus/gorm-extension)
- [miniyus/goworker](https://github.com/miniyus/goworker)