package dynamock

import (
	"context"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// ToTable - method for set Table expectation
func (e *UpdateItemExpectation) ToTable(table string) *UpdateItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *UpdateItemExpectation) WithKeys(keys map[string]types.AttributeValue) *UpdateItemExpectation {
	e.key = keys
	return e
}

// Updates - method for set Updates expectation
func (e *UpdateItemExpectation) Updates(attrs map[string]types.AttributeValueUpdate) *UpdateItemExpectation {
	e.attributeUpdates = attrs
	return e
}

// WillReturns - method for set desired result
func (e *UpdateItemExpectation) WillReturns(res dynamodb.UpdateItemOutput) *UpdateItemExpectation {
	e.output = &res
	return e
}

// UpdateItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *params.TableName {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *params.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, params.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %s but found key %s", x.key, params.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, params.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %s but found key %s", x.attributeUpdates, params.AttributeUpdates)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item Expectation Not Found")
}
