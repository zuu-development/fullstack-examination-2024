# Question

Please answer each question.

In addition to answering the questions, the following items will also be considered for extra points:

- Fixing test code
- Refactoring (organizing and optimizing code)
- Code formatting (structuring code)
- Adding or updating documentation(swagger etc.)

## No.1

Implement the `TestTodoHandler_FindAll` function in `internal/handler/todo_test.go`.

- Execute the `make test-backend` command and ensure all tests pass.
- Call the `FindAll` function in `internal/handler/todo.go` and perform the necessary operations.

## No.2

Enable searching by `task` and `status` values at the GET `/api/v1/todos` endpoint.

- Modify the endpoint to accept `task` and `status` query parameters.
- Execute the `make test-backend` command and ensure all tests pass.
- Update the Swagger API documentation to reflect the changes.

## No.3

Add a task sorting feature.

- Implement functionality to sort tasks by priority.

## No.4

Display completed tasks (`done` tasks) separately from incomplete tasks, at the bottom.

- Ensure tasks with a `status` of `done` are displayed separately from other tasks.
- Sorting functionality is not required for the `done` list.
