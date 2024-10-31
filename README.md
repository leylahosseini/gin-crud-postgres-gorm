# gin-crud-postgres-gorm

This project is basically made for Golang, which uses Gin Gorm and Postgres ,which includes the CRUD operation.

###
curl -X POST localhost:8080/users -d '{"name":"test1" , "email":"test1@test.com" }' --header "Content-Type: application/json"
###
curl -X PUT localhost:8080/users/:id -d '{"name":"test1" , "email":"test1@test.com" }' --header "Content-Type: application/json"
###
curl -X DELETE localhost:8080/users/:id
###
curl -X GET localhost:8080/users

### 
Support Swaggers in 
localhost:8080/docs/index.html
