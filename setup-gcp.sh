#!/bin/bash

# Setup script for Google Cloud Platform deployment
# Flip Fullstack Test - Automated GCP Configuration

set -e

echo "🚀 Flip Fullstack Test - GCP Setup"
echo "===================================="
echo ""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;36m'
NC='\033[0m'

# Configuration
PROJECT_ID="flip-fullstack-test"
REGION="asia-southeast1"
REPOSITORY="flip-fullstack-test"
SERVICE_ACCOUNT_NAME="github-actions"

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo -e "${RED}❌ gcloud CLI is not installed${NC}"
    echo "Install it from: https://cloud.google.com/sdk/docs/install"
    exit 1
fi

echo -e "${GREEN}✅ gcloud CLI found${NC}"

# Check if logged in
if ! gcloud auth list --filter=status:ACTIVE --format="value(account)" &> /dev/null; then
    echo -e "${RED}❌ Not logged into gcloud${NC}"
    echo "Run: gcloud auth login"
    exit 1
fi

ACTIVE_ACCOUNT=$(gcloud auth list --filter=status:ACTIVE --format="value(account)")
echo -e "${GREEN}✅ Logged in as: $ACTIVE_ACCOUNT${NC}"
echo ""

# Check if project exists
echo "📦 Checking project..."
if gcloud projects describe $PROJECT_ID &> /dev/null; then
    echo -e "${YELLOW}⚠️  Project '$PROJECT_ID' already exists${NC}"
    echo "Using existing project"
else
    echo "Creating project '$PROJECT_ID'..."
    gcloud projects create $PROJECT_ID --name="Flip Fullstack Test" || {
        echo -e "${RED}❌ Failed to create project${NC}"
        echo "You may need to:"
        echo "  1. Enable billing: https://console.cloud.google.com/billing"
        echo "  2. Or use a different project ID"
        exit 1
    }
    echo -e "${GREEN}✅ Project created${NC}"
fi

# Set default project
gcloud config set project $PROJECT_ID
echo ""

# Check billing (skip check, just warn)
echo "💳 Billing check..."
echo -e "${YELLOW}⚠️  Please ensure billing is enabled for this project${NC}"
echo "   Visit: https://console.cloud.google.com/billing/linkedaccount?project=$PROJECT_ID"
echo ""
read -p "Press Enter to continue if billing is enabled, or Ctrl+C to exit..."
echo ""

# Enable APIs
echo "🔧 Enabling required APIs..."
gcloud services enable \
  run.googleapis.com \
  cloudbuild.googleapis.com \
  artifactregistry.googleapis.com \
  containerregistry.googleapis.com \
  iam.googleapis.com \
  --quiet

echo -e "${GREEN}✅ APIs enabled${NC}"
echo ""

# Create Artifact Registry
echo "📦 Checking Artifact Registry..."
if gcloud artifacts repositories describe $REPOSITORY --location=$REGION &> /dev/null; then
    echo -e "${YELLOW}⚠️  Repository '$REPOSITORY' already exists${NC}"
else
    echo "Creating Artifact Registry repository..."
    gcloud artifacts repositories create $REPOSITORY \
      --repository-format=docker \
      --location=$REGION \
      --description="Flip Fullstack Test Docker images" \
      --quiet
    echo -e "${GREEN}✅ Artifact Registry created${NC}"
fi
echo ""

# Create Service Account
echo "👤 Checking service account..."
SA_EMAIL="${SERVICE_ACCOUNT_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

if gcloud iam service-accounts describe $SA_EMAIL &> /dev/null; then
    echo -e "${YELLOW}⚠️  Service account already exists${NC}"
else
    echo "Creating service account..."
    gcloud iam service-accounts create $SERVICE_ACCOUNT_NAME \
      --display-name="GitHub Actions Deployment" \
      --quiet
    echo -e "${GREEN}✅ Service account created${NC}"
fi
echo ""

