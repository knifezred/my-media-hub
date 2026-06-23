# API Contract Rules

Version: 1.0

Applies To:

* Go Backend
* Vue Frontend
* Trae
* Cursor
* Claude Code
* GitHub Copilot
* All AI Agents

---

# Purpose

This document defines the API contract specification.

Goals:

* Stable API responses
* Predictable frontend behavior
* AI-generated code consistency
* Backward compatibility

If conflicts exist:

docs/API.md

defines business APIs.

This document defines response contracts.

---

# Core Principles

## Contract First

Frontend and backend must follow the same response schema.

---

## Stability First

Never remove fields.

Never rename fields.

Never change field meaning.

Prefer additive changes.

---

## Consistency First

All APIs must follow the same response format.

---

# Standard Success Response

All successful APIs must return:

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

# Success Response Types

## Empty Response

Used for:

* create
* update
* delete
* mark viewed
* clear history
* hide media

Example:

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## Object Response

Example:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "Example"
  }
}
```

---

## List Response

Example:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## Pagination Response

Example:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [],
    "total": 0,
    "page": 1,
    "page_size": 20
  }
}
```

---

# Standard Error Response

All failed APIs must return:

```json
{
  "code": 10001,
  "message": "media not found",
  "data": {}
}
```

---

# Forbidden Error Formats

Forbidden:

```json
{
  "error": "media not found"
}
```

Forbidden:

```json
{
  "success": false
}
```

Forbidden:

```json
{
  "msg": "media not found"
}
```

---

# Error Code Rules

Error codes are numeric integers.

`code: 0` always means success.

Non-zero codes follow module-based ranges:

| Range | Category |
|-------|----------|
| 0 | Success |
| 10001-19999 | Media errors |
| 20001-29999 | User behavior errors |
| 30001-39999 | Search errors |
| 40001-49999 | Request errors |
| 50001-59999 | Scanner errors |
| 90001-99999 | System errors |

See the full error code registry in:

docs/API.md

---

# Nullable Rules

## Arrays

Never:

```json
{
  "items": null
}
```

Always:

```json
{
  "items": []
}
```

---

## String

Never return:

```json
{
  "title": null
}
```

Always:

```json
{
  "title": ""
}
```

---

## Number

Never return:

```json
{
  "rating": null
}
```

Always:

```json
{
  "rating": 0
}
```

---

## Boolean

Never return:

```json
{
  "favorite": null
}
```

Always:

```json
{
  "favorite": false
}
```

---

## Object

Never return:

```json
{
  "metadata": null
}
```

Always:

```json
{
  "metadata": {}
}
```

---

# JSON Field Rules

All fields must be stable.

Do not omit fields because values are empty.

Bad:

```json
{
  "title": "Example"
}
```

Good:

```json
{
  "title": "Example",
  "description": "",
  "tags": [],
  "metadata": {}
}
```

---

# Go DTO Rules

## Forbidden

```go
Title string `json:"title,omitempty"`
```

```go
Items []Media `json:"items,omitempty"`
```

```go
Metadata map[string]string `json:"metadata,omitempty"`
```

---

## Required

```go
Title string `json:"title"`
```

```go
Items []Media `json:"items"`
```

```go
Metadata map[string]string `json:"metadata"`
```

---

# Slice Rules

Before response serialization:

Bad:

```go
var items []Media
```

Result:

```json
{
  "items": null
}
```

Forbidden.

---

Required:

```go
items := make([]Media, 0)
```

Result:

```json
{
  "items": []
}
```

---

# Pagination Rules

All list APIs should support:

```json
{
  "page": 1,
  "page_size": 20
}
```

Response:

```json
{
  "items": [],
  "total": 0,
  "page": 1,
  "page_size": 20
}
```

---

# Time Format Rules

Use:

RFC3339

Example:

```json
{
  "created_at": "2026-06-23T12:00:00Z"
}
```

---

# ID Rules

Use:

```json
{
  "id": 1
}
```

Type:

```text
INTEGER
```

Never expose internal rowid.

---

# API Versioning

Current version:

```text
/v1
```

Example:

```http
/api/v1/media
```

---

# Backward Compatibility

Allowed:

* Add fields
* Add endpoints
* Add query parameters

---

Forbidden:

* Remove fields
* Rename fields
* Change field type
* Change response structure

Without explicit approval.

---

# Frontend Compatibility Rules

Frontend should assume:

* fields always exist
* arrays always exist
* objects always exist

Frontend should not need:

```ts
if (data?.items)
```

Instead:

```ts
data.items.map(...)
```

must always be safe.

---

# AI Development Rules

AI agents must:

* follow docs/API.md
* follow this contract

AI agents must not:

* invent response formats
* invent pagination formats
* introduce nullable arrays
* introduce omitempty

without approval.

---

# Response DTO

All API responses must use DTO objects.

Do not return database entities directly.

Do not expose GORM/SQL models directly.

Repository Entity
    ↓
Service DTO
    ↓
API Response DTO

# Final Rule

API responses are contracts.

Contracts must remain stable.

A stable API is more important than a perfect API.
