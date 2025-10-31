# Backend - Flip Bank Statement API

Golang REST API built with Fiber, SQLite, and Domain-Driven Design.

## 🚀 Quick Start

```bash
cd backend

# Install dependencies & tools
make install

# Copy environment configuration
cp .env.example .env

# Run development server (with hot reload)
air

# Or without hot reload
make run
```

Server runs on: **http://localhost:9000**

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

**Full API documentation:** See root [README.md](../README.md#-api-contract) or [API.md](./API.md)

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

- [Full API Reference](./API.md) - Complete endpoint documentation
- [Root README](../README.md) - Architecture & deployment guide
- [Makefile Targets](./Makefile) - All available commands

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

For questions, see the [root README](../README.md) or full [API documentation](./API.md).
