package quarz

import "time"

var Ticker = time.NewTicker(time.Second / CPUSpeed)
