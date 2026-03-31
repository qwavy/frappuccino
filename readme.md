# Frappuccino: Advanced Coffee Shop Management System

## Learning Objectives
- RESTful API Design & Layered Software Architecture
- PostgreSQL & Advanced SQL (Aggregations, Full-Text Search)
- Database Design (ERD, Normalization)
- Containerization (Docker, Docker Compose)
- Transaction Management & Concurrency

## Abstract
In this project, you will build a robust, scalable **coffee shop management system**.

You will create a backend application using Go and PostgreSQL that simulates a real-world ordering system. The application will allow staff to manage orders, menu items, and inventory. Unlike a simple file-based storage system, this project requires a production-ready relational database, structured with proper constraints, relationships, and data types.

## Context
Coffee shops rely on sophisticated management systems that coordinate orders, inventory, menu items, and customer preferences in real-time. As businesses grow, their systems must scale.

You are tasked with building a modernized backend. By implementing a **three-layered architecture** and a **PostgreSQL database**, you ensure the code is maintainable, scalable, and capable of handling complex analytical queries. You will design the database schema, write complex SQL queries for reporting, and handle concurrent bulk orders using database transactions.

---

## General Criteria
- Your code MUST be written in accordance with [gofumpt](https://github.com/mvdan/gofumpt).
- Your program MUST compile successfully and MUST NOT exit unexpectedly (no panics like `nil-pointer dereference`, `index out of range`, etc.).
- Only built-in packages are allowed, **except for the PostgreSQL driver** (e.g., `pq` or `pgx`). Using any other external packages will result in a grade of 0.
- The project MUST be containerized and run using the following command in the project's root directory:
  ```shell
  $ docker compose up
  ```

---

## Architecture Requirements

Implement the application using a **three-layered architecture**:

1. **Presentation Layer (Handlers):**
    - Parses HTTP requests (JSON) and formats HTTP responses.
    - Validates input data.
2. **Business Logic Layer (Services):**
    - Contains core business rules (e.g., checking if enough ingredients exist before creating an order).
    - Manages transactions for bulk operations.
3. **Data Access Layer (Repositories):**
    - Executes SQL queries to interact with the PostgreSQL database.
    - Strictly separate from the business logic.

---

## Database Requirements (PostgreSQL)

### 1. ERD (Entity-Relationship Diagram)
Before coding, design an ERD based on the core entities. The schema must utilize the following specific PostgreSQL data types:
* **JSONB:** e.g., customization options, special instructions.
* **Arrays:** e.g., tags, allergens.
* **ENUM:** e.g., order statuses (`open`, `closed`, `rejected`), item sizes.
* **Timestamp with time zone:** e.g., order dates, inventory update logs.

### 2. Core Tables
Your database must include at least the following tables, properly related via Foreign Keys:
* `orders`: Main order info, timestamps, customer details.
* `order_items`: Junction connecting orders to menu items (with snapshot of price/quantity).
* `menu_items`: Products for sale, pricing, categorization.
* `menu_item_ingredients`: Recipe junction (menu items <-> ingredients).
* `inventory`: Tracks available ingredients, units, and stock levels.
* `order_status_history`: Tracks order state changes over time.
* `price_history`: Tracks menu item price changes.
* `inventory_transactions`: Records all inventory movements (usage, restock).

### 3. Initialization (`init.sql`)
You must create an `init.sql` file placed in the root folder. When `docker compose up` runs, this file must automatically:
- Define ENUMs and types.
- Create all tables with Primary/Foreign keys.
- Create at least **2 indexes** over all tables (e.g., for Full-Text Search or frequently queried columns).
- Insert **Mock Data**:
    - 10+ menu items, 20+ inventory items, 30+ orders.
    - Historical data (price history, status transitions) to allow testing of the analytical endpoints.

---

## API Endpoints & Business Logic

All data MUST be returned and accepted in JSON format. Use standard HTTP status codes (`200 OK`, `201 Created`, `400 Bad Request`, `404 Not Found`, `500 Internal Server Error`).

### Core CRUD Endpoints
* **Orders:**
    * `POST /orders`: Create a new order. *(Must deduct required ingredients from inventory. If insufficient, return error and abort).*
    * `GET /orders`, `GET /orders/{id}`, `PUT /orders/{id}`, `DELETE /orders/{id}`
    * `POST /orders/{id}/close`: Close an order.
* **Menu Items:**
    * `POST /menu`, `GET /menu`, `GET /menu/{id}`, `PUT /menu/{id}`, `DELETE /menu/{id}`
* **Inventory:**
    * `POST /inventory`, `GET /inventory`, `GET /inventory/{id}`, `PUT /inventory/{id}`, `DELETE /inventory/{id}`

### Advanced & Analytical Endpoints

**1. Aggregations (Basic)**
* `GET /reports/total-sales`: Get the total sales amount.
* `GET /reports/popular-items`: Get a list of popular menu items.

**2. Number of Ordered Items**
* `GET /orders/numberOfOrderedItems?startDate={YYYY-MM-DD}&endDate={YYYY-MM-DD}`
* Returns a list of items and their quantities ordered within the specified period. If omitted, returns all-time data.

**3. Full Text Search Report**
* `GET /reports/search?q={query}&filter={orders|menu|all}&minPrice={val}&maxPrice={val}`
* Search through orders, menu items, and customers with partial matching and ranking (using PostgreSQL FTS features like `tsvector`).

**4. Ordered Items by Period**
* `GET /reports/orderedItemsByPeriod?period={day|month}&month={month}&year={year}`
* Returns the number of orders grouped by day (within a month) or by month (within a year).

**5. Get Leftovers**
* `GET /inventory/getLeftOvers?sortBy={price|quantity}&page={page}&pageSize={pageSize}`
* Returns paginated and sorted inventory leftovers. Includes `currentPage`, `hasNextPage`, `pageSize`, and `totalPages` in the response.

**6. Bulk Order Processing**
* `POST /orders/batch-process`
* Processes an array of multiple orders simultaneously.
* **Crucial:** Must handle concurrent operations and maintain data integrity using **SQL Transactions**. If one item in a bulk order fails due to inventory, handle the rejection cleanly without crashing the accepted orders.

---

## Logging
Use Go's `log/slog` package. Log significant events (order creation, inventory updates) and errors with appropriate levels (`Info`, `Warn`, `Error`). Include contextual data (e.g., `orderID`).

## Deployment / Docker Guide
1. Create `init.sql` with your schema and mock data.
2. Use the provided `Dockerfile` and `docker-compose.yml`.
3. Ensure your Go application connects to the DB using these credentials:
    - **Host:** db
    - **Port:** 5432
    - **User:** latte
    - **Password:** latte
    - **Database:** frappuccino
4. Start the system via `docker compose up`. The API will be available at `localhost:8080`.

---

Как тебе такой вариант? Хочешь, чтобы мы сейчас вместе набросали структуру базы данных (ERD) или начали писать `init.sql` со всеми нужными ENUM и таблицами?