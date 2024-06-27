Jika ingin menjalankan tanpa docker
1. Buat ```.env``` file dari ```.env.example``` file
2. Masuk ke directory project
3. Jalankan ```go mod download```
4. Install golang-migrate ```go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
5. Lakukan migration ```migrate -database "mysql://YOUR_MYSQL_USER:YOUR_MYSQL_PASSWORD@tcp(YOUR_MYSQL_ADDRESS:YOUR_MYSQL_PORT)/YOUR_DATABASE" -path db/migrations/ up```
6. Jalankan perintah ```go run main.go```
  
Jika ingin menjalankan dengan docker
1. Buat ```.env``` file dari ```.env.example``` file
2. Jalankan perintah ```docker-compose up --build```

Informasi:
1. Default admin account
   1. Email : admin@mail.com password : admin@mail.com
2. Pada Postman untuk folder User, hanya menampilkan data khusus user yang sudah terautentikasi dan tidak akan berpengaruh terhadap user yang lain
3. Untuk request “detail”, “update”, dan “delete” (kecuali untuk request detail transaction) url parameter menggunakan id
   1. Contoh : http://localhost:8080/api/admin/categories/1
4. Untuk request “detail” pada transaction url parameter menggunakan transaction_code
   1. Contoh : http://localhost:8080/api/admin/transactions/qwioeu1234

