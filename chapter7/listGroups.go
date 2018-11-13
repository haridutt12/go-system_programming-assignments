package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {

	var u *user.User
	var err error
	arguments := os.Args
	if len(arguments) == 1 {
		u, err = user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		u, err = user.Lookup(arguments[1])
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	groupsid, _ := u.GroupIds()
	for _, gid := range groupsid {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s(%s) ", group.Gid, group.Name)
	}
	fmt.Println()

}
