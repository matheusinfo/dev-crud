
# Dev CRUD API

This API allows you to manage developers. You can create, list, retrieve by ID, and remove developers.

## Technologies Used

- **Go (Golang)**
- **Gin Framework**
- **GORM (ORM)**
- **MySQL** (or another configured database)
- **Swagger** (for API documentation)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/dev-crud.git
    ```

2. Navigate to the project directory:
    ```bash
    cd dev-crud
    ```

3. Install dependencies:
    ```bash
    go mod tidy
    ```

4. Set up your `.env` file for environment variables:
    ```bash
    PORT=8080
    DB_STRING="your-database-connection-string"
    ```

5. Run the application:
    ```bash
    go run main.go
    ```

## API Documentation

### Base URL

```
/api/v1
```

### 1. **Create a Dev**
- **URL:** `/dev`
- **Method:** `POST`
- **Description:** Create a new developer.
- **Request Body:**
  ```json
  {
    "id": "string",
    "username": "string",
    "email": "string"
  }
  ```
- **Response:**
  - **201 Created**
    ```json
    {
      "id": "string",
      "username": "string",
      "email": "string"
    }
    ```
  - **400 Bad Request**: If the request body is incorrect.
    ```json
    {
      "error": "error message"
    }
    ```
  - **500 Internal Server Error**: If an error occurs while saving the developer.
    ```json
    {
      "error": "error message"
    }
    ```

### 2. **Get All Devs**
- **URL:** `/dev`
- **Method:** `GET`
- **Description:** Retrieve all developers.
- **Response:**
  - **200 OK**
    ```json
    [
      {
        "id": "string",
        "username": "string",
        "email": "string"
      }
    ]
    ```
  - **500 Internal Server Error**: If an error occurs while retrieving the developers.
    ```json
    {
      "error": "error message"
    }
    ```

### 3. **Get Dev by ID**
- **URL:** `/dev/:id`
- **Method:** `GET`
- **Description:** Retrieve a developer by their ID.
- **Response:**
  - **200 OK**
    ```json
    {
      "id": "string",
      "username": "string",
      "email": "string"
    }
    ```
  - **404 Not Found**: If the developer with the specified ID does not exist.
    ```json
    {
      "error": "Developer not found"
    }
    ```

### 4. **Remove Dev by ID**
- **URL:** `/dev/:id`
- **Method:** `DELETE`
- **Description:** Remove a developer by their ID.
- **Response:**
  - **204 No Content**
    ```json
    {}
    ```
  - **404 Not Found**: If the developer with the specified ID does not exist.
    ```json
    {
      "error": "Developer not found"
    }
    ```
  - **500 Internal Server Error**: If an error occurs while deleting the developer.
    ```json
    {
      "error": "error message"
    }
    ```

## Running Tests

To run the tests for this API, you can use the following command:

```bash
go test ./...
```
