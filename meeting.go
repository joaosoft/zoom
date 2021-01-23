package zoom

import (
	"fmt"
	"net/http"
)

func CreateMeeting(opts *CreateMeetingOptions) (*Meeting, error) {
	var ret = &Meeting{}
	return ret, defaultClient.request(&requestOptions{
		Method:         http.MethodPost,
		Path:           fmt.Sprintf(createMeetingPath, opts.UserID),
		DataParameters: &opts,
		Ret:            ret,
	})
}

func DeleteMeeting(opts *DeleteMeetingOptions) error {
	return defaultClient.request(&requestOptions{
		Method:        http.MethodDelete,
		Path:          fmt.Sprintf(deleteMeetingPath, opts.MeetingID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
