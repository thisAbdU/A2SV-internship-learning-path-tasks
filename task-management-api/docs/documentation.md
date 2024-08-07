## API Documentation

### Authentication Routes

#### Register
- **Endpoint**: `POST /auth/register`
- **Description**: Registers a new user.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "email": "string"
  }
  ```
- **Response**:
  - **Success (201 Created)**: 
    ```json
    {
      "message": "User registered successfully"
    }
    ```
  - **Error (400 Bad Request)**: 
    ```json
    {
      "error": "Detailed error message"
    }
    ```

#### Login
- **Endpoint**: `POST /auth/login`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "token": "string"
    }
    ```
  - **Error (401 Unauthorized)**: 
    ```json
    {
      "error": "Invalid username or password"
    }
    ```

### Task Management Routes

#### Get Tasks
- **Endpoint**: `GET /task/`
- **Description**: Retrieves a list of all tasks.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    [
      {
        "id": "string",
        "title": "string",
        "description": "string",
        "status": "string"
      }
    ]
    ```
  - **Error (500 Internal Server Error)**: 
    ```json
    {
      "error": "Detailed error message"
    }
    ```

#### Create Task
- **Endpoint**: `POST /task/`
- **Description**: Creates a new task.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "status": "string"
  }
  ```
- **Response**:
  - **Success (201 Created)**: 
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```
  - **Error (400 Bad Request)**: 
    ```json
    {
      "error": "Detailed error message"
    }
    ```

#### Get Task by ID
- **Endpoint**: `GET /task/:id`
- **Description**: Retrieves a task by its ID.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "Task not found"
    }
    ```

#### Update Task
- **Endpoint**: `PATCH /task/:id`
- **Description**: Updates an existing task by its ID.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "status": "string"
  }
  ```
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "string"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "Task not found"
    }
    ```

#### Delete Task
- **Endpoint**: `DELETE /task/:id`
- **Description**: Deletes a task by its ID.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "message": "Task deleted successfully"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "Task not found"
    }
    ```

### User Management Routes

#### Get Users
- **Endpoint**: `GET /`
- **Description**: Retrieves a list of all users.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    [
      {
        "id": "string",
        "username": "string",
        "email": "string"
      }
    ]
    ```
  - **Error (500 Internal Server Error)**: 
    ```json
    {
      "error": "Detailed error message"
    }
    ```

#### Get User by ID
- **Endpoint**: `GET /:id`
- **Description**: Retrieves a user by their ID.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "id": "string",
      "username": "string",
      "email": "string"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "User not found"
    }
    ```

#### Update User
- **Endpoint**: `PATCH /:id`
- **Description**: Updates an existing user by their ID.
- **Request Body**:
  ```json
  {
    "username": "string",
    "email": "string"
  }
  ```
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "id": "string",
      "username": "string",
      "email": "string"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "User not found"
    }
    ```

#### Delete User
- **Endpoint**: `DELETE /:id`
- **Description**: Deletes a user by their ID.
- **Response**:
  - **Success (200 OK)**: 
    ```json
    {
      "message": "User deleted successfully"
    }
    ```
  - **Error (404 Not Found)**: 
    ```json
    {
      "error": "User not found"
    }
    ```

### Middleware

- **AuthMiddleware**: This middleware ensures that the user is authenticated before accessing certain routes. It is used for routes that require user authentication.
