diesel::table! {
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