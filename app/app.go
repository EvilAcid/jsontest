package app

import (
	"fmt"
	compliance2 "jsontest/models/compliance"
	"jsontest/utils"
	"log"
)

func Start() {
	// Create all Databases and member map
	/*
		ComplianceDB
		membersDB
		TeamDB
		DashboardDB
		DetectorDB
	*/
	var CDB = compliance2.ComplianceDB{}
	//var members map[string]string
	//.....
	//----------------------------------------------------------

	//Load all Databases
	/*
		Load ComplianceDB from file
		fill all other DBs with the given data from the API
	*/
	if err := CDB.Load(); err != nil {
		log.Fatalln(err)
	}
	//.....
	//----------------------------------------------------------

	// Fill up members map
	/*
		fill all members into map (for loop)
	*/

	//.....
	//----------------------------------------------------------

	// Fill up members map
	/*
		fill all members into map (for loop)
	*/

	//.....
	//----------------------------------------------------------

	//
	// Below is only for Testing
	//

	//utils.MakeData(&CDB, 2)
	/*
		updateDate := compliance2.Dashboards{
			ID:         "xyz",
			Name:       "Dashboardname",
			Deletetime: 120,
		}

	*/

	//compliance.AddDashboard(&CDB, "1", &updateDate)

	if err := CDB.Save(); err != nil {
		log.Fatalln(err)
	}

	test, err := utils.BuildEmail(CDB.Users[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(test)

}
