# System Architecture – Web Proxy

This document outlines the architecture of the `web-proxy` microservice system. The project is designed to be a high-performance, extensible proxy/load balancer using Golang and follows microservice principles.

---

## 🧱 Overview

The system is composed of independent microservices, each responsible for a specific capability such as proxying, metrics, authentication, and configuration management.

### Core Microservices:

- **proxy-service**: Handles incoming HTTP/HTTPS/HTTP2 requests, load balancing, TLS termination, and request forwarding.
- **auth-service** (optional): Validates API keys or tokens before allowing access to protected proxy endpoints.
- **config-service** (optional): Manages dynamic configuration and distributes updates to services.
- **metrics-service**: Exposes Prometheus-style metrics for monitoring traffic and service health.

---

## 🔁 Request Flow

1. **Client Request** hits the `proxy-service` endpoint (e.g. via `https://proxy.example.com`)
2. If enabled, the request passes through `auth-service` middleware for token/API key validation
3. Proxy selects a backend using a load balancing strategy (e.g., round-robin)
4. If backend is available and circuit is closed, request is forwarded
5. Response is passed through any middleware (e.g., logging, transformation)
6. Metrics are emitted to `metrics-service` and logs stored
7. If a rate limit is hit, a `429 Too Many Requests` response is returned

---

## ⚙️ Service Responsibilities

### 1. `proxy-service`
- Accepts HTTP, HTTPS, HTTP/2
- Supports reverse proxy routing
- Implements load balancing algorithms (e.g., round-robin, least-connections)
- TLS termination using provided certs
- Request/response middleware support
- Circuit breaking + rate limiting middleware

### 2. `auth-service`
- Validates incoming tokens or API keys
- Optional: integrates with external identity provider
- Lightweight and stateless

### 3. `config-service`
- Serves proxy configuration from JSON, ENV, or dynamic source (e.g. Redis, DB)
- Pushes config updates via HTTP or PubSub
- Provides versioning and rollback for config changes

### 4. `metrics-service`
- Uses Prometheus-compatible metrics format
- Aggregates connection count, errors, request duration, rate-limiting stats
- Enables alerting integration (e.g., with Grafana + Alertmanager)

---

## 🌐 Protocols & Interfaces

| Protocol       | Used For                    | Services              |
|----------------|-----------------------------|-----------------------|
| HTTP/1.1, 2     | Proxying requests            | proxy-service         |
| TLS (HTTPS)    | Encryption/termination       | proxy-service         |
| REST (JSON)    | Internal API communication   | config-service, auth  |
| Prometheus     | Metrics scraping             | metrics-service       |

---

## 🔒 Security

- Each exposed endpoint supports API key or token validation (via `auth-service`)
- Rate limiting and circuit breaking prevent abuse and overload
- TLS termination is supported on the proxy layer
- Environment-based configuration for secrets and credentials

---

## 🔄 Configuration Model

All services support centralized config loading via:
- `.env` files
- ENV variables
- Remote config from `config-service`

---

## 🔧 Extensibility

You can extend the system by:
- Adding more load balancing strategies
- Introducing caching logic at the proxy
- Integrating DDoS mitigation middleware
- Connecting a UI dashboard for runtime control

---

## 📌 Deployment Strategy

Each service is deployed independently via Render.

- **Build Command:** None (Go is auto-detected)
- **Start Command:** `go run ./cmd/main.go`
- **ENV Variables:** Managed in Render Dashboard


