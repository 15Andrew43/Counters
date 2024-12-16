# Counters

## Overview

The **Counters** project is a web service that tracks banner clicks and provides click statistics over specified time intervals. This tool is useful for analytics and managing advertising campaigns efficiently.

---

## Features

### 1. Click Counter (`/counter/<bannerID>`)
- **Method**: `GET`
- **Description**: Increments the click count for a specified banner.
- **Parameters**:
  - `bannerID`: Identifier of the banner.
- **Example Request**:
  ```http
  GET /counter/123
  ```
- **Response**: HTTP 200 status with confirmation of the increment.

### 2. Statistics Retrieval (`/stats/<bannerID>`)
- **Method**: `POST`
- **Description**: Returns click statistics for a specified banner over a time range.
- **Parameters** (in the request body):
  - `tsFrom`: Start of the time range (ISO 8601 timestamp).
  - `tsTo`: End of the time range (ISO 8601 timestamp).
  - `bannerID`: Identifier of the banner.
- **Example Request**:
  ```http
  POST /stats/123
  {
    "tsFrom": "2024-01-01T00:00:00Z",
    "tsTo": "2024-01-02T00:00:00Z"
  }
  ```
- **Response**: HTTP 200 status with the total number of clicks in the specified interval.

---

## Technical Details

### Technologies
- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL (default) or MongoDB (alternative)
- **Architecture**: REST API

### Scalability Levels
- **Junior**: Supports 10–50 `/counter` requests per second.
- **Middle+**: Supports 100–500 `/counter` requests per second.

---

## Solution Design

1. **Click Tracking**:
   - Implements an in-memory cache to handle high-frequency `/counter` requests.
   - Periodically flushes cached clicks to the database at configurable intervals.

2. **Statistics Retrieval**:
   - Queries the database to calculate total clicks for a banner over a given time range.

3. **Concurrency**:
   - Uses `sync.Mutex` for thread-safe cache access.

4. **Database Schema**:
   - A `clicks` table stores `banner_id`, `timestamp`, and `count`.
   - Ensures unique records per `banner_id` and `timestamp` using a unique constraint.

5. **Load Testing**:
   - Utilizes `k6` for stress testing with predefined traffic scenarios for both endpoints.

---

## Deployment

### Docker Setup
- **Services**:
  - `click_counter`: The Go service.
  - `postgres`: The PostgreSQL database.
- **Environment Variables**:
  - `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Database connection details.
  - `FLUSH_INTERVAL`: Cache flush interval in seconds.
- **Network**: Uses a shared bridge network for service communication.

### Example Docker Compose Configuration
```yaml
version: "3.8"
services:
  click_counter:
    build:
      context: .
      dockerfile: ./click_counter/Dockerfile
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=clicks_db
      - FLUSH_INTERVAL=10
    depends_on:
      postgres:
        condition: service_healthy
    networks:
      - click_network

  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: clicks_db
    ports:
      - "5432:5432"
    volumes:
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 2s
      timeout: 5s
      retries: 5
    networks:
      - click_network

networks:
  click_network:
    driver: bridge
```

---

## Database Schema
```sql
CREATE TABLE IF NOT EXISTS clicks (
    id SERIAL PRIMARY KEY,
    banner_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    count INT NOT NULL,
    UNIQUE (banner_id, timestamp)
);
```

---

## Stress Testing
### Load Test Configuration (`k6`)
- **Click Counter Endpoint**:
  - Simulates up to 4000 requests per second.
  - Validates response status and latency.

- **Statistics Endpoint**:
  - Simulates up to 2000 requests per second.
  - Validates response status and latency for dynamic time ranges.

---

## Quick Start
1. Build the Docker containers:
   ```bash
   docker-compose up --build
   ```
2. Access the service:
   - Increment clicks: `http://localhost:8080/counter/<bannerID>`
   - Get stats: `http://localhost:8080/stats/<bannerID>` (POST request with JSON body).
3. Run stress tests:
   ```bash
   k6 run ./click_counter/tests/load/get_test.js
   k6 run ./click_counter/tests/load/post_test.js
   ```

