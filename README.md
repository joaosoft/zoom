Zoom
================

[![Build Status](https://travis-ci.org/joaosoft/zoom.svg?branch=master)](https://travis-ci.org/joaosoft/zoom) | [![codecov](https://codecov.io/gh/joaosoft/zoom/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/zoom) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/zoom)](https://goreportcard.com/report/github.com/joaosoft/zoom) | [![GoDoc](https://godoc.org/github.com/joaosoft/zoom?status.svg)](https://godoc.org/github.com/joaosoft/zoom)

A simple client for zoom api.

## With support:

### Meetings
* Create
* Delete

### Users
* Create
* Delete

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependency Management
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/zoom
```

## Usage 
This examples are available in the project at [zoom/examples](https://github.com/joaosoft/zoom/tree/master/examples)

### Code
```go
import (
	"time"
	"github.com/joaosoft/zoom"
)

var (
	ApiKey      = "<API_KEY>"
	ApiSecret   = "<API_SECRET>"
	Enabled     = true
	UserGroupID = "<USER_GROUP_ID"
)

func init() {
	zoom.Init(&zoom.Config{
		ApiKey:    ApiKey,
		ApiSecret: ApiSecret,
		Enabled:   Enabled,
	})
}

func main() {
	var err error

	// create user
	var user *zoom.User
	if user, err = createUser(); err != nil {
		panic(err)
	}

	// associate user to group
	if err = associateUserToGroup(user, UserGroupID); err != nil {
		panic(err)
	}

	// create meeting
	var meeting *zoom.Meeting
	now := time.Now()
	if meeting, err = createMeeting(user, "First meeting", &now); err != nil {
		panic(err)
	}

	// delete meeting
	if err = deleteMeeting(meeting); err != nil {
		panic(err)
	}

	// delete user
	if err = deleteUser(user); err != nil {
		panic(err)
	}
}

func createUser() (*zoom.User, error) {
	return zoom.CreateUser(&zoom.CreateUserOptions{
		UserInfo: zoom.UserInfo{
			FirstName: "João",
			LastName:  "Ribeiro",
			Email:     "joaosoft@gmail.com",
			Type:      zoom.TypeBasic,
		},
		Action: zoom.ActionCustCreate,
	})
}

func associateUserToGroup(user *zoom.User, groupId string) error {
	return zoom.AddUserToGroup(&zoom.AddUserToGroupOptions{
		GroupID: groupId,
		Members: []zoom.Member{
			zoom.Member{
				ID: user.ID,
			},
		},
	})
}

func createMeeting(user *zoom.User, topic string, time *time.Time) (*zoom.Meeting, error) {
	return zoom.CreateMeeting(&zoom.CreateMeetingOptions{
		UserID:    user.ID,
		Type:      zoom.MeetingTypeScheduled,
		Duration:  1,
		StartTime: time,
		Topic:     topic,
		Settings: zoom.MeetingSettings{
			Audio:        zoom.AudioBoth,
			ApprovalType: zoom.ApprovalTypeAutomaticallyApprove,
		},
	})
}

func deleteMeeting(meeting *zoom.Meeting) error {
	return zoom.DeleteMeeting(&zoom.DeleteMeetingOptions{
		MeetingID: meeting.ID,
	})
}

func deleteUser(user *zoom.User) error {
	return zoom.DeleteUser(&zoom.DeleteUserOptions{
		UserID: user.ID,
	})
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
