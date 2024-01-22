### TigerHall Kittens


### MySQL Database Structure:
- Database Name: `tigerkittens_db`
- Tables:
  - `user`: Stores user data.
  - `tiger`: Stores tiger list data.
  - `tiger_sightings`: Stores tiger sightings data.


### Database Migrations:
- Database migrations are enabled to manage schema changes over time.

### Dockerization:
- The application is Dockerized, ensuring consistency across environments.

### Code Architecture:
- Follows a Service, Repository, and Controller-based architecture for modularity and extensibility.

### Authentication:
- Token-based JWT authentication is implemented.
- Tokens are set as cookies and are valid for CookieExpiryTime seconds.
- Applied to create endpoints and user APIs (delete, find by user ID).

### Image Resizing:
- Tiger sighting images are resized to 250x250.

### Distance Calculation:
- Uses the Haversian formula to calculate distances within 5 km of tiger sightings.

### Pagination:
- Pagination is implemented for all GET APIs.

### Notification:
- Email notifications are sent using a Golang channel-based message queue.
- Basic implementation with simulated email client messages.

### Swagger Documentation:
- Swagger is integrated to expose API documentation for frontend developers.

### Planned Improvements:
- Auto ID generation.
- Integration with Uber logger (Zap).
- Default values for paginated APIs and optional fields in APIs.
- End-to-end code tests and coverage.
- Integration with email clients (e.g., Twilio, Mailchimp).
- Bug fixes and code cleanup.



Swagger :
http://localhost:8090/swagger/index.html#/


Use testing_instructions.md for sample input