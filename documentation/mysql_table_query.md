
CREATE TABLE IF NOT EXISTS tiger (
    tiger_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date_of_birth VARCHAR(10) NOT NULL,
    last_seen_timestamp VARCHAR(19) NOT NULL, -- Assuming the timestamp format is "yyyy-mm-dd HH:ii:ss"
    last_seen_coordinates_lat FLOAT NOT NULL,
    last_seen_coordinates_lon FLOAT NOT NULL
);

drop table tiger_sighting_data;


CREATE TABLE IF NOT EXISTS tiger_sighting_data (
    sighting_id VARCHAR(255) PRIMARY KEY,
    tiger_id VARCHAR(255) NOT NULL,
    latitude DOUBLE NOT NULL,
    longitude DOUBLE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    timestamp VARCHAR(19) NOT NULL,
    sighting_image TEXT,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE IF NOT EXISTS user (
    user_id VARCHAR(255) PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);





