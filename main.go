package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-ole/go-ole"
	_ "github.com/mattn/go-adodb"
)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	f := "./testdata/mme_code.mdb"
	//provider := "Microsoft.ACE.OLEDB.12.0"
	provider := "Microsoft.Jet.OLEDB.4.0"

	db, err := sql.Open("adodb", "Provider="+provider+";Data Source="+f+";")
	if err != nil {
		fmt.Println("open", err)
		return
	}
	defer db.Close()

	// Main Location
	query := "select" +
		" mme_transducer_main_locations.[type]," +
		" mme_transducer_main_locations.[trans_main_loc]," +
		" mme_transducer_main_locations.[text_l1]," +
		" sys_data_history.dataversion," +
		" mme_transducer_main_locations.[reamrks]" +
		" from sys_data_history inner join mme_transducer_main_locations" +
		" on sys_data_history.[id] = mme_transducer_main_locations.[version]" +
		" order by mme_transducer_main_locations.[type]," +
		" mme_transducer_main_locations.sortkey;"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("select", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var locationType string
		var code string
		var name string
		var version string
		var remarks *string
		err = rows.Scan(&locationType, &code, &name, &version, &remarks)
		if err != nil {
			fmt.Println("scan", err)
			return
		}
		s := strings.ReplaceAll(name, " ", "_")
		s = strings.ReplaceAll(name, ",", "_")
		fmt.Printf("%s,  // %s\n", strings.ToUpper(s), version)
	}
}
