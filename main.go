package main

import (
	"deputy/RepositoryApis/roles/linear"
	"deputy/RepositoryApis/roles/roleMap"
	"deputy/RepositoryApis/roles/sortedArray"
	"deputy/RepositoryApis/users"
	"deputy/data/smallRepository"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Use Staff ID as a command line argument")
		return
	}
	userIdString := os.Args[1]
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Println("Cannot parse user id string to int")
		return
	}

	roleId, err := users.GetRoleIdForUserId(userId)
	if err != nil {
		log.Println("Cannot obtain Role Id for this user")
		return
	}

	// linear
	subordinates, err := linear.GetSubordinatesForRoleId(roleId)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(subordinates)

	// roleMap
	r := roleMap.RoleMap{}
	r.InitRoleMapById(smallRepository.Roles)
	subordinates, err = users.GetUsersByRoleIdsFromSimpleRepository(r.GetAllSubordinateIdsById(roleId))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(subordinates)

	// sortedArray
	r2 := sortedArray.RoleSortedArray{}
	r2.InitSortedById(smallRepository.Roles)
	subordinateIds, err := r2.GetAllSubordinateIdsById(roleId)
	if err != nil {
		log.Println(err)
		return
	}
	subordinates, err = users.GetUsersByRoleIdsFromSimpleRepository(subordinateIds)
	fmt.Println(subordinates)

}
