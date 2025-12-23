# Secure User Profile & Access Control System – Backend

## Project Overview

### Assignment 1 – Secure User Profile & Access Control System (Backend)

This backend service implements an **Identity Management Microservice** responsible for secure user authentication and profile management. It exposes RESTful APIs for **user registration, login, and user data retrieval**.

### Implementation Approach

* **Authentication:** Stateless authentication using **JWT (JSON Web Tokens)**.
* **Data Security:** Sensitive user fields such as **Aadhaar/ID Number** are encrypted at rest using **AES-256**.
* **Authorization:** Protected routes are accessible only via valid JWT tokens.
* **Architecture:** Layered structure to ensure maintainability and testability.
* **Testing Focus:** Encryption/decryption logic and token validation utilities are unit-tested.

## Setup / Run Instructions

### Prerequisites

Ensure the following are installed on your system:

* **Git**
* **Go** (>= 1.23) 
* **Package Manager:** go mod
* **Database:** SQLite (>= 3.x.x)

### Backend Setup

1. Clone the repository:

```bash
  git clone https://github.com/Alpharius397/Assignment
  cd Assignment/backend
```

2. Install dependencies:

```bash
  go mod tidy
```

3. Configure environment variables:

Create a `.env` file in the `backend` directory:

```env
  JWT_SECRET=<JWT-secret-here>
  AES_KEY=<AES-key-here>
```

4. Start the backend server:

```bash
  go run .
```

The backend will be available at: `http://localhost:8081`

## API Documentation

This backend exposes RESTful APIs documented using **Swagger (OpenAPI 2.0)**. The following section is derived directly from the Swagger specification used in this project.

### Base URL

```
http://localhost:8081
```

### Authentication APIs

#### POST `/login`

* **Description:** Validates user credentials and generates JWT access and refresh tokens.
* **Request Body:**

```json
{
  "user_name": "user_1",
  "password": "Pass1"
}
```

* **Responses:**

  * `200 OK` – Returns access and refresh tokens
  * `401 Unauthorized` – Invalid credentials
  * `500 Internal Server Error`

#### POST `/register`

* **Description:** Registers a new user. Aadhaar number is encrypted before being stored.
* **Request Body:**

```json
{
  "user_name": "user_1",
  "email": "user1@example.com",
  "password": "Pass1",
  "confirm_password": "Pass1",
  "aadhar": "946720527727"
}
```

* **Responses:**

  * `200 OK` – User registered successfully
  * `400 Bad Request` – Validation error
  * `500 Internal Server Error`

#### POST `/refresh`

* **Description:** Issues a new access token using a valid refresh token.
* **Headers:**

```
Authorization: Bearer <refresh_token>
```

* **Responses:**

  * `200 OK` – New access token returned
  * `401 Unauthorized`
  * `500 Internal Server Error`

### Profile APIs

#### GET `/profile`

* **Description:** Returns the authenticated user's profile information. Aadhaar is decrypted before response.
* **Headers:**

```
Authorization: Bearer <access_token>
```

* **Responses:**

  * `200 OK` – Returns user profile data
  * `401 Unauthorized`
  * `500 Internal Server Error`


### User Data APIs

#### GET `/get-data`

* **Description:** Returns a paginated list of user profiles.
* **Headers:**

```
Authorization: Bearer <access_token>
```

* **Query Parameters:**

  * `offset` (number) – Pagination offset
  * `limit` (number) – Pagination limit
  * `raw` (boolean) – If true, returns decrypted aadhar ID else returns encrypted aadhar ID

* **Responses:**

  * `200 OK` – Returns user list and total count
  * `401 Unauthorized`
  * `500 Internal Server Error`

## Database Schema

SQLite3 was used due to its ease of use and lightweight setup. 

The users table stores user credentials and identity details, enforces uniqueness on email and user_name, automatically tracks creation and update timestamps, and supports soft deletion through the optional deleted_at field.

### Table Structure of users

|   Column   |            Type             | Nullable |      Default
|------------|-----------------------------|----------|-------------------
| user_name  | text                        | not null |
| email      | text                        | not null |
| aadhar     | text                        | not null |
| password   | text                        | not null |
| created_at | datetime                    | not null | CURRENT_TIMESTAMP
| updated_at | datetime                    |          | CURRENT_TIMESTAMP
| deleted_at | datetime                    |          |

### Indexes on users

| Index Name          | Type   | Columns   |
| ------------------- | ------ | --------- |
| users_email_key     | UNIQUE | email     |
| users_user_name_key | UNIQUE | user_name |



## AI Tool Usage Log

### AI-Assisted Tasks

AI tools were used **only for test-related work**, as required by the assignment:

* Generated unit test cases for AES-256 encryption and decryption utilities.
* Assisted in writing test cases for JWT token validation and refresh logic.
* Assisted in documentation process for better readability

No AI tools were used for core API logic, encryption implementation, or database design.

### Effectiveness Score

#### Score: 3 / 5

AI tools provided moderate efficiency gains by accelerating the creation of unit test cases for AES-256 encryption/decryption utilities, JWT token validation logic, and improving documentation readability. However, AI-generated suggestions did not fully account for low-level cryptographic edge cases. For example, in the AES-CBC decryption helper, the AI did not flag that mode.CryptBlocks(dst, src) can panic if the input is not block-aligned or if buffer assumptions are violated, requiring manual code review and defensive validation.

**Reference: cipher.BlockMode.CryptBlocks Behavior (Go Documentation)**

```go
func (cipher.BlockMode) CryptBlocks(dst []byte, src []byte)
```
*CryptBlocks encrypts or decrypts a number of blocks. The length of src must be a multiple of the block size. Dst and src must overlap entirely or not at all.*

*If len(dst) < len(src), CryptBlocks should panic. It is acceptable to pass a dst bigger than src, and in that case, CryptBlocks will only update dst[:len(src)] and will not touch the rest of dst.*
