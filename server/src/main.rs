
#[macro_use] extern crate rocket;
#[macro_use] extern crate rocket_sync_db_pools;
#[macro_use] extern crate diesel;

mod house;

// #[database("db")]
// struct Db(diesel::PgConnection);

#[launch]
fn rocket() -> _ {
    // @todo some kind of security fairing required on requests
    rocket::build()
        .attach(house::stage())
}
