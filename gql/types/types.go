package types

import (
    "os"
    "github.com/graphql-go/graphql"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/bmonikraj/goql/model"
)

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

var FieldType = graphql.Fields{
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
                db, err := sql.Open("mysql", os.Args[4]+":"+os.Args[5]+"@tcp("+os.Args[3]+")/"+os.Args[6])

                if err != nil {
                    panic(err.Error())
                }

                defer db.Close()

                results, err := db.Query("SELECT * FROM data WHERE id = ?", id)
                if err != nil {
                    panic(err.Error())
                }
                
                for results.Next() {
                    var record model.Records
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
                db, err := sql.Open("mysql", os.Args[4]+":"+os.Args[5]+"@tcp("+os.Args[3]+")/"+os.Args[6])

                if err != nil {
                    panic(err.Error())
                }

                defer db.Close()

                results, err := db.Query("SELECT * FROM data LIMIT ?", num)
                if err != nil {
                    panic(err.Error())
                }
                
                var recordsList []model.Records
                for results.Next() {
                    var record model.Records
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

var MutationType = graphql.NewObject(
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
                    record := model.Records{
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
                    db, err := sql.Open("mysql", os.Args[4]+":"+os.Args[5]+"@tcp("+os.Args[3]+")/"+os.Args[6])

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
