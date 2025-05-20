# Design Decisions – Web Proxy Project

This document outlines key technical and architectural decisions made during the development of the `web-proxy` project, along with justifications for each choice.

---

## 1. Microservice Architecture

**Decision:** Use a microservice architecture instead of a monolith.

**Rationale:**
- Encourages separation of concerns (proxy, auth, metrics, config)
- Enables independent scaling and deployment of services
- Facilitates better fault isolation
- Makes the system more flexible for future integrations (e.g., admin dashboards)

---

## 2. Programming Language: Go (Golang)

**Decision:** Write all services in Go.

**Rationale:**
- Go is highly performant and suitable for networking-heavy applications
- Concurrency primitives (goroutines, channels) make it ideal for proxies and rate limiting
- Easy deployment (compile to single binary)
- First-class HTTP and TLS support in the standard library

---

## 3. Render for Deployment

**Decision:** Deploy to [Render](https://render.com) instead of using Docker/Kubernetes.

**Rationale:**
- Simple, fast setup with automatic builds from GitHub
- No need for container orchestration complexity at this stage
- Supports multiple services and environment variables
- Ideal for startups, prototypes, or small production workloads

---

## 4. TLS Termination at the Proxy

**Decision:** Perform TLS termination inside `proxy-service`.

**Rationale:**
- Allows fine-grained control over certificate handling
- Supports self-managed and automated cert rotation in the future
- Enables future features like SNI-based routing or custom cert chains

---

## 5. Custom Load Balancing Strategies

**Decision:** Implement internal support for multiple load balancing strategies (`round-robin`, `least-connections`, etc.)

**Rationale:**
- Provides flexibility over built-in reverse proxies (like NGINX)
- Helps adapt to dynamic backends in real-time
- Lays the foundation for plugin-based strategy injection later

---

## 6. Prometheus-Compatible Metrics

**Decision:** Use Prometheus metrics format via `metrics-service`.

**Rationale:**
- Compatible with most observability stacks (Grafana, Alertmanager, etc.)
- Allows detailed request tracing and performance profiling
- Easy to integrate with Render and cloud-native tools

---

## 7. Environment-Based Configuration

**Decision:** Use `.env` files and environment variables for config.

**Rationale:**
- Twelve-Factor App standard
- Secure and easy to override per environment
- Works seamlessly with Render and other cloud platforms

---

## 8. Optional Services (Auth, Config, Metrics)

**Decision:** Modularize optional functionality into their own services.

**Rationale:**
- Keeps core proxy lightweight and fast
- Allows teams to work independently on isolated features
- Optional services can be disabled or replaced without breaking the system

---

## Future Design Considerations

- **Plugin system** for routing, authentication, and transformation
- **Admin API** for runtime control of routing rules and backend pools
- **Distributed backend discovery** (e.g., via Consul, etcd)
- **gRPC support** for internal service communication (if scaling becomes complex)

---

## Related Docs

- [architecture.md](./architecture.md)
- [config-reference.md](./config-reference.md)
- [deployment-guide.md](./deployment-guide.md)

