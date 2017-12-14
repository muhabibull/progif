# progif
Tugas Pemrograman Integratif 2017

Web service ini merupakan program pencari ruangan di Labtek V yang tidak dipakai untuk kelas mata kuliah.

Seluruh code terdapat di file main.go

Database kelas terdapat di folder data.

Untuk mencoba web service akses pada web browser halaman 167.205.67.227:8181/kelas_kosong (dengan internet ITB atau menggunakan VPN)

Untuk mencoba di local computer :
 - import data kelas_labtekv.sql ke database MySQL
 - ubah alamat database pada program main.go pada kedua fungsi Get : {username}:{password}@tcp({localhost}:{port})/{nama database}
 - jalankan program Anda dan akses pada web browser halaman localhost:{port}/kelas_kosong
 
 
