package dynamock

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
)

// Name - method for set Name expectation
func (e *CreateTableExpectation) Name(table string) *CreateTableExpectation {
	e.table = &table
	return e
}

// KeySchema - method for set KeySchema expectation
func (e *CreateTableExpectation) KeySchema(keySchema []*types.KeySchemaElement) *CreateTableExpectation {
	e.keySchema = keySchema
	return e
}

// WillReturns - method for set desired result
func (e *CreateTableExpectation) WillReturns(res dynamodb.CreateTableOutput) *CreateTableExpectation {
	e.output = &res
	return e
}

// CreateTable - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) CreateTable(ctx context.Context, parmas *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	if len(e.dynaMock.CreateTableExpect) > 0 {
		x := e.dynaMock.CreateTableExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *parmas.TableName {
				return &dynamodb.CreateTableOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *parmas.TableName)
			}
		}

		if x.keySchema != nil {
			if !reflect.DeepEqual(x.keySchema, parmas.KeySchema) {
				return &dynamodb.CreateTableOutput{}, fmt.Errorf("Expect keySchema %s but found keySchema %s", x.keySchema, parmas.KeySchema)
			}
		}

		// delete first element of expectation
		e.dynaMock.CreateTableExpect = append(e.dynaMock.CreateTableExpect[:0], e.dynaMock.CreateTableExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.CreateTableOutput{}, dynamo.ErrNotFound
}
