# Go `net/http` Package

The [`net/http`](https://pkg.go.dev/net/http) package provides HTTP client and server implementations for Go.  
It is widely used for building web servers, REST APIs, and making HTTP requests.

## Common Types and Functions

- `http.Server`: Defines parameters for running an HTTP server.
- `http.NewServeMux()`: Creates a new HTTP request multiplexer, i.e router.
- `http.HandleFunc(pattern, handler)`: Registers the handler function for the given route.
- `http.Get(url)`: Issues a GET request to the specified URL.
- `http.Post(url, contentType, body)`: Issues a POST request to the specified URL.
- `http.Request` / `http.Response`: Represent a HTTP request/response.

---

[Go Back](../../README.md)