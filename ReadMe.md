# Go Application with PostgreSQL

This README provides instructions on how to set up and run a Go application that connects to a PostgreSQL database using Docker Compose. The application and database are containerized, making it easy to deploy and manage.

## Prerequisites

- Docker: Ensure Docker is installed on your machine.
- Docker Compose: Ensure Docker Compose is installed on your machine.

## Getting Started

1. **Clone the Repository**: Clone this repository to your local machine.

```bash
git clone <repository-url> cd <repository-name>
```

2. **Build and Run the Application**: Use Docker Compose to build and run the application along with the PostgreSQL database.

```bash
docker-compose up --build
```

## Accessing the Application

Once the containers are up and running, you can access the Go application by navigating to `http://localhost:8080` in your web browser.

## Database Access

The PostgreSQL database is accessible at `localhost:5432`.

## Stopping the Application

To stop the application and the PostgreSQL database, run:

```bash
docker-compose down
```

This command will stop and remove the containers, networks, and volumes defined in the `docker-compose.yml` file.
