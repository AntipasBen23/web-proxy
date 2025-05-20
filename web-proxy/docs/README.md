# Web Proxy – High-Performance Go Proxy & Load Balancer

`web-proxy` is a high-performance, microservice-based reverse proxy and load balancer written in Go. It supports HTTP, HTTPS, and HTTP/2 protocols with features like TLS termination, load balancing algorithms, request/response transformation, rate limiting, and circuit breaking.

---

## Features

- HTTP/HTTPS/HTTP2 support
- Load balancing strategies (round-robin, least-connections, etc.)
- TLS termination
- Rate limiting and circuit breaking
- Request/Response transformation hooks
- Prometheus-compatible metrics
- Modular microservices (Proxy, Auth, Config, Metrics)
- Render-compatible deployment (no Docker required)

---

## Architecture Overview

This project follows a **microservice architecture**. Key components:

| Service          | Description                                       |
|------------------|---------------------------------------------------|
| `proxy-service`  | Core reverse proxy with load balancing and TLS    |
| `auth-service`   | Validates API tokens or keys (optional)           |
| `config-service` | Centralized dynamic config loader (optional)      |
| `metrics-service`| Exposes Prometheus metrics for monitoring         |

For technical details, see [`docs/architecture.md`](./docs/architecture.md).

---

## Folder Structure

