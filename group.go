package zoom

import (
	"fmt"
	"net/http"
)

func AddUserToGroup(opts *AddUserToGroupOptions) error {
	return defaultClient.request(&requestOpts{
		Method:         http.MethodPost,
		Path:           fmt.Sprintf(addUserToGroupPath, opts.GroupID),
		DataParameters: &opts,
		HeadResponse:   true,
	})
}
