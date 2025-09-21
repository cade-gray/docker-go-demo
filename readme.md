# docker-go-demo

A Go-based REST API for querying US license plates, jokes, and trips, built with Gin and GORM, containerized for easy deployment.  
This project is designed for learning, prototyping, and as a foundation for a production API on VPS 2.0.

## Features

- RESTful endpoints for plates, jokes, and trips
- PostgreSQL integration via GORM ORM
- Dockerized for local and remote deployment
- Environment-based configuration (dev/prod)
- Ready for deployment to GitHub Container Registry and DigitalOcean VPS

## Endpoints

- `GET /plates` — List all plates
- `GET /plates/:id` — Get plate by ID
- `GET /jokes` — List jokes
- `GET /trips` — List trips
- `POST /ping` — Test endpoint

## Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- PostgreSQL database

### Local Development

1. Clone the repo:

   ```sh
   git clone https://github.com/USERNAME/docker-go-demo.git
   cd docker-go-demo
   ```

2. Edit `compose.dev.yaml` to set your development database credentials:

   ```yaml
   services:
     server:
       build:
         context: .
         target: final
       ports:
         - 4269:4269
       environment:
         - DB_HOST=localhost
         - DB_USER=youruser
         - DB_PASSWORD=yourpass
         - DB_NAME=yourdb
         - DB_PORT=5432
         - DB_SSLMODE=disable
   ```

3. Build and run with Docker Compose:

   ```sh
   docker compose -f compose.dev.yaml up --build
   ```

4. Access the API at [http://localhost:4269](http://localhost:4269)

### Production Deployment

1. Build and push your image to GitHub Container Registry:

   ```sh
   docker build -t ghcr.io/USERNAME/docker-go-demo:latest .
   docker push ghcr.io/USERNAME/docker-go-demo:latest
   ```

2. On your VPS, edit `compose.prod.yaml` to set your production database credentials:

   ```yaml
   services:
     server:
       image: ghcr.io/USERNAME/docker-go-demo:latest
       ports:
         - 4269:4269
       environment:
         - DB_HOST=prod-db-host
         - DB_USER=produser
         - DB_PASSWORD=prodpass
         - DB_NAME=proddb
         - DB_PORT=5432
         - DB_SSLMODE=disable
       restart: always
   ```

3. Pull the image and run with Docker Compose:

   ```sh
   docker compose -f compose.prod.yaml up -d
   ```

4. Your API will be available at your VPS IP on port 4269.

## Database Setup

Create a PostgreSQL user and database:

```sql
CREATE DATABASE platefind_db;
CREATE USER gorm WITH PASSWORD 'your_secure_password';
GRANT ALL PRIVILEGES ON DATABASE platefind_db TO gorm;
```

## Project Structure

```
main.go                # Entry point
routes/                # API route handlers
Dockerfile             # Build instructions
compose.yaml           # Docker Compose config
compose.dev.yaml       # Dev Compose config
compose.prod.yaml      # Prod Compose config
.env                   # Environment variables
```

## Future Plans

- Harden for production (security, logging, error handling)
- Add authentication and rate limiting
- Expand API endpoints
- Automate deployment and scaling

---
