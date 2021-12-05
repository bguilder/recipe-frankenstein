package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	tableName = "test-table"
)

type testRecipe struct {
	TestKey    string
	RecipeName string
	Ingredient string
	CookTime   int
}

func main() {
	client, err := newDynamoClient()
	if err != nil {
		fmt.Println("error new client", err.Error())
	}
	_, err = client.CreateTable(newSchema())
	if err != nil {
		fmt.Println("error creating table", err.Error())
	}

	dbKey := "testkey"

	recipe1 := &testRecipe{TestKey: dbKey, Ingredient: "tomato", RecipeName: "blah", CookTime: 10}

	dbItem, err := dynamodbattribute.MarshalMap(recipe1)
	if err != nil {
		fmt.Println("error creating attribute", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      dbItem,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(input)
	if err != nil {
		fmt.Println("error putting item", err.Error())
	}

	getInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"TestKey": {
				S: aws.String(dbKey),
			},
		},
		TableName: aws.String(tableName),
	}

	item, err2 := client.GetItem(getInput)
	if err2 != nil {
		fmt.Println("error getting item", err2.Error())
	}

	result := &testRecipe{}
	fmt.Printf("item: %+v\n\n", item)
	if err := dynamodbattribute.UnmarshalMap(item.Item, result); err != nil {
		fmt.Println("error unmarshalling item", err.Error())
	}
	fmt.Printf("complete! result: %+v", result)
}

func newDynamoClient() (*dynamodb.DynamoDB, error) {
	awsCfg := aws.NewConfig().
		WithRegion("us-east-1").
		WithEndpoint("http://localhost:8000")

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess, awsCfg), nil
}

func newSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String("TestKey"),
			AttributeType: aws.String("S"),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("TestKey"),
			KeyType:       aws.String("HASH"),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
}
