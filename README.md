## Needed /Tech stacks
    + go
    + mysql
    + dep
    
## To get started follow this checklist:
    + create schema golang
    + dep ensure
    + dep ensure -add github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux
    + dep ensure -add "github.com/jinzhu/gorm"
    + dep ensure -add "github.com/jinzhu/gorm/dialects/mysql"
    + dep ensure -add "github.com/go-sql-driver/mysql"
    + go test -v ./service ./util  -cover
    + go build
    + go run go-mux-gorm (for production use supervisor)

note: route are in main.go just look at the file
