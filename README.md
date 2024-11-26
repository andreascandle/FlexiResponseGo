# **FlexiResponseGo**
[![Go Report Card](https://goreportcard.com/badge/github.com/andreascandle/FlexiResponseGo)](https://goreportcard.com/report/github.com/andreascandle/FlexiResponseGo)
[![GoDoc](https://pkg.go.dev/badge/github.com/andreascandle/FlexiResponseGo.svg)](https://pkg.go.dev/github.com/andreascandle/FlexiResponseGo)
[![Coverage Status](https://coveralls.io/repos/github/andreascandle/FlexiResponseGo/badge.svg)](https://coveralls.io/github/andreascandle/FlexiResponseGo)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/v/release/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo/releases)
[![Issues](https://img.shields.io/github/issues/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo/pulls)
[![Contributors](https://img.shields.io/github/contributors/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo/graphs/contributors)
[![Code Size](https://img.shields.io/github/languages/code-size/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo)
[![Lines of Code](https://img.shields.io/tokei/lines/github/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo)
[![Last Commit](https://img.shields.io/github/last-commit/andreascandle/FlexiResponseGo)](https://github.com/andreascandle/FlexiResponseGo/commits)

## **Streamline API Response Management in Go**

FlexiResponseGo is a robust and extensible shared library for managing HTTP responses in Go projects. It standardizes response structures, integrates structured logging, and supports multiple web frameworks, including Fiber, Gin, Echo, and native HTTP. The library also includes built-in support for observability with Prometheus metrics and OpenTelemetry tracing.

## **Features**

- **Standardized Responses**:
  - Unified success and error response formats across all frameworks.
  - Supports metadata, field-specific validation errors, and trace IDs.

- **Logging Integration**:
  - Structured logging using `zap`.
  - Automatic logging of incoming requests and outgoing responses.

- **Framework Support**:
  - Adapters for popular Go frameworks:
    - Fiber
    - Gin
    - Echo
    - Native HTTP

- **Observability**:
  - Metrics collection with Prometheus.
  - Distributed tracing with OpenTelemetry.

- **Extensibility**:
  - Easy integration into new frameworks.
  - Centralized configuration and dynamic updates.

---

## **Features**
- **Standardized Responses**:
  - Unified success and error response formats across all frameworks.
  - Supports metadata, field-specific validation errors, and trace IDs.
- **Logging Integration**:
  - Structured logging using `zap`.
  - Automatic logging of incoming requests and outgoing responses.
- **Framework Agnostic**: Seamless integration with:
  - [Fiber](https://github.com/gofiber/fiber)
  - [Gin](https://github.com/gin-gonic/gin)
  - [Echo](https://github.com/labstack/echo)
  - [net/http](https://pkg.go.dev/net/http)
- **Observability**:
  - Metrics collection with Prometheus.
  - Distributed tracing with OpenTelemetry.
- **Extensibility**:
  - Easy integration into new frameworks.
  - Centralized configuration and dynamic updates.
- **Advanced Error Management**:
  - Categorized errors with extensible codes and sanitization.
- **Performance Optimized**:
  - High-speed JSON serialization using [json-iterator/go](https://github.com/json-iterator/go).
- **Dynamic Configuration**:
  - Runtime settings for metadata, logging levels, and behaviors.

---

## **Installation**

Install FlexiResponseGo via `go get`:

```bash
go get github.com/andreascandle/FlexiResponseGo
```

## **Installation**

```bash
import "github.com/andreascandle/FlexiResponseGo"
```

## **Usage**
### 1. Setup Configuration
#### Update global settings for metadata and logging:
```bash
import "github.com/andreascandle/FlexiResponseGo/config"

func main() {
    conf := config.GetConfig()
    conf.UpdateMetadata("serviceName", "MyService")
    conf.UpdateLogLevel("debug")
}
```
### 2. Framework-Specific Adapters
#### Fiber Example:
```bash
import (
    "github.com/andreascandle/FlexiResponseGo/adapters"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return adapters.FiberSuccessResponse(c, "Welcome to FlexiResponseGo!", map[string]string{"status": "ok"})
    })

    app.Listen(":3000")
}
```
#### Gin Example:
```bash
import (
    "github.com/andreascandle/FlexiResponseGo/adapters"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        adapters.GinSuccessResponse(c, "Welcome to FlexiResponseGo!", map[string]string{"status": "ok"})
    })

    r.Run(":3000")
}
```
#### Echo Example:
```bash
import (
    "github.com/andreascandle/FlexiResponseGo/adapters"
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    e := echo.New()

    e.GET("/success", func(c echo.Context) error {
        return adapters.EchoSuccessResponse(c, "Operation successful", map[string]string{"example": "echo"})
    })

    e.Start(":3000")
}
```
#### net/http
```bash
import (
    "net/http"
    "response/adapters"
)

func main() {
    http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
        return adapters.EchoSuccessResponse(c, "Operation successful", map[string]string{"example": "echo"})
    })

    http.ListenAndServe(":8080", nil)
}
```

### Observability
- **Distributed Tracing:** Add tracing using OpenTelemetry.
- **Metrics Tracking:** Export metrics to Prometheus for better API monitoring.
#### Prometheus Metrics
Expose metrics at /metrics:
```bash
import (
    "github.com/andreascandle/FlexiResponseGo/observability"
    "net/http"
)

func main() {
    http.Handle("/metrics", observability.HTTPHandlerForMetrics())
    http.ListenAndServe(":9090", nil)
}
```
#### OpenTelemetry Tracing
Initialize tracing for your service:
```bash
import "github.com/andreascandle/FlexiResponseGo/observability"

func main() {
    shutdown := observability.InitTracer("MyService")
    defer shutdown()
}
```
### Testing
Run the test suite using:
```bash
go test ./tests/... -v
```
Ensure you have all dependencies installed for full testing.

### Contributing
We welcome contributions to FlexiResponseGo! If you’d like to report an issue, suggest a feature, or submit a code change, follow the guidelines below.

### How to Contribute
- Fork the repository.
- Create a feature branch.
- Submit a pull request with detailed descriptions of changes.

### License
This project is licensed under the [MIT License](https://github.com/andreascandle/FlexiResponseGo/blob/main/LICENSE).

### Acknowledgments
Special thanks to the Go developer community and contributors to:
- Gorilla Mux
- Fiber
- Gin
- Echo
- OpenTelemetry
- Prometheus
- less