# Grant IAM roles
echo "🔐 Granting IAM roles..."
ROLES=(
    "roles/run.admin"
    "roles/artifactregistry.writer"
    "roles/iam.serviceAccountUser"
)

for ROLE in "${ROLES[@]}"; do
    gcloud projects add-iam-policy-binding $PROJECT_ID \
      --member="serviceAccount:${SA_EMAIL}" \
      --role="$ROLE" \
      --condition=None \
      --quiet > /dev/null 2>&1 || true
done

echo -e "${GREEN}✅ IAM roles granted${NC}"
echo ""

# Generate service account key
echo "🔑 Checking service account key..."
KEY_FILE="github-actions-key.json"

if [ -f "$KEY_FILE" ]; then
    echo -e "${YELLOW}⚠️  Key file already exists: $KEY_FILE${NC}"
    read -p "Regenerate key? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm $KEY_FILE
        gcloud iam service-accounts keys create $KEY_FILE \
          --iam-account=$SA_EMAIL \
          --quiet
        echo -e "${GREEN}✅ New key generated${NC}"
    else
        echo "Using existing key file"
    fi
else
    gcloud iam service-accounts keys create $KEY_FILE \
      --iam-account=$SA_EMAIL \
      --quiet
    echo -e "${GREEN}✅ Key generated: $KEY_FILE${NC}"
fi
echo ""

# Summary
echo "===================================="
echo -e "${GREEN}🎉 Setup Complete!${NC}"
echo "===================================="
echo ""
echo -e "${BLUE}📋 Configuration Summary:${NC}"
echo ""
echo "  Project ID:     $PROJECT_ID"
echo "  Region:         $REGION"
echo "  Repository:     $REPOSITORY"
echo "  Service Account: $SA_EMAIL"
echo ""
echo -e "${BLUE}🔐 GitHub Secrets to Add:${NC}"
echo ""
echo "Go to: https://github.com/YOUR_USERNAME/flip-fullstack-test/settings/secrets/actions"
echo ""
echo "Add these secrets:"
echo ""
echo "  GCP_PROJECT_ID = $PROJECT_ID"
echo "  GCP_REGION = $REGION"
echo "  GCP_SA_KEY = (copy entire content of $KEY_FILE)"
echo ""
echo -e "${BLUE}📝 To view the key:${NC}"
echo ""
echo "  cat $KEY_FILE"
echo ""
echo -e "${BLUE}🚀 Or use GitHub CLI:${NC}"
echo ""
echo "  gh secret set GCP_PROJECT_ID -b \"$PROJECT_ID\""
echo "  gh secret set GCP_REGION -b \"$REGION\""
echo "  gh secret set GCP_SA_KEY < $KEY_FILE"
echo ""
echo -e "${YELLOW}⚠️  IMPORTANT:${NC}"
echo "  1. Add secrets to GitHub before pushing code"
echo "  2. Delete $KEY_FILE after adding to GitHub:"
echo "     rm $KEY_FILE"
echo ""
echo -e "${BLUE}📖 Next Steps:${NC}"
echo ""
echo "  1. Add GitHub secrets (see above)"
echo "  2. Deploy backend:"
echo "     git add backend/"
echo "     git commit -m \"deploy: backend\""
echo "     git push origin main"
echo ""
echo "  3. After backend deploys, get URL:"
echo "     gcloud run services describe flip-fullstack-test-backend \\"
echo "       --region=$REGION \\"
echo "       --format='value(status.url)'"
echo ""
echo "  4. Add backend URL as secret:"
echo "     gh secret set NEXT_PUBLIC_API_URL -b \"https://your-backend-url.run.app\""
echo ""
echo "  5. Deploy frontend:"
echo "     git add frontend/"
echo "     git commit -m \"deploy: frontend\""
echo "     git push origin main"
echo ""
echo "Full guide: DEPLOYMENT.md"
echo ""
