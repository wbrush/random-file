package process

import (
	"strings"
	"errors"
)

func StripNames(in_names []string) (out_names []string, err error) {

	for _, item := range in_names {
		if (strings.Contains(item, "_")) {
			items := strings.Split(item, "_")
			if (len(items) == 2 && item[0] >= '0' && item[0] <= '9') {
				item = items[1]
			} else if (len(items) > 2 || item[0] < '0' || item[0] > '9') {
				err = errors.New("Extra underscore in this file name: "+item)
				return
			}
		}
		out_names = append(out_names, item)
	}

	return
}