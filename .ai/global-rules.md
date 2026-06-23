# Global Development Rules

These rules apply to all code.

---

# Core Principle

Prefer consistency over creativity.

Prefer existing patterns over new abstractions.

---

# Scope Control

Implement only the current milestone.

Do not implement future roadmap items.

---

# Simplicity

Prefer:

- simple code
- clear code
- maintainable code

Avoid:

- over-engineering
- premature optimization
- speculative design

---

# Naming

Use domain language.

Preferred:

Media
Tag
Category
Favorite
Rating
Profile
Recommendation

Avoid:

Helper
Util
Tool
Manager
Common

---

# Error Handling

All errors must be handled.

Never silently ignore errors.

---

# Logging

Log:

- startup
- shutdown
- fatal errors

Avoid excessive logging.

---

# Testing

New business logic should be testable.

Critical logic should include tests.

---

# Performance

Target platform:

Intel N100
8GB RAM

Optimize for:

- memory efficiency
- query performance
- startup speed

Avoid unnecessary allocations.

---

# Security

Validate all external input.

Never trust request data.

Always use parameterized queries.

---

# Documentation

Code should explain itself.

Comments should explain WHY, not WHAT.