## Clone this repo into folder github.com by cmd:
> git clone https://github.com/tsrnd/goweb4.git

## How to run app:
> 1. cd goweb4
> 2. run command `go build` to build source
> 3. run command ` ./goweb4`
> 4. open `http://localhost:8080/` to see the index page
> 5. copy `.env.example` and rename it into `.env`

## If you have added new package, please modify this readme.md file by adding command to get this package here
List packages:
1. "github.com/gorilla/mux"
```
go get -u github.com/gorilla/mux
```
2. "github.com/lib/pq"
```
go get -u github.com/lib/pq
```
3. "http://jinzhu.me/gorm/"
```
go get -u github.com/jinzhu/gorm
```
4. "github.com/joho/godotenv"
```
go get github.com/joho/godotenv
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
