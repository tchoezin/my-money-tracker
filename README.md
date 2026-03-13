# My Money Tracker

Simple app for practicing full-stack development and taking better care of my finances.

## Backend API

Base URL: http://localhost:8080

### Run backend

1. Go to backend folder
2. Run: go run .

### Transaction JSON shape

All API requests/responses use lowercase JSON keys:

- id
- amount
- date
- type
- description
- category

Example request body:

{"amount": 1200, "type": "income", "description": "Salary", "category": "Work"}

Rules:

- type must be income or expense
- amount must be greater than 0

### Endpoints

- GET /transactions
- GET /transactions/:id
- POST /transactions
- PUT /transactions/:id
- DELETE /transactions/:id

### curl examples

Create transaction:

curl -X POST http://localhost:8080/transactions -H "Content-Type: application/json" -d '{"amount":1200,"type":"income","description":"Salary","category":"Work"}'

Get all transactions:

curl http://localhost:8080/transactions

Get transaction by id:

curl http://localhost:8080/transactions/1

Update transaction:

curl -X PUT http://localhost:8080/transactions/1 -H "Content-Type: application/json" -d '{"amount":100,"type":"expense","description":"Groceries","category":"Food"}'

Delete transaction:

curl -X DELETE http://localhost:8080/transactions/1