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


migrate create -ext sql -dir db/migration -seq your_migration_name
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

# git
```powershell
git config --global http.proxy 127.0.0.1:7890
git config --global https.proxy 127.0.0.1:7890
git config --global --get http.proxy
git config --global --get https.proxy
git config --global --unset http.proxy
git config --global --unset https.proxy
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
