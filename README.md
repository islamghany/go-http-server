# How I Build HTTP Servers in Go

## 1. Introduction

In my approach to projects, the primary focus isn't on how to write a RESTful HTTP server. After all, a RESTful HTTP server is just one layer of an application. If I concentrate solely on that, I risk losing sight of the bigger picture. The main objective is to develop code that is clean and maintainable.

For example, if I decide to add support for gRPC or GraphQL in the future, I should be able to implement it without overhauling the entire application.

To facilitate this, I concentrate on the business logic and the domain model, treating the HTTP server as a layer that connects the domain model to the outside world.

## 2. Project Architecture Patterns

When starting a new project, I begin by establishing the project structure because it helps me conceptualize the project at a high level.

I strive to avoid over-engineering the project structure; my goal is to keep it simple and easy to understand.

With that in mind, I enjoy working with two main project structures:

### **HSR**: Handler, Service, Repository structure:

The Handler-Service-Repository (HSR) pattern is a common architectural approach that separates your application into three distinct layers:

- **Handler**: The handler layer is responsible for handling incoming HTTP requests and returning responses to the client. It is the entry point to your application and is responsible for parsing incoming requests, validating input, and returning responses to the client.

- **Service**: The service layer contains the business logic of your application. It is responsible for processing data, enforcing business rules, and interacting with the repository layer to persist data.

- **Repository**: The repository layer is responsible for interacting with the database. It abstracts away the details of the database implementation and provides a clean interface for the service layer to interact with.

**Pros**:

- Separation of concerns: The HSR pattern separates your application into distinct layers, each with its own responsibilities. This makes your code easier to understand, test, and maintain.
- Testability: By separating your application into distinct layers, you can easily test each layer in isolation. This makes it easier to write unit tests and integration tests for your application.
- Scalability: The HSR pattern makes it easier to scale your application by allowing you to swap out components (e.g., the database) without affecting other parts of your application.
- Maintainability: The HSR pattern makes your codebase more maintainable by enforcing a clear separation of concerns. This makes it easier to make changes to your application without affecting other parts of the codebase.

**Cons**:

- Complexity: The HSR pattern can add complexity to your application, especially for small projects. If your application is simple and does not require a lot of business logic, the HSR pattern may be overkill.

For me I tend to use the HSR pattern in large-sizes projects, because it helps me to separate the concerns and make the codebase more maintainable.
but it will be overkill for small projects and for some medium-sized projects.

### **Fat-Service**:

The Fat-Service pattern is a common architectural approach that combines the service and repository layers into a single layer called the service layer. This pattern is often used in small to medium-sized projects where the separation of concerns provided by the HSR pattern is not necessary.

Instead of separating data access (repositories) and business logic (services), the Fat Service pattern combines them. Each service is responsible for both the business logic and directly interacting with the database or storage system.

By reducing the number of layers, the architecture becomes simpler, which can be beneficial for small to medium-sized projects.

**Pros**:

- Simplicity: The Fat-Service pattern is simpler than the HSR pattern because it combines the service and repository layers into a single layer. This makes it easier to understand and work with, especially for small to medium-sized projects.

- Speed of development: The Fat-Service pattern can speed up development because it reduces the number of layers in your application. This can make it faster to implement new features and make changes to your application.

- Flexibility: The Fat-Service pattern is more flexible than the HSR pattern because it does not enforce a strict separation of concerns. This can be beneficial for projects where the separation of concerns is not necessary.

**Cons**:

- Future Growth: The Fat-Service pattern can make it harder to scale your application in the future. As your application grows, you may find it more difficult to maintain and extend your codebase.
- Team Collaboration: The Fat-Service pattern can make it harder for teams to collaborate on a project. Without a clear separation of concerns, it can be more difficult to work on different parts of the codebase in parallel.

> **Note**: I will use the Fat-Service pattern in this project because it is a small project and I don't need to separate the concerns.

## 3. Project Structure

```bash
your-app/
├── cmd/
│ └── your-app/
│ └── main.go
├── internal/
│ ├── api/
│ │ ├── api.go
│ │ └── server.go
├── middleware/
│ │ ├── error_middleware.go
│ │ └── ...
│ ├── config/
│ │ └── config.go
│ ├── handlers/
│ │ ├── user_handler.go
│ │ └── ...
│ ├── services/
│ │ ├── user_service.go
│ │ └── ...
│ ├── repository/
│ │ ├── user_repository.go
│ │ └── ...
│ ├── models/
│ │ └── user_model.go
│ ├── validator/
│ │ └── validator.go
│ └── web/
│ └── web.go
├── pkg/
│ └── (optional shared packages)
├── db/
│ ├── migrations/
│ │ └── (database migration files)
│ └── sqlc.yaml
├── docs/
│ └── (API documentation)
├── go.mod
└── go.sum
```

- **cmd**: Contains the main applications for your project.
- **internal**: Holds internal packages that are not intended to be used by external applications or libraries.
- **pkg**: Contains library code that's ok to use by external applications.
- **api**: Contains the API layer, which is responsible for handling HTTP requests and returning responses.
- **middleware**: Contains middleware functions that can be used by the API layer.
- **config**: Contains configuration logic.
- **handlers**: Contains handler functions that are called by the API layer to process HTTP requests.
- **services**: Contains service functions that contain the business logic of your application.
- **repository**: Contains repository functions that interact with the database.
- **models**: Contains data models used by your application.
- **validator**: Contains validation logic.
- **web**: Contains web server logic.
- **db**: Contains database schema and migration files.
- **docs**: Contains API documentation.
