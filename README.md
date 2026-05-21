# Contact Management API

A RESTful Contact Management API built using Golang, Fiber, PostgreSQL, and GORM.

The project was developed as part of a backend developer assessment with a focus on clean architecture, maintainable code structure, request validation, and scalable API design.

---

## Features

### Contact Management
- Create a new contact
- Fetch all contacts
- Fetch contact by ID
- Update contact details
- Soft delete contacts

### Query Support
- Pagination
- Search by name or email
- Filter by status
- Sorting support

### Validation & Error Handling
- Required field validation
- Email format validation
- Duplicate active email restriction
- Standardized API responses
- Proper HTTP status codes

### Additional Features
- Swagger/OpenAPI documentation
- Request logging middleware
- Rate limiting middleware
- Global error handling
- Environment-based configuration
- Dockerized PostgreSQL setup

---

## Tech Stack

- Golang
- Fiber
- PostgreSQL
- GORM
- Swagger
- Docker

---

## Project Structure

```text
cmd/
    server/
        main.go

config/
    config.go
    database.go

internal/
    contact/
        delivery/http/
        dto/
        entity/
        repository/
        service/
        validator/

    middleware/

    shared/
        pagination/
        response/

docs/

The project follows a layered architecture approach to keep business logic, database operations, request handling, and validations properly separated.

Getting Started
1. Clone the Repository
git clone <your-repository-url>
cd Contact-Management-API
2. Install Dependencies
go mod tidy
3. Start PostgreSQL Container
docker compose up -d
4. Configure Environment Variables

Create a .env file in the project root.

Example:

APP_NAME=contact-management-api
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=contact_db
DB_SSLMODE=disable
5. Generate Swagger Docs
swag init -g cmd/server/main.go
6. Run the Application
go run cmd/server/main.go

Server runs on:

http://localhost:8080
Swagger Documentation

Swagger UI:

http://localhost:8080/swagger/index.html
API Endpoints
Method	Endpoint	Description
POST	/api/v1/contacts	Create contact
GET	/api/v1/contacts	List contacts
GET	/api/v1/contacts/{id}	Get contact by ID
PUT	/api/v1/contacts/{id}	Update contact
DELETE	/api/v1/contacts/{id}	Soft delete contact
Query Parameters
List Contacts
Parameter	Description
page	Page number
limit	Number of records per page
search	Search by first name, last name, or email
status	Filter by Active / Inactive
sort_by	Sort field
order	asc / desc

Example:

GET /api/v1/contacts?page=1&limit=10&search=anvar&status=Active&sort_by=created_at&order=desc
Sample Request
Create Contact
POST /api/v1/contacts

Request Body:

{
  "first_name": "Anvar",
  "last_name": "Sha",
  "email": "anvar@example.com",
  "phone_number": "9876543210",
  "company_name": "Moveon",
  "status": "Active"
}
Sample Success Response
{
  "success": true,
  "message": "contact created successfully",
  "data": {
    "id": 1,
    "first_name": "Anvar",
    "last_name": "Sha",
    "email": "anvar@example.com"
  }
}
Sample Error Response
{
  "success": false,
  "message": "active contact with this email already exists"
}
Technical Decisions
Clean architecture was used to separate handlers, services, repositories, and DTOs.
Soft delete support was implemented using GORM's DeletedAt.
DTOs were used to avoid coupling API payloads directly with database entities.
Email normalization and duplicate active email validation were handled at the service layer.
Query parameters were validated before database operations.
Sorting fields were restricted to avoid unsafe query construction.
Future Improvements
Unit tests and integration tests
JWT authentication
Redis caching
CI/CD pipeline
Multi-stage Docker builds
Database indexing optimizations
Estimated Time Spent

Approximately 10-12 hours.