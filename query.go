package dynamock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/guregu/dynamo/v2"
)

// Table - method for set Table expectation
func (e *QueryExpectation) Table(table string) *QueryExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *QueryExpectation) WillReturns(res dynamodb.QueryOutput) *QueryExpectation {
	e.output = &res
	return e
}

// QueryWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	if len(e.dynaMock.QueryExpect) > 0 {
		x := e.dynaMock.QueryExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.QueryOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.QueryExpect = append(e.dynaMock.QueryExpect[:0], e.dynaMock.QueryExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.QueryOutput{}, dynamo.ErrNotFound
}
