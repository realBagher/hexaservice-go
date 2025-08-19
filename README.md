# hexaservice-go

**hexaservice-go** is a Go-based microservices architecture demonstration that implements an academic publishing system using hexagonal architecture (ports and adapters pattern). The project consists of two main services: a **Journal service** that manages academic journals (with properties like name, description, and impact factor) and an **Article service** that manages research articles (with title, abstract, author, and journal references). The Journal service exposes a gRPC API for inter-service communication, while both services support multiple storage adapters (in-memory and MySQL repositories) that can be swapped without changing the core business logic. The Article service communicates with the Journal service via gRPC to fetch journal information, demonstrating a clean separation between business logic (core), external interfaces (ports), and infrastructure implementations (adapters) - making the system highly testable, maintainable, and adaptable to different storage backends or communication protocols.

## Architecture

The project follows the **Hexagonal Architecture** (also known as Ports and Adapters) pattern:

- **Core**: Contains business logic and domain models
- **Ports**: Define interfaces for external interactions
- **Adapters**: Implement the ports for specific technologies (MySQL, in-memory storage, gRPC)

## Services

### Journal Service
- Manages academic journals with impact factors
- Supports both in-memory and MySQL storage

### Article Service  
- Manages research articles
- Communicates with Journal service via gRPC
- Supports both in-memory and MySQL storage

## Prerequisites

- Go 1.19 or higher
- MySQL (optional - services will fall back to in-memory storage)
- Protocol Buffers compiler (if you want to compile prtoc yourself)

## Running the Services

### Option 1: Using In-Memory Storage (Default)

1. **Start the Journal Service:**
   ```bash
   cd journal
   go run main.go
   ```

2. **Start the Article Service (in a new terminal):**
   ```bash
   cd article  
   go run main.go
   ```

### Option 2: Using MySQL Storage

1. **Set up MySQL database and set the environment variable:**
   ```bash
   export MYSQL_DSN="username:password@tcp(localhost:3306)/database_name"
   ```

2. **Start the Journal Service:**
   ```bash
   cd journal
   go run main.go
   ```

3. **Start the Article Service (in a new terminal):**
   ```bash
   cd article
   go run main.go  
   ```

## What Happens When You Run

1. **Journal Service** starts a gRPC server on port 50051 and demonstrates CRUD operations
2. **Article Service** demonstrates article operations and communicates with the Journal service via gRPC to fetch journal information
3. Both services will show demo output in the console, displaying created and retrieved records
4. If MySQL is configured, both services will use persistent storage; otherwise, they fall back to in-memory storage



