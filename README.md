# Payments API

A simple REST API for creating and managing payments built with Go, Gin,
and PostgreSQL.

This project demonstrates basic backend concepts such as HTTP routing,
database integration, and idempotent payment creation.

------------------------------------------------------------------------

## Features

-   REST API built with Gin
-   PostgreSQL database integration
-   Payment creation endpoint
-   Idempotency key support to prevent duplicate payments
-   Simple project structure separating routes and database logic

------------------------------------------------------------------------

## Tech Stack

-   Go
-   Gin (HTTP framework)
-   PostgreSQL
-   database/sql

------------------------------------------------------------------------

## Project Structure

    payments-api/
    ├── cmd/
    │   └── main.go          # application entry point
    │
    ├── internal/
    │   ├── db/              # database connection and queries
    │   └── routes/          # API route definitions
    │
    ├── go.mod
    └── README.md

------------------------------------------------------------------------

## Running the Application

### 1. Clone the repository

``` bash
git clone https://github.com/YOUR_USERNAME/payments-api
cd payments-api
```

### 2. Install dependencies

``` bash
go mod tidy
```

### 3. Configure PostgreSQL

Create a database and a payments table.

Example schema:

``` sql
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    amount INTEGER NOT NULL,
    currency TEXT NOT NULL,
    idempotency_key TEXT UNIQUE
);
```

### 4. Run the server

``` bash
go run cmd/main.go
```

The API will start at:

    http://localhost:8080

------------------------------------------------------------------------

## Example Endpoint

### Create Payment

    POST /payments

Example request body:

``` json
{
  "amount": 1000,
  "currency": "USD",
  "idempotency_key": "abc123"
}
```

Example response:

``` json
{
  "id": 1,
  "amount": 1000,
  "currency": "USD"
}
```

------------------------------------------------------------------------

## Purpose

This project was created as a small backend exercise to practice
building REST APIs in Go and integrating them with a relational
database.
