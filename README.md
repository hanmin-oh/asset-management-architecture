# Scalable Asset Management Architecture

Designing a scalable backend system for an asset management platform with ML-ready and search-integrated architecture.

---

## 🎯 Project Goals

- Design for 1M+ MAU scalability
- Ensure strong data consistency for financial transactions
- Optimize read performance using caching strategies
- Integrate search capabilities using Elasticsearch
- Prepare architecture for future ML-based recommendation features
- Document trade-offs and architectural decisions

---

## 🏗 System Overview

This project simulates a production-ready asset management backend.

The system is designed with:

- Strong consistency for transaction handling
- Event-driven architecture for extensibility
- Search-optimized read models
- Distributed-ready components
- Observability and scalability considerations

The focus is not just feature implementation, but architectural reasoning.

---

## ⚙️ Architecture Highlights

- Monolithic core with modular boundaries (future MSA-ready)
- PostgreSQL as source of truth
- Redis for read optimization
- Elasticsearch for search queries
- MongoDB for event storage & reporting views
- Event-driven indexing and reporting pipeline
- Cache invalidation strategy
- Failure scenario considerations

---

## 🧠 What This Portfolio Demonstrates

- System design thinking
- Trade-off analysis
- Scalability considerations
- Data consistency handling
- Search integration strategy
- Distributed system fundamentals
- ML-ready architectural mindset

---

## 🛠 Tech Stack

- Go (Golang)
- PostgreSQL
- MongoDB
- Elasticsearch
- Redis
- Docker
- GitHub Actions (CI)

---

## 📂 Documentation

Detailed architectural documents are available in the `/docs` directory.

