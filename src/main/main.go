package main

import(
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	_ "github.com/go-sql-driver/mysql"
)

	type Tim struct {
		Rank int
		Himpunan string
		Tanding int 
		Menang int
		Seri int
		Kalah int
		GM int
		GK int
		SG int
		Poin int
	}
	
	

	
	
func main() {
        
        port := 8181

        http.HandleFunc("/klasemen/", makeHandler(viewHandler))

        log.Printf("Server starting on port %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
        
  }

  func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
					m := validPath.FindStringSubmatch(r.URL.Path)
					if m == nil {
							http.NotFound(w, r)
							return
					}
					fn(w, r, m[2])
			}
  }

  var validPath = regexp.MustCompile("^/(klasemen)/([a-zA-Z0-9]+)$")

 func viewHandler(w http.ResponseWriter, r *http.Request, opsi string) {
    
        if (opsi == "all") {
				GetKlasemen(w,r)           
        } else if (opsi == "sortbyGM") {
                GetAllSortedByGM(w,r)
        } else if (opsi == "sortbyGK") {
                GetAllSortedByGK(w,r)
        } else if (opsi == "sortbySG") {
                GetAllSortedBySG(w,r)
        } else {
                GetHimpunan(w,r,opsi)
        }
  }

//GetKlasemen 
//Mengirimkan keluaran Klasemen himpunan peserta Liga, diurutkan sesuai Ranking tertinggi, dari Poin terbesar
func GetKlasemen(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/bolaitb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}

	rows, err:=db.Query("SELECT * FROM liga ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.Rank, &team.Himpunan, &team.Tanding, &team.Menang, &team.Seri, &team.Kalah, &team.GM, &team.GK, &team.SG, &team.Poin)
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
//Mengirimkan data hasil liga suatu tim bola HMJ sesuai masukan
func GetHimpunan (w http.ResponseWriter, r *http.Request, hmj string) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/bolaitb")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	team := Tim{}
    rows, err := db.Query("SELECT * FROM liga WHERE Himpunan LIKE?", hmj)
    if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.Rank, &team.Himpunan, &team.Tanding, &team.Menang, &team.Seri, &team.Kalah, &team.GM, &team.GK, &team.SG, &team.Poin)
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

//GetAllSortedByGM (Gol Memasukkan Terbanyak)
//Menampilkan data klasemen seluruh tim, diurutkan dari yang memasukkan gol dengan jumlah terbanyak
func GetAllSortedByGM(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/bolaitb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}

	rows, err:=db.Query("SELECT * FROM liga ORDER BY GM DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.Rank, &team.Himpunan, &team.Tanding, &team.Menang, &team.Seri, &team.Kalah, &team.GM, &team.GK, &team.SG, &team.Poin)
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

//GetAllSortedByGK (Gol Kemasukan Terbanyak)
//Menampilkan data klasemen seluruh tim, diurutkan dari yang kemasukan/kebobolan gol dengan jumlah terbanyak
func GetAllSortedByGK(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/bolaitb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}

	rows, err:=db.Query("SELECT * FROM liga ORDER BY GK DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.Rank, &team.Himpunan, &team.Tanding, &team.Menang, &team.Seri, &team.Kalah, &team.GM, &team.GK, &team.SG, &team.Poin)
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

//GetAllSortedBySG (Selisih gol terbesar)
//Menampilkan data klasemen seluruh tim, diurutkan dari yang selisih memasukkan dan kemasukan gol terbesar
func GetAllSortedBySG(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
			"root:@tcp(127.0.0.1:3306)/bolaitb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	team := Tim{}

	rows, err:=db.Query("SELECT * FROM liga ORDER BY SG DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err:= rows.Scan(&team.Rank, &team.Himpunan, &team.Tanding, &team.Menang, &team.Seri, &team.Kalah, &team.GM, &team.GK, &team.SG, &team.Poin)
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

	
	
