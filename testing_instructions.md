{
  "date_of_birth": "12-12-2001",
  "last_seen_coordinates_lat": 22.2,
  "last_seen_coordinates_lon": 23.2,
  "last_seen_timestamp": "2023-02-02 10:00:02",
  "name": "Tiger Bengal",
  "tiger_id": "Tiger001"
}



login
{
  "email": "abc125@gmail.com",
  "password": "abc125"
}


signup
{
  "email": "abc125@gmail.com",
  "password": "abc125",
  "user_id": "abc125",
  "user_name": "abc125"
}


//create tiger sigthing

curl -X 'POST' \
  'http://localhost:8090/tigerSighting/v1/create_new' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'photo=@tigerimg.jpg;type=image/jpeg' \
  -F 'latitude=22112.1' \
  -F 'longitude=221.1' \
  -F 'sighting_id=TIGER_SIGH011' \
  -F 'tiger_id=Tiger003' \
  -F 'user_id=abc123' \
  -F 'timestamp=2023-02-01 10:00:02'


  create new tigerid each time and change the latitilutes  each time keep the userid same or change it upload a random image