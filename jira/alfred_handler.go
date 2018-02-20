package jira

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"alfred-go/src/feedback"
	"encoding/xml"
)

type Handler struct {
	Url         string
	Credentials string
}

func (handler *Handler) GetTicketData(ticket string) (string) {

	request, newRequestError := http.NewRequest("GET", handler.Url+"/rest/api/2/issue/"+ticket, nil)
	request.Header.Add("Authorization", "Basic "+handler.Credentials)

	if newRequestError != nil {
		return MakeError(newRequestError)
	}

	client := &http.Client{}

	response, responseError := client.Do(request)

	if responseError != nil {
		return MakeError(responseError)
	}

	defer response.Body.Close()

	responseBytes, ioError := ioutil.ReadAll(response.Body)

	if ioError != nil {
		return MakeError(ioError)
	}

	res := Response{}

	unMarshalError := json.Unmarshal(responseBytes, &res)

	if unMarshalError != nil {
		return MakeError(unMarshalError)
	}

	fb := feedback.Feedback{}

	fb.AddItem(MakeItem(res, handler.Url))

	fbString, err := xml.Marshal(fb)

	if err != nil {
		return MakeError(err)
	}

	return string(fbString)

}

func MakeItem(response Response, baseUrl string) feedback.Item {

	title := response.Key + " | " + response.Fields.Summary
	description := "Status: " + response.Status.Name + " | Assignee: " + response.Assignee.DisplayName
	url := baseUrl + "/browse/" + response.Key
	return feedback.Item{response.Key, url, "", "", title, description, ""}

}

func MakeError(error error) string {
	return `<items><item arg=""><title>Error occurred while fetching jira ticket info.</title><subtitle>` + error.Error() + `</subtitle></item></items>`
}
