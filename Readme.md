
# TigerHall Kittens Documentation

## MySQL Database Structure:
- **Database Name:** `tigerkittens_db`
- **Tables:**
  - `user`: Stores user data.
  - `tiger`: Stores tiger list data.
  - `tiger_sightings`: Stores tiger sightings data.

## Database Migrations:
- Database migrations are enabled to manage schema changes over time.

## Dockerization:
- The application is Dockerized, ensuring consistency across environments.

## Code Architecture:
- Follows a modular architecture with Service, Repository, and Controller components.
- Designed for modularity and extensibility.

## Authentication:
- Token-based JWT authentication is implemented.
- Tokens are set as cookies and remain valid for a specified `CookieExpiryTime`.
- Applied to create endpoints and user APIs (delete, find by user ID).

## Image Resizing:
- Tiger sighting images are automatically resized to 250x250 for consistent display.

## Distance Calculation:
- Utilizes the Haversian formula to calculate distances within a 5 km radius of tiger sightings.

## Pagination:
- Implements pagination for all GET APIs, enhancing data retrieval efficiency.

## Notification:
- Email notifications are triggered using a Golang channel-based message queue.
- Basic implementation with simulated email client messages.

## Swagger Documentation:
- Swagger is integrated to provide API documentation for frontend developers.
- Access Swagger documentation at [http://localhost:8090/swagger/index.html#/](http://localhost:8090/swagger/index.html#/).

## Error Handling:
- Handles invalid image uploads or other format issues with error code 4001.
- Error message: "Error during photo upload and resizing."
{
  "error_code": 4001,
  "error_message": "Error during photo upload and resizing"
}
{
  "error_code": 4002,
  "error_message": "Tiger sighting within 5 km"
}

## Setup:
1. Initialize Docker daemon.
2. Run the following command to build and start the application:
    ```bash
    docker-compose up --build
    ```
   
## Planned Improvements:
- Auto ID generation.
- Integration with Uber logger (Zap).
- Default values for paginated APIs and optional fields.
- Implement end-to-end code tests and coverage.
- Integration with additional email clients (e.g., Twilio, Mailchimp).
- Ongoing bug fixes and code cleanup.

---
