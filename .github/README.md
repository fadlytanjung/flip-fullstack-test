# GitHub Actions Workflows

This directory contains CI/CD workflows for automated deployment to Google Cloud Run.

## Workflows

### 🎨 Frontend Deployment
**File:** `workflows/deploy-frontend.yml`

[![Deploy Frontend](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/deploy-frontend.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/deploy-frontend.yml)

**Triggers:**
- Push to `main` or `develop` with changes in `frontend/`
- Manual workflow dispatch

**What it does:**
1. Builds Docker image for Next.js app
2. Pushes to Google Artifact Registry
3. Deploys to Cloud Run service: `bank-viewer-frontend`

---

### ⚙️ Backend Deployment
**File:** `workflows/deploy-backend.yml`

[![Deploy Backend](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/deploy-backend.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/deploy-backend.yml)

**Triggers:**
- Push to `main` or `develop` with changes in `backend/`
- Manual workflow dispatch

**What it does:**
1. Builds Docker image for Go app
2. Pushes to Google Artifact Registry
3. Deploys to Cloud Run service: `bank-viewer-backend`

---

## Required Secrets

Configure these in: **Settings > Secrets and variables > Actions**

| Secret Name | Description |
|-------------|-------------|
| `GCP_PROJECT_ID` | Your Google Cloud project ID |
| `GCP_REGION` | Deployment region (e.g., `asia-southeast1`) |
| `GCP_SA_KEY` | Service account JSON key |
| `NEXT_PUBLIC_API_URL` | Backend API URL (for frontend) |

---

## Manual Deployment

### Via GitHub UI
1. Go to **Actions** tab
2. Select workflow (Frontend or Backend)
3. Click **Run workflow**
4. Choose branch
5. Click **Run workflow**

### Via GitHub CLI
```bash
# Deploy frontend
gh workflow run deploy-frontend.yml

# Deploy backend
gh workflow run deploy-backend.yml
```

---

## Monitoring Deployments

### View Running Workflows
```bash
gh run list
```

### View Workflow Details
```bash
gh run view <run-id>
```

### View Logs
```bash
gh run view <run-id> --log
```

---

## Workflow Configuration

### Path Filters

The workflows only trigger when relevant files change:

**Frontend triggers on:**
- `frontend/**` (any file in frontend directory)
- `.github/workflows/deploy-frontend.yml` (workflow file itself)

**Backend triggers on:**
- `backend/**` (any file in backend directory)
- `.github/workflows/deploy-backend.yml` (workflow file itself)

### Branch Strategy

| Branch | Behavior |
|--------|----------|
| `main` | Deploy to production |
| `develop` | Deploy to staging/develop |
| PR to `main` | Build test only (no deploy) |
| Other branches | No action |

---

## Customization

### Add Staging Environment

Create separate workflows or add environment logic:

```yaml
# In deploy-frontend.yml
env:
  SERVICE_NAME: ${{ github.ref == 'refs/heads/main' && 'bank-viewer-frontend' || 'bank-viewer-frontend-staging' }}
```

### Add Tests Before Deploy

Add a job before deployment:

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm test

  deploy:
    needs: test  # Only deploy if tests pass
    runs-on: ubuntu-latest
    # ... deployment steps
```

### Add Slack Notifications

Add notification step:

```yaml
- name: Notify Slack
  if: always()
  uses: slackapi/slack-github-action@v1
  with:
    payload: |
      {
        "text": "Deployment ${{ job.status }}"
      }
  env:
    SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK }}
```

---

## Troubleshooting

### Workflow Not Triggering?

**Check:**
- Files changed are in the correct path filter
- Branch is `main` or `develop`
- Workflow file has no syntax errors

**Validate workflow:**
```bash
# Install actionlint
brew install actionlint

# Validate workflow files
actionlint .github/workflows/*.yml
```

### Deployment Fails?

**Check:**
1. **GitHub Actions logs** - Detailed error messages
2. **Secrets** - All 4 secrets are set correctly
3. **GCP permissions** - Service account has correct roles
4. **Artifact Registry** - Repository exists in correct region

### Useful Commands

```bash
# View recent workflow runs
gh run list --limit 5

# View specific run
gh run view <run-id>

# Re-run failed workflows
gh run rerun <run-id>

# Watch a running workflow
gh run watch
```

---

## Best Practices

1. **Always test in staging first**
   - Use `develop` branch for testing
   - Promote to `main` when stable

2. **Use semantic commit messages**
   ```bash
   feat: add new feature
   fix: fix bug
   chore: update dependencies
   ```

3. **Review deployment logs**
   - Check GitHub Actions summary
   - Verify service URL works

4. **Monitor costs**
   - Check GCP billing regularly
   - Set up budget alerts

5. **Keep secrets secure**
   - Never commit secrets to repo
   - Rotate service account keys quarterly
   - Use least privilege permissions

---

## Adding More Workflows

### Database Migration Workflow

```yaml
name: Run Migrations
on:
  workflow_dispatch:

jobs:
  migrate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run migrations
        run: |
          # Migration commands
```

### Scheduled Jobs

```yaml
name: Cleanup Old Images
on:
  schedule:
    - cron: '0 0 * * 0'  # Weekly

jobs:
  cleanup:
    # ... cleanup steps
```

---

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Workflow Syntax](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions)
- [Action Marketplace](https://github.com/marketplace?type=actions)

---

**Note:** Replace `YOUR_USERNAME/YOUR_REPO` in the badge URLs with your actual GitHub username and repository name.

