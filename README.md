# Glofox API

This Glofox API is built using a layered, modular design with the [Gin](https://gin-gonic.com/) web framework.

## Architecture

The application is organized into the following packages, each with a specific responsibility:

* **`models`**: Defines the core data structures used throughout the application, such as `Class` and `Booking`.
* **`datastore`**: Provides an abstraction layer for data access. Currently, it uses an in-memory implementation (`inmemory.go`) optimized for concurrent read operations using `sync.RWMutex`. This allows multiple read requests to be processed simultaneously without blocking.
* **`services`**: Contains the application's business logic. These services orchestrate the interaction between the handlers and the data store, performing tasks like creating classes and handling bookings.
* **`handlers`**: Responsible for handling incoming HTTP requests using the Gin framework. They parse request data, perform basic validation, and delegate the core business logic to the `services` layer. Finally, they format the responses sent back to the client.
* **`main`**: The entry point of the application. It sets up the Gin router, initializes the necessary dependencies (data store, services, and handlers), and defines the API routes.

## Data Flow

The general flow of a request through the application is as follows:

1.  **Request In:** An HTTP request is received by the Gin router defined in `main.go` and handled by the appropriate function within the `handlers` package.
2.  **Business Logic:** The handler calls the corresponding method in the `services` package to perform the core business operation.
3.  **Data Access:** The service interacts with the `datastore` (currently the in-memory store) to retrieve or persist data.
4.  **Response Out:** The service returns the result to the handler, which then formats it into an HTTP response (typically JSON) using Gin.

This layered and modular design promotes a clear separation of concerns, making the codebase more organized, maintainable, and testable.
