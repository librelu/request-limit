package trackers

import "fmt"

func getCacheKey(ip string) string {
	return fmt.Sprintf(ipKeyPattern, ip)
}
