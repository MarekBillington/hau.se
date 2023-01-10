#[macro_use] extern crate rocket;

#[rocket::get("/world")]
fn world() -> &'static str {
    "Hello, world!"
}

#[rocket::launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index])
}
