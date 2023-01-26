
use rocket::fairing::AdHoc;
use rocket::serde::{Serialize, Deserialize, json::Json};
use rocket_sync_db_pools::diesel;

use self::diesel::prelude::*;

#[database("db")]
pub struct Db(diesel::PgConnection);

#[derive(Serialize, Deserialize, Debug, Queryable)]
#[serde(crate = "rocket::serde")]
#[diesel(table_name = house)]
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

table! {
    house (id) {
        id -> Int4,
        created -> Timestamptz,
        updated -> Timestamptz,
        address -> Varchar,
        bedroom -> Int4,
        bathroom -> Int4,
        garage -> Int4,
        landarea -> Float4,
        floorarea -> Float4,
        misc -> Nullable<Text>,
    }
}

#[rocket::get("/house")]
async fn all_houses(db: Db) -> Json<Vec<House>> {
    db.run(|conn| house::table.select(house::id).load(conn))
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

pub fn stage() -> AdHoc {
    AdHoc::on_ignite("House Stage", |rocket| async {
        rocket.mount("/", routes![all_houses])
    })
}