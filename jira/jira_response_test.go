package jira

import (
	"testing"
	"encoding/json"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func Test_it_can_parse_a_jira_json_response(t *testing.T){

	fakeResponseJson := `{
    "expand": "renderedFields,names,schema,operations,editmeta,changelog,versionedRepresentations",
    "id": "1090718",
    "self": "https://jira.bancvue.com/rest/api/2/issue/1090718",
    "key": "KL-1379",
    "fields": {
        "resolution": null,
        "lastViewed": "2018-02-12T21:30:49.561-0600",
        "labels": [
            "Reversals"
        ],
        "aggregatetimeoriginalestimate": null,
        "issuelinks": [
            {
                "id": "154749",
                "self": "https://jira.bancvue.com/rest/api/2/issueLink/154749",
                "type": {
                    "id": "10010",
                    "name": "3-Related task",
                    "inward": "6-is related task of (Related)",
                    "outward": "5-related task (Related)",
                    "self": "https://jira.bancvue.com/rest/api/2/issueLinkType/10010"
                },
                "outwardIssue": {
                    "id": "1081745",
                    "key": "KL-1351",
                    "self": "https://jira.bancvue.com/rest/api/2/issue/1081745",
                    "fields": {
                        "summary": "Implement Replay Processor",
                        "status": {
                            "self": "https://jira.bancvue.com/rest/api/2/status/10144",
                            "description": "",
                            "iconUrl": "https://jira.bancvue.com/images/icons/statuses/generic.png",
                            "name": "In Production",
                            "id": "10144",
                            "statusCategory": {
                                "self": "https://jira.bancvue.com/rest/api/2/statuscategory/3",
                                "id": 3,
                                "key": "done",
                                "colorName": "green",
                                "name": "Done"
                            }
                        },
                        "priority": {
                            "self": "https://jira.bancvue.com/rest/api/2/priority/6",
                            "iconUrl": "https://jira.bancvue.com/images/icons/custom/Low.png",
                            "name": "4 - Low",
                            "id": "6"
                        },
                        "issuetype": {
                            "self": "https://jira.bancvue.com/rest/api/2/issuetype/15",
                            "id": "15",
                            "description": "gh.issue.story.desc",
                            "iconUrl": "https://jira.bancvue.com/secure/viewavatar?size=xsmall&avatarId=14935&avatarType=issuetype",
                            "name": "Story",
                            "subtask": false,
                            "avatarId": 14935
                        }
                    }
                }
            }
        ],
        "assignee": {
            "self": "https://jira.bancvue.com/rest/api/2/user?username=juan.contreras",
            "name": "juan.contreras",
            "key": "juan.contreras",
            "emailAddress": "juan.contreras@kasasa.com",
            "avatarUrls": {
                "48x48": "https://jira.bancvue.com/secure/useravatar?ownerId=juan.contreras&avatarId=15947",
                "24x24": "https://jira.bancvue.com/secure/useravatar?size=small&ownerId=juan.contreras&avatarId=15947",
                "16x16": "https://jira.bancvue.com/secure/useravatar?size=xsmall&ownerId=juan.contreras&avatarId=15947",
                "32x32": "https://jira.bancvue.com/secure/useravatar?size=medium&ownerId=juan.contreras&avatarId=15947"
            },
            "displayName": "Juan Contreras",
            "active": true,
            "timeZone": "America/New_York"
        },
        "customfield_14094": "kloans-loan",
        "components": [],
        "customfield_14087": {
            "self": "https://jira.bancvue.com/rest/api/2/user?username=juan.contreras",
            "name": "juan.contreras",
            "key": "juan.contreras",
            "emailAddress": "juan.contreras@kasasa.com",
            "avatarUrls": {
                "48x48": "https://jira.bancvue.com/secure/useravatar?ownerId=juan.contreras&avatarId=15947",
                "24x24": "https://jira.bancvue.com/secure/useravatar?size=small&ownerId=juan.contreras&avatarId=15947",
                "16x16": "https://jira.bancvue.com/secure/useravatar?size=xsmall&ownerId=juan.contreras&avatarId=15947",
                "32x32": "https://jira.bancvue.com/secure/useravatar?size=medium&ownerId=juan.contreras&avatarId=15947"
            },
            "displayName": "Juan Contreras",
            "active": true,
            "timeZone": "America/New_York"
        },
        "subtasks": [],
        "reporter": {
            "self": "https://jira.bancvue.com/rest/api/2/user?username=juan.contreras",
            "name": "juan.contreras",
            "key": "juan.contreras",
            "emailAddress": "juan.contreras@kasasa.com",
            "avatarUrls": {
                "48x48": "https://jira.bancvue.com/secure/useravatar?ownerId=juan.contreras&avatarId=15947",
                "24x24": "https://jira.bancvue.com/secure/useravatar?size=small&ownerId=juan.contreras&avatarId=15947",
                "16x16": "https://jira.bancvue.com/secure/useravatar?size=xsmall&ownerId=juan.contreras&avatarId=15947",
                "32x32": "https://jira.bancvue.com/secure/useravatar?size=medium&ownerId=juan.contreras&avatarId=15947"
            },
            "displayName": "Juan Contreras",
            "active": true,
            "timeZone": "America/New_York"
        },
        "customfield_10043": "juan.contreras(juan.contreras)",
        "progress": {
            "progress": 0,
            "total": 0
        },
        "votes": {
            "self": "https://jira.bancvue.com/rest/api/2/issue/KL-1379/votes",
            "votes": 0,
            "hasVoted": false
        },
        "worklog": {
            "startAt": 0,
            "maxResults": 20,
            "total": 0,
            "worklogs": []
        },
        "issuetype": {
            "self": "https://jira.bancvue.com/rest/api/2/issuetype/15",
            "id": "15",
            "description": "gh.issue.story.desc",
            "iconUrl": "https://jira.bancvue.com/secure/viewavatar?size=xsmall&avatarId=14935&avatarType=issuetype",
            "name": "Story",
            "subtask": false,
            "avatarId": 14935
        },
		"project": {
            "self": "https://jira.bancvue.com/rest/api/2/project/13593",
            "id": "13593",
            "key": "KL",
            "name": "K-Loans",
            "avatarUrls": {
                "48x48": "https://jira.bancvue.com/secure/projectavatar?pid=13593&avatarId=16124",
                "24x24": "https://jira.bancvue.com/secure/projectavatar?size=small&pid=13593&avatarId=16124",
                "16x16": "https://jira.bancvue.com/secure/projectavatar?size=xsmall&pid=13593&avatarId=16124",
                "32x32": "https://jira.bancvue.com/secure/projectavatar?size=medium&pid=13593&avatarId=16124"
            }
        },
        "resolutiondate": null,
        "watches": {
            "self": "https://jira.bancvue.com/rest/api/2/issue/KL-1379/watchers",
            "watchCount": 1,
            "isWatching": true
        },
        "updated": "2018-02-05T13:30:27.000-0600",
        "description": "Need to determine what the best way to lock the loan is, the process lock that [~andrew.chinn] is working on is one option, however my concern with that is that it may be too slow and it adds a new call to every change to a loan. I think that a database field like locked_at or is_locked would do the job. \r\n\r\nThe code to lock and unlock needs to happen as part of the ReplayProcess process, it needs to lock at the beginning and unlock once it is done.\r\n\r\nAlso we need update every endpoint that accepts changes to the the loan to find out if its locked and return a 409 Conflict if its locked.",
        "summary": "Loan should be locked when the replay process starts.",
        "comment": {
            "comments": [],
            "maxResults": 0,
            "total": 0,
            "startAt": 0
        },
        "priority": {
            "self": "https://jira.bancvue.com/rest/api/2/priority/6",
            "iconUrl": "https://jira.bancvue.com/images/icons/custom/Low.png",
            "name": "4 - Low",
            "id": "6"
        },
        "versions": [],
        "status": {
            "self": "https://jira.bancvue.com/rest/api/2/status/10114",
            "description": "",
            "iconUrl": "https://jira.bancvue.com/images/icons/statuses/inprogress.png",
            "name": "Developing",
            "id": "10114",
            "statusCategory": {
                "self": "https://jira.bancvue.com/rest/api/2/statuscategory/4",
                "id": 4,
                "key": "indeterminate",
                "colorName": "yellow",
                "name": "In Progress"
            }
        },
        "creator": {
            "self": "https://jira.bancvue.com/rest/api/2/user?username=juan.contreras",
            "name": "juan.contreras",
            "key": "juan.contreras",
            "emailAddress": "juan.contreras@kasasa.com",
            "avatarUrls": {
                "48x48": "https://jira.bancvue.com/secure/useravatar?ownerId=juan.contreras&avatarId=15947",
                "24x24": "https://jira.bancvue.com/secure/useravatar?size=small&ownerId=juan.contreras&avatarId=15947",
                "16x16": "https://jira.bancvue.com/secure/useravatar?size=xsmall&ownerId=juan.contreras&avatarId=15947",
                "32x32": "https://jira.bancvue.com/secure/useravatar?size=medium&ownerId=juan.contreras&avatarId=15947"
            },
            "displayName": "Juan Contreras",
            "active": true,
            "timeZone": "America/New_York"
        },
        "created": "2018-02-05T10:38:23.000-0600"

    }
}`

	res := Response{}
	json.Unmarshal([]byte(fakeResponseJson), &res)

	fmt.Println(res)

	assert.Equal(t, "juan.contreras", res.Assignee.Name)
	assert.Equal(t, "Juan Contreras", res.Assignee.DisplayName)
	assert.Equal(t, "juan.contreras", res.Reporter.Name)
	assert.Equal(t, "Juan Contreras", res.Reporter.DisplayName)
	assert.Equal(t, "KL-1379", res.Key)
	assert.Equal(t, "Loan should be locked when the replay process starts.", res.Fields.Summary)
	assert.Equal(t, "Developing", res.Fields.Status.Name)

	//assert.Len(t, res.Jobs, 2)

}
