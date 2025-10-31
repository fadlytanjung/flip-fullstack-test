<div align="center">

# ✅ Full Stack Take‑Home Assignment

Bank Statement Viewer — Go + React/Next.js

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Node](https://img.shields.io/badge/Node-20+-339933?logo=node.js&logoColor=white)](https://nodejs.org)
[![License](https://img.shields.io/badge/License-MIT-purple.svg)](#)
[![Timebox](https://img.shields.io/badge/Timebox-4–6h-blue)](#)

</div>

---

## 📚 Table of Contents

- [Overview](#overview)
- [Input Format](#input-format)
- [Requirements](#requirements)
- [API Contract](#api-contract)
- [Quick Start](#quick-start)
- [Development](#development)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Testing](#testing)
- [Deployment](#deployment)
- [Documentation](#documentation)

---

## 📌 Overview

Build a small full‑stack app that lets users upload a bank statement CSV, view insights, and inspect transaction issues. The application follows **Domain-Driven Design** with **Clean Architecture** principles.

**Tech Stack:**
- **Backend**: Go 1.23 with Fiber framework and SQLite
- **Frontend**: Next.js with React
- **Database**: SQLite (in-memory capable)
- **Architecture**: Domain-Driven Design (DDD)

---

## 📥 Input Format

```csv
# Example CSV payload with decimal amounts
1624507883,JOHN DOE,DEBIT,250000.50,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
1624512883,COMPANY A,CREDIT,12000000.99,SUCCESS,salary
1624615065,E-COMMERCE B,DEBIT,150000.25,PENDING,clothes

# Format
timestamp,name,type,amount,status,description
```

**Notes:**
- `timestamp` is a Unix epoch (seconds).
- `type` is one of `CREDIT` | `DEBIT`.
- `status` is one of `SUCCESS` | `FAILED` | `PENDING`.
- `amount` can be integer (250000) or decimal (250000.50) - stored as cents internally

---

## ✅ Requirements

### Backend (Golang)

Build REST APIs with in‑memory storage. Clean architecture is a plus (handler → service → repository).

**✅ Implemented:**
- Domain-Driven Design architecture
- SQLite database with automatic schema migration
- Comprehensive input validation
- Filtering, sorting, and pagination
- CORS support
- Error handling

### Frontend (React or Next.js)

Provide a simple UI to upload CSV, show computed end balance, and list non‑successful transactions with visual status.

**✅ Implemented:**
- Upload interface with drag-and-drop
- Dashboard with balance display (from `/api/balance`)
- Transaction list with filtering, sorting, pagination (using `/api/transactions`)
- Issues count from `/api/issues` endpoint
- Status indicators (SUCCESS/FAILED/PENDING)
- Empty state UI for tables
- Real-time API integration with error handling via notifications
- Skeleton/loader states driven by API responses
- Non-blocking file upload with concurrent processing
- **NEW:** Dropzone loading spinner with opacity overlay during upload
- **NEW:** Three-state sorting (ASC → DESC → RESET)
- **NEW:** Debounced search (300ms) for better performance
- **NEW:** Custom Select dropdown component for filters

### Extras (Optional)

Dockerfile, request validation & error handling, CI (GitHub Actions), clean structure, and tests.

**✅ Implemented:**
- ✅ Dockerfile (optimized multi-stage)
- ✅ Input validation (file and field validators)
- ✅ Error handling (comprehensive)
- ✅ Clean structure (DDD)
- ✅ Unit tests and coverage

---

## 🔌 API Contract

| Method | Endpoint     | Description                                                                       |
|-------:|--------------|-----------------------------------------------------------------------------------|
|  POST  | `/api/upload`    | Accepts CSV file upload, parses it, stores transactions in memory                |
|   GET  | `/api/balance`   | Returns total balance = credits − debits (from SUCCESS transactions only)        |
|   GET  | `/api/transactions` | Returns all transactions with filtering, sorting, and pagination             |
|   GET  | `/api/issues`    | Returns non‑successful transactions (`FAILED` + `PENDING`) with filtering/sorting |
|   GET  | `/api/health`    | Health check endpoint                                                             |
| DELETE | `/api/clear`     | Clear all transaction data                                                        |

**API Features:**
- ✅ **Decimal Amount Support** - CSV accepts decimal values (e.g., `1234.56`) stored as cents
- ✅ **Duplicate Detection** - Automatically detects and skips duplicate transactions
- ✅ Filtering by status, type, amount, date range
- ✅ Searching by name/description
- ✅ Sorting by any field (ASC/DESC, no default sort if not specified)
- ✅ Pagination with navigation links
- ✅ Comprehensive error responses

Response format matches frontend contract with pagination metadata and filters.

---

## 🚀 Quick Start

### Prerequisites
- Go 1.23.10+
- Node.js 20+
- Docker (optional)
- make (optional)

### Backend Setup (5 minutes)

```bash
cd backend

# Install tools
make install

# Copy environment configuration
cp .env.example .env

# Run server
make run
```

Server runs on `http://localhost:9000`

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev
```

Frontend runs on `http://localhost:3000`

---

## 🔐 Environment Variables

### Local Development

**Backend** (`backend/.env`)
```bash
# No configuration needed for local development
# Database defaults to in-memory SQLite
```

**Frontend** (`.env.local` or `frontend/.env.local`)
```bash
# Local development (default if not set)
NEXT_PUBLIC_API_URL=http://localhost:9000/api
```

### Production (Google Cloud Run)

**Frontend Environment Variable:**
```bash
NEXT_PUBLIC_API_URL=https://flip-fullstack-test-backend-xxx.run.app/api
```

**How it's configured:**

1. **GitHub Secrets** - Store sensitive values:
   ```bash
   gh secret set NEXT_PUBLIC_API_URL -b "https://your-backend-url.run.app"
   ```

2. **Docker Build** - Injected during image build:
   ```dockerfile
   ARG NEXT_PUBLIC_API_URL=http://localhost:9000/api
   ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL
   ```

3. **Cloud Run Deployment** - Set in service environment:
   ```bash
   gcloud run deploy flip-fullstack-test-frontend \
     --set-env-vars "NEXT_PUBLIC_API_URL=https://your-backend-url.run.app"
   ```

**Backend Environment Variable:**
```bash
ENV=production
PORT=9000
```

See [DEPLOYMENT.md](./docs/DEPLOYMENT.md) for complete production setup guide.

---

## 🔧 Development

### 🚀 Run Both Backend + Frontend with Hot Reload

From the **root** directory:

```bash
npm install
npm run dev
```

This starts:
- **Backend**: Go server with Air (hot reload) on `http://localhost:9000`
- **Frontend**: Next.js dev server on `http://localhost:3000`

Both will reload automatically when you make changes.

### 🔧 Individual Commands

**Backend Only:**

```bash
cd backend

# Development with hot reload (air)
make dev

# Or production mode (standard go run)
make run

# Install tools first
make install
```

**Frontend Only:**

```bash
cd frontend
npm run dev
```

### 📝 Available npm Scripts (from root)

```bash
npm run install:all    # Install all dependencies
npm run dev            # Development with hot reload (both)
npm run dev:backend    # Backend only with hot reload
npm run dev:frontend   # Frontend only
npm run build          # Build both backend & frontend
npm run build:backend  # Build backend binary
npm run build:frontend # Build frontend
npm run start          # Start both (production)
npm run start:backend  # Start backend (production)
npm run start:frontend # Start frontend (production)
npm run test           # Run backend tests
npm run coverage       # Generate coverage report
npm run lint           # Lint both backend & frontend
npm run lint:backend   # Lint backend only
npm run lint:frontend  # Lint frontend only
npm run format         # Format backend code
npm run tidy           # Tidy backend dependencies
```

### 📝 Available Make Commands (from `backend/` directory)

```bash
make help          # Show all available commands
make dev           # Development with hot reload (air)
make run           # Run server (production mode)
make build         # Build binary
make lint          # Run code linter
make format        # Format code
make test          # Run tests
make coverage      # Run tests with coverage report
make tidy          # Tidy dependencies
make install       # Install development tools
```

### 🌍 Environment Configuration

```bash
cd backend
cp .env.example .env
```

Example `.env`:
```env
PORT=9000
ENV=development
LOG_LEVEL=debug
DATABASE_URL=sqlite:transactions.db
MAX_FILE_SIZE=10485760
ALLOWED_FILE_TYPES=csv
DEFAULT_PAGE_SIZE=10
MAX_PAGE_SIZE=100
```

See `backend/.env.example` for all options.

### 🔄 Hot Reload

- **Backend**: Changes to Go files automatically rebuild with Air
- **Frontend**: Changes to Next.js files automatically reload

### Database

- **Type**: SQLite (file-based)
- **File**: `transactions.db` (auto-created)
- **Migration**: Automatic schema creation on startup

Reset database:
```bash
rm transactions.db
make run
```

---

## 🏗️ Project Structure

```
.
├── backend/                         # Go backend (Port 9000)
│   ├── cmd/server/
│   │   ├── main.go                 # Entry point
│   │   └── bootstrap.go            # DI & routing
│   ├── domain/
│   │   ├── transaction/            # Transaction domain
│   │   │   ├── handler/
│   │   │   ├── repository/
│   │   │   ├── schemas/
│   │   │   ├── use_case/
│   │   │   └── *_test.go           # Unit tests
│   │   └── upload/                 # Upload domain
│   │       ├── handler/
│   │       ├── repository/
│   │       ├── use_case/
│   │       └── *_test.go           # Unit tests
│   ├── pkg/
│   │   ├── db/                     # Database setup
│   │   ├── validator/              # Input validation
│   │   └── response/               # Generic response wrapper
│   ├── mocks/                      # Mock implementations for testing
│   ├── .env.example                # Environment template
│   ├── Dockerfile                  # Production image
│   ├── Makefile                    # Build commands
│   └── README.md                   # Backend docs
│
├── frontend/                        # Next.js frontend (Port 3000)
│   ├── app/
│   │   ├── (dashboard)/
│   │   ├── components/
│   │   └── layout.tsx
│   ├── package.json
│   └── README.md
│
├── docs/
│   └── DEPLOYMENT.md              # Deployment guide
├── .env.example                    # Root env example
├── README.md                       # This file
└── setup-gcp.sh                    # GCP deployment script
```

---

## 🏛️ Architecture

### Domain-Driven Design

The project follows DDD principles with clean architecture separation:

```
Handler → Use Case → Repository → Database
```

#### Transaction Domain
- Query transactions
- Calculate balance (credits - debits from SUCCESS transactions)
- Filter and sort issues
- Validate transaction data

#### Upload Domain
- Parse CSV files with validation
- Handle multipart file uploads
- Store transactions in database
- Return upload statistics

### Package Naming Convention

Package names match folder names:
```
Folder: domain/transaction/use_case/
Package: package use_case (not usecase)
Import: "github.com/.../use_case"
Alias: transactionUseCase "github.com/.../use_case"
```

### Validators

- **CSV Validator**: File extension, size, format
- **Field Validator**: Timestamp, name, type, amount, status, description

---

## 🧪 Testing

### Run Tests

```bash
# Run all tests
make test

# Run tests with coverage report
make coverage
```

### Test Coverage

Coverage reports include:
- HTML report: `coverage/coverage.html`
- Text summary in console output

### Unit Tests Included

Tests are provided for:
- **Transaction Domain**: Balance calculation, filtering, sorting
- **Upload Domain**: CSV parsing, validation, error handling
- **Validators**: Field and file validation

---

## 🔍 API Examples

### Upload CSV

```bash
curl -X POST http://localhost:9000/api/upload \
  -F "file=@transactions.csv"
```

### Get Balance

```bash
curl http://localhost:9000/api/balance
```

### Get Issues with Filters

```bash
# Filter by status
curl "http://localhost:9000/api/issues?status=FAILED"

# Filter by type
curl "http://localhost:9000/api/issues?type=DEBIT"

# Search
curl "http://localhost:9000/api/issues?search=restaurant"

# Sort and paginate
curl "http://localhost:9000/api/issues?sort_by=amount&sort_order=DESC&page=1&page_size=10"

# Combined filters
curl "http://localhost:9000/api/issues?type=DEBIT&status=PENDING&sort_by=amount&sort_order=DESC"
```

---

## 🐳 Docker

### Build Backend Image

```bash
docker build -t flip-bank-backend:latest -f backend/Dockerfile backend/
```

### Run Container

```bash
docker run -p 9000:9000 \
  -e PORT=9000 \
  -e ENV=production \
  flip-bank-backend:latest
```

### Docker Compose

```bash
docker compose up -d
```

---

## 🚀 Deployment

### Google Cloud Run

Quick deploy to Cloud Run (15 minutes):

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

See [docs/DEPLOYMENT.md](./docs/DEPLOYMENT.md) for complete guide.

---

## 📖 Documentation

### Backend Documentation

- **[backend/README.md](./backend/README.md)** - Architecture, setup, and development guide
- **[backend/API.md](./backend/API.md)** - Complete API reference with examples
- **[backend/.env.example](./backend/.env.example)** - Environment configuration template

### Frontend Documentation

- **[frontend/README.md](./frontend/README.md)** - Frontend setup and component guide

### Deployment Documentation

- **[docs/DEPLOYMENT.md](./docs/DEPLOYMENT.md)** - Complete deployment guide
- **[setup-gcp.sh](./setup-gcp.sh)** - Automated GCP setup script

---

## ✅ Development Checklist

### New Developer Setup

- [ ] Clone repository
- [ ] Backend setup: `cd backend && make install && cp .env.example .env && make dev`
- [ ] Frontend setup: `cd frontend && npm install && npm run dev`
- [ ] Combined setup: `cd frontend && npm install && npm run dev:all`
- [ ] Test backend: `curl http://localhost:9000/api/health`
- [ ] Test frontend: Visit `http://localhost:3000`
- [ ] Read [backend/README.md](./backend/README.md)
- [ ] Start developing!

### Before Committing

- [ ] Run `make format` in backend
- [ ] Run `make lint` in backend
- [ ] Run `make test` in backend
- [ ] Check `make coverage` report
- [ ] Frontend: `npm run lint` (if available)

---

## 🎨 UI Status Styles

- ⚠️ PENDING → warning style (yellow)
- ❌ FAILED → danger/red style
- ✅ SUCCESS → success/green style

**Bonus UI Features:**
- Pagination with smart page numbers
- Reusable component library
- Pure CSS (no Tailwind/UI library)
- TypeScript support
- Responsive design

---

## 📦 Technologies

### Backend
- **Language**: Go 1.23.10
- **Framework**: Fiber v2.52.5
- **Database**: SQLite with GORM ORM
- **Validation**: Custom validators
- **Architecture**: Domain-Driven Design

### Frontend
- **Framework**: Next.js
- **Language**: TypeScript
- **Styling**: CSS Modules
- **Components**: Reusable component system

### DevOps
- **Containerization**: Docker
- **Orchestration**: Docker Compose
- **CI/CD**: GitHub Actions
- **Deployment**: Google Cloud Run
- **Development**: Concurrently, Air (Go hot reload)

---

## 🚨 Troubleshooting

### Backend

**Port already in use:**
```bash
PORT=3000 make run
```

**Build fails:**
```bash
go clean -modcache
make install
make build
```

**Database issues:**
```bash
rm transactions.db
make run
```

### Frontend

**Dependencies issue:**
```bash
rm -rf node_modules package-lock.json
npm install
```

---

## 📋 Notes

- This is a demonstration project for a fullstack take-home assignment
- All code follows best practices and clean architecture principles
- The project is designed to be production-ready but simplified for demo purposes
- Time-boxed: 4-6 hours for implementation

---

## 📝 License

MIT License - See repository for details

---

**Ready to get started? See [Quick Start](#quick-start) or [Development](#development) above!** 🚀

For questions or issues, refer to the backend or frontend README files.
