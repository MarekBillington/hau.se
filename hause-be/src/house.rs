#[rocket::get("/house")]
pub fn all_houses() -> &'static str {
    "all houses"
}

#[rocket::get("/house/<id>")]
pub fn house_by_id(id: i64) -> String {
    let output: String = id.to_string();
    output
}
