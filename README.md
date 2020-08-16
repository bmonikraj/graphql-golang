# graphql-golang
GraphQL based Golang Service for Data Querying

## Build Binary from source

`cd graphql-golang`

`go build -o qogl main.go`

## Run the Binary as executable
**Pre-requisite : "Build Binary from source"**

`cd graphql-golang`

`./goql <host> <port> <mysql host:port> <mysql user> <mysql pwd> <mysql db-name>`

Example : 

`./goql 127.0.0.1 8080 localhost:3306 user-dummy pwd-dummy db-dummy`

## Dummy Data Sample

```
mysql> SELECT * FROM data LIMIT 5;
+---------+----------+---------+--------------------------------------+-----------+------------+--------+---------------------+------------------------------------------------------------------------------------+
| ID      | SERVER   | NETWORK | REQ_ID                               | REGION_IN | REGION_OUT | CLOAD  | REQ_TIME            | HASH_ID                                                                            |
+---------+----------+---------+--------------------------------------+-----------+------------+--------+---------------------+------------------------------------------------------------------------------------+
| 1827471 | YAMUNA   | S1      | f7f36bf7-1402-40ff-b389-4cdc56828ad6 | SYS9      | SYS1       |      0 | 2020-07-20 17:55:03 | 0x5b40cf43860bb69bca3ea38e53a6a4ad94d263c44a0e445b83954fbb6505033eb0a4138518cca1e9 |
| 1827472 | NARMADA  | S1      | f74ebc99-e2b0-4bfa-a56f-52aa7dfe948d | SYS6      | SYS8       | 430828 | 2020-07-20 17:55:06 | 0x0ce67e709ce0650e943bec7493ac5b75efdada2f29c323b194d7177854d5bcfd24e295663956aa89 |
| 1827473 | YAMUNA   | S1      | 6f38edb3-67a5-47e6-9b79-8a191432bd4c | SYS3      | SYS8       |      0 | 2020-07-20 17:55:08 | 0x183bc7e040e2d8f482749acc5766da73c2c8bb307ceb3b24c57c0be19eed4ff72997c56608dd32f9 |
| 1827474 | GANGA_S2 | S2      | ba805b63-a813-4e2a-9629-9c2e253a24e8 | SYS2      | SYS6       |      0 | 2020-07-20 17:55:10 | 0x1a991ad2e53ca8bfe8e29217bffc3cd7f01cbe995d8ca51c05618758481e7446936328cc695fe394 |
| 1827475 | NARMADA  | S1      | 3b91e81d-20ba-4e67-88aa-ea26b234e971 | SYS7      | SYS5       | 343204 | 2020-07-20 17:55:14 | 0xc8f0852469f72bf1aabdfa75843988a2433fba28b60e4a480f186716f59d2df433a9be655fe03c75 |
+---------+----------+---------+--------------------------------------+-----------+------------+--------+---------------------+------------------------------------------------------------------------------------+
5 rows in set (0.00 sec)
```