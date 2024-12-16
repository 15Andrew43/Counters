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

```bash
➜  click_counter git:(develop) ✗ k6 run tests/load/get_test.js


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/get_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 4000 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 4000 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✓ status is 200
     ✗ response time < 200ms
      ↳  99% — ✓ 91162 / ✗ 199

     checks.........................: 99.89% 182523 out of 182722
     data_received..................: 12 MB  300 kB/s
     data_sent......................: 8.1 MB 200 kB/s
     http_req_blocked...............: avg=30.51µs min=1µs   med=4µs    max=31.1ms   p(90)=12µs   p(95)=49µs   
     http_req_connecting............: avg=18.34µs min=0s    med=0s     max=30.21ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=6.46ms  min=288µs med=1.74ms max=302.22ms p(90)=9.66ms p(95)=24.46ms
       { expected_response:true }...: avg=6.46ms  min=288µs med=1.74ms max=302.22ms p(90)=9.66ms p(95)=24.46ms
     http_req_failed................: 0.00%  0 out of 91361
     http_req_receiving.............: avg=44.75µs min=10µs  med=26µs   max=29.6ms   p(90)=59µs   p(95)=86µs   
     http_req_sending...............: avg=49.17µs min=4µs   med=9µs    max=37.9ms   p(90)=53µs   p(95)=108µs  
     http_req_tls_handshaking.......: avg=0s      min=0s    med=0s     max=0s       p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=6.37ms  min=265µs med=1.68ms max=302.11ms p(90)=9.44ms p(95)=24.15ms
     http_reqs......................: 91361  2241.9002/s
     iteration_duration.............: avg=1s      min=1s    med=1s     max=1.3s     p(90)=1.01s  p(95)=1.03s  
     iterations.....................: 91361  2241.9002/s
     vus............................: 120    min=57               max=3993
     vus_max........................: 4000   min=4000             max=4000


running (0m40.8s), 0000/4000 VUs, 91361 complete and 0 interrupted iterations
default ✓ [======================================] 0000/4000 VUs  40s
➜  click_counter git:(develop) ✗ k6 run tests/load/post_test.js 


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/post_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 2000 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 2000 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✗ status is 200
      ↳  99% — ✓ 20401 / ✗ 191
     ✗ response time < 200ms
      ↳  53% — ✓ 10974 / ✗ 9618

     checks.........................: 76.18% 31375 out of 41184
     data_received..................: 2.5 MB 51 kB/s
     data_sent......................: 4.3 MB 88 kB/s
     http_req_blocked...............: avg=175.96µs min=2µs    med=5µs     max=382.58ms p(90)=76.7µs p(95)=368µs
     http_req_connecting............: avg=160.05µs min=0s     med=0s      max=382.48ms p(90)=0s     p(95)=290µs
     http_req_duration..............: avg=1.44s    min=1.05ms med=93.71ms max=31.27s   p(90)=3.92s  p(95)=6.7s 
       { expected_response:true }...: avg=1.4s     min=1.05ms med=87.34ms max=31.16s   p(90)=3.88s  p(95)=6.53s
     http_req_failed................: 0.92%  191 out of 20592
     http_req_receiving.............: avg=56.87µs  min=16µs   med=48µs    max=9.9ms    p(90)=87µs   p(95)=111µs
     http_req_sending...............: avg=33.7µs   min=6µs    med=20µs    max=29.24ms  p(90)=54µs   p(95)=76µs 
     http_req_tls_handshaking.......: avg=0s       min=0s     med=0s      max=0s       p(90)=0s     p(95)=0s   
     http_req_waiting...............: avg=1.44s    min=1ms    med=93.63ms max=31.27s   p(90)=3.92s  p(95)=6.7s 
     http_reqs......................: 20592  417.057641/s
     iteration_duration.............: avg=2.45s    min=1s     med=1.09s   max=32.27s   p(90)=4.92s  p(95)=7.7s 
     iterations.....................: 20592  417.057641/s
     vus............................: 33     min=33             max=1991
     vus_max........................: 2000   min=2000           max=2000


running (0m49.4s), 0000/2000 VUs, 20592 complete and 0 interrupted iterations
default ✓ [======================================] 0000/2000 VUs  40s
```

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

