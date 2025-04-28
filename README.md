# Person Enrichment API

**Person Enrichment API** is a REST API service developed in Go that provides functionality for working with person data. The service allows creating, reading, updating, and deleting person records, as well as automatically enriching them with additional information such as age, gender, and nationality.

## Technologies Used
- **Go**: Main programming language for API development
- **Gin**: Web framework for creating REST API
- **PostgreSQL**: Database for storing person information
- **Swagger**: API documentation generation and display
- **Slog**: Structured logging
- **Docker**: Application containerization
- **Godotenv**: Environment variables reading and processing
- **Validator**: Input data validation
- **Migrate**: Database migrations management

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/soloda1/person-enrichment-api.git
    cd person-enrichment-api
    ```

2. **Set up environment variables**:
    Create a `.env` file in the root directory using `example.env` as a template:
    ```sh
    cp example.env .env
    ```

3. **Run using Docker**:
    ```sh
    docker-compose up --build
    ```

## API Endpoints

### People
- `POST /create` - Create a new person
- `GET /person/{id}` - Get person information by ID
- `GET /persons` - Get list of people with filtering
- `PUT /update` - Update person information
- `DELETE /delete/{id}` - Delete person by ID

### Health Check
- `GET /ping` - API health check

## API Documentation
API documentation is available at:
```
http://localhost:8080/swagger/index.html
```

## Project Structure
```
.
├── cmd/                 # Application entry point
├── config/             # Application configuration
├── docs/               # Generated Swagger documentation
├── external/           # External services and integrations
├── internal/           # Internal application logic
│   ├── api/           # API routing and middleware
│   ├── handlers/      # HTTP request handlers
│   ├── models/        # Data models
│   ├── repository/    # Database operations
│   ├── service/       # Business logic
│   └── utils/         # Utility functions
├── migrations/         # Database migrations
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile         # Docker configuration
├── example.env        # Environment variables example
├── go.mod             # Go dependencies
└── go.sum             # Dependency checksums
```

## Open Source
This project is open source. You are free to use, study, and contribute to its development.

## Project Status
The project is under active development and will continue to improve. 