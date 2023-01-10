#[macro_use] extern crate rocket;

#[rocket::get("/house/<id>")]
fn house(id: i64) -> & 'static str {
    let output: String = id.to_string();
    
}

#[rocket::launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![house])
}
