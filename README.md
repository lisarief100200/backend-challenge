# Backend-Challenge

Setting DB in main.go
```
dsn := "host=localhost user=DBUSERNAME password=DBPASSWORD dbname=DBNAME port=5432 sslmode=disable TimeZone=Asia/Jakarta"
db, err := sql.Open("postgres", dsn)
```

## API

#### /customers
* `GET` : Get all customers
* `POST` : Create a new customer

#### /customers/:id
* `PUT` : Update a customer
* `DELETE` : Delete a customer

#Post Params
```
{
  "customer_name": "John Daw",
  "customer_cont_no": "123",
  "customer_address": "Sudirman",
  "total_buy": 10000,
  "creator_id": 1
}
```