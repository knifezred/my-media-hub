# AGENTS.md

# My Media Hub Agent Instructions

This repository follows a strict Single Source of Truth model.

---

# Step 1

Read:

.ai/project-governance.md

before reading any other files.

Governance rules override all other instructions.

---

# Documentation Authority

Project knowledge is stored exclusively in:

docs/

These documents are authoritative.

Examples:

docs/PRD.md

docs/SAD.md

docs/ERD.md

docs/API.md

docs/RecommendationEngine.md

docs/Roadmap.md

docs/UI.md

---

# AI Documents

The .ai directory contains:

* governance
* coding rules
* API rules
* milestone status
* task status

The .ai directory does NOT define project knowledge.

---

# Required Reading Order

1. .ai/project-governance.md

2. AI_README.md

3. .ai/project-current-milestone.md

4. .ai/project-task-breakdown.md

5. docs/Roadmap.md

6. Relevant project documents

---

# Architecture

Follow docs/SAD.md

---

# Database

Follow docs/ERD.md

---

# API

Follow docs/API.md

.ai/api-contract.md

---

# Recommendation

Follow docs/RecommendationEngine.md

---

# Current Scope

Current implementation scope is defined by:

.ai/project-current-milestone.md

Do not implement features outside the active milestone.

---

# Task Scope

Current executable tasks are defined by:

.ai/project-task-breakdown.md

---

# Documentation Rules

Human-owned:

docs/*

AI may analyze but must not modify.

---

# AI-owned

.ai/project-current-milestone.md

.ai/project-task-breakdown.md

AI may update automatically.

---

# Final Rule

Project knowledge belongs to humans.

Project status belongs to AI.
