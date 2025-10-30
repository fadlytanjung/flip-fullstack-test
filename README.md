<div align="center">

# ✅ Full Stack Take‑Home Assignment

Bank Statement Viewer — Go + React/Next.js

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Node](https://img.shields.io/badge/Node-18+-339933?logo=node.js&logoColor=white)](https://nodejs.org)
[![License](https://img.shields.io/badge/License-MIT-purple.svg)](#)
[![Timebox](https://img.shields.io/badge/Timebox-4–6h-blue)](#)

</div>

---

### 📚 Table of Contents

- **Overview**
- **Input Format**
- **Requirements**
  - Backend (Go)
  - Frontend (React/Next.js)
  - Extras
- **API Contract**
- **Getting Started**
- **Frontend Components** ⭐ NEW
- **Documentation**
- **Submission**

---

### 📌 Overview

Build a small full‑stack app that lets users upload a bank statement CSV, view insights, and inspect transaction issues. No vibe coding allowed to answer this question.

---

### 📥 Input Format

```csv
# Example CSV payload
1624507883 , JOHN DOE, DEBIT, 250000 , SUCCESS, restaurant
1624608050 , E-COMMERCE A, DEBIT, 150000 , FAILED, clothes
1624512883 , COMPANY A, CREDIT, 12000000 , SUCCESS, salary
1624615065 , E-COMMERCE B, DEBIT, 150000 , PENDING, clothes

# Format
timestamp, name, type, amount, status, description
```

Notes:
- `timestamp` is a Unix epoch (seconds).
- `type` is one of `CREDIT` | `DEBIT`.
- `status` is one of `SUCCESS` | `FAILED` | `PENDING`.

---

### ✅ Requirements

#### Backend (Golang)

Build REST APIs with in‑memory storage. Clean architecture is a plus (handler → service → repository).

#### Frontend (React or Next.js)

Provide a simple UI to upload CSV, show computed end balance, and list non‑successful transactions with visual status.

#### Extras (Optional)

Dockerfile, request validation & error handling, CI (GitHub Actions), clean structure, and tests.

---

### 🔌 API Contract

| Method | Endpoint     | Description                                                                       |
|-------:|--------------|-----------------------------------------------------------------------------------|
|  POST  | `/upload`    | Accepts CSV file upload, parses it, stores transactions in memory                |
|   GET  | `/balance`   | Returns total balance = credits − debits (from SUCCESS transactions only)        |
|   GET  | `/issues`    | Returns non‑successful transactions (`FAILED` + `PENDING`)                        |

Response examples are encouraged but not required.

---

### 🎨 UI Status Styles

- ⚠ PENDING → warning style
- ❌ FAILED → danger/red style

Bonus UI features:
- Pagination or sorting
- Reusable components
- Pure CSS (no Tailwind/UI library)
- TypeScript support

---

### 🚀 Getting Started

#### Local Development

**Backend (Go):**
```bash
cd backend
go mod download
go run ./cmd/server
```

**Frontend (Next.js):**
```bash
cd frontend
npm install
npm run dev
```

#### Deployment to Google Cloud Run

**Quick Deploy (15 minutes):**
```bash
# 1. Run automated setup (5 min)
./setup-gcp.sh

# 2. Add GitHub secrets (3 min)
# Follow script output instructions

# 3. Deploy! (2 min each)
git push origin main  # Auto-deploys via GitHub Actions
```

**Deployment Features:**
- ✅ Separate Cloud Run services for frontend & backend
- ✅ Automated CI/CD via GitHub Actions
- ✅ Path-based triggers (only deploy what changed)
- ✅ Docker multi-stage builds for optimization
- ✅ Auto-scaling (0-10 instances) and HTTPS included
- ✅ ~$0-40/month cost (free tier available)

**Documentation:**
- [DEPLOYMENT.md](./DEPLOYMENT.md) - Complete deployment guide
- [setup-gcp.sh](./setup-gcp.sh) - Automated setup script

---

### 🎨 Frontend Components

The application includes a comprehensive set of UI components:

#### Alert Component
Static inline alerts for persistent feedback (success, error, warning, info)

#### Notification System
Toast notifications with global context for temporary messages

#### Smart Pagination
Intelligent page number display with ellipsis patterns

**Try them out:** Visit `/demo` after starting the app!

**Documentation:**
- [Alert & Notification Guide](./frontend/ALERT_NOTIFICATION_GUIDE.md)
- [Pagination Guide](./frontend/PAGINATION_GUIDE.md)
- [Component Showcase](./frontend/COMPONENT_SHOWCASE.md)
- [Quick Start Guide](./frontend/NEW_COMPONENTS.md)

---

### 📖 Documentation

| Type | Document | Description |
|------|----------|-------------|
| ⚡ Quick Start | [QUICK_REFERENCE.md](./QUICK_REFERENCE.md) | All commands in one place |
| 🚀 Deployment | [DEPLOYMENT.md](./DEPLOYMENT.md) | Complete deployment guide |
| 🛠️ Setup | [setup-gcp.sh](./setup-gcp.sh) | Automated GCP setup script |

---

### 📦 Submission

Provide:
- Link to a public GitHub repository
- README including setup instructions and architecture decisions
- Please finish within 4–6 hours, max 2 days after you received the email

Presentation: up to 5 minutes explaining how your solution works.
