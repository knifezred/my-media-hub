# Project Governance

This document defines project governance rules.

All AI agents must read this file before reading any other project documents.

Applies to:

- Trae
- Cursor
- Claude Code
- Cline
- Roo Code
- GitHub Copilot
- Future AI coding agents

---

# Governance Model

The project follows a:

Single Source of Truth

model.

Project knowledge exists only in:

docs/

AI execution rules and status exist only in:

.ai/

---

# Documentation Ownership

## Human-Owned

The following directories are human-maintained:

docs/

Files under docs/ define:

- requirements
- architecture
- database design
- API design
- recommendation design
- roadmap
- UI design

These files are authoritative.

AI agents may:

- read
- analyze
- explain
- review
- suggest improvements

AI agents must NOT:

- rewrite
- regenerate
- overwrite

without explicit approval.

---

## AI-Owned

The following files are AI-maintained:

.ai/project-current-milestone.md

.ai/project-task-breakdown.md

AI agents may update these files automatically.

Purpose:

- track project status
- track task progress
- track implementation progress

---

# Documentation Authority

When conflicts exist:

docs/*

always overrides

.ai/*

---

# Bootstrap Rule

If either file is empty:

.ai/project-current-milestone.md

.ai/project-task-breakdown.md

AI agents should:

1. Read docs/Roadmap.md
2. Read project documentation
3. Generate initial milestone state
4. Generate task breakdown

before implementation begins.

---

# Milestone Governance

Source of Truth:

docs/Roadmap.md

AI agents must NOT:

- create milestones
- delete milestones
- reorder milestones

without approval.

---

# Architecture Governance

Source of Truth:

docs/SAD.md

AI agents must not:

- change architecture style
- introduce new layers
- redesign project structure

without approval.

---

# Database Governance

Source of Truth:

docs/ERD.md

AI agents must not:

- redesign schema
- rename entities
- modify relationships

without approval.

---

# API Governance

Source of Truth:

docs/API.md

AI agents must also follow:

.ai/api-contract.md

API compatibility must be preserved.

---

# Recommendation Governance

Source of Truth:

docs/RecommendationEngine.md

AI agents must not redesign recommendation strategy without approval.

---

# Final Principle

Project knowledge belongs to humans.

Project execution belongs to AI.