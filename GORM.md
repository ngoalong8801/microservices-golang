# Gorm Introduction

## Get the library
### Connect DB
```shell
    go get "gorm.io/driver/postgres"
    go get "gorm.io/gorm"
```
### Code gen
```shell
    go get "gorm.io/gen"
```

## How to connect ?
1. Config database

    [Example config database with port, user ...](server/config/env.conf.go)
2. Connect 

    [Example connect gorm with database](server/dal/psql.go)

After that u can use Module in file psql.go which contains <strong>db *gorm.DB</strong> used to get data.

## Declaring model
### How to declare model in gorm ?
1. Just simply define struct type with name equal with name of table in your database. For example, you have database named <strong> users</strong>. Let's declare model like this
```sql
    create table users (
                           id bigserial primary key,
                           name varchar(128)
    )
```

```go
        type User struct {
	        ID        int64
	        Name      string
}
```
### Associations
#### Declare
In this project i have implemented fully the association of database with gorm: [link model database](server/dal/model)
- users with ordertabs:<strong> one to many </strong>
- products - categories: <strong> many to many </strong>
- orderdetail - products - ordertabs: <strong> one to one </strong>

Gorm provide some tags helping define clearly relationship
- `gorm:"primaryKey"`: indicate that this field is primary key
- `gorm:"foreignkey:ProductID;references:ProductID"` indicate references to table in which it depends on
- `gorm:"many2many:product_category;foreignKey:CategoryID;joinForeignKey:CategoryID;References:ProductID;joinReferences:ProductID"`: many to many indicate the name table in database

#### Loading eager vs lazy
By default, gorm will not automatically load all associations of model \
If you need load model along with its associations. U have to use Preload to load it:\
For example: [example link](server/utils/test/query.test.go)

## Transaction 
[link reference](https://gorm.io/docs/transactions.html)\
[link example](server/utils/test/transaction.test.go)
## Query
### Query builder
1. Config
   - gen.NewGenerator: config path folder codegen
   - g.ApplyBasic: generate query for database
   - g.Execute: executing generate\
   [link config code gen](server/utils/gen/gen.go)
2. Usage\
  [link reference how  to use query with code gen](https://github.com/go-gorm/gen#query)
### Native query
[reference](https://gorm.io/docs/sql_builder.html)
### ORM-specific query
[link doc query](https://gorm.io/docs/query.html)
[link example](server/utils/test/query.test.go)

## Database Migration
### Install
#### Download pre-built migrate-cli binary (Windows, MacOS, or Linux)
```shell
 https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```
Make sure that u have assigned it to your $PATH
#### Usage
Using make file like [makefile](server/Makefile) with option:

Migrate up: it will automatically run file migration defined in your PATH to execute to generate your database. It depends on the last version that u have run in the past. If having any errors migration will be dirty, and you need manually to fix that.\
Migrate down: define your own function to roll back your version\
Migrate force: forcing version migrate to version that u want to get there\
Migrate create: create file up and down with option (seq, timezone, etc ...)

[file migration example](server/db/migration)\
[link reference](https://github.com/golang-migrate/migrate)

