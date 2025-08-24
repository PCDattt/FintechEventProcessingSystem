# FintechEventProcessingSystem

Go-based backend for fintech flows (**user signup**, **deposit**, **withdraw**, **payment**) using:
- **Gin** lightweight HTTP framework used to expose REST APIs for user management.
- **gRPC** efficient, low-latency communication between internal services.
- **PostgreSQL** relational database for storing user accounts, balances, and transaction history with ACID guarantees.
- **RabbitMQ** message broker for asynchronous task processing and event-driven workflows.
- **Grafana** visualization layer for logs and metrics dashboards.
- **Prometheus** time-series database used to scrape, store, and query application metrics.
- **Loki** log aggregation system, integrated with Grafana for centralized log search.

## Prerequisites

Before develop the project, ensure the following tools are installed:

- **[Docker](https://docs.docker.com/get-docker/)** – to run PostgreSQL, RabbitMQ, Grafana, Prometheus, Loki.
- **[Docker Compose](https://docs.docker.com/compose/)** – to start all containers with one command.
- **[Make](https://www.gnu.org/software/make/)** – to run project build and setup commands.
- **[protoc](https://grpc.io/docs/protoc-installation/)** - protocol buffers compiler.
- **[golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)** – database migrations.
- **[sqlc](https://sqlc.dev/)** – generates type-safe Go code from SQL queries.

## Sequence Diagram
![Sequence Diagram](images/sequence_diagram.png)
