package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	waktuInit := time.Date(2017, 12, 11, 0, 0, 0, 0, t.Location()) //acuan hari untuk mendapatkan hari dari waktuNow
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
			fmt.Println(101)
		case 8:
			fmt.Println(102)
		case 9:
			fmt.Println(103)
		case 10:
			fmt.Println(104)
		case 11:
			fmt.Println(105)
		case 12:
			fmt.Println(106)
		case 13:
			fmt.Println(107)
		case 14:
			fmt.Println(108)
		case 15:
			fmt.Println(109)
		case 16:
			fmt.Println(110)
		case 17:
			fmt.Println(111)
		default:
			fmt.Println(199)
		}
	case 2: //selasa
		switch j {
		case 7:
			fmt.Println(201)
		case 8:
			fmt.Println(202)
		case 9:
			fmt.Println(203)
		case 10:
			fmt.Println(204)
		case 11:
			fmt.Println(205)
		case 12:
			fmt.Println(206)
		case 13:
			fmt.Println(207)
		case 14:
			fmt.Println(208)
		case 15:
			fmt.Println(209)
		case 16:
			fmt.Println(210)
		case 17:
			fmt.Println(211)
		default:
			fmt.Println(299)
		}
	case 3: //rabu
		switch j {
		case 7:
			fmt.Println(301)
		case 8:
			fmt.Println(302)
		case 9:
			fmt.Println(303)
		case 10:
			fmt.Println(304)
		case 11:
			fmt.Println(305)
		case 12:
			fmt.Println(306)
		case 13:
			fmt.Println(307)
		case 14:
			fmt.Println(308)
		case 15:
			fmt.Println(309)
		case 16:
			fmt.Println(310)
		case 17:
			fmt.Println(311)
		default:
			fmt.Println(399)
		}
	case 4: //kamis
		switch j {
		case 7:
			fmt.Println(401)
		case 8:
			fmt.Println(402)
		case 9:
			fmt.Println(403)
		case 10:
			fmt.Println(404)
		case 11:
			fmt.Println(405)
		case 12:
			fmt.Println(406)
		case 13:
			fmt.Println(407)
		case 14:
			fmt.Println(408)
		case 15:
			fmt.Println(409)
		case 16:
			fmt.Println(410)
		case 17:
			fmt.Println(411)
		default:
			fmt.Println(499)
		}
	case 5: //jumat
		switch j {
		case 7:
			fmt.Println(501)
		case 8:
			fmt.Println(502)
		case 9:
			fmt.Println(503)
		case 10:
			fmt.Println(504)
		case 11:
			fmt.Println(505)
		case 12:
			fmt.Println(506)
		case 13:
			fmt.Println(507)
		case 14:
			fmt.Println(508)
		case 15:
			fmt.Println(509)
		case 16:
			fmt.Println(510)
		case 17:
			fmt.Println(511)
		default:
			fmt.Println(599)
		}
	}
}
