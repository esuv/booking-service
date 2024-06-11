<img align="right" width="35%" src="./images/gopher.png">

# Booking service
## Description
Basic application for booking hotel rooms

## API Reference

#### Create an order

```sh
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hotel_id": "reddison",
    "room_id": "lux",
    "email": "guest@mail.ru",
    "from": "2024-01-02T00:00:00Z",
    "to": "2024-01-04T00:00:00Z"
}'
```

## Deployment

To run this project

```bash
    $ git clone ...
    $ go run cmd/app/main.go
```
