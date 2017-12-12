package main

import(
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

	type Tim struct {
		rank int
		himpunan string
		tanding int 
		menang int
		seri int
		kalah int
		gm int
		gk int
		sg int
		poin int
	}
	
	
	//GET seluruh klasemen (sorted by poin/rank) (GetKlasemen)

	//GET sorted by GM (jumlah gol memasukkan) (GetSortGM)
	
	//GET sorted by GK (jumlah gol kemasukan) (GetSortGK)
	
	//GET detail spesifik suatu tim bola himpunan (GetHimpunan)
	
func main(){
	port:=8181
	http.HandleFunc("/klasemen", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				s:=r.URL.Query().Get("himpunan")
				if (s!="") {
					GetHimpunan(w,r,s)
				} else {
					GetKlasemen(w,r)
				}
			default:
				http.Error(w,"invalid",405)
		}
	})
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

//GetKlasemen
func GetKlasemen(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/ligaITB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}

	rows, err:=db.Query("select *from klasemen")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.rank, &team.himpunan, &team.tanding, &team.menang, &team.seri, &team.kalah, &team.gm, &team.gk, &team.sg, &team.poin)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&team)
	}
	err=rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//GetHimpunan
func GetHimpunan (w http.ResponseWriter, r *http.Request, hmj string) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/ligaITB")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}
	rows, err:=db.Query("select * from klasemen where himpunan like?", hmj)
    if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.rank, &team.himpunan, &team.tanding, &team.menang, &team.seri, &team.kalah, &team.gm, &team.gk, &team.sg, &team.poin)
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err=rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

	
	
