# GitHub Actions CI/CD Setup Guide

This guide will help you configure GitHub Actions to automatically build and deploy your Todo API to Azure Container Instances.

## Prerequisites

- Azure subscription
- Azure Container Registry (ACR) created
- Azure Container Instance or Web App
- GitHub repository

## Step 1: Get Azure Credentials

### Get ACR Credentials

```powershell
# Get your ACR credentials
az acr credential show --name gotutacr --resource-group gotut-rg
```

Copy the **username** and **password**.

### Create Azure Service Principal

```powershell
# Get your subscription ID
az account show --query id -o tsv

# Create service principal for GitHub Actions
az ad sp create-for-rbac --name "gotut-github-actions" `
  --role contributor `
  --scopes /subscriptions/<YOUR-SUBSCRIPTION-ID>/resourceGroups/gotut-rg `
  --sdk-auth
```

This will output JSON like:

```json
{
  "clientId": "...",
  "clientSecret": "...",
  "subscriptionId": "...",
  "tenantId": "...",
  ...
}
```

**Copy the entire JSON output!**

## Step 2: Add GitHub Secrets

Go to your GitHub repository:

1. Click **Settings**
2. Click **Secrets and variables** → **Actions**
3. Click **New repository secret**

Add these three secrets:

### Secret 1: AZURE_ACR_USERNAME

- **Name**: `AZURE_ACR_USERNAME`
- **Value**: `gotutacr` (from Step 1)

### Secret 2: AZURE_ACR_PASSWORD

- **Name**: `AZURE_ACR_PASSWORD`
- **Value**: The password from Step 1

### Secret 3: AZURE_CREDENTIALS

- **Name**: `AZURE_CREDENTIALS`
- **Value**: The entire JSON output from the `az ad sp create-for-rbac` command

## Step 3: Install Git Hooks (Local Development)

On your local machine, install the commit message validation hook:

```powershell
# Windows PowerShell
.\scripts\setup-hooks.ps1

# Or Git Bash
bash scripts/setup-hooks.sh
```

This ensures your commits follow Angular convention.

## Step 4: Test the Workflow

### Create a feature branch:

```bash
git checkout -b feature/test-cicd
```

### Make a change and commit with Angular convention:

```bash
git add .
git commit -m "feat: add CI/CD pipeline"
```

If your commit message is invalid, you'll see an error. Fix it and try again.

### Push to GitHub:

```bash
git push origin feature/test-cicd
```

### Check GitHub Actions:

1. Go to your GitHub repository
2. Click the **Actions** tab
3. You should see a workflow run for your commit
4. Click on it to see the build progress

## What Happens When You Push?

### Commit Message Check

- ✅ Commits starting with `feat:` or `fix:` → Build Docker image
- ⏭️ Other commit types (`docs:`, `chore:`, etc.) → Skip build

### Build Process

1. Checkout code
2. Set up Docker Buildx
3. Login to Azure Container Registry
4. Build Docker image
5. Tag with commit SHA and `latest`
6. Push to ACR

### Deployment

- **Push to `develop` branch** → Deploy to **staging** environment
- **Push to `main` branch** → Deploy to **production** environment

## Workflow Diagram

```
Commit with feat: or fix:
        ↓
   Build Docker Image
        ↓
   Push to Azure ACR
        ↓
    ┌─────────────┴─────────────┐
    ↓                           ↓
develop branch              main branch
    ↓                           ↓
Deploy to Staging       Deploy to Production
```

## Environment URLs

After deployment, your API will be available at:

**Staging** (develop branch):

- API: `http://gotut-api-staging.aehpfwbgdhgahdgx.germanywestcentral.azurecontainer.io:8080`
- Swagger: `http://gotut-api-staging.aehpfwbgdhgahdgx.germanywestcentral.azurecontainer.io:8080/swagger/index.html`

**Production** (main branch):

- API: `http://gotut-api.aehpfwbgdhgahdgx.germanywestcentral.azurecontainer.io:8080`
- Swagger: `http://gotut-api.aehpfwbgdhgahdgx.germanywestcentral.azurecontainer.io:8080/swagger/index.html`

## Commit Message Examples

✅ **Valid commits that trigger builds:**

```bash
git commit -m "feat: add user authentication endpoint"
git commit -m "fix: resolve CORS configuration issue"
git commit -m "feat(api): implement todo filtering"
git commit -m "fix(handlers): correct validation logic"
```

✅ **Valid commits that skip builds:**

```bash
git commit -m "docs: update README with deployment instructions"
git commit -m "style: format code with gofmt"
git commit -m "chore: update dependencies"
git commit -m "refactor: simplify handler logic"
```

❌ **Invalid commits (will be rejected by git hook):**

```bash
git commit -m "added new feature"           # No type prefix
git commit -m "feat:add authentication"     # Missing space after colon
git commit -m "FEAT: new feature"           # Type must be lowercase
git commit -m "feature: something"          # Wrong type (should be 'feat')
```

## Troubleshooting

### Workflow Not Triggering

**Check:**

1. Secrets are configured correctly in GitHub
2. Commit message starts with `feat:` or `fix:`
3. Branch is `main`, `develop`, or starts with `feature/` or `fix/`

### Build Fails

**Check:**

1. `AZURE_ACR_USERNAME` and `AZURE_ACR_PASSWORD` secrets are correct
2. Azure Container Registry name in workflow matches your ACR
3. Docker build succeeds locally: `docker build -t test -f build/Dockerfile .`

### Deployment Fails

**Check:**

1. `AZURE_CREDENTIALS` secret contains valid service principal JSON
2. Service principal has contributor role on resource group
3. Container instance name and resource group match workflow

### View Workflow Logs

1. Go to **Actions** tab in GitHub
2. Click on the failed workflow run
3. Click on the failed job
4. Expand the failed step to see error details

## Manual Deployment

If you need to deploy manually:

```powershell
# Build and push
docker build -t gotut-api -f build/Dockerfile .
docker tag gotut-api gotutacr.azurecr.io/gotut-api:latest
docker push gotutacr.azurecr.io/gotut-api:latest

# Deploy to Azure
.\scripts\deploy-azure.ps1
```

## Security Best Practices

✅ **Do:**

- Use GitHub Secrets for all credentials
- Rotate service principal credentials regularly
- Use minimal required permissions
- Enable branch protection rules

❌ **Don't:**

- Commit credentials to repository
- Share secrets publicly
- Use admin credentials for CI/CD
- Disable security checks

## Next Steps

1. ✅ Set up branch protection rules (require PR for `main`)
2. ✅ Add status checks (require CI to pass before merge)
3. ✅ Set up staging environment for testing
4. ✅ Configure custom domain for production
5. ✅ Add monitoring and alerts

## Support

For issues:

- Check GitHub Actions logs
- Review Azure Container logs: `az container logs --resource-group gotut-rg --name gotut-api`
- Verify secrets are configured correctly
- Ensure service principal has correct permissions
