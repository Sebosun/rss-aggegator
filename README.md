RSS Aggregator

# API endpoints

### GET /v1/healthz

Healthcheck, should return JSON with OK status

### GET /v1/err

Err endpoint, should return 404 with err message JSON

### GET /v1/user

Expects Authorization Bearer token

Returns user data

### POST /v1/user

Account creation endpoint
Expect following JSON Structure:

```json
name: string
```

Returns follwing json

```json
id:             string
created_at:     string
updated_at:     string
name:           string
api_key:        string
```
