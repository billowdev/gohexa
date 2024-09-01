# WORK IN PROGRESS

# gohexa - Golang Hexagonal Directory Template

- hexagonal

# Directory
hexagonal/
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   └── core/
└── README.md


# Hexagonal 

This project follows the Hexagonal Architecture (also known as Ports and Adapters Architecture), which aims to create a clear separation between the core business logic and external systems. The structure of the project is designed to ensure maintainability, testability, and scalability by decoupling various components.

## Project Structure

`cmd`
Contains the main application entry points. This is where the application starts and is responsible for setting up the application context, configuration, and running the server.

`internal`
The core of the application is divided into several packages:
- `adapters`: Contains the implementations for interacting with external systems.
  -  `app`: Application-level adapters.
  -  `database`: Database-related adapters, including models and seeders.
     - `models`: Database models used for ORM.
     - `seeders`: Scripts for populating the database with initial data.
   - `http`: HTTP-related adapters, including handlers and routers.
     - `handlers`: HTTP request handlers, organized by domain (e.g., documents, orders, system_fields).
     - `routers`: Routing configurations for handling HTTP requests and responses.
   - `repositories`: Data access repositories, organized by domain (e.g., documents, orders, system_fields).
 - `core`: Contains the business logic and domain-specific code.
	- `domain`: Domain models and business entities, organized by domain (e.g., documents, orders, system_fields).
	- `ports`: Interfaces that define the interactions between the core domain and the external systems.
    	- documents, master_files, orders, system_fields: Ports for different domains.

  - `services`: Business logic and use cases, organized by domain (e.g., documents, orders, system_fields).


`pkg`
Contains shared utilities and helpers that can be used across the application.

- `configs`: Configuration management.
- `helpers`: Utility functions for various tasks.
- `filters`: Helpers for filtering data.
- `pagination`: Helpers for pagination.
- `utils`: Additional utility functions, such as Argon2ID for password hashing



## Hexagonal Architecture
### Core Concepts
- Domain: The core business logic and rules are encapsulated within the core package. This includes the domain models, services, and ports.
- Ports: Interfaces in the core/ports package define the required interactions between the domain and external systems.
Adapters: The internal/adapters package contains implementations that interact with external systems (e.g., HTTP handlers, database repositories). These adapters convert data to and from the domain format.

### How It Works
- Application Initialization: The cmd package initializes the application, setting up necessary configurations and dependencies.

- Business Logic: The core business logic is contained in the core package. It remains independent of external systems and is only influenced by the interfaces (ports) it defines.

- Interaction with External Systems: Adapters in the internal/adapters package implement the interfaces defined in the core/ports package. These adapters handle data persistence (e.g., in databases) and external communication (e.g., HTTP requests).

- Routing and Handlers: The http/routers package sets up the routing for incoming HTTP requests. The http/handlers package contains the logic for processing these requests and invoking the appropriate business logic.

- Frontend: The views package contains the front-end code, including static assets and UI components.


## Read more


- Kyodo Tech [Pseudonym]. (2019, August 20). Hexagonal Architecture, Ports and Adapters in Go. Medium. https://medium.com/@kyodo-tech/hexagonal-architecture-ports-and-adapters-in-go-f1af950726b
- Glushach, R. (2020, August 12). Hexagonal Architecture, the Secret to Scalable and Maintainable Code for Modern Software. Medium. https://romanglushach.medium.com/hexagonal-architecture-the-secret-to-scalable-and-maintainable-code-for-modern-software-d345fdb47347   
- Wikipedia. (2024, August 25). Hexagonal architecture (software). https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)   
- Hexagonal Architecture. AWS Prescriptive Guidance. https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/hexagonal-architecture.html
Wom, F. (n.d.). go-hexagonal. GitHub. Retrieved [date you accessed the repository], from https://github.com/felipewom/go-hexagonal
- Bagash Izuddin (bagashiz). (2021, September 29). Building RESTful API with Hexagonal Architecture in Go. DEV Community. https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij
- Part 1. Medium. https://medium.com/@pthtantai97/hexagonal-architecture-with-golang-part-1-7f82a364b29
