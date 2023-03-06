# hau.se

When making changes, use the `/scripts/setup.sh` script to build the client/server and reload their related docker containers

## Front end - Qwik (React)

> Required: node 18.13.0

cd ./hau.se/client & npm run

Local testing and building goes to localhost:5173

> https://qwik.builder.io/docs/overview/
>
> https://github.com/wmertens/styled-vanilla-extract
>
> https://vanilla-extract.style/documentation/getting-started/

```
client /
 src /
  routes /
   | index.tsx
   house /
    | index.tsx
   tenants /
   ...
 components /
  house /
   interfaces /
    | house.tsx
   screens /
    <..name of area>/
     | panel.tsx
    <.. name of area> /
     | form.tsx
  utility /
   api /
    | api.tsx
   inputs /
    | button.tsx
    | text.tsx
    | ...
```

## Back End - Gin (Go)

> Required: golang 1.19.5

cd ./hau.se/server & go run . -isDev=true

Local port 8002
Prod port 8001

> https://pkg.go.dev/github.com/gin-gonic/gin#section-documentation
>
> https://gorm.io/docs/index.html

Struct:
```
server \
 | main.go
 | go.sum
 | go.mod
 auth \
  | controller.go
  dto \
   | login.go
   | ...
  middleware \
   | jwt_auth.go
  repository \
   | create_user.go
   | ...
  service \
   | handler.go
   | login_user.go
   | ...
 entity \
 | migration.go
 | user.go
```

Debuging, use to see cli output `fmt.Fprintf(os.Stdout, "%+v", val)`

## Postgres (RDS)

Will setup automatically from docker compose

## Nginx

Front end is distributed at http://dev.hau.se/

Back end is distributed at http://dev.hau.se/api/

Required, add `120.0.0.1 dev.hau.se` to `hostsfile`
