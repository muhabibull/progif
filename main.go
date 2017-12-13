package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type kelas struct {
	Id_ruangan  string
	Id_time     int
	Mata_kuliah string
}

type ruangan struct {
	Id_ruangan string
	Nama       string
	Kapasitas  int
	Fasilitas  string
}

type timeslot struct {
	Id_time int
	Hari    string
	Jam     string
}

func main() {
	port := 8181
	http.HandleFunc("/kelas_kosong/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			time := getTimeslot()
			if time != 999 {
				GetHariJamNow(w, r, time)
				GetKelasKosong(w, r, time)
			} else {
				log.Println("Saat ini tidak ada jam kuliah!")
			}
		default:
			http.Error(w, "Invalid request method", 405)
		}
	})
	log.Printf("Server berjalan pada port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func GetHariJamNow(w http.ResponseWriter, r *http.Request, t int) {
	db, err := sql.Open("mysql",
		"root:@tcp(167.205.67.251:3306)/habib_labtekv")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hariJamNow := timeslot{}

	ts := strconv.Itoa(t) //convert int to string
	rows, err := db.Query("select hari, jam from timeslot where id_time = ?;", ts)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&hariJamNow.Hari, &hariJamNow.Jam)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&hariJamNow)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetKelasKosong(w http.ResponseWriter, r *http.Request, t int) {
	db, err := sql.Open("mysql",
		"root:@tcp(167.205.67.251:3306)/habib_labtekv")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ruangan_kosong := ruangan{}

	ts := strconv.Itoa(t) //convert int to string
	rows, err := db.Query("select * from ruangan where id_ruangan = any ( select id_ruangan from kelas where mata_kuliah = '' and id_time = ?);", ts)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ruangan_kosong.Id_ruangan, &ruangan_kosong.Nama, &ruangan_kosong.Kapasitas, &ruangan_kosong.Fasilitas)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&ruangan_kosong)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func getTimeslot() int {
	t := time.Now()
	waktuInit := time.Date(2017, 12, 10, 0, 0, 0, 0, t.Location()) //acuan hari untuk mendapatkan nama hari dari waktuNow
	waktuNow := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	var waktuSelisih time.Duration = waktuNow.Sub(waktuInit) //hitung selisih hariInit dan hariNow
	var i int = (int(waktuSelisih.Hours()/24) % 7)           //cari modulus selisih hariInit dan hariNow utk dapat index hari
	var j int = (int(t.Hour()))                              //cari jam sekarang
	switch i {
	case 1: //senin
		if (j > 6) && (j < 18) {
			return j + 94
		} else {
			return 999
		}
	case 2: //selasa
		if (j > 6) && (j < 18) {
			return j + 194
		} else {
			return 999
		}
	case 3: //rabu
		if (j > 6) && (j < 18) {
			return j + 294
		} else {
			return 999
		}
	case 4: //kamis
		if (j > 6) && (j < 18) {
			return j + 394
		} else {
			return 999
		}
	case 5: //jumat
		if (j > 6) && (j < 18) {
			return j + 494
		} else {
			return 999
		}
	default: //sabtu minggu
		return 999
	}
}
