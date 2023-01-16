# hau.se
 
## Front end - Qwik (React)

cd ./hau.se/hause-fe & npm run

localhost:5173

node_modules ~ 

## Back End - Rocket (Rust)

cd ./hau.se/hause-be & cargo run

localhost:8000

docker run --rm -it -p 8000:8000/tcp --mount type=bind,source=/home/marek/www/hau.se/server/target/debug/hause-be,target=/user/src/app/hause-be hause:latest

## Postgres (RDS)

postgres://postgres:postgrespw@localhost:49153