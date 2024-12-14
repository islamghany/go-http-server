# Service Layer (`services` package)

- Receives data from the handler.
- Performs business logic, validation, and any necessary computations.
- Calls the appropriate method on the repository layer.
- Returns the result back to the handler.

This is considered a FAT SERVICE pattern, where the service layer is responsible for the majority of the business logic. This is in contrast to a thin service pattern, where the service layer is responsible for only coordinating the business logic.
