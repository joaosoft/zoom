package main

import (
	"time"
	"zoom"
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
			Type:      zoom.UserTypeBasic,
		},
		Action: zoom.UserCreateActionCustCreate,
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