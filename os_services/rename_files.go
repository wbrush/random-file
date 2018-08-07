package os_services

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

func RenameFiles(path string, orig, mid, names []string, wflag bool) {

	for i := range orig {

		if (!strings.Contains(orig[i], "_")) { continue }

		oname := path+string(filepath.Separator)+orig[i]
		mname := path+string(filepath.Separator)+mid[i]

		if (wflag) {
			err := os.Rename( oname, mname )
			if (err != nil) {
				fmt.Println("Error writing ", mid[i], ": ", err.Error())
			}
		} else {
			fmt.Println("Renaming ", oname, " to ", mname)
		}
	}

	for i := range mid {

		mname := path+string(filepath.Separator)+mid[i]
		nname := path+string(filepath.Separator)+names[i]

		if (wflag) {
			err := os.Rename( mname, nname )
			if (err != nil) {
				fmt.Println("Error writing ", mid[i], ": ", err.Error())
			}
		} else {
			fmt.Println("Renaming ", mname, " to ", nname)
		}
	}

	return
}