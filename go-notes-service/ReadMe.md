curl запрос для получения всех заметок

```bash
curl -X GET http://localhost:8080/notes \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY"
```

curl запрос для создания заметки

```bash
curl -X POST http://localhost:8080/notes \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY" \
-H "Content-Type: application/json" \
-d '{
  "text": "This is a new note",
  "is_public": true
}'
```

curl запрос для получения заметки

```bash
curl -X GET http://localhost:8080/notes/3 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY"
```

curl запрос для изменения заметки

```bash
curl -X PATCH http://localhost:8080/notes/7 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY" \
-H "Content-Type: application/json" \
-d '{
  "text": "Updated note content",
  "is_public": true
}'
```

curl запрос для удаления заметки

```bash
curl -X DELETE http://localhost:8080/notes/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcxMzg2MTA0Nn0.hldYR4VxZKTd-8yQe40D8AnvtypPChBOiBivoqqOcQY"
```
