#[macro_use] extern crate rocket;

mod house;
#[database("db")]
struct Db(diesel::PgConnection);

#[launch]
fn rocket() -> _ {
    let routes:Vec<rocket::Route> = routes![
        house::all_houses,
    ];
    
    // @todo some kind of security fairing required on requests
    rocket::build()
        .attach(Db::fairing())
        .mount("/", routes)
}
