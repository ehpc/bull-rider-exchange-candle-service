package transport

import (
	"fmt"
	"sort"
)

// RequestParams are request parameters for Receive
type RequestParams map[string]string

// Hash returns hash representation of RequestParams
func (rp *RequestParams) Hash() string {
	hash := ""
	rpMap := map[string]string(*rp)
	keys := make([]string, len(rpMap))
	i := 0
	for k := range rpMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		hash += fmt.Sprintf("%s:%s#", k, rpMap[k])
	}
	if len(hash) == 0 {
		return hash
	}
	return hash[:len(hash)-1]
}
