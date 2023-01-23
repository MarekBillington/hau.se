
use chrono::{DateTime, Utc};
use diesel::{Queryable, Insertable, table};
use rocket::database;
use rocket::serde::{Serialize, Deserialize, json::Json};

#[database("db")]
pub struct Db(diesel::PgConnection);

#[derive(Serialize, Deserialize, Queryable, Debug, Insertable)]
#[serde(crate = "rocket::serde")]
struct House {
    id: i64,
    // created: chrono::DateTime<Utc>,
    // updated: chrono::DateTime<Utc>,
    address: String,
    bedroom: i32,
    bathroom: i32,
    garage: i32,
    landarea: f32,
    floorarea: f32,
    misc: String,
}

#[derive(Deserialize)]
#[serde(crate = "rocket::serde")]
struct NewHouse {
    address: String,
}

#[rocket::get("/house")]
pub async fn all_houses(conn: Db) -> Json<Vec<House>> {
    conn.run(|c| house::table.load(c))
        .await
        .map(Json)
        .expect("Failed to fetch houses")
}

// #[rocket::get("/house/<id>")]
// pub async fn house_by_id(id: i64) {
    
//     Ok(())
// }

// #[rocket::post("/house", format = "application/json", data = "<house>")]
// pub async fn create_house(house: Json<NewHouse>) {
    
//     Ok(())
// }
