# Changelog

All notable changes to the `web-proxy` project will be documented in this file.

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

### Added
- Initial microservice architecture with `proxy-service`, `auth-service`, `config-service`, and `metrics-service` folders
- Documentation scaffolding: `README.md`, `architecture.md`, `api-spec.md`, `deployment-guide.md`, etc.
- CLI setup script (`scripts/setup-env.sh`)
- Render deployment configs for each service
- Basic file structure for each service including `cmd/main.go` and `go.mod`

---

## [0.1.0] – 2025-05-20

### Added
- Initial scaffolding of `proxy-service`
- `GET /healthz` endpoint
- `POST /reload-config` placeholder
- Basic rate limiting and request logging middleware
- Load balancing interface with round-robin strategy

### Changed
- N/A

### Fixed
- N/A

---

## [0.1.1] – *Coming Soon*

### Planned
- Implement circuit breaker logic
- Add Prometheus metrics via `metrics-service`
- Implement full config reload via `config-service`
- API key validation in `auth-service`
