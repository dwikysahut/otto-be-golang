# Go API - Voucher & Redemption System

This project is a simple API for managing brands, vouchers, and transaction redemptions. It includes endpoints for creating brands, vouchers, retrieving vouchers, and making redemptions.

## Prerequisites

Before running the API, ensure you have the following installed:

- Go (1.16+)
- PostgreSQL database
- Go modules for dependency management

## Setup

1. **Clone the repository:**

```bash
git clone https://github.com/your-repo/go-api.git
cd go-api
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Set up your PostgreSQL database**
   Update the dsn (Data Source Name) in your main.go file with your database credentials:

```bash
dsn := "host=localhost user=user password=user dbname=otto_api port=5432 sslmode=disable"

```

4. **run db migrations**

```bash
migrate -database "postgres://user:password@localhost:5432/go_api?sslmode=disable" -path ./migrations up
```

5. **Start API Server**

```bash
go run main.go
```

## End Point

**POST**

- `/brand`

body

```bash
{
    "name":"bear brabnd2"
}
```

response

```bash
{
    "status": "success",
    "data": {
        "id": 7,
        "name": "bear brabnd2"
    }
}
```

- `/voucher`

body

```bash
{
    "brand_id":1,
    "name":"voucher-1",
    "cost_in_point":5000

}
```

response

```bash
{
    "status": "success",
    "data": {
        "id": 2,
        "brand_id": 1,
        "name": "voucher-1",
        "cost_in_point": 5000
    }
}
```

- `/transaction/redemption`

body

```bash
{
  "customer_name": "Dwiky",
  "details": [
    {
      "voucher_id": 2,
      "quantity": 2
    }
  ]
}
```

response

```bash
{
    "status": "success",
    "data": {
        "id": 4,
        "customer_name": "Dwiky",
        "total_points": 10000,
        "details": [
            {
                "id": 3,
                "transaction_id": 4,
                "voucher_id": 2,
                "quantity": 2,
                "total_cost": 0,
                "voucher": {
                    "id": 0,
                    "brand_id": 0,
                    "name": "",
                    "cost_in_point": 0,
                    "brand": {
                        "id": 0,
                        "name": ""
                    }
                }
            }
        ]
    }
}
```

**GET**

- `/voucher?id={voucher_id}`

response

```bash
{
    "status": "success",
    "data": {
        "id": 2,
        "brand_id": 1,
        "name": "voucher-1",
        "cost_in_point": 5000,
        "brand": {
            "id": 1,
            "name": "bear brabnd"
        }
    }
}
```

- `/voucher/brand?id={brand_id}`

response

```bash
{
    "status": "success",
    "data": [
        {
            "id": 2,
            "brand_id": 1,
            "name": "voucher-1",
            "cost_in_point": 5000
        },
        {
            "id": 3,
            "brand_id": 1,
            "name": "voucher-1",
            "cost_in_point": 5000
        },
        {
            "id": 4,
            "brand_id": 1,
            "name": "voucher-1",
            "cost_in_point": 5000
        }
    ]
}
```

- `/transaction/redemption?transactionId={transactionId}`

response

```bash
{
    "status": "success",
    "data": {
        "id": 4,
        "customer_name": "Dwiky",
        "total_points": 10000,
        "details": [
            {
                "id": 3,
                "transaction_id": 4,
                "voucher_id": 2,
                "quantity": 2,
                "total_cost": 0,
                "voucher": {
                    "id": 2,
                    "brand_id": 1,
                    "name": "voucher-1",
                    "cost_in_point": 5000,
                    "brand": {
                        "id": 1,
                        "name": "bear brabnd"
                    }
                }
            },
            {
                "id": 4,
                "transaction_id": 4,
                "voucher_id": 2,
                "quantity": 2,
                "total_cost": 10000,
                "voucher": {
                    "id": 2,
                    "brand_id": 1,
                    "name": "voucher-1",
                    "cost_in_point": 5000,
                    "brand": {
                        "id": 1,
                        "name": "bear brabnd"
                    }
                }
            }
        ]
    }
}
```

**Author by Dwiky Satria Hutomo**
