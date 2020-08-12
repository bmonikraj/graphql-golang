package main

import (
    "encoding/json"
    "fmt"
    "log"
    "github.com/graphql-go/graphql"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Records struct {
	ID int `json:"ID"`
	SERVER string `json:"SERVER"`
	NETWORK string `json:"NETWORK"`
	REQ_ID string `json:"REQ_ID"`
	REGION_IN string `json:"REGION_IN"`
	REGION_OUT string `json:"REGION_OUT"`
	CLOAD int `json:"CLOAD"`
	REQ_TIME string `json:"REQ_TIME"`
	HASH_ID string `json:"HASH_ID"`
}

var recordType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Records",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "server": &graphql.Field{
                Type: graphql.String,
            },
            "network": &graphql.Field{
                Type: graphql.String,
            },
            "req_id": &graphql.Field{
                Type: graphql.String,
            },
            "region_in": &graphql.Field{
                Type: graphql.String,
            },
            "region_out": &graphql.Field{
                Type: graphql.String,
            },
            "cload": &graphql.Field{
                Type: graphql.Int,
            },
            "req_time": &graphql.Field{
                Type: graphql.String,
            },
            "hash_id": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var fieldType = graphql.Fields{
    "records": &graphql.Field{
        Type: recordType,
        Description: "Get Record By ID",
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
                Type: graphql.Int,
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            id, ok := p.Args["id"].(int)
            if ok {
                db, err := sql.Open("mysql", "root:Root123#@tcp(13.233.86.58:3306)/graphql")

                if err != nil {
                    panic(err.Error())
                }

                defer db.Close()

                results, err := db.Query("SELECT * FROM data WHERE id = ?", id)
                if err != nil {
                    panic(err.Error())
                }
                
                for results.Next() {
                    var record Records
                    err = results.Scan(&record.ID, &record.SERVER, &record.NETWORK, &record.REQ_ID, &record.REGION_IN, &record.REGION_OUT, &record.CLOAD, &record.REQ_TIME, &record.HASH_ID)
                    if err != nil {
                        panic(err.Error())
                    }                    
                    return record, nil
                }
            }
            return nil, nil
        },
    },
    "list": &graphql.Field{
        Type: graphql.NewList(recordType),
        Description: "Get Top N Records List",
        Args: graphql.FieldConfigArgument{
            "num": &graphql.ArgumentConfig{
                Type: graphql.Int,
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            num, ok := p.Args["num"].(int)
            if ok {
                db, err := sql.Open("mysql", "root:Root123#@tcp(13.233.86.58:3306)/graphql")

                if err != nil {
                    panic(err.Error())
                }

                defer db.Close()

                results, err := db.Query("SELECT * FROM data LIMIT ?", num)
                if err != nil {
                    panic(err.Error())
                } else {
                    fmt.Println("results fetched")
                }
                
                var recordsList []Records
                for results.Next() {
                    var record Records
                    err = results.Scan(&record.ID, &record.SERVER, &record.NETWORK, &record.REQ_ID, &record.REGION_IN, &record.REGION_OUT, &record.CLOAD, &record.REQ_TIME, &record.HASH_ID)
                    if err != nil {
                        panic(err.Error())
                    }
                    recordsList = append(recordsList, record)
                }
                return  recordsList, nil
            }
            return  nil, nil
        },
    },
}

var mutationType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "RecordMutation",
        Fields: graphql.Fields{
            "create": &graphql.Field{
                Type:        recordType,
                Description: "Create a new Record",
                Args: graphql.FieldConfigArgument{
                    "id" : &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Int),
                    },
                    "server": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "network": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "req_id": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "region_in": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "region_out": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "cload": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Int),
                    },
                    "req_time": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "hash_id": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                },
                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                    record := Records{
                        ID: params.Args["id"].(int),
                        SERVER: params.Args["server"].(string),
                        NETWORK: params.Args["network"].(string),
                        REQ_ID: params.Args["req_id"].(string),
                        REGION_IN: params.Args["region_in"].(string),
                        REGION_OUT: params.Args["region_out"].(string),
                        CLOAD: params.Args["cload"].(int),
                        REQ_TIME: params.Args["req_time"].(string),
                        HASH_ID: params.Args["hash_id"].(string),
                    }
                    db, err := sql.Open("mysql", "root:Root123#@tcp(13.233.86.58:3306)/graphql")

                    if err != nil {
                        panic(err.Error())
                    }

                    defer db.Close()

                    insertQuery, err := db.Prepare("INSERT INTO data(ID, SERVER, NETWORK, REQ_ID, REGION_IN, REGION_OUT, CLOAD, REQ_TIME, HASH_ID) VALUES(?,?,?,?,?,?,?,?,?)")

                    if err != nil {
                        panic(err.Error())
                    }

                    insertQuery.Exec(record.ID, record.SERVER, record.NETWORK, record.REQ_ID, record.REGION_IN, record.REGION_OUT, record.CLOAD, record.REQ_TIME, record.HASH_ID)

                    return record, nil
                },
            },
        },
    },
)


func main() {
    
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fieldType}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationType}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error : %v", err)
    }

    /*
    query := `
        {
            list (num:5){
                id
                req_time
                cload
            }
        }
    `
    */
    
    
    query := `
        {
            records (id:1){
                id
                req_time
                cload
                server
            }
        }
    `
    

    /*
    query := `
        mutation {
            create(id:1, server: "YAMUNA", network: "S2", req_id: "f74ebc99-e2b0-4bfa-a56f-52aa7dfe948d", region_in: "SYS10", region_out: "SYS3", cload: 1003, req_time: "2020-08-12 17:30:00", hash_id: "0x5b40cf43860bb69bca3ea38e53a6a4ad94d263c4dd0e445b83954fbb6505033eb0a4138518cca1e9") {
                id
                cload
            }
        }
    `
    */

    
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON)
}