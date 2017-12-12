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
	id_ruangan  string
	id_time     int
	mata_kuliah string
}

type ruangan struct {
	id_ruangan string
	nama       string
	kapasitas  int
	fasilitas  string
}

type timeslot struct {
	id_time int
	hari    string
	jam     string
}

func main() {
	port := 8181
	http.HandleFunc("/kelas_kosong/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			time := getTimeslot() //511
			if time != 999 {
				GetKelasKosong(w, r, time)
			} else {
				log.Println("Saat ini tidak jam kuliah!")
			}
		//case "POST":
		//case "DELETE":
		default:
			http.Error(w, "Invalid request method", 405)
		}
	})
	log.Printf("Server berjalan pada port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func GetKelasKosong(w http.ResponseWriter, r *http.Request, t int) {
	db, err := sql.Open("mysql", "root:yoyomam@@tcp(127.0.0.1:3306)/habib_labtekv")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `
		select * from ruangan
		where id_ruangan = any (
			select id_ruangan from kelas
			where mata_kuliah = '' and
			id_time = ?);
	`
	ts := strconv.Itoa(t) //convert int to string
	rows, err := db.Query(q, ts)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	ruangan_kosong := ruangan{}

	for rows.Next() {
		err := rows.Scan(&ruangan_kosong.id_ruangan, &ruangan_kosong.nama, &ruangan_kosong.kapasitas, &ruangan_kosong.fasilitas)
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
	//log.Printf("Oi %f", waktuSelisih.Hours()) //buat debug
	//log.Printf("%f", (waktuSelisih.Hours() / 24)) //buat debug
	var i int = (int(waktuSelisih.Hours()/24) % 7) //cari modulus selisih hariInit dan hariNow utk dapat index hari
	//fmt.Println(i) //buat debug
	var j int = (int(t.Hour())) //cari jam sekarang
	switch i {
	case 1: //senin
		switch j {
		case 7:
			return 101
		case 8:
			return 102
		case 9:
			return 103
		case 10:
			return 104
		case 11:
			return 105
		case 12:
			return 106
		case 13:
			return 107
		case 14:
			return 108
		case 15:
			return 109
		case 16:
			return 110
		case 17:
			return 111
		default:
			return 999 //di luar jam kuliah
		}
		switch j {
		case 7:
			return 201
		case 8:
			return 202
		case 9:
			return 203
		case 10:
			return 204
		case 11:
			return 205
		case 12:
			return 206
		case 13:
			return 207
		case 14:
			return 208
		case 15:
			return 209
		case 16:
			return 210
		case 17:
			return 211
		default:
			return 999 //di luar jam kuliah
		}
	case 3: //rabu
		switch j {
		case 7:
			return 301
		case 8:
			return 302
		case 9:
			return 303
		case 10:
			return 304
		case 11:
			return 305
		case 12:
			return 306
		case 13:
			return 307
		case 14:
			return 308
		case 15:
			return 309
		case 16:
			return 310
		case 17:
			return 311
		default:
			return 999 //di luar jam kuliah
		}
	case 4: //kamis
		switch j {
		case 7:
			return 401
		case 8:
			return 402
		case 9:
			return 403
		case 10:
			return 404
		case 11:
			return 405
		case 12:
			return 406
		case 13:
			return 407
		case 14:
			return 408
		case 15:
			return 409
		case 16:
			return 410
		case 17:
			return 411
		default:
			return 999 //di luar jam kuliah
		}
	case 5: //jumat
		switch j {
		case 7:
			return 501
		case 8:
			return 502
		case 9:
			return 503
		case 10:
			return 504
		case 11:
			return 505
		case 12:
			return 506
		case 13:
			return 507
		case 14:
			return 508
		case 15:
			return 509
		case 16:
			return 510
		case 17:
			return 511
		default:
			return 999 //di luar jam kuliah
		}
	default: //sabtu minggu
		return 999
	}
}
