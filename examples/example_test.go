package examples

import (
	"testing"

	dynamock "github.com/sardp/go-dynamock/v2"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var mock *dynamock.DynaMock

func init() {
	Dyna = new(MyDynamo)
	Dyna.Db, mock = dynamock.New()
}

func TestGetName(t *testing.T) {
	expectKey := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: "1",
		},
	}

	expectedResult := "jaka"
	result := dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{
				Value: expectedResult,
			},
		},
	}

	//lets start dynamock in action
	mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturns(result)

	actualResult, _ := GetName("1")
	if *actualResult != expectedResult {
		t.Errorf("Test Fail")
	}
}

func TestGetTransactGetItems(t *testing.T) {
	databaseOutput := dynamodb.TransactWriteItemsOutput{}

	mock.ExpectTransactWriteItems().Table("wrongTable").WillReturns(databaseOutput)

	err := GetTransactGetItems("")

	if err == nil {
		t.Errorf("Test failed")
	}
}
