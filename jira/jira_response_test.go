package jira

import (
	"testing"
	"encoding/json"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func Test_it_can_parse_a_jira_json_response(t *testing.T){

	fakeResponseJson := GetResponseStub()

	res := Response{}
	json.Unmarshal([]byte(fakeResponseJson), &res)

	fmt.Println(res)

	assert.Equal(t, "juan.contreras", res.Assignee.Name)
	assert.Equal(t, "Juan Contreras", res.Assignee.DisplayName)
	assert.Equal(t, "juan.contreras", res.Reporter.Name)
	assert.Equal(t, "Juan Contreras", res.Reporter.DisplayName)
	assert.Equal(t, "A-123", res.Key)
	assert.Equal(t, "The Summary", res.Fields.Summary)
	assert.Equal(t, "Developing", res.Fields.Status.Name)

	//assert.Len(t, res.Jobs, 2)

}
