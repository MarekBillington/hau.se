use rocket::serde::{Serialize, Deserialize, json::Json};

#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
pub struct House {
    id: i64,
    address: String,  
}

#[derive(Deserialize)]
#[serde(crate = "rocket::serde")]
pub struct NewHouse {
    address: String,
}

#[rocket::get("/house")]
pub fn all_houses() -> Json<Vec<House>> {
    let h = House {id: 1, address: String::from("abc")};
    let h2 = House {id: 2, address: String::from("abc")};

    let mut vec = Vec::new();
    vec.push(h);
    vec.push(h2);
    
    Json(vec)
}

#[rocket::get("/house/<id>")]
pub fn house_by_id(id: i64) -> Json<House> {
    let h3 = House {id: id, address: String::from("123 test road")};
    Json(h3)
}

#[rocket::post("/house", format = "application/json", data = "<house>")]
pub fn create_house(house: Json<NewHouse>) -> Json<House> {
    let h = House {id: 1, address: String::from(&house.address)};
    Json(h)
}
