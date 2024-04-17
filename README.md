# init
```powershell
go mod init github.com/LKarrie/learn-go-project
go mod tidy
```

# install make
```powershell
choco install make
```

# install migrate
## open powershell(admin)
```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser 
iex "& {$(irm get.scoop.sh)} -RunAsAdmin"
scoop install migrate
scoop update migrate

migrate create -ext sql -dir db/migration -seq your_migration_name

run migrate in code 
https://github.com/golang-migrate/migrate

```

# fix make missing separator error
```markdown
Delete the spaces at the beginning of the line.
Insert a tab character at the beginning of the line.
```

# show lock
```sql
SELECT blocked_locks.pid     AS blocked_pid,
        blocked_activity.usename  AS blocked_user,
        blocking_locks.pid     AS blocking_pid,
        blocking_activity.usename AS blocking_user,
        blocked_activity.query    AS blocked_statement,
        blocking_activity.query   AS current_statement_in_blocking_process
  FROM  pg_catalog.pg_locks         blocked_locks
  JOIN pg_catalog.pg_stat_activity blocked_activity  ON blocked_activity.pid = blocked_locks.pid
  JOIN pg_catalog.pg_locks         blocking_locks 
      ON blocking_locks.locktype = blocked_locks.locktype
      AND blocking_locks.database IS NOT DISTINCT FROM blocked_locks.database
      AND blocking_locks.relation IS NOT DISTINCT FROM blocked_locks.relation
      AND blocking_locks.page IS NOT DISTINCT FROM blocked_locks.page
      AND blocking_locks.tuple IS NOT DISTINCT FROM blocked_locks.tuple
      AND blocking_locks.virtualxid IS NOT DISTINCT FROM blocked_locks.virtualxid
      AND blocking_locks.transactionid IS NOT DISTINCT FROM blocked_locks.transactionid
      AND blocking_locks.classid IS NOT DISTINCT FROM blocked_locks.classid
      AND blocking_locks.objid IS NOT DISTINCT FROM blocked_locks.objid
      AND blocking_locks.objsubid IS NOT DISTINCT FROM blocked_locks.objsubid
      AND blocking_locks.pid != blocked_locks.pid

  JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
  WHERE NOT blocked_locks.granted;

SELECT  a.application_name,
         l.relation::regclass,
         l.transactionid,
         l.mode,
         l.locktype,
         l.GRANTED,
         a.usename,
         a.query,
         a.pid
FROM pg_stat_activity a
JOIN pg_locks l ON l.pid = a.pid
ORDER BY a.pid;
```

# sqlc
```markdown
https://github.com/sqlc-dev/sqlc
https://docs.sqlc.dev/en/latest/

version 1

version: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: false
    emit_interface: true
    emit_exact_table_names: false
    emit_empty_slices: true
```

# git
```powershell
git config --global http.proxy 127.0.0.1:7890
git config --global https.proxy 127.0.0.1:7890
git config --global --get http.proxy
git config --global --get https.proxy
git config --global --unset http.proxy
git config --global --unset https.proxy
```

# testify
> https://github.com/stretchr/testify
```powershell
go get github.com/stretchr/testify
```

# GIN
> https://gin-gonic.com/zh-cn/docs/quickstart/
```powershell
go get -u github.com/gin-gonic/gin
```

# Viper
> https://github.com/spf13/viper
```powershell
go get github.com/spf13/viper
```

# Gomock
> https://github.com/uber-go/mock
```powershell
go install go.uber.org/mock/mockgen@latest
mocken -version
```
```markdown
error
mockgen -destination db/mock/store.go github.com/LKarrie/learn-go-project/db/sqlc Store
prog.go:12:2: no required module provides package go.uber.org/mock/mockgen/model; to add it:
        go get go.uber.org/mock/mockgen/model
prog.go:12:2: no required module provides package go.uber.org/mock/mockgen/model; to add it:
        go get go.uber.org/mock/mockgen/model
prog.go:14:2: no required module provides package github.com/LKarrie/learn-go-project/db/sqlc: go.mod file not found in current directory or any parent directory; see 'go help modules'      
prog.go:12:2: no required module provides package go.uber.org/mock/mockgen/model: go.mod file not found in current directory or any parent directory; see 'go help modules'
2023/11/11 16:36:13 Loading input failed: exit status 1

fix:

import _ "github.com/golang/mock/mockgen/model"
mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/LKarrie/learn-go-project/db/sqlc Store

```

