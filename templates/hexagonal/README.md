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



# The design pattern for repositories adapters
Three different approaches to handling transactions across multiple repositories
let's example business context in `Order` and `Document`
- Abstract Repository Combining Two/Multiple Repositories
- Transactor Repository (TransactorRepo)
- Transaction at the Repository Level (BeginTransaction Method)
Let's explore the pros and cons of each, and I'll provide guidance on which might be the most suitable...

## Approach 1: Abstract Repository Combining Two/Multiple Repositories (OrderDocumentRepo)
In this approach, you create an OrderDocumentRepo that combines both the OrderRepo and DocumentRepo into a single abstract repository, managing the transaction within this repository.

`Pros`:

- Single Responsibility: Encapsulates the entire transaction logic related to both orders and documents within a single repository, making it easy to manage.
- Consistency: Ensures that both order and document operations are tightly coupled within the same transaction scope, reducing the risk of partial failures.

`Cons`:

- Complexity: The repository becomes more complex as it handles more than one entity, potentially leading to a violation of the Single Responsibility Principle.
Tight Coupling: Tightly couples OrderRepo and DocumentRepo, making it harder to reuse or test them independently.
- Scalability: If more repositories need to be added to the transaction, the OrderDocumentRepo would have to be modified, reducing flexibility.

`Best Practice`:

- This approach is suitable when the operations on orders and documents are always tied together and never operate independently.
- It can be useful when you need to ensure consistency between two repositories that are closely related and where changes are always made in tandem.

## Approach 2: Transactor Repository (TransactorRepo)
In this approach, you define a TransactorRepo that provides a transaction context that can be passed to multiple repositories. The transaction is managed at the service level.


`Pros`:

- Loose Coupling: Keeps repositories independent from each other. The transaction management is handled separately, allowing repositories to remain focused on their own responsibilities.
- Flexibility: Easily extendable to support more repositories without modifying existing ones.
Testability: Easier to test individual repositories independently, as they don't depend on each other.

`Cons`:

- Complexity: Slightly more complex as it requires managing transaction context explicitly in the service layer.
- Potential for Misuse: Requires careful management of transaction context to ensure it's passed correctly between repositories, which can be error-prone.

`Best Practice`:

- This approach is a good fit for a Hexagonal Architecture as it aligns well with dependency injection and separation of concerns. It allows you to inject the TransactorRepo into your service layer, making it flexible and scalable.
- It’s suitable for more complex systems where multiple repositories need to be part of a transaction but should remain loosely coupled.


## Approach 3: Transaction at the Repository Level (BeginTransaction Method)

Here, each repository has its own BeginTransaction method, and the service manages the transaction by passing the transaction context to the repositories.

`Pros`:

- Granularity: Allows fine-grained control over transactions, enabling different repositories to manage their own transactions.
- Direct Control: Gives the service layer direct control over the transaction process, which can be useful in certain scenarios.

`Cons`:
- Code Duplication: Each repository needs to implement its own transaction management logic, leading to potential code duplication.
- Complexity: Increases the complexity of the service layer, as it must coordinate transaction management across multiple repositories.
Risk of Inconsistency: Higher risk of inconsistencies if the transaction context is not managed properly across repositories.

`Best Practice`:

This approach is less common in Hexagonal Architecture because it mixes concerns and adds complexity. It's usually better to centralize transaction management rather than spreading it across repositories.
Suitable only in cases where each repository truly needs its own independent transaction management.



## Recommendation

Because this is the hexagonal architecture tutorial we will use the Approach 2 because of this below ...

Given your use of Hexagonal Architecture and dependency injection, Approach 2 (Transactor Repository) is generally the best practice:

- Decoupling: It keeps your repositories focused on their primary responsibilities without adding the burden of transaction management.
- Flexibility and Scalability: The transaction logic is centralized, making it easier to manage and extend to additional repositories if needed.
- Ease of Testing: By isolating transaction logic, you can more easily test your repositories and service logic independently.

This approach aligns well with the principles of Hexagonal Architecture, where the business logic (services) is kept separate from infrastructure concerns (like database transactions), and dependencies are injected where needed.

In summary, Approach 2 offers the right balance of flexibility, maintainability, and scalability, making it the most suitable choice for your scenario.


### Transactor

- `WithTransactionContextTimeout` Function
  
```go
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// WithTransactionContextTimeout executes a function within a transaction with a specified context timeout.
// The transaction is committed if successful, or rolled back if an error occurs or the context times out.
func (d *TransactorImpl) WithTransactionContextTimeout(ctx context.Context, timeout time.Duration, tFunc func(ctx context.Context) error) error {
	// Create a new context with timeout
	transactionCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Start a new transaction
	tx, err := d.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Ensure that the transaction is finalized properly
	defer func() {
		select {
		case <-transactionCtx.Done():
			// Rollback if the transaction context is done (timeout or cancel)
			if rollbackErr := d.RollbackTransaction(tx); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		default:
			// Commit if no error and context is still valid
			if commitErr := tx.Commit().Error; commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
				err = commitErr
			}
		}
	}()

	// Run the callback function with the transaction context
	err = tFunc(InjectTx(transactionCtx, tx))
	if err != nil {
		tx.Error = err // Mark the transaction as needing a rollback
		return err
	}

	return nil
}
```

