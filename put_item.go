package dynamock

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// ToTable - method for set Table expectation
func (e *PutItemExpectation) ToTable(table string) *PutItemExpectation {
	e.table = &table
	return e
}

// WithItems - method for set Items expectation
func (e *PutItemExpectation) WithItems(item map[string]types.AttributeValue) *PutItemExpectation {
	e.item = item
	return e
}

// WillReturns - method for set desired result
func (e *PutItemExpectation) WillReturns(res dynamodb.PutItemOutput) *PutItemExpectation {
	e.output = &res
	return e
}

// PutItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	if len(e.dynaMock.PutItemExpect) > 0 {
		x := e.dynaMock.PutItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *params.TableName {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *params.TableName)
			}
		}

		if x.item != nil {
			if !reflect.DeepEqual(x.item, params.Item) {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect item %s but found item %s", x.item, params.Item)
			}
		}

		// delete first element of expectation
		e.dynaMock.PutItemExpect = append(e.dynaMock.PutItemExpect[:0], e.dynaMock.PutItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.PutItemOutput{}, fmt.Errorf("Put Item Expectation Not Found")
}
