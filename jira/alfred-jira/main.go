package main

import (
	"os"
	"fmt"
	"alfred-jira/jira"
)

func main(){

	args := os.Args[1:4]
	url := args[0]
	credentials := args[1]
	query := args[2]

	h := jira.Handler{url, credentials}

	response := h.GetTicketData(query)

	fmt.Println(response);

}
