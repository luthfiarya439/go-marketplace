If you want to run this project without docker
1. Go to this project directory
2. go mod download
3. Install golang-migrate
  1. go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
4. Run migration
  ```migrate -database "mysql://your_mysql_user:your_mysql_password@tcp(your_mysql_address:your_mysql_port)/go-marketplace" -path db/migrations/ up```
5. Run
   1. go run main.go
