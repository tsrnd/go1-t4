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

## If you have added new package, please modify this readme.md file by adding command to get this package here
List packages:
1. "github.com/gorilla/mux"
2. "github.com/joho/godotenv"
```
go get -u github.com/gorilla/mux
```

### Please make sure that you push code to `new branch`
### Create new `Pull Request` and add `2 persons` into this pull request to review your code
