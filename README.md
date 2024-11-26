# **FlexiResponseGo**
[![Go Report Card](https://goreportcard.com/badge/github.com/andreascandle/FlexiResponseGo)](https://goreportcard.com/report/github.com/andreascandle/FlexiResponseGo)
[![Build Status](https://github.com/andreascandle/FlexiResponseGo/actions/workflows/go.yml/badge.svg)](https://github.com/andreascandle/FlexiResponseGo/actions)
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

FlexiResponseGo is a versatile and extensible response handling library designed for Go developers. It simplifies API response management across multiple frameworks, ensuring clean, standardized, and maintainable responses. With built-in support for logging, tracing, error handling, and metrics tracking, FlexiResponseGo empowers developers to create high-performance APIs with ease.

---

## **Features**
- **Framework Agnostic**: Seamless integration with:
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [Fiber](https://github.com/gofiber/fiber)
  - [Gin](https://github.com/gin-gonic/gin)
  - [Echo](https://github.com/labstack/echo)
  - [net/http](https://pkg.go.dev/net/http)
- **Standardized Response Structures**:
  - Unified handling for success, error, and validation responses.
- **Observability**:
  - Distributed tracing with [OpenTelemetry](https://opentelemetry.io/).
  - Metrics tracking with [Prometheus](https://prometheus.io/).
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

## **Usage**
### 1. Framework Integration
#### Gorilla Mux
```bash
import (
    "net/http"
    "response/adapters"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
        adapters.MuxSuccessResponse(w, r, "Operation successful", map[string]string{"example": "mux"})
    }).Methods(http.MethodGet)

    http.ListenAndServe(":8080", router)
}

```
#### Fiber
```bash
import (
    "response/adapters"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/success", func(c *fiber.Ctx) error {
        return adapters.FiberSuccessResponse(c, "Operation successful", map[string]string{"example": "fiber"})
    })

    app.Listen(":8080")
}
```
#### Gin
```bash
import (
    "response/adapters"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/success", func(c *gin.Context) {
        adapters.GinSuccessResponse(c, "Operation successful", map[string]string{"example": "gin"})
    })

    router.Run(":8080")
}

```
#### Echo
```bash
import (
    "response/adapters"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/success", func(c echo.Context) error {
        return adapters.EchoSuccessResponse(c, "Operation successful", map[string]string{"example": "echo"})
    })

    e.Start(":8080")
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
        adapters.HTTPSuccessResponse(w, r, "Operation successful", map[string]string{"example": "http"})
    })

    http.ListenAndServe(":8080", nil)
}
```

### Observability
- **Distributed Tracing:** Add tracing using OpenTelemetry.
- **Metrics Tracking:** Export metrics to Prometheus for better API monitoring.

### Folder Structure
```bash
response/
  core/               # Core response handling logic
  logger/             # Adaptive logger
  config/             # Dynamic configuration management
  utils/              # Utility functions (trace ID generation, sanitization)
  adapters/           # Adapters for various frameworks
    fiber_adapter.go  # Fiber adapter
    gin_adapter.go    # Gin adapter
    http_adapter.go   # net/http adapter
    echo_adapter.go   # Echo adapter
  observability/      # Observability tools (tracing, metrics)
  tests/              # Tests for all components
```
### Contributing
Contributions are welcome! To contribute:

- Fork the repository.
- Create a feature branch.
- Submit a pull request with detailed descriptions of changes.
### License
This project is licensed under the MIT License.

### Acknowledgments
Special thanks to the Go developer community and contributors to:
- Gorilla Mux
- Fiber
- Gin
- Echo
- OpenTelemetry
- Prometheus
- less

