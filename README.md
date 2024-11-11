# Email Verification Service

This project is a simple API for handling email verification with token-based authentication. It includes endpoints for creating emails, generating tokens, verifying emails, and checking readiness.

## Table of Contents

- [Installation](#installation)
- [API Endpoints](#api-endpoints)
  - [GET /ready](#get-ready)
  - [GET /err](#get-err)
  - [POST /email](#post-email)
  - [GET /email/{id}](#get-emailid)
  - [POST /email/{id}/token](#post-emailidtoken)
  - [GET /verify](#get-verify)
- [Running the Application](#running-the-application)
- [Testing the API](#testing-the-api)

## Installationx

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Build the application:**

   ```bash
   go build
   ```

4. **Run the application:**

   ```bash
   go run main.go
   ```

   The API will be running on `http://localhost:PORT`.

## API Endpoints

### **GET /ready**

This endpoint checks the readiness of the API.

- **Description**: Returns a simple readiness check.
- **Response**:
  - `200 OK`: The API is ready.

### **GET /err**

This endpoint simulates an error for testing purposes.

- **Description**: Used to trigger an error for testing.
- **Response**:
  - `500 Internal Server Error`: Something went wrong.

### **POST /email**

This endpoint creates a new email record in the database.

- **Request body**: 
  - `email` (string): The email address to be created.
  
- **Response**:
  - `201 Created`: The email has been successfully created.
  - `400 Bad Request`: Invalid email address or input data.

### **GET /email/{id}**

This endpoint retrieves the details of an email by its unique identifier or the email string.

- **Parameters**:
  - `id` (UUID/string): The ID of the email or an email string.
  
- **Response**:
  - `200 OK`: The email details are returned in JSON format.
  - `400 Bad Request`: Invalid email ID.
  - `404 Not Found`: Email not found.

### **POST /email/{id}/token**

This endpoint creates a verification token for a given email.

- **Parameters**:
  - `id` (UUID): The ID of the email record.
  
- **Response**:
  - `201 Created`: A new verification token has been created.
  - `400 Bad Request`: Invalid email ID or error during token creation.

### **GET /verify**

This endpoint verifies the email based on the provided token.

- **Query Parameters**:
  - `token` (string): The token sent to the user's email for verification.

- **Response**:
  - `200 OK`: The email has been successfully verified.
  - `400 Bad Request`: Invalid or expired token.
  - `404 Not Found`: Token not found in the database.

## Running the Application

Once the application is running, you can access the API at `http://localhost:PORT`.

You can use `curl` or Postman to interact with the API. Here's how you can test each endpoint.

### Test Example using `curl`:

#### **1. Create an Email**:
```bash
curl -X POST http://localhost:PORT/email -d '{"email": "user@example.com"}' -H "Content-Type: application/json"
```

#### **2. Get Email by ID**:
```bash
curl http://localhost:PORT/email/{id}
```

#### **3. Create Token for Email**:
```bash
curl -X POST http://localhost:PORT/email/{id}/token
```

#### **4. Verify Email**:
```bash
curl http://localhost:PORT/verify?token={token}
```
