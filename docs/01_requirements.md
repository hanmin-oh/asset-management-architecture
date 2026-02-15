
# 01. Requirements

## 1️⃣ Functional Requirements

### 1. Asset Management
- Create, update, delete assets
- Track asset categories (cash, stock, crypto, real estate, etc.)
- Maintain balance consistency

### 2. Transaction Handling
- Record deposits and withdrawals
- Support atomic balance updates
- Prevent duplicate transaction processing (idempotency)

### 3. Asset Summary
- Provide aggregated balance per user
- Provide category-based breakdown
- Optimize read-heavy queries

### 4. Search & Filtering
- Full-text search on transaction descriptions
- Filter by date range, category, amount
- Sorting & pagination

### 5. Reporting
- Monthly spending reports
- Category-based analysis
- Historical trend tracking

---

## 2️⃣ Non-Functional Requirements

### Scalability
- Designed for 1M+ MAU
- Assume peak concurrent users: 5,000+
- Read-heavy workload optimization

### Consistency
- Strong consistency for financial transactions
- Eventual consistency for search indexing and reporting

### Performance
- Target API latency: p95 < 200ms
- Cached endpoints < 50ms

### Reliability
- System should tolerate:
  - Cache failure (fallback to DB)
  - Search index failure (degraded search only)
- Financial data must remain safe during partial failures

### Observability
- Structured logging
- Metrics collection (latency, error rate)
- Clear failure detection strategy

Observability is included to ensure the system can be operated and monitored in production environments.
---

## 3️⃣ Assumptions

- Single region deployment
- Initial monolithic deployment
- Horizontal scaling considered for future phase
- Authentication system is external or simplified

