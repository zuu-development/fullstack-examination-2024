# Question

Please answer each question.

In addition to answering the questions, the following items will also be considered for extra points:

- Refactoring (organizing and optimizing code)
- Code formatting (structuring code)
- Adding or updating documentation (Swagger, etc.)

## No.1

Enable searching by `task` and `status` values at the GET `/api/v1/todos` endpoint.

- Modify the endpoint to accept `task` and `status` query parameters.
- Update the Swagger API documentation to reflect the changes.

## No.2

Add a task sorting feature. Currently, tasks are not sorted in any meaningful way in the UI. Implement sorting by priority.

- Add a field to the Todo model to represent priority, as it currently lacks one.
- Implement functionality to sort tasks by priority.
- Ensure higher priority tasks are sorted to the top.

## No.3

Display completed tasks (`done` tasks) separately from incomplete tasks, at the bottom.

- Ensure tasks with a `status` of `done` are displayed separately from other tasks.
- Sorting functionality is not required for the `done` list.
- Tasks changed to `done` status should be displayed in the bottom list.
- Tasks changed from `done` to another status should be displayed in the top list.
