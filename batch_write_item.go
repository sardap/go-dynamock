package dynamock

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// WithRequest - method for set Request expectation
func (e *BatchWriteItemExpectation) WithRequest(input map[string][]*types.WriteRequest) *BatchWriteItemExpectation {
	e.input = input
	return e
}

// WillReturns - method for set desired result
func (e *BatchWriteItemExpectation) WillReturns(res dynamodb.BatchWriteItemOutput) *BatchWriteItemExpectation {
	e.output = &res
	return e
}

// BatchWriteItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	if len(e.dynaMock.BatchWriteItemExpect) > 0 {
		for i, x := range e.dynaMock.BatchWriteItemExpect {
			if x.input != nil {
				if reflect.DeepEqual(x.input, params.RequestItems) {
					e.dynaMock.BatchWriteItemExpect = append(e.dynaMock.BatchWriteItemExpect[:i], e.dynaMock.BatchWriteItemExpect[i:]...)
					return x.output, nil
				}
			}
		}
	}

	return &dynamodb.BatchWriteItemOutput{}, fmt.Errorf("Batch Write Item Expectation Failed. Expected one of %s to equal %s", e.dynaMock.BatchWriteItemExpect, params.RequestItems)
}
