curl запрос для регистрации пользователя

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser","password":"testpassword"}' http://localhost:5000/register
```

curl запрос для авторизации

```bash
curl -X POST \
  http://localhost:5000/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "testuser",  
    "password": "testpassword"
}'
```

curl запрос для получения информации об авторизованном пользователе

```bash
curl -X GET http://localhost:5000/me -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY"
```

curl запрос для изменения имени пользователя

```bash
curl -X PATCH http://localhost:5000/me -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg1MzE2N30.sB-L1xUvEvs6c22klsMIL3rWpVgSnKitfEu1h4rb528" -H "Content-Type: application/json" -d '{"new_username": "testuser"}'
```

curl запрос для изменения пароля пользователя

```bash
curl -X PATCH http://localhost:5000/me/password -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY" -H "Content-Type: application/json" -d '{"new_password": "password"}'
```
