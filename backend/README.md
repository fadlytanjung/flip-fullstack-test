# Backend - Flip Bank Statement API

Golang REST API built with Fiber, SQLite, and Domain-Driven Design with enterprise-grade configuration and logging.

## 🚀 Quick Start

```bash
cd backend

# Install dependencies & tools
make install

# Option 1: Run with .env file
cp .env.example .env  # Create your .env file
./app

# Option 2: Run with environment variables
PORT=8080 LOG_LEVEL=debug ./app

# Option 3: Run development server (with hot reload)
air
```

Server runs on: **http://localhost:9000** (or your custom port)

### Environment Configuration

**Configuration Priority (high to low):**
1. **Environment variables** (highest priority) - Used in production
2. **`.env` file** - Used for local development
3. **Default values** (lowest priority) - Fallback for development

**Local Development:**
```bash
# Option 1: Create .env file
cat > .env << EOF
ENV=development
PORT=9000
LOG_LEVEL=debug
DATABASE_PATH=transactions.db
CORS_ALLOW_ORIGINS=*
EOF

# Option 2: Use defaults (no .env needed)
./app  # Uses concrete defaults from config/defaults.go
```

**Production Deployment:**
- Cloud Run automatically provides: `PORT` (injected by platform)
- GitHub Actions sets: `ENV=production` (always)
- Other env vars set via `--set-env-vars` in deployment:
  - `LOG_LEVEL=info`
  - `DATABASE_PATH=/tmp/transactions.db`
  - `CORS_ALLOW_ORIGINS=*`

**Available Environment Variables:**
| Variable | Default (Dev) | Production | Description |
|----------|---------------|------------|-------------|
| `ENV` | `development` | `production` | Environment mode (MUST be production in deployment) |
| `PORT` | `9000` | Auto (Cloud Run) | Server port |
| `LOG_LEVEL` | `info` | `info` | Logging level: debug/info/warn/error |
| `DATABASE_PATH` | `transactions.db` | `/tmp/transactions.db` | SQLite database path |
| `CORS_ALLOW_ORIGINS` | `*` | `*` | CORS allowed origins |
| `LOG_HOST_IP` | `""` | - | UDP log server IP (optional) |
| `LOG_HOST_PORT` | `0` | - | UDP log server port (optional) |

See [docs/CONFIG.md](docs/CONFIG.md) for full configuration guide.

---

## 📋 API Endpoints

| Method | Endpoint     | Description |
|--------|--------------|-------------|
| GET    | `/api/health` | Health check |
| POST   | `/api/upload` | Upload CSV file (supports decimal amounts) |
| GET    | `/api/balance` | Get account balance |
| GET    | `/api/transactions` | Get all transactions with filtering, sorting, pagination |
| GET    | `/api/issues` | List non-successful transactions |
| DELETE | `/api/clear` | Clear all data |

**Full API documentation:** See root [README.md](../README.md#-api-contract)

### API Features

- ✅ **Decimal Amount Support**: CSV can use decimal values (e.g., `1234.56`) - stored as cents internally
- ✅ **Duplicate Detection**: Automatically detects and skips duplicate transactions
- ✅ **Filtering**: By status, type, amount, date range
- ✅ **Searching**: By name/description
- ✅ **Sorting**: ASC/DESC by any field (no default sort applied when not specified)
- ✅ **Pagination**: With navigation links
- ✅ **Error Handling**: Comprehensive validation and error responses

---

## 🛠️ Development

### Available Commands

```bash
make help          # Show all commands
make run           # Run server
make build         # Build binary
make test          # Run tests
make coverage      # Generate coverage report
make lint          # Lint code
make format        # Format code
make tidy          # Tidy dependencies
make install       # Install tools
```

### Project Structure

```
backend/
├── cmd/server/              # Application entry point
├── domain/                  # DDD modules
│   ├── transaction/        # Balance & issues
│   └── upload/             # CSV upload
├── pkg/                    # Shared packages
│   ├── db/                # Database setup
│   ├── validator/         # Input validation
│   └── response/          # Generic response wrapper
├── mocks/                 # Mock implementations for testing
├── Makefile              # Build commands
└── .env.example          # Configuration template
```

### Architecture

- **Handler**: HTTP request handlers
- **Use Case**: Business logic
- **Repository**: Data access layer
- **Schemas**: Domain models

---

## 📦 Technology Stack

- **Language**: Go 1.23+
- **Framework**: Fiber v2
- **Database**: SQLite with GORM
- **Testing**: Go testing + testify
- **Validation**: Custom validators

---

## 🧪 Testing

```bash
# Run all tests
make test

# Generate coverage report (HTML + summary)
make coverage

# View coverage
open coverage/coverage.html
```

**Coverage**: 55%+ of critical code paths

---

## 🐳 Docker

```bash
# Build image
docker build -t flip-bank-backend:latest .

# Run container
docker run -p 9000:9000 flip-bank-backend:latest
```

---

## 🌍 Environment Variables

```env
PORT=9000              # Server port
ENV=development        # Environment
LOG_LEVEL=debug        # Log level
DATABASE_URL=sqlite:transactions.db
MAX_FILE_SIZE=10485760 # 10MB
```

See `.env.example` for all options.

---

## 📖 Additional Resources

- [Root README](../README.md) - Full API contract, architecture & deployment guide
- [Frontend README](../frontend/README.md) - Frontend setup and component guide
- [Deployment Guide](../docs/DEPLOYMENT.md) - Production deployment instructions

---

## 🚨 Troubleshooting

**Port already in use**
```bash
PORT=3000 make run
```

**Build fails**
```bash
go clean -modcache
make install
make build
```

**Database locked**
```bash
rm transactions.db
make run
```

---

**Happy coding! 🎉**

For questions, see the [root README](../README.md) for full API documentation.
