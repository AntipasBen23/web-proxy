# Testing Guide – Web Proxy Project

This document outlines the testing strategies and tools used in the `web-proxy` microservice architecture. The goal is to ensure correctness, reliability, and performance across all services.

---

## Test Layers Overview

| Type              | Scope                                | Purpose                             |
|-------------------|--------------------------------------|-------------------------------------|
| Unit Tests        | Individual functions/methods         | Verify logic in isolation           |
| Integration Tests | End-to-end service interactions      | Test real-world flows and failures  |
| Load Tests        | Proxy throughput and stress handling | Measure performance under pressure  |
| Smoke Tests       | Basic health checks after deploy     | Ensure services start and respond   |

---

## 1. Unit Testing (Go Test)

### Tool: `go test`

**Scope:**  
- All internal service logic (`proxy`, `balancer`, `middleware`, etc.)
- Each package should include its own `*_test.go` files

**Example:**

```go

func TestHandleRequest(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    handler := NewProxyHandler()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", w.Code)
    }
}
