# 03. Data Model Design

## 1️⃣ Core Data Store (PostgreSQL)

PostgreSQL is used as the source of truth for financial data.

Strong consistency is required for:
- Transactions
- Balance updates

---

### Tables

#### users
- id (PK)
- email
- created_at

#### assets
- id (PK)
- user_id (FK)
- category
- balance
- created_at
- updated_at

Indexes:
- (user_id)
- (user_id, category)

Reason:
- Frequent queries by user
- Category-based breakdown optimization

---

#### transactions
- id (PK)
- user_id (FK)
- asset_id (FK)
- type (deposit/withdraw)
- amount
- description
- created_at

Indexes:
- (user_id, created_at DESC)
- (asset_id)
- (user_id, amount)

Reason:
- Recent transaction lookup
- Asset-based filtering
- Amount-based filtering

---

## 2️⃣ Search Model (Elasticsearch)

Elasticsearch stores denormalized transaction documents.

Why separate from PostgreSQL?

- Full-text search
- Complex filtering
- Sorting performance
- Reduced load on primary DB

Consistency model:
- Event-driven indexing
- Eventual consistency acceptable

---

## 3️⃣ Reporting / Event Store (MongoDB)

MongoDB stores:

- Transaction events (raw JSON)
- Monthly report documents
- Aggregated user statistics

Why MongoDB?

- Flexible schema
- Efficient document-based reporting
- Easy historical snapshot storage

Consistency model:
- Eventual consistency
- Rebuildable from source data

---

## 4️⃣ Caching Layer (Redis)

Used for:

- User summary cache
- Frequently accessed asset balances

Cache strategy:
- Key pattern: summary:{user_id}
- TTL: 60 seconds
- Invalidation on transaction commit

Failure handling:
- Fallback to PostgreSQL
