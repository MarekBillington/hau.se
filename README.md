# hau.se
 
## Front end - Qwik (React)

cd ./hau.se/client & npm run

localhost:5173

node_modules ~ 

## Back End - Rocket (Rust)
SEE: https://itnext.io/creating-a-rust-web-app-with-rocket-and-diesel-58f5f6cacd27#856adi

cd ./hau.se/server & cargo run

localhost:8000

`cargo install diesel_cli@2.0.1 --no-default-features --features postgres`
>> if error: linking with `cc` failed: exit status: 1 
>> Do `sudo apt-get install libpq-dev`


## Postgres (RDS)

postgres://postgres:postgres@localhost:5432