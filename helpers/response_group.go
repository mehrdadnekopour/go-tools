package helpers

import "github.com/liip/sheriff"

// GroupMarshal ...
func GroupMarshal(groups []string, obj interface{}) (interface{}, error) {
	o := &sheriff.Options{
		Groups: groups,
	}

	response, err := sheriff.Marshal(o, obj)

	return response, err
}
