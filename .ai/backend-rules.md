# Backend Rules

Stack:

- Go
- Gin
- SQLite
- FTS5

---

# Architecture

Mandatory layers:

API

↓

Service

↓

Repository

↓

Database

---

Forbidden

API → SQL

API → Database

Service → SQL

Repository → Business Logic

---

# Package Design

One package = one responsibility.

Avoid giant packages.

---

# Dependency Direction

Outer layers depend on inner abstractions.

Never reverse dependency direction.

---

# Context

Pass context.Context through:

API
Service
Repository

---

# Errors

Use wrapped errors.

Example:

fmt.Errorf("create media: %w", err)

Never return raw database errors directly to clients.

---

# Repository

Repository handles:

- queries
- persistence

Repository must not contain business rules.

---

# Service

Service handles:

- recommendation logic
- search logic
- user behavior logic

---

# API

API handles:

- request validation
- response serialization

No business logic.

---

# Database

Prefer indexed queries.

Avoid table scans.

FTS must be used for text search.

---

# JSON

Follow:

.ai/api-contract.md