# UUID
```powershell
go get github.com/google/uuid
```

# JWT
> https://github.com/golang-jwt/jwt
```powershell
go get github.com/dgrijalva/jwt-go
```

# PASETO
> https://github.com/o1egl/paseto
```powershell
go get github.com/o1egl/paseto
```

# GIT
```powershell
git checkout -b ft/docker
git status
git add .
git commit -m "update readme"
```

# Docker 
```powershell
docker network create bank-network
docker network connect bank-network postgres12
docker network inspect bank-network 
docker container inspect postgres12

docker desktop WSL error
run 
netsh winsock reset  
```

# Other
```shell
openssl rand -hex 64 | head -c 32
```

```markdown
sh: ./wait-for.sh: not found
CRLF TO LF

Makefile:29: *** missing separator.  Stop.
use tab in makefile not space
```

```markdwon
npm install -g dbdocs
dbdocs login
dbdocs build doc/db.dbml
dbdocs password --set secret --project simple_bank

npm install -g @dbml/cli
dbml2sql --postgres -o doc/schema.sql doc/db.dbml

```

# grpc
```markdown
4 types of grpc
unary grpc
client streaming grpc
server streaming grpc
bidirectional streaming grpc

https://grpc.io/docs/languages/go/quickstart/

install protobuf
https://github.com/protocolbuffers/protobuf/releases

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version

https://protobuf.dev/programming-guides/proto3/

test grpc client 
https://github.com/ktr0731/evans
https://github.com/ktr0731/evans/releases

package pb
show service
service LearnGo
call CreateUser

```

# grpc gateway
```markdown
gRPC Gateway
A plugin of protobuf compiler
Generate proxy code from protobuf
Translate HTTP JSON calls to gRPC
* in-process translation: only for unary
* Separate proxy server: both unary and streaming

https://github.com/grpc-ecosystem/grpc-gateway
https://grpc-ecosystem.github.io/grpc-gateway/
https://github.com/grpc-ecosystem/grpc-gateway/blob/main/examples/internal/proto/examplepb/a_bit_of_everything.proto

// +build tools

package tools

import (
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
    _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc-gen-grpc-gateway -help

```

# swagger
```markdown
change swagger-initializer.js

```

# Statik
```
statik allows you to embed a directory of static files into your Go binary to be later served from an http.FileSystem.

https://github.com/rakyll/statik

tools file add github.com/rakyll/statik
and
go install github.com/rakyll/statik

statik -help
-ns      The namespace where assets will exist, "default" by default.

```

# zerolog 
```markdown
https://github.com/rs/zerolog
go get -u github.com/rs/zerolog/log

zerolog allows for logging at the following levels (from highest to lowest):

panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)

```

# Anynq
```markdwon
https://github.com/hibiken/asynq
Features
Guaranteed at least one execution of a task
Scheduling of tasks
Retries of failed tasks
Automatic recovery of tasks in the event of a worker crash
Weighted priority queues
Strict priority queues
Low latency to add a task since writes are fast in Redis
De-duplication of tasks using unique option
Allow timeout and deadline per task
Allow aggregating group of tasks to batch multiple successive operations
Flexible handler interface with support for middlewares
Ability to pause queue to stop processing tasks from the queue
Periodic Tasks
Support Redis Cluster for automatic sharding and high availability
Support Redis Sentinels for high availability
Integration with Prometheus to collect and visualize queue metrics
Web UI to inspect and remote-control queues and tasks
CLI to inspect and remote-control queues and tasks

go get -u github.com/hibiken/asynq
go mod tidy

```

# email
```markdown
https://github.com/jordan-wright/email

go get github.com/jordan-wright/email

```

# pgx
go get github.com/jackc/pgx/v5/pgxpool