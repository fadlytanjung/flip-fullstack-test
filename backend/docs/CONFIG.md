# 🔧 Configuration Guide

Comprehensive guide for configuring the backend application.

---

## 📋 Table of Contents

1. [Overview](#overview)
2. [Configuration System](#configuration-system)
3. [Environment Variables](#environment-variables)
4. [Local Development](#local-development)
5. [Production Deployment](#production-deployment)
6. [Configuration Priority](#configuration-priority)
7. [Adding New Config](#adding-new-config)

---

## Overview

The backend uses a **three-tier configuration system**:

1. **Concrete Defaults** - Hardcoded values in `pkg/config/defaults.go`
2. **`.env` File** - Optional file for local development
3. **System Environment Variables** - Used in production (Cloud Run)

**Key Principle:** 
- **Development** uses defaults (with optional `.env` override)
- **Production** uses system environment variables (no `.env` needed)

---

## Configuration System

### Architecture

```
pkg/config/
├── types.go       # Config struct definitions
├── defaults.go    # Concrete default values
└── config.go      # Loading and validation logic
```

### How It Works

```go
// 1. Load defaults
loadDefaults()

// 2. Try to read .env file (optional)
viper.AddConfigPath(".")
viper.SetConfigName(".env")
viper.ReadInConfig()  // Ignores error if file not found

// 3. Bind environment variables (override .env)
viper.AutomaticEnv()
viper.BindEnv("ENV")
viper.BindEnv("PORT")
// ... etc

// 4. Unmarshal into struct
config := &GlobalConfig{}
viper.Unmarshal(config)
```

---

## Environment Variables

### Complete List

| Variable | Type | Default (Dev) | Production | Description |
|----------|------|---------------|------------|-------------|
| `ENV` | string | `development` | `production` | Environment mode |
| `DEBUG` | bool | `false` | `false` | Debug mode (computed from LOG_LEVEL) |
| `PORT` | int | `9000` | Auto (Cloud Run) | Server port |
| `LOG_LEVEL` | string | `info` | `info` | Logging level: debug/info/warn/error |
| `SERVICE_NAME` | string | `flip-fullstack-test-backend` | Same | Service name for logs |
| `SERVICE_VERSION` | string | From `VERSION` file | Same | Service version |
| `DATABASE_PATH` | string | `transactions.db` | `/tmp/transactions.db` | SQLite database file path |
| `LOG_HOST_IP` | string | `""` | - | Optional UDP log server IP |
| `LOG_HOST_PORT` | int | `0` | - | Optional UDP log server port |
| `CORS_ALLOW_ORIGINS` | string | `*` | `*` | CORS allowed origins |

### Required vs Optional

**Required (must have values):**
- `ENV` - Always set (default: development)
- `PORT` - Always set (default: 9000, Cloud Run provides it)
- `SERVICE_NAME` - Always set (default: flip-fullstack-test-backend)

**Optional (can be empty):**
- `LOG_HOST_IP` - Only needed for UDP log shipping
- `LOG_HOST_PORT` - Only needed for UDP log shipping

---

## Local Development

### Option 1: Use Defaults (Recommended)

No configuration needed! Just run:

```bash
cd backend
./app
```

This uses concrete defaults from `pkg/config/defaults.go`:
- `ENV=development`
- `PORT=9000`
- `LOG_LEVEL=info`
- `DATABASE_PATH=transactions.db`
- `CORS_ALLOW_ORIGINS=*`

### Option 2: Create `.env` File

For custom local configuration:

```bash
# Create .env file
cat > .env << EOF
ENV=development
PORT=8080
LOG_LEVEL=debug
DATABASE_PATH=./data/transactions.db
CORS_ALLOW_ORIGINS=http://localhost:3000
EOF

# Run (will use .env values)
./app
```

### Option 3: Environment Variables

Override specific values:

```bash
# Override port
PORT=8080 ./app

# Override log level
LOG_LEVEL=debug ./app

# Multiple overrides
PORT=8080 LOG_LEVEL=debug DATABASE_PATH=./custom.db ./app
```

---

## Production Deployment

### Cloud Run Configuration

**Environment variables are set in GitHub Actions:**

```yaml
# .github/workflows/deploy-cloudrun.yml
gcloud run deploy flip-fullstack-test-backend \
  --set-env-vars "ENV=production,LOG_LEVEL=info,DATABASE_PATH=/tmp/transactions.db,CORS_ALLOW_ORIGINS=*"
```

**Key Points:**

1. **`ENV=production`** is ALWAYS set in deployment
2. **`PORT`** is automatically provided by Cloud Run (injected at runtime)
3. **`DATABASE_PATH=/tmp/transactions.db`** uses Cloud Run's writable temp storage
4. **No `.env` file** is included in the Docker image or needed in production
5. **All config comes from system environment variables**

### Why `/tmp/transactions.db`?

Cloud Run containers have:
- **Read-only filesystem** - Can't write to `/app/transactions.db`
- **Writable `/tmp`** - Only location for temporary file storage
- **Ephemeral storage** - Data is lost when container stops (OK for demo/test)

For production with persistent data, use Cloud SQL or Cloud Storage.

---

## Configuration Priority

**Order of precedence (highest to lowest):**

```
1. System Environment Variables (export ENV=production)
   ↓ overrides
2. .env File (ENV=development)
   ↓ overrides
3. Default Values (viper.SetDefault("ENV", "development"))
```

**Examples:**

```bash
# Scenario 1: No .env, no env vars
# Result: ENV=development (default)

# Scenario 2: Has .env with ENV=staging
# Result: ENV=staging (.env overrides default)

# Scenario 3: Has .env with ENV=staging, but ENV=production exported
# Result: ENV=production (env var overrides .env)
```

---

## Adding New Config

### Step 1: Add to `types.go`

```go
type GlobalConfig struct {
    // ... existing fields ...
    
    // New field
    NewFeatureEnabled bool   `mapstructure:"NEW_FEATURE_ENABLED"`
    NewFeatureTimeout int    `mapstructure:"NEW_FEATURE_TIMEOUT"`
}
```

### Step 2: Add default in `defaults.go`

```go
func loadDefaults() {
    // ... existing defaults ...
    
    // New defaults
    viper.SetDefault("NEW_FEATURE_ENABLED", false)
    viper.SetDefault("NEW_FEATURE_TIMEOUT", 30)
}
```

### Step 3: Use in code

```go
cfg, _ := config.LoadConfig()

if cfg.NewFeatureEnabled {
    timeout := time.Duration(cfg.NewFeatureTimeout) * time.Second
    // ... use config ...
}
```

### Step 4: Update deployment (if needed for production)

```yaml
# .github/workflows/deploy-cloudrun.yml
--set-env-vars "ENV=production,...,NEW_FEATURE_ENABLED=true"
```

---

## Best Practices

### ✅ DO

- Use concrete defaults for local development
- Use system env vars for production
- Keep `.env` in `.gitignore` (already done)
- Document all config in this file
- Use meaningful default values
- Set `ENV=production` in all deployments

### ❌ DON'T

- Don't commit `.env` files
- Don't use `.env` in production
- Don't hardcode secrets in code
- Don't assume config exists (always use LoadConfig())
- Don't change defaults without documenting
- Don't deploy without `ENV=production`

---

## Troubleshooting

### Config not loading?

```go
cfg, err := config.LoadConfig()
if err != nil {
    log.Fatal("Failed to load config:", err)
}
```

### Check what config is loaded:

```go
cfg := config.GetConfig()
fmt.Printf("Environment: %s\n", cfg.Environment)
fmt.Printf("Port: %d\n", cfg.Port)
fmt.Printf("Log Level: %s\n", cfg.LogLevel)
```

### Verify environment variables:

```bash
# Check in container
docker exec -it <container> env | grep ENV
docker exec -it <container> env | grep PORT

# Check in Cloud Run logs
gcloud run services logs tail flip-fullstack-test-backend
```

---

## References

- [Viper Documentation](https://github.com/spf13/viper)
- [Cloud Run Environment Variables](https://cloud.google.com/run/docs/configuring/environment-variables)
- [12-Factor App Config](https://12factor.net/config)
