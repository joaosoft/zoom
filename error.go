package zoom

import (
	"encoding/json"
	"fmt"
)

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("Zoom API error %d: \"%s\"", e.Code, e.Message)
}

func checkError(body []byte) error {
	var c errorContainer
	if err := json.Unmarshal(body, &c); err != nil {
		return err
	}

	// need to explicitly return nil or it will register as an error
	if c.APIError == nil {
		return nil
	}

	return c.APIError
}
