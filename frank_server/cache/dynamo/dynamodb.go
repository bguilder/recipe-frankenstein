package dynamo

import (
	"frank_server/cache"
	"frank_server/scraper"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "live-recipe-cache"
const tableKey = "SearchKey"

type dynamoStore struct {
	client *dynamodb.DynamoDB
}

type RecipeDocument struct {
	SearchKey string
	Recipes   []*scraper.Recipe
}

func NewDynamoStore(env string) cache.Store {
	client := newDynamoClient(env)
	return &dynamoStore{client: client}
}

// TODO: per best practice we should create this outside of the handle function to re-use connections
// https://aws.amazon.com/blogs/database/building-enterprise-applications-using-amazon-dynamodb-aws-lambda-and-golang/
func newDynamoClient(env string) *dynamodb.DynamoDB {
	if env == "live" {
		dbSession := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		return dynamodb.New(dbSession)
	}
	// TODO: separate out into config
	log.Printf("here?? : %s", tableName)
	awsCfg := aws.NewConfig().
		WithRegion("us-east-1").
		WithEndpoint("http://localhost:8000").
		WithCredentials(credentials.NewStaticCredentials("test", "test", "test"))

	sess, err := session.NewSession(awsCfg)

	if err != nil {
		log.Panic(err)
	}

	return dynamodb.New(sess, awsCfg)
}

func (d *dynamoStore) PutRecipes(searchKey string, recipes []*scraper.Recipe) error {
	doc := &RecipeDocument{SearchKey: searchKey, Recipes: recipes}
	dbItem, err := dynamodbattribute.MarshalMap(doc)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      dbItem,
		TableName: aws.String(tableName),
	}

	_, err = d.client.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoStore) GetRecipes(searchKey string) ([]*scraper.Recipe, error) {
	getInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			tableKey: {
				S: aws.String(searchKey),
			},
		},
		TableName: aws.String(tableName),
	}

	item, err := d.client.GetItem(getInput)
	if err != nil {
		return nil, err
	}

	result := &RecipeDocument{}
	if err := dynamodbattribute.UnmarshalMap(item.Item, result); err != nil {
		return nil, err
	}

	return result.Recipes, nil
}
