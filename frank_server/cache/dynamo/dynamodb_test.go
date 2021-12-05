package dynamo

import (
	"fmt"
	"frank_server/models"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const testEnv = "test"

func TestDynamo(t *testing.T) {
	initTest()
	test := NewDynamoStore(testEnv)
	key := "test"
	title := "title"
	recipes := []*models.Recipe{{Title: title}}
	err := test.PutRecipes(key, recipes)
	if err != nil {
		log.Panic(err)
	}
	res, err := test.GetRecipes(key)
	if err != nil {
		log.Panic(err)
	}
	if len(res) != 1 {
		t.Fail()
	}

	if res[0].Title != title {
		t.Fail()
	}
}

func initTest() {
	client := newDynamoClient(testEnv)
	_, err := client.CreateTable(newSchema())
	if err != nil {
		fmt.Printf("error creating table: %s, err: %s", tableName, err.Error())
	}
}

func newSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String(tableKey),
			AttributeType: aws.String("S"),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String(tableKey),
			KeyType:       aws.String("HASH"),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
}
