# Credit Card Validation API

## Description

This project is an API for validating credit card numbers using the Luna algorithm and Golang.

## Running project

1. Build the Docker image:

   ```bash
   docker build .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 imagename
   ```

## Sending Requests

Send requests to the following URL:

```
http://localhost:8080/validate
```

### Example request

The request should be sent in JSON format:

```json
{
  "cardnumber": "4111111111111111",
  "expirationyear": "2028",
  "expirationmonth": "11"
}
```

### Expected response

The response will contain the result of the validation and may look like this:

```json
{
  "valid": true
}
```

or

```json
{
  "valid": false,
  "error": {
    "code": "some_error",
    "message": "some_message"
  }
}
```
