If you want to run this project without docker
1. make ```.env``` file from ```.env-example``` file
2. Go to this project directory
3. ```go mod download```
4. Install golang-migrate
   1. ```go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
5. Run migration
   1. ```migrate -database "mysql://your_mysql_user:your_mysql_password@tcp(your_mysql_address:your_mysql_port)/go-marketplace" -path db/migrations/ up```
6. Run
   1. ```go run main.go```
  
If you want to run this project with docker
1. make ```.env``` file from ```.env-example``` file
2. Run ```docker-compose up --build```

More information about this project visit : https://docs.google.com/document/d/1CTi-PGJ9yxbF_XM1jLXKZnaxJlX3cE7KcgL_iuKS9F4/edit?usp=sharing
