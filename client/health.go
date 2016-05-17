package client

func (c *client) HealthCheck() error {
	return c.get("/api/health", nil)
}
