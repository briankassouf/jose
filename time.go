package jose

import "time"

// Now returns the current time in UTC.
var Now = func() time.Time { return time.Now().UTC() }
