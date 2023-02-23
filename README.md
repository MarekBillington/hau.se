# hau.se

When making changes, use the `/scripts/setup.sh` script to build the FE and BE and reload the docker containers

## Front end - Qwik (React)

> Required: node 18.13.0

cd ./hau.se/client & npm run

Local testing and building goes to localhost:5173

https://qwik.builder.io/docs/overview/

https://github.com/wmertens/styled-vanilla-extract

https://vanilla-extract.style/documentation/getting-started/

## Back End - Gin (Go)

> Required: golang 1.19.5

cd ./hau.se/server & go run . -isDev=true

https://pkg.go.dev/github.com/gin-gonic/gin#section-documentation

Local port 8002
Prod port 8001

Struct:
```
server (hause)
    \ user              > split by route group for folder, same name for package
        \ endpoints.go  > contains routes and controllers
        \ main.go       > for setup of gin, aggregates routes and db setup for structs
        \ model.go      > contains model functions to work with gorm
        \ struct.go     > contains structs used for endpoints and model
```

Debug `fmt.Fprintf(os.Stdout, "%+v", val)`

## Postgres (RDS)


## Nginx

Front end is distributed at http://dev.hau.se/

Back end is distributed at http://dev.hau.se/api/

Required, add `120.0.0.1 dev.hau.se` to `hostsfile`