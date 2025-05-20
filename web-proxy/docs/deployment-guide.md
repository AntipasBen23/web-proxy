# Deployment Guide – Web Proxy Microservices on Render

This document outlines how to deploy each microservice in the `web-proxy` architecture using [Render](https://render.com), a fully managed cloud platform.

---

## Prerequisites

- A [Render account](https://dashboard.render.com/register)
- GitHub repository containing the full `web-proxy` project
- Each microservice placed in `services/<service-name>/` with:
  - `go.mod`
  - `cmd/main.go`
  - Optional `render.yaml` (local or root-level)

---

## 1. Deploying `proxy-service`

### Location:
`services/proxy-service/`

### Steps:

1. Go to [Render dashboard](https://dashboard.render.com/)
2. Click **"New Web Service"**
3. Connect your GitHub repo and choose the branch (`main`)
4. Set the following:
   - **Name**: `proxy-service`
   - **Root Directory**: `services/proxy-service`
   - **Environment**: `Go`
   - **Build Command**: (leave blank)
   - **Start Command**:  
     ```bash
     go run ./cmd/main.go
     ```
   - **Plan**: Choose **Free** or **Pro**

5. Set environment variables manually or through a `.env` file (recommended).
6. Click **Create Web Service**

---

## 2. Deploying `auth-service`

### Location:
`services/auth-service/`

### Start Command:
```bash
go run ./cmd/main.go
