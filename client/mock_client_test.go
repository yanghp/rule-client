package client

import (
	"testing"
)

func TestMockEngine(t *testing.T) {
	// Arrange
	//rule := NewMockRule(func(pl *dto.Payload) dto.Data {
	//	return dto.Data{"foo": "bar"}
	//})
	//engine := NewMockEngine(map[string]Tenanter{
	//	"conf-prod": rule,
	//})
	//
	//// Act
	//c, err := engine.Of("conf-prod").Payload(nil)
	//
	//// asserts
	//assert.NoError(t, err)
	//assert.Equal(t, "bar", c.Get("foo"))
}
