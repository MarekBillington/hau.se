-- Your SQL goes here
CREATE TABLE house (
    id SERIAL PRIMARY KEY,
    created TIMESTAMPTZ NOT NULL,
    updated TIMESTAMPTZ NOT NULL,
    address VARCHAR NOT NULL,
    bedroom INT NOT NULL,
    bathroom INT NOT NULL,
    garage INT NOT NULL,
    landarea REAL NOT NULL,
    floorarea REAL NOT NULL,
    misc TEXT DEFAULT ''
)
