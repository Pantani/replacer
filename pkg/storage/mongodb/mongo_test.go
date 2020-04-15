package mongodb

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type testCase struct {
	Name    string `json:"name" bson:"name"`
	Id      int    `json:"id" bson:"id"`
	Snippet string `json:"snippet" bson:"snippet"`
}

var testCaseBson1 = bson.M{
	"name":    "Danilo",
	"id":      1,
	"snippet": "<a></a>",
}

var testCaseBson2 = bson.M{
	"name":    "Marcos",
	"id":      3,
	"snippet": "<b></b>",
}

var testCaseBsonInvalid = bson.M{
	"test":     nil,
	"invalid":  50,
	"provider": "pantani.eth",
}

var testCase1 = testCase{
	Name:    "Danilo",
	Id:      1,
	Snippet: "<a></a>",
}

var testCase2 = testCase{
	Name:    "Marcos",
	Id:      3,
	Snippet: "<b></b>",
}

func Test_getResult(t *testing.T) {
	type args struct {
		doc    []interface{}
		result interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"test asset 1",
			args{doc: []interface{}{testCase1}, result: &[]testCase{}},
			&[]testCase{testCase1},
			false,
		}, {
			"test asset 2",
			args{doc: []interface{}{testCaseBson2}, result: &[]testCase{}},
			&[]testCase{testCase2},
			false,
		}, {
			"test 2 assets",
			args{doc: []interface{}{testCaseBson1, testCaseBson2}, result: &[]testCase{}},
			&[]testCase{testCase1, testCase2},
			false,
		}, {
			"want error",
			args{doc: []interface{}{testCaseBsonInvalid}, result: &[]testCase{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := getResult(tt.args.doc, tt.args.result)
			if tt.wantErr {
				assert.NotNil(t, tt.args.result)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, tt.args.result)
		})
	}
}

func TestNewMongoDbClientError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{"error", "..trust.."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoDbClient(tt.uri)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}
