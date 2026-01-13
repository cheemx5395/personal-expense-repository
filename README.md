# Personal Expense Tracker (Golang Assignment)

A simple **Personal Expense Tracker** built using **Golang**, providing all basic **CRUD** operations for managing expenses.
The project is designed to be minimal, readable, and production-oriented while demonstrating clean backend practices.

## Tech Stack

* **Language**: Go (Golang)
* **Routing**: `gorilla/mux`
* **Database**: SQLite3
* **Migrations**: Goose
* **Query Generation**: sqlc (type-safe Go code from SQL)
* **Architecture**: Layered (handler → service → repository)

SQLite was chosen for simplicity since this is a small, single-node application. The design remains portable to other databases like PostgreSQL with minimal changes.

## Features

* Health check endpoint
* Create, read, update, and delete expenses
* Type-safe database access using sqlc
* Schema versioning with Goose migrations
* RESTful API design

## API Endpoints

### Health Check

**GET `/health`**

Check whether the server is running.

**Response**

* Status Code: `200 OK`

```json
{
  "message": "Healthy Connection Established!"
}
```

### Create Expense

**POST `/expense`**

Create a new expense entry.

**Request Body**

```json
{
  "title": "lunch",
  "amount": 60,
  "description": "tiffin payment done"
}
```

**Response**

* Status Code: `201 Created`

```json
{
  "id": 1,
  "title": {
    "String": "lunch",
    "Valid": true
  },
  "amount": 60,
  "description": {
    "String": "tiffin payment done",
    "Valid": true
  },
  "created_at": {
    "Time": "2026-01-13T08:26:27Z",
    "Valid": true
  },
  "updated_at": "2026-01-13T08:26:27Z"
}
```

### Get All Expenses

**GET `/expenses`**

Fetch all recorded expenses.

**Response**

* Status Code: `200 OK`

```json
[
  {
    "id": 1,
    "title": {
      "String": "lunch",
      "Valid": true
    },
    "amount": 60,
    "description": {
      "String": "tiffin payment done",
      "Valid": true
    },
    "created_at": {
      "Time": "2026-01-13T08:26:27Z",
      "Valid": true
    },
    "updated_at": "2026-01-13T08:26:27Z"
  }
]
```

### Get Expense by ID

**GET `/expenses/{id}`**

Fetch a single expense using its ID.

**Response**

* Status Code: `200 OK`

```json
{
  "id": 1,
  "title": {
    "String": "lunch",
    "Valid": true
  },
  "amount": 60,
  "description": {
    "String": "tiffin payment done",
    "Valid": true
  },
  "created_at": {
    "Time": "2026-01-13T08:26:27Z",
    "Valid": true
  },
  "updated_at": "2026-01-13T08:26:27Z"
}
```

### Update Expense

**PUT `/expenses/{id}`**

Update an existing expense.

**Request Body**

```json
{
  "title": "tiffin",
  "amount": 60,
  "description": "tiffin payment done"
}
```

**Response**

* Status Code: `200 OK`

```json
{
  "id": 1,
  "title": {
    "String": "tiffin",
    "Valid": true
  },
  "amount": 60,
  "description": {
    "String": "tiffin payment done",
    "Valid": true
  },
  "created_at": {
    "Time": "2026-01-13T08:26:27Z",
    "Valid": true
  },
  "updated_at": "2026-01-13T08:26:27Z"
}
```

### Delete Expense

**DELETE `/expenses/{id}`**

Delete an expense by ID.

**Response**

* Status Code: `200 OK`

```json
{
  "message": "expense deleted successfully"
}
```

## Postman Collection

A Postman collection containing all API requests is available here:

**Postman Collection Link**: [collection](https://warped-meadow-913182.postman.co/workspace/New-Team-Workspace~850b93a7-4078-4f7e-bcb5-331e137d6e73/collection/32759292-3a085ffd-da7c-4fc0-9cb0-e0c40c1b78ac?action=share&source=copy-link&creator=32759292)

## Conclusion

Kept it as simple as possible since the assignment doesn't specify anything more!