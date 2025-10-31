# Frontend - Flip Bank Statement Dashboard

Next.js React application for uploading and viewing bank statement transactions.

## 🚀 Quick Start

```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev
```

App runs on: **http://localhost:3000**

---

## 📋 Features

### Dashboard
- ✅ **Balance Card** - Total balance from `/api/balance` endpoint
- ✅ **Issues Card** - Count of failed/pending transactions from `/api/issues` endpoint
- ✅ **Transaction Table** - Full transaction list from `/api/transactions` endpoint
- ✅ **Upload Section** - CSV file upload with drag-and-drop support
- ✅ **Real-time Loading States** - Skeleton loaders and spinners during data fetch
- ✅ **Error Handling** - Toast notifications for all errors

### Upload Features
- ✅ **Drag & Drop** - Easy file upload
- ✅ **Loading Spinner** - Shows progress with primary color overlay during upload
- ✅ **File Reset** - Dropzone auto-resets after successful upload
- ✅ **Validation** - Client-side file type validation
- ✅ **Status Feedback** - Success/error notifications with upload statistics

### Table Features
- ✅ **Filtering** - By type (CREDIT/DEBIT) and status (SUCCESS/FAILED/PENDING)
- ✅ **Sorting** - Three-state sort (ASC → DESC → RESET)
- ✅ **Search** - Debounced search by name/description (300ms)
- ✅ **Pagination** - Navigate through transaction pages
- ✅ **Status Indicators** - Visual badges for transaction status
- ✅ **Empty States** - User-friendly message when no data
- ✅ **Scroll Preservation** - Maintains scroll position during filtering/sorting

### UI Components
- ✅ **Custom Select Dropdown** - For type/status filters
- ✅ **Dropzone** - With loading state and file preview
- ✅ **DataTable** - Responsive table with sorting icons
- ✅ **Notifications** - Toast alerts for success/error/warning
- ✅ **Skeleton Loaders** - Placeholder during data load
- ✅ **Responsive Design** - Mobile-friendly layout

---

## 🛠️ Development

### Available Scripts

```bash
npm run dev           # Start development server
npm run build         # Build production bundle
npm run start         # Start production server
npm run lint          # Lint code
```

### Project Structure

```
frontend/
├── app/
│   ├── (dashboard)/              # Dashboard pages
│   │   ├── demo/
│   │   ├── reports/
│   │   ├── settings/
│   │   ├── transactions/
│   │   └── upload/
│   ├── components/               # Reusable components
│   │   ├── Alert/
│   │   ├── Button/
│   │   ├── DataTable/            # Main transaction table
│   │   ├── Dropzone/             # File upload with loading state
│   │   ├── Input/
│   │   ├── Loader/
│   │   ├── Notification/
│   │   ├── Pagination/
│   │   ├── Search/               # Debounced search
│   │   ├── Select/               # Custom dropdown
│   │   ├── Sidebar/
│   │   └── Table/
│   ├── contexts/                 # React contexts
│   │   ├── NotificationContext
│   │   └── SidebarContext
│   ├── ui/
│   │   ├── Dashboard/
│   │   │   ├── DashboardView.tsx # Main dashboard logic
│   │   │   └── api.ts            # API client
│   │   └── hooks/
│   │       ├── useTable.tsx      # Table state & logic
│   │       └── useLoading.tsx    # Loading state management
│   ├── layout.tsx
│   ├── page.tsx
│   └── globals.css
├── package.json
└── README.md (this file)
```

### Hooks

**useTable** - Manages table state:
- Pagination (current page, page size)
- Sorting (three-state: ASC → DESC → NULL)
- Filtering (type, status)
- Search (debounced)
- Data fetching

**useLoading** - Manages loading states:
- `tableLoading` - Transaction table loading
- `statsLoading` - Balance/issues card loading
- `uploading` - File upload loading

### API Integration

All API calls use the centralized `api.ts` client:

```typescript
// From frontend/app/ui/Dashboard/api.ts
export async function fetchTransactions(params: {...}): Promise<...>
export async function fetchBalance(): Promise<...>
export async function fetchIssuesCount(): Promise<...>
export async function uploadFile(file: File): Promise<...>
```

Environment variable: `NEXT_PUBLIC_API_URL`
- Local: `http://localhost:9000/api`
- Production: Set via GitHub secret

---

## 🌍 Environment Variables

### Local Development
```bash
NEXT_PUBLIC_API_URL=http://localhost:9000/api
```

### Production (Google Cloud Run)
```bash
NEXT_PUBLIC_API_URL=https://flip-fullstack-test-backend-xxx.run.app/api
```

See [root README](../README.md#-environment-variables) and [DEPLOYMENT.md](../docs/DEPLOYMENT.md) for full setup.

---

## 📦 Technology Stack

- **Framework**: Next.js 15+
- **Language**: TypeScript
- **Styling**: CSS Modules
- **State Management**: React hooks
- **Icons**: lucide-react
- **HTTP Client**: Fetch API (built-in)

---

## 🚀 Deployment

### Build Production Bundle

```bash
npm run build
npm run start
```

### Docker

```bash
docker build -t flip-bank-frontend:latest .
docker run -p 3000:3000 \
  -e NEXT_PUBLIC_API_URL=https://your-backend-url.run.app/api \
  flip-bank-frontend:latest
```

### Google Cloud Run

Automatic deployment via GitHub Actions. See [DEPLOYMENT.md](../docs/DEPLOYMENT.md).

---

## 🚨 Troubleshooting

**Can't connect to backend?**
```bash
# Check NEXT_PUBLIC_API_URL is set
echo $NEXT_PUBLIC_API_URL

# Verify backend is running
curl http://localhost:9000/api/health
```

**Build fails?**
```bash
rm -rf node_modules package-lock.json
npm install
npm run build
```

**Port already in use?**
```bash
PORT=3001 npm run dev
```

---

**Happy coding! 🎉**

For more info, see [root README](../README.md) or [DEPLOYMENT.md](../docs/DEPLOYMENT.md).
