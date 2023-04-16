
#curl -v -X 'GET' 'localhost:8081/api/v1/auth' -H 'accept: application/json'

curl -v -X 'POST' 'localhost:8081/api/v1/login' -H 'accept: application/json' -H 'Cookie: _csrf=9Bj8Hd8HJEzZrsMr2L8lujbG4lJf6HDq; Expires=Mon, 17 Apr 2023 18:42:35 GMT' -H 'X-Csrf-Token: 9Bj8Hd8HJEzZrsMr2L8lujbG4lJf6HDq' -H 'Content-Type: application/json' -d '{"email": "string@mail.ru","password": "12345678"}'
