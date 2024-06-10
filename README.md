<img align="right" width="35%" src="./images/gopher.png">

# Booking service
## Description
Basic application for booking hotel rooms

## API Reference

#### Create order

```http
  POST /orders
```

| Parameter | Type     | Description                |
|:----------|:---------|:---------------------------|
| `api_key` | `string` | **Required**. Your API key |

## Deployment

To run this project

```bash
    $ git clone ...
    $ go run cmd/app/main.go
```
