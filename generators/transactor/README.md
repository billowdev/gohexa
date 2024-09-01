## Transactor File Generator

### Overview
The Transactor File Generator is a command-line tool that creates a transactor.go file in the specified output directory. This file contains a set of transaction management utilities for use with the Gorm ORM in a hexagonal architecture.

### Features
- Generates a transactor.go file with pre-defined transaction management functions.
- Supports injecting, extracting, and managing database transactions within a context.
- Provides functions for handling transactions with timeout contexts.

### Usage
#### Flags
`-output <OutputDirectory>`: The output directory where the transactor.go file will be generated.

### Command
To generate the transactor.go file, run the following command:

```bash
go run init_transactor.go -generate transactor -output <OutputDirectory>
```

### Example Command
Generate the transactor.go file in the ./database directory:

```bash
go run init_transactor.go -generate transactor -output ./database
```

### Output
The tool generates a file named transactor.go in the specified directory. The generated file includes:

- Transaction Management: Functions for beginning, committing, and rolling back transactions.
- Context Injection/Extraction: Utilities for injecting and extracting the transaction from the context.
- Timeout Handling: Functions for executing transactions with a context timeout.

### Generated File Structure
The transactor.go file includes the following structure:

- Package: database
- Transaction Utilities:
	- InjectTx(ctx context.Context, tx *gorm.DB) context.Context
	- ExtractTx(ctx context.Context) *gorm.DB
	- HelperExtractTx(ctx context.Context, db *gorm.DB) *gorm.DB

- Transaction Management:
	- BeginTransaction() (*gorm.DB, error)
	- RollbackTransaction(tx *gorm.DB) error
	- WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	- WithTransactionContextTimeout(ctx context.Context, timeout time.Duration, tFunc func(ctx context.Context) error) error
- Interfaces:
	- IDatabaseTransactor

### Notes
Ensure that the output directory exists or the tool will create it.
The generated transactor.go file is designed to work with the Gorm ORM in a hexagonal architecture.