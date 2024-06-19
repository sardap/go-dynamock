package dynamock

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
)

// WithRequest - method for set Request expectation
func (e *BatchGetItemExpectation) WithRequest(input map[string]*types.KeysAndAttributes) *BatchGetItemExpectation {
	e.input = input
	return e
}

// WillReturns - method for set desired result
func (e *BatchGetItemExpectation) WillReturns(res dynamodb.BatchGetItemOutput) *BatchGetItemExpectation {
	e.output = &res
	return e
}

// BatchGetItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	if len(e.dynaMock.BatchGetItemExpect) > 0 {
		x := e.dynaMock.BatchGetItemExpect[0] //get first element of expectation

		if x.input != nil {
			if !reflect.DeepEqual(x.input, input.RequestItems) {
				return &dynamodb.BatchGetItemOutput{}, fmt.Errorf("Expect input %s but found input %s", x.input, input.RequestItems)
			}
		}

		// delete first element of expectation
		e.dynaMock.BatchGetItemExpect = append(e.dynaMock.BatchGetItemExpect[:0], e.dynaMock.BatchGetItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.BatchGetItemOutput{}, dynamo.ErrNotFound
}
