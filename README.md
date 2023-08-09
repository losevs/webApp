# webApp
CRUD Web App using fiber framework

example of possible cmds:
curl -X GET http://localhost/
curl -X GET http://localhost/incr
curl -X GET http://localhost/decr
curl -X GET http://localhost/tab
curl -d "name=Chris&email=Chris42@gmail.com" -X POST http://localhost/task
curl -d "name=noone&email=none" -X PATCH http://localhost/patch/1
curl -X DELETE http://localhost/del/1