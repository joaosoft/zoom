package zoom

// Option ...
type Option func(client *Client)

// Reconfigure ...
func (c *Client) Reconfigure(options ...Option) {
	for _, option := range options {
		option(c)
	}
}