- Explanation
1. Context with Timeout:
- Creates a new context with a specified timeout using `context.WithTimeout`. This context will be used for the transaction operations.
2. Start Transaction:
- Begins a new transaction with `BeginTransaction`.
3. Deferred Finalization:

- Uses `defer` to ensure the transaction is finalized correctly:
  - Rollback on Timeout or Cancellation: Rolls back the transaction if the context is canceled or times out.
  - Commit on Success: Commits the transaction if it completes successfully within the timeout period.
4. Function Execution:
  Executes the provided function `tFunc` within the transaction context, passing the transaction as part of the context using `InjectTx`.
5. Error Handling:
- Sets the transaction error to indicate a rollback is needed if tFunc returns an error.
- Logs errors during commit and rollback operations.


- Usage Example
```go
// CreateSeaPort attempts to create a new sea port entry within a transaction.
// It uses the WithTransactionContextTimeout function to ensure that the transaction
// is managed properly with a specified timeout.
func (s *SeaPortServicesImpl) CreateSeaPort(ctx context.Context, payload domain.SeaPortDomain) utils.APIV2Response {
    var result utils.APIV2Response

    // Use the WithTransactionContextTimeout function to handle the transaction
    err := s.transactor.WithTransactionContextTimeout(ctx, 5*time.Second, func(txCtx context.Context) error {
        // Convert the domain payload to a model suitable for database operations
        data := domain.ToSeaPortModel(payload)
        
        // Create the sea port entry in the database using the transaction context
        if err := s.repo.CreateSeaPort(txCtx, data); err != nil {
            return err // Return error to trigger rollback
        }
        
        // If no errors occurred, prepare a successful response
        result = utils.APIV2Response{
            StatusCode:    utils.API_SUCCESS_CODE,
            StatusMessage: "Success",
            Data:          domain.ToSeaPortDomain(data),
        }
        return nil // Indicate success
    })

    // Check if there was an error during the transaction
    if err != nil {
        // Prepare an error response if something went wrong
        result = utils.APIV2Response{
            StatusCode:    utils.API_ERROR_CODE,
            StatusMessage: "Error",
            Data:          err,
        }
    }

    // Return the result of the transaction
    return result
}
```


#### Using WithinTransaction:

```
func (s *SeaPortServicesImpl) CreateSeaPort(ctx context.Context, payload domain.SeaPortDomain) utils.APIV2Response {
	var result utils.APIV2Response

	// Create a context with a timeout for the transaction
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Use the WithinTransaction function to handle the transaction
	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		// Convert the domain payload to a model suitable for database operations
		data := domain.ToSeaPortModel(payload)

		// Create the sea port entry in the database using the transaction context
		if err := s.repo.CreateSeaPort(txCtx, data); err != nil {
			return err // Return error to trigger rollback
		}

		// If no errors occurred, prepare a successful response
		result = utils.APIV2Response{
			StatusCode:    utils.API_SUCCESS_CODE,
			StatusMessage: "Success",
			Data:          domain.ToSeaPortDomain(data),
		}
		return nil // Indicate success
	})

	// Check if there was an error during the transaction
	if err != nil {
		// Prepare an error response if something went wrong
		result = utils.APIV2Response{
			StatusCode:    utils.API_ERROR_CODE,
			StatusMessage: "Error",
			Data:          err,
		}
	}

	// Return the result of the transaction
	return result
}
```


#### Key Differences WithinTransaction & WithTransactionContextTimeout


1. Timeout Management:

- `WithTransactionContextTimeout`: Manages both transaction and context timeout internally.
- `WithinTransaction`: Requires separate context timeout management; transaction timeout is handled externally.
2. Flexibility:
- `WithTransactionContextTimeout`: Less flexible for cases where you need different timeout values for transactions and overall request processing.
- `WithinTransaction`: Provides more flexibility by allowing you to set a timeout for the transaction context separately.
3. Complexity:
- `WithTransactionContextTimeout`: Combines timeout handling with transaction management, which may be simpler but less flexible.
- `WithinTransaction`: Keeps transaction management separate, requiring explicit timeout management but offering more control.
4. Error Handling:
- Both methods handle errors similarly by triggering rollbacks on failures. The primary difference is how the timeouts are managed and handled.
Summary
If you need to manage timeouts for transactions separately from the overall request timeout, using `WithinTransaction` with an explicitly managed context timeout allows more flexibility and control. If you prefer to encapsulate timeout management within the transaction handling logic itself, `WithTransactionContextTimeout` may simplify your code but with less flexibility.


