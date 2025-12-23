# Secure User Profile & Access Control System

**Full Stack Assignment (Assignment 1)**

A full-stack identity management system implementing secure user authentication, encrypted storage of sensitive user data, and protected profile access using a Go backend and a React frontend.


## Project Overview

This project implements an **Identity Management Microservice** that supports secure user registration, authentication, and profile access. Sensitive fields such as Aadhaar/ID numbers are encrypted at rest, and all protected resources require JWT-based authentication.

The system is divided into:

* **Backend:** Authentication, encryption, authorization, and persistence
* **Frontend:** Secure UI for interacting with authenticated APIs


## Technology Stack

### Backend

* Go
* JWT (Access & Refresh Tokens)
* AES-256 encryption
* SQLite3
* Swagger (OpenAPI 2.0)

### Frontend

* React + Vite
* Tailwind CSS, shadcn/ui
* React Query
* Axios
* localStorage
* Swagger/OpenAPI as API contract


## Repository Structure

```
root
  ├── backend/
  ├── frontend/
  └── README.md
```


## Setup Instructions

### Backend

**Prerequisites:** Go (>= 1.23), SQLite3, Git

```bash
git clone https://github.com/Alpharius397/Assignment
cd Assignment/backend
go mod tidy
```

Create `.env`:

```env
JWT_SECRET=<JWT-secret>
AES_KEY=<AES-key>
```

Run:

```bash
go run .
```

Backend runs at `http://localhost:8081`


### Frontend

**Prerequisites:** Node.js, npm or bun

```bash
cd Assignment/frontend
npm install
npm run dev
```
or
```bash
cd Assignment/frontend
bun install
bun run dev
```

Frontend runs at `http://localhost:5173`


## API Overview

### Authentication

* `POST /login`
* `POST /register`
* `POST /refresh`

### Protected APIs

* `GET /profile`
* `GET /get-data` (paginated)

All protected endpoints require:

```
Authorization: Bearer <access_token>
```


## Database Schema

SQLite3 is used for its simplicity and lightweight setup.

### users Table

| Column     | Type     | Nullable | Default           |
| ---------- | -------- | -------- | ----------------- |
| user_name  | text     | not null |                   |
| email      | text     | not null |                   |
| aadhar     | text     | not null |                   |
| password   | text     | not null |                   |
| created_at | datetime | not null | CURRENT_TIMESTAMP |
| updated_at | datetime |          | CURRENT_TIMESTAMP |
| deleted_at | datetime |          |                   |

**Indexes**

* `users_email_key` (UNIQUE)
* `users_user_name_key` (UNIQUE)



## AI Tool Usage Log

### AI-Assisted Tasks

AI tools were used **selectively** and **only where appropriate**.

**Backend**

* Unit test generation for AES-256 encryption/decryption utilities
* Assistance with JWT token validation and refresh test cases
* Documentation clarity improvements

**Frontend**

* Validation schema boilerplate
* Basic UI scaffolding and layout templates

### Effectiveness Score

**Score:** **3 / 5**

**Justification:**
AI tools provided moderate productivity gains by reducing time spent on repetitive scaffolding, test generation, and documentation. However, AI-generated suggestions failed to account for critical edge cases.

On the backend, AI did not identify that `cipher.BlockMode.CryptBlocks` can panic if buffer alignment guarantees are violated, requiring manual defensive validation.

On the frontend, AI-generated pagination patterns failed under aggressive React Compiler memoization, causing desynchronization between table state and server-paginated data. Resolving this required manual debugging and explicit state control.

All security-sensitive logic, pagination behavior, and state synchronization were therefore implemented and reviewed manually to ensure correctness and robustness.