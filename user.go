package zoom

import (
	"fmt"
	"net/http"
)

func GetUser(opts *GetUserOpts) (*User, error) {
	var ret = &User{}
	return ret, defaultClient.request(&requestOpts{
		Method:        http.MethodGet,
		Path:          fmt.Sprintf(getUserPath, opts.EmailOrID),
		URLParameters: opts,
		Ret:           ret,
	})
}

func CreateUser(opts *CreateUserOptions) (*User, error) {
	var ret = &User{}
	return ret, defaultClient.request(&requestOpts{
		Method:         http.MethodPost,
		Path:           createUserPath,
		DataParameters: &opts,
		Ret:            ret,
	})
}

func DeleteUser(opts *DeleteUserOptions) error {
	return defaultClient.request(&requestOpts{
		Method:        http.MethodDelete,
		Path:          fmt.Sprintf(deleteUserPath, opts.UserID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
