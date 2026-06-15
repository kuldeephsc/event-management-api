# Event Management REST API

A REST API built using **Go**, **Gin**, **JWT Authentication**, and **SQLite/MySQL** for managing events and user registrations.

---

## Features

- User Signup
- User Login with JWT Authentication
- Create Events
- Update Events
- Delete Events
- Get All Events
- Get Single Event
- Register for Events
- Unregister from Events
- Password Hashing using bcrypt
- Middleware-based Authentication

---

## Tech Stack

- Go
- Gin Framework
- JWT
- bcrypt
- SQLite / MySQL
- REST API

---

## Project Structure

```bash
.
├── main.go
├── routes/
├── models/
├── middlewares/
├── utils/
├── db/
├── create-event.http
├── login.http
├── create-user.http
└── update-event.http
```

---

## Installation

### Clone Repository

```bash
git clone <your-repository-url>
cd api
```

### Install Dependencies

```bash
go mod tidy
```

### Run Server

```bash
go run .
```

Server will start on:

```bash
http://localhost:8085
```

---

## API Endpoints

### Authentication

#### Signup

```http
POST /signup
```

Request Body:

```json
{
  "email": "test@example.com",
  "password": "123456"
}
```

---

#### Login

```http
POST /login
```

Request Body:

```json
{
  "email": "test@example.com",
  "password": "123456"
}
```

Response:

```json
{
  "message": "Login successful",
  "token": "jwt-token"
}
```

---

## Event APIs

### Get All Events

```http
GET /events
```

---

### Get Single Event

```http
GET /events/:id
```

---

### Create Event

```http
POST /events
Authorization: <jwt-token>
```

Request Body:

```json
{
  "name": "Go Workshop",
  "description": "Learning Go",
  "location": "Bangalore",
  "dateTime": "2026-01-01T15:30:00.000Z"
}
```

---

### Update Event

```http
PUT /events/:id
Authorization: <jwt-token>
```

---

### Delete Event

```http
DELETE /events/:id
Authorization: <jwt-token>
```

---

### Register for Event

```http
POST /events/:id/register
Authorization: <jwt-token>
```

---

### Unregister from Event

```http
DELETE /events/:id/register
Authorization: <jwt-token>
```

---

## JWT Authentication

JWT token is generated during login and validated using middleware.

Protected routes:

- Create Event
- Update Event
- Delete Event
- Register for Event
- Unregister from Event

---

## Password Security

Passwords are hashed using bcrypt before storing them in the database.

---

## Sample Authorization Header

```http
Authorization: your-jwt-token
```

---

## Future Improvements

- Refresh Tokens
- Role-Based Authorization
- Docker Support
- Swagger Documentation
- Unit Testing
- PostgreSQL Support
- Environment Variables for Secrets

---

## Author

Kuldeep Maurya

Built using Go, JWT
