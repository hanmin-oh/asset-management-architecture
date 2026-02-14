# 02. Architecture

## 1) Architecture Principles

- Start monolithic for speed and maintainability
- Define clear module boundaries to enable future extraction
- Keep financial data strongly consistent (source of truth)
- Use eventual consistency for derived read models (search/report)
- Design for failure: degrade non-critical features first

---

## 2) High-level Components

- API (Go): HTTP endpoints, auth, request validation
- Core domain modules: asset, transaction, user
- Data stores:
  - PostgreSQL: source of truth (balances, transactions)
  - Redis: caching for read-heavy endpoints
  - Elasticsearch: search index for transactions/assets
  - MongoDB: events/logs & reporting views
- Async pipeline:
  - Event bus + workers for indexing/reporting

---

## 3) Data Flow Overview

(여기에 mermaid 다이어그램 넣기)

Key flows:
- Write path: transaction -> DB commit -> publish event -> async indexing/reporting
- Read path: summary -> cache -> fallback to DB
- Search path: query -> Elasticsearch (degraded if ES unavailable)

---

## 4) Consistency Model

- Strong consistency:
  - Transaction creation
  - Balance updates
- Eventual consistency:
  - Elasticsearch indexing
  - Monthly reports (MongoDB)

---

## 5) Failure Scenarios & Degradation

- Redis down:
  - Read falls back to DB
  - Rate limit / circuit breaker considered
- Elasticsearch down:
  - Search endpoints degrade (return minimal results or 503)
  - Core transaction flows remain unaffected
- Worker backlog:
  - Index/report delays tolerated
  - Monitoring & alert thresholds defined

---

## 6) MSA Evolution Plan (Future)

This project starts as a modular monolith. Candidate modules for extraction:
- Search Indexer Service
- Reporting Service
- Notification Service

When to split:
- Independent scaling needs
- Deployment frequency divergence
- Ownership boundaries / team scaling
- Reliability isolation requirements

How to split:
- Keep DB as source of truth initially
- Use events as integration contract
- Gradually move to service-owned data where needed
