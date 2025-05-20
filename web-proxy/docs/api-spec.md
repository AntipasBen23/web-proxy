# API Specification – Web Proxy Service

This document describes the RESTful API endpoints exposed by the Web Proxy microservices.

---

## ✅ Base URL
https://your-proxy-service.onrender.com


---

## 🩺 1. Health Check

### `GET /healthz`

**Description:**  
Check if the proxy service is alive and responding.

**Response:**

```json
{
  "status": "ok",
  "timestamp": "2025-05-20T14:00:00Z"
}

2. Reload Configuration
POST /reload-config
Description:
Reloads the proxy configuration (e.g., updated backend pool, TLS certs, etc.)
May be secured behind admin token.

Headers:
Authorization: Bearer <admin-token>

Response:
{
  "message": "Configuration reloaded successfully"
}

3. Metrics Endpoint
GET /metrics
Description:
Exposes Prometheus-style metrics for performance and traffic monitoring.

Response:
Prometheus-formatted metrics text.

Example:
http_requests_total{method="GET", path="/"} 12345
proxy_active_connections 14

4. Rate Limit Status (Optional)
GET /rate-limit-status
Description:
Returns the current rate limit status for the requesting IP or API key.

Headers (optional):
X-API-KEY: abc123

Response:
{
  "limit": 1000,
  "remaining": 823,
  "reset": "2025-05-20T15:00:00Z"
}

5. Test Endpoint (Development Only)
GET /debug/echo
Description:
Returns request headers and basic info for debugging proxy behaviors.

Response:
{
  "method": "GET",
  "headers": {
    "User-Agent": "PostmanRuntime/7.32.3",
    "Accept": "*/*"
  },
  "remoteIP": "92.101.84.56"
}

