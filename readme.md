## Clone this repo into folder github.com by cmd:
> git clone https://github.com/tsrnd/goweb4.git

## How to run app:
> 1. cd goweb4
> 2. create `.env` file (password and db_user depend on your environment)
```
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=goweb4
DB_HOST=127.0.0.1
PORT=3306
```
> 2. run command `go build` to build source
> 3. run command ` ./goweb4`
> 4. open `http://localhost:8080/` to see the index page
> 5. copy `.env.example` and rename it into `.env`

## If you have added new package, please modify this readme.md file by adding command to get this package here
List packages:
1. "github.com/gorilla/mux"
2. "github.com/joho/godotenv"
```
$ go get -u github.com/gorilla/mux
```
2. "github.com/lib/pq"
```
$ go get -u github.com/lib/pq
```
3. "http://jinzhu.me/gorm/"
```
$ go get -u github.com/jinzhu/gorm
```
4. "github.com/joho/godotenv"
```
$ go get github.com/joho/godotenv
```
5. "http://www.gorillatoolkit.org/pkg/schema"
`Package gorilla/schema fills a struct with form values`
```
$ go get github.com/gorilla/schema
```
6. "github.com/nfnt/resize"
`Package nfnt/resize resize a picture`
```
$ go get github.com/nfnt/resize
```
7. "github.com/stretchr/objx"
```
$ go get github.com/stretchr/objx
```
## Install PostgreSQL
Ubuntu:
```
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```
### How to use Postgre:
Use this cmd to access Postgres:
```
sudo -u postgres psql
```

### Please make sure that you push code to `new branch`
### Create new `Pull Request` and add `2 persons` into this pull request to review your code
