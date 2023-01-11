#[macro_use] extern crate rocket;

mod house;

#[launch]
fn rocket() -> _ {
    let routes:Vec<rocket::Route> = routes![
        house::all_houses,
        house::house_by_id
    ];
    
    // @todo some kind of security fairing required on requests
    rocket::build()
        .mount("/", routes)
}
