# RestTestify

**RestTestify** is a Golang-based tool designed to simplify REST API testing using YAML files. Specify API endpoints, headers, request bodies, and expected responses in a YAML format, and RestTestify will automatically perform tests for you.

## Features

- Simple YAML configuration for defining REST API tests.
- Supports HTTP methods like GET, POST, PUT, DELETE.
- Customizable headers and request bodies.
- Validates response status codes and content.
- CLI-based for easy integration into CI/CD pipelines.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/rest-testify.git
   ```

## Usage

1. **Create a YAML test file(e.g., tests.yaml)**

```yaml
tests:
  - name: "Get User Data"
    method: "GET"
    endpoint: "https://jsonplaceholder.typicode.com/posts/1"
    headers:
      Content-Type: "application/json"
    expected_status: 200
    expected_body_contains: "userId"

  - name: "Create New Post"
    method: "POST"
    endpoint: "https://jsonplaceholder.typicode.com/posts"
    headers:
      Content-Type: "application/json"
    body:
      title: "foo"
      body: "bar"
      userId: 1
    expected_status: 201
    expected_body_contains: "id"
```

### **Configuration Options**

- method: HTTP method (GET, POST, PUT, DELETE).
- endpoint: API URL endpoint.
- headers: Custom HTTP headers as key-value pairs.
- body: Request body for POST/PUT methods.
- expected_status: Expected HTTP status code.
- expected_body_contains: String that should appear in the response body.

2. **Run the tests**

```sh
./rest-testify -file tests.yaml
```
