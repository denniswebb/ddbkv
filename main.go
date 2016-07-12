package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"flag"
	"strings"
	"log"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	table, key, value string
)

type Record struct {
	Env 	string	`dynamodbav:"env"`
	Value 	string	`dynamodbav:"value"`
}

func main() {
	flag.StringVar(&table, "table", "", "DynamoDB table name.")
	flag.StringVar(&key, "key", "", "Key")
	flag.StringVar(&value, "value", "", "Value")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage: ddbkv -table <table> -key <key> -value <value>\n\n")
		flag.PrintDefaults()
	}

	if strings.Trim(table, " ") == "" { log.Fatal("Error: table must be set.") }
	if strings.Trim(key, " ") == "" { log.Fatal("Error: key must be set.") }
	if strings.Trim(value, " ") == "" { log.Fatal("Error: value must be set.") }

	r := Record{Env: key, Value: value}

	ddb := dynamodb.New(session.New())

	item, err := dynamodbattribute.MarshalMap(r)

	if err != nil {
		log.Fatal("Failed to marshall.", err)
	}

	log.Println(item)

	_, err = ddb.PutItem(&dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(table),
	})

	if err != nil {
		log.Fatal("Failed to put item.", err)
	}
}
