# Configuration Reference

This document lists all environment variables and configuration options used across the `web-proxy` microservices.

---

## Common Environment Variables

These apply to **all services** unless otherwise stated:

| Variable Name       | Type     | Default        | Description                                |
|---------------------|----------|----------------|--------------------------------------------|
| `ENV`               | string   | `development`  | Application environment (`production`, `development`, `staging`) |
| `PORT`              | integer  | `8080`         | Port the service listens on                |
| `LOG_LEVEL`         | string   | `info`         | Logging verbosity: `debug`, `info`, `warn`, `error` |
| `CONFIG_SOURCE`     | string   | `env`          | Where to load config from: `env`, `file`, `remote` |

---

## proxy-service Configuration

| Variable Name              | Type     | Default        | Description                                |
|----------------------------|----------|----------------|--------------------------------------------|
| `PROXY_TLS_CERT_PATH`      | string   | `""`           | Path to TLS certificate for HTTPS         |
| `PROXY_TLS_KEY_PATH`       | string   | `""`           | Path to TLS private key                   |
| `PROXY_ENABLE_TLS`         | bool     | `false`        | Enable TLS termination                    |
| `PROXY_BACKEND_POOL`       | string   | `""`           | Comma-separated backend URLs (e.g., `http://localhost:9000,http://localhost:9001`) |
| `PROXY_BALANCER_STRATEGY`  | string   | `round_robin`  | Load balancing strategy (`round_robin`, `least_conn`) |
| `RATE_LIMIT_PER_MINUTE`    | int      | `1000`         | Max allowed requests per minute per IP    |
| `ENABLE_CIRCUIT_BREAKER`   | bool     | `true`         | Enable circuit breaker protection         |

---

## auth-service Configuration

| Variable Name          | Type     | Default   | Description                                |
|------------------------|----------|-----------|--------------------------------------------|
| `AUTH_SECRET_KEY`      | string   | `""`      | Secret key used to validate tokens or API keys |
| `AUTH_HEADER_NAME`     | string   | `Authorization` | Header where the token/API key is expected |
| `AUTH_STRATEGY`        | string   | `token`   | Strategy: `token`, `api_key`               |

---

## config-service Configuration

| Variable Name            | Type     | Default   | Description                                |
|--------------------------|----------|-----------|--------------------------------------------|
| `CONFIG_STORAGE`         | string   | `file`    | Source of config: `file`, `env`, `remote` |
| `CONFIG_FILE_PATH`       | string   | `./config/config.yaml` | Local config file location        |
| `CONFIG_REMOTE_URL`      | string   | `""`      | URL to pull remote config from (if enabled) |
| `CONFIG_REFRESH_INTERVAL`| int      | `60`      | Time (in seconds) to check for config changes |

---

## metrics-service Configuration

| Variable Name       | Type     | Default   | Description                                |
|---------------------|----------|-----------|--------------------------------------------|
| `METRICS_PORT`      | int      | `9090`    | Port to expose Prometheus metrics          |
| `METRICS_PATH`      | string   | `/metrics`| Endpoint path for metrics                  |
| `METRICS_NAMESPACE` | string   | `webproxy`| Prometheus namespace prefix for metrics    |

---

## Example `.env` for proxy-service

```dotenv
ENV=production
PORT=8080
LOG_LEVEL=info

PROXY_ENABLE_TLS=true
PROXY_TLS_CERT_PATH=./certs/cert.pem
PROXY_TLS_KEY_PATH=./certs/key.pem
PROXY_BACKEND_POOL=http://localhost:9000,http://localhost:9001
PROXY_BALANCER_STRATEGY=round_robin

RATE_LIMIT_PER_MINUTE=1000
ENABLE_CIRCUIT_BREAKER=true
