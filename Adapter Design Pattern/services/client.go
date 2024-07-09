package services

// client
type Client struct {
}

func (c *Client) ChargeMobile(mob Mobile) {
	mob.ChargeAppleMobile()
}
