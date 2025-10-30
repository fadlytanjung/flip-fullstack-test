# 🚀 Deployment Guide - Google Cloud Run

Deploy the Flip Fullstack Test application (frontend + backend) to Google Cloud Run with automated CI/CD via GitHub Actions.

---

## 📋 Table of Contents

1. [Prerequisites](#prerequisites)
2. [Quick Setup (5 minutes)](#quick-setup-5-minutes)
3. [GitHub Configuration](#github-configuration)
4. [Deploy](#deploy)
5. [Verify Deployment](#verify-deployment)
6. [Troubleshooting](#troubleshooting)
7. [How It Works](#how-it-works)

---

## Prerequisites

- ✅ Google Cloud account with billing enabled
- ✅ `gcloud` CLI installed and logged in
- ✅ GitHub repository admin access

---

## Quick Setup (5 minutes)

Run the automated setup script:

```bash
./setup-gcp.sh
```

This script will:
1. Check if you're logged into gcloud
2. Create or use existing project `flip-fullstack-test`
3. Enable required APIs
4. Create Artifact Registry repository
5. Create service account for GitHub Actions
6. Generate authentication key

**After running, you'll get:**
- Project ID
- Region
- Service account key (JSON file)

---

## GitHub Configuration

### Add Secrets to GitHub

Go to: **Settings > Secrets and variables > Actions** and add:

| Secret Name | Value | Where to get it |
|-------------|-------|-----------------|
| `GCP_PROJECT_ID` | `flip-fullstack-test` | From setup script output |
| `GCP_REGION` | `asia-southeast1` | From setup script output |
| `GCP_SA_KEY` | `{...JSON...}` | Copy entire `github-actions-key.json` content |

**Using GitHub CLI:**
```bash
gh secret set GCP_PROJECT_ID -b "flip-fullstack-test"
gh secret set GCP_REGION -b "asia-southeast1"
gh secret set GCP_SA_KEY < github-actions-key.json
```

**⚠️ IMPORTANT:** Delete the key file after adding to GitHub:
```bash
rm github-actions-key.json
```

---

## Deploy

### Step 1: Deploy Backend First

```bash
# Make sure backend code is ready
git add backend/
git commit -m "deploy: backend to cloud run"
git push origin main
```

**Wait for deployment** (check GitHub Actions tab)

**Get backend URL:**
```bash
gcloud run services describe flip-fullstack-test-backend \
  --region=asia-southeast1 \
  --format='value(status.url)'
```

Copy this URL (e.g., `https://flip-fullstack-test-backend-xxx.run.app`)

### Step 2: Configure Frontend with Backend URL

```bash
# Add backend URL to GitHub secrets
gh secret set NEXT_PUBLIC_API_URL -b "https://your-backend-url.run.app"

# Or via GitHub UI: Add secret NEXT_PUBLIC_API_URL
```

### Step 3: Deploy Frontend

```bash
git add frontend/
git commit -m "deploy: frontend to cloud run"
git push origin main
```

**Get frontend URL:**
```bash
gcloud run services describe flip-fullstack-test-frontend \
  --region=asia-southeast1 \
  --format='value(status.url)'
```

### 🎉 Done! Your app is live!

- **Frontend:** `https://flip-fullstack-test-frontend-xxx.run.app`
- **Backend:** `https://flip-fullstack-test-backend-xxx.run.app`

---

## Verify Deployment

### Check Services
```bash
gcloud run services list --region=asia-southeast1
```

### Test Backend
```bash
# Health check
curl https://your-backend-url.run.app/health
```

### Test Frontend
Open the frontend URL in your browser and verify:
- [ ] Homepage loads
- [ ] Components render correctly
- [ ] No console errors

### View Logs
```bash
# Frontend logs
gcloud run services logs tail flip-fullstack-test-frontend --region=asia-southeast1

# Backend logs
gcloud run services logs tail flip-fullstack-test-backend --region=asia-southeast1
```

---

## Troubleshooting

### Deployment Failed?

**1. Check GitHub Actions logs:**
- Go to: `https://github.com/YOUR_USERNAME/flip-fullstack-test/actions`
- Click on failed workflow
- Review error messages

**2. Common issues:**

**Authentication error:**
```bash
# Verify secrets are set
gh secret list

# Should show: GCP_PROJECT_ID, GCP_REGION, GCP_SA_KEY
```

**Service account permissions:**
```bash
# Verify service account has correct roles
gcloud projects get-iam-policy flip-fullstack-test \
  --flatten="bindings[].members" \
  --filter="bindings.members:serviceAccount:github-actions@*"
```

**Artifact Registry not found:**
```bash
# Check if repository exists
gcloud artifacts repositories describe flip-fullstack-test \
  --location=asia-southeast1
```

### Frontend Can't Reach Backend?

```bash
# 1. Verify NEXT_PUBLIC_API_URL is set
gh secret list

# 2. Check backend is running
gcloud run services describe flip-fullstack-test-backend \
  --region=asia-southeast1

# 3. Test backend directly
curl https://your-backend-url.run.app/health
```

### View Recent Deployments
```bash
# List recent GitHub Actions runs
gh run list --limit 5

# View specific run
gh run view RUN_ID --log
```

---

## How It Works

### Architecture

```
┌─────────────────────────────────────────────────────────┐
│  GitHub Repository (flip-fullstack-test)                │
│  ├── frontend/  → Cloud Run: flip-fullstack-test-frontend
│  └── backend/   → Cloud Run: flip-fullstack-test-backend
└─────────────────────────────────────────────────────────┘
```

### CI/CD Workflow

```
1. Push code to main branch
         ↓
2. GitHub detects changes in frontend/ or backend/
         ↓
3. Triggers corresponding workflow:
   - deploy-frontend.yml (if frontend/ changed)
   - deploy-backend.yml (if backend/ changed)
         ↓
4. Workflow builds Docker image
         ↓
5. Pushes to Artifact Registry
         ↓
6. Deploys to Cloud Run
         ↓
7. Service is live with auto-scaling!
```

### Service Configuration

Both services are configured with:
- **Memory:** 512 Mi
- **CPU:** 1 vCPU
- **Min instances:** 0 (scale to zero when idle)
- **Max instances:** 10
- **Timeout:** 300s
- **Port:** 3000 (frontend), 8080 (backend)

### Cost Estimate

- **Free tier:** 2M requests/month, 360k GB-seconds
- **Typical cost:** $0-40/month depending on traffic
- **Scale to zero:** $0 when no traffic

---

## Useful Commands

### Service Management

```bash
# List all services
gcloud run services list --region=asia-southeast1

# Get service URL
gcloud run services describe SERVICE_NAME \
  --region=asia-southeast1 \
  --format='value(status.url)'

# Update environment variable
gcloud run services update SERVICE_NAME \
  --region=asia-southeast1 \
  --set-env-vars "KEY=VALUE"

# Delete service
gcloud run services delete SERVICE_NAME \
  --region=asia-southeast1
```

### Logs & Monitoring

```bash
# Real-time logs
gcloud run services logs tail SERVICE_NAME --region=asia-southeast1

# Historical logs (last 100 lines)
gcloud run services logs read SERVICE_NAME \
  --region=asia-southeast1 \
  --limit=100
```

### GitHub Actions

```bash
# List recent runs
gh run list

# View specific run
gh run view RUN_ID

# Watch running workflow
gh run watch

# Trigger manual deployment
gh workflow run deploy-frontend.yml
gh workflow run deploy-backend.yml
```

---

## Ongoing Deployments

After initial setup, deployments are automatic:

```bash
# Just commit and push!
git add .
git commit -m "feat: new feature"
git push origin main

# GitHub Actions will:
# - Detect which service changed
# - Build and deploy automatically
# - No manual steps needed!
```

### Deploy Only Frontend
```bash
git add frontend/
git commit -m "feat: update UI"
git push origin main
# Only frontend deploys
```

### Deploy Only Backend
```bash
git add backend/
git commit -m "feat: new API endpoint"
git push origin main
# Only backend deploys
```

---

## Manual Deployment (Optional)

If you need to deploy manually without GitHub Actions:

### Backend
```bash
cd backend
gcloud run deploy flip-fullstack-test-backend \
  --source . \
  --region=asia-southeast1 \
  --allow-unauthenticated \
  --port=8080
```

### Frontend
```bash
cd frontend
gcloud run deploy flip-fullstack-test-frontend \
  --source . \
  --region=asia-southeast1 \
  --allow-unauthenticated \
  --port=3000 \
  --set-env-vars "NEXT_PUBLIC_API_URL=https://your-backend-url.run.app"
```

---

## Cleanup

To delete all resources:

```bash
# Delete services
gcloud run services delete flip-fullstack-test-frontend --region=asia-southeast1 --quiet
gcloud run services delete flip-fullstack-test-backend --region=asia-southeast1 --quiet

# Delete Artifact Registry repository
gcloud artifacts repositories delete flip-fullstack-test --location=asia-southeast1 --quiet

# Delete service account
gcloud iam service-accounts delete github-actions@flip-fullstack-test.iam.gserviceaccount.com --quiet

# Delete project (optional - deletes everything)
gcloud projects delete flip-fullstack-test --quiet
```

---

## Resources

- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Next.js Deployment](https://nextjs.org/docs/deployment)
- [Go Deployment](https://go.dev/doc/)

---

**Need help?** Check the troubleshooting section above or review GitHub Actions logs for detailed error messages.
