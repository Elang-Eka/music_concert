# Music Concert 2022

## Tools
1. Go (Gin Framework)
2. MySQL
3. sqlx
4. viper
5. swagger

## Run Program
```
cd src
go run main.go
```

## API Server
1. ping

Request : 
```
GET http://localhost:8080/event
```
Response :
"200"
```
{
  "message": {
    "title": "OK",
    "body": "Request successful"
  },
  "metadata": {
    "path": "localhost:8080/ping",
    "statusCode": 200,
    "status": "OK",
    "message": "GET /ping [200] OK",
    "timestamp": "2022-04-09T13:44:14+07:00"
  },
  "data": "PONG!"
}
```
Image :
![image](https://user-images.githubusercontent.com/54503473/162560274-8ac119a6-d3c0-4f48-b194-5c2dec67012b.png)

2. Get List Event
Request : 
```
GET http://localhost:8080/event
```
Response :
"200"
```
{
    "message": {
        "title": "OK",
        "body": "Request successful"
    },
    "metadata": {
        "path": "localhost:8080/event",
        "statusCode": 200,
        "status": "OK",
        "message": "GET /event [200] OK",
        "timestamp": "2022-04-09T13:46:57+07:00"
    },
    "data": [
        {
            "id": 1,
            "name": "Metalica",
            "price": 150000,
            "location": "Kediri",
            "date": "2022-04-09T00:00:00Z",
            "time": "21:00:00",
            "organizer": "PT. Gudang Garam Tbk"
        }
    ]
}
```
Image :
![image](https://user-images.githubusercontent.com/54503473/162560361-786368fe-898e-4d4e-8c8d-3972d5eee2bf.png)

3. Booking Ticket
Request : 
```
POST http://localhost:8080/booking
```
Response :
"200"
```
{
    "message": {
        "title": "OK",
        "body": "Request successful"
    },
    "metadata": {
        "path": "localhost:8080/booking",
        "statusCode": 200,
        "status": "OK",
        "message": "POST /booking [200] OK",
        "timestamp": "2022-04-09T13:54:49+07:00"
    },
    "data": {
        "id": 6,
        "event_id": 1,
        "transaction_date": "0001-01-01T00:00:00Z",
        "quantity": 3,
        "total_price": 450903,
        "payment_method": "BNI",
        "action": "",
        "code": 82241,
        "User": [
            {
                "id": 10,
                "name": "Endy",
                "age": 22,
                "gender": "male",
                "email": "endy@gmail.com",
                "transaction_id": 6
            }
        ]
    }
}
```
Image :
![image](https://user-images.githubusercontent.com/54503473/162560689-8217d00e-53a9-410d-b59a-51b599ae05bb.png)
![image](https://user-images.githubusercontent.com/54503473/162560743-43019b47-ec4d-4188-8756-ee27bac61f42.png)

4. Pay ticket

Request : 
```
PUT http://localhost:8080/payment
```
Response :
"200"
```
{
    "message": {
        "title": "OK",
        "body": "Request successful"
    },
    "metadata": {
        "path": "localhost:8080/payment",
        "statusCode": 200,
        "status": "OK",
        "message": "PUT /payment [200] OK",
        "timestamp": "2022-04-09T13:59:20+07:00"
    },
    "data": "successful payment"
}
```
Response : "500"
```
{
    "message": {
        "title": "Error",
        "body": "Request failure"
    },
    "metadata": {
        "path": "localhost:8080/payment",
        "statusCode": 500,
        "status": "Internal Server Error",
        "message": "PUT /payment [500] Internal Server Error",
        "timestamp": "2022-04-09T14:05:42+07:00",
        "error": {
            "code": 500,
            "message": "Event Ticket Expired! Please order again"
        }
    }
}
```
### note!
pembayaran error jika lebih dari 30 menit

Image :
"200"
![image](https://user-images.githubusercontent.com/54503473/162560780-b4ec3893-dc7e-4c13-8f4a-0630bf01e9f9.png)
"500"
![image](https://user-images.githubusercontent.com/54503473/162561053-9c19a8b9-3de2-47b4-98f7-39d93e56f39c.png)

