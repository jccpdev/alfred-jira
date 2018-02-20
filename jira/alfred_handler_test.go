package jira

import (
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_it_will_get_ticket_info_from_jira(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://jira.someCompany.com/rest/api/2/issue/A-123",
		httpmock.NewStringResponder(200, GetResponseStub()))

	handler := Handler{
		Url:         "https://jira.someCompany.com",
		Credentials: "usernamePassword"}

	response := handler.GetTicketData("A-123")

	assert.Equal(t, `<items><item uid="A-123" arg="https://jira.someCompany.com/browse/A-123"><title>A-123 | The Summary</title><subtitle>Status: Developing | Assignee: Juan Contreras</subtitle></item></items>`, response)
}
