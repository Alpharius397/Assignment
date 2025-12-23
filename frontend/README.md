# Secure User Profile & Access Control System – Frontend

## Project Overview

### Assignment 1 – Secure User Profile & Access Control System (Frontend)

This frontend application provides the **user interface** for the Identity Management system built in Assignment 1. It enables users to **register, log in, and securely view their profile data**, while communicating with a JWT-secured backend.

The frontend is designed to consume the backend’s **Swagger (OpenAPI) specification** and interact with authenticated APIs using **access and refresh tokens**.


## Technology Stack

* **Framework:** React
* **Build Tool:** Vite
* **Styling:** Tailwind CSS
* **UI Components:** shadcn/ui
* **Data Fetching & Caching:** React Query
* **HTTP Client:** Axios
* **State Persistence:** localStorage
* **API Contract:** Swagger / OpenAPI (served by backend)


## Architecture & Approach

* **Authentication Flow**

  * Access & refresh tokens stored in `localStorage`
  * Axios interceptor automatically refreshes access tokens on non-auth API failures

* **Data Fetching**

  * React Query handles server state, caching, retries, and loading states

* **UI**

  * Tailwind CSS for utility-first styling
  * shadcn/ui for accessible, reusable components

* **Security**

  * Protected routes require valid access tokens

* **Separation of Concerns**

  * API logic isolated from UI components
  * Axios instance configured globally


## Setup / Run Instructions

### Prerequisites

Ensure the following are installed:

* **Node.js**
* **npm / bun**


### Installation

1. Clone the repository and navigate to the frontend directory:

```bash
git clone https://github.com/Alpharius397/Assignment
cd Assignment/frontend
```

2. Install dependencies:

```bash
npm install
```
or
```bash
bun install
```


### Run the Application

```bash
npm run dev
```
or
```bash
bun run dev
```

The frontend will be available at:

```
http://localhost:5173
```


## Authentication & Token Handling

* **Access Token**

  * Stored in `localStorage`
  * Sent with every authenticated request via Axios interceptor

* **Refresh Token**

  * Stored in `localStorage`
  * Automatically used to fetch a new access token when API requests fail due to token expiry

* **Non-auth Paths**

  * Token refresh logic is skipped for `/login` and `/register` endpoints

This ensures a **seamless user experience** without manual re-login on token expiry.


## API Integration

### Backend Contract

* The backend exposes its Swagger (OpenAPI 2.0) specification
* The frontend uses this specification as the **source of truth** for API structure

### Core APIs Used

* `POST /login` – User authentication
* `POST /register` – User registration
* `POST /refresh` – Access token refresh
* `GET /profile` – Fetch authenticated user profile
* `GET /get-data` – Fetch paginated user list

All authenticated requests include:

```
Authorization: Bearer <access_token>
```


## Pages & Features

### Authentication

* Login page
* Registration page
* Client-side form validation
* Error handling for invalid credentials

### Profile Dashboard

* Displays authenticated user information
* Securely renders decrypted Aadhaar/ID number
* Handles loading and error states gracefully

### Error Handling

* API errors surfaced clearly to users
* Automatic token refresh on expiry
* Graceful logout on refresh failure


## AI Tool Usage Log (MANDATORY)

### AI-Assisted Tasks

AI tools were used **only** to generate:

* Validation schemas boilerplate
* Basic UI scaffolding / layout templates


### Effectiveness Score

#### Score: 3 / 5

AI tools provided moderate productivity gains by accelerating validation schema generation and basic UI scaffolding. However, AI-generated patterns did not adequately account for framework-level optimizations introduced by the React Compiler.

In the data listing view, the pagination component (`DataTablePagination`) receives the table instance created by `useReactTable`. Due to aggressive memoization by **React Compiler**, the table instance could become stale relative to the latest server-paginated data, leading to desynchronization between the table’s internal pagination state and the actual data rendered in the UI. Resolving this issue required manual debugging, careful control of pagination state, and an in-depth understanding of React’s rendering and memoization behavior.
