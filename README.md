Nama  : D'sharlendita Febianda Aurelia
NIM   : 2311102069
Kelas : IF - 11 - 03

# Modul 2 Review Struktur Kontrol

Modul ini membahas dasar-dasar pengendalian alur program yang sangat penting dalam membangun logika program yang kompleks. Struktur kontrol memungkinkan kita untuk menentukan bagaimana suatu program merespons kondisi tertentu dan bagaimana program harus melanjutkan eksekusinya berdasarkan kondisi tersebut.

## Code Explanation

Soal Latihan Modul 2A

1.
Program ini pada dasarnya melakukan penukaran atau rotasi terhadap nilai-nilai string yang diberikan oleh pengguna. Jika diperhatikan, urutan string setelah penukaran adalah:
Nilai dari satu berpindah ke dua,
Nilai dari dua berpindah ke tiga,
Nilai dari tiga berpindah ke satu.

2.
Program ini merupakan program yang digunakan untuk menentukan apakah suatu tahun merupakan tahun kabisat atau bukan. Program ini menerima input berupa tahun dari pengguna, kemudian mengevaluasi apakah tahun tersebut memenuhi kriteria tahun kabisat menggunakan aturan yang sudah ditentukan, dan akhirnya menampilkan hasilnya.
Program ini melakukan pengecekan tahun kabisat berdasarkan aturan umum:
Tahun kabisat adalah tahun yang habis dibagi 400, atau tahun yang habis dibagi 4 tetapi tidak habis dibagi 100.
Setelah menentukan apakah tahun yang dimasukkan memenuhi syarat kabisat, program akan menampilkan hasil apakah tahun tersebut kabisat atau bukan kabisat.

3.
Program ini bertujuan untuk menghitung volume dan luas permukaan bola berdasarkan jari-jari yang dimasukkan oleh pengguna. Kemudian, berdasarkan input tersebut, program menghitung:
Volume bola, dengan rumus 4/3 * phi * r * r * r
Luas permukaan bola, dengan rumus 4 * phi * r * r

4.
Program ini merupakan program konversi suhu. Program akan menerima masukan suhu dalam satuan Celcius dari pengguna yang akan digunakan untuk menghitung serta menampilkan suhu tersebut dalam tiga satuan suhu lainnya: Fahrenheit, Reamur, dan Kelvin. Proses perhitungannya adalah sebagai berikut:
Fahrenheit dihitung dengan rumus F = 9/5 * C + 32
Reamur dihitung dengan rumus R = 4/5 * C
Kelvin dihitung dengan rumus K = 5/9 * ( F + 459.67)
Kemudian nilai Kelvin dibulatkan ke bilangan bulat dengan menggunakan int()

5.
Program ini menerima input dari user yang berupa lima angka dan tiga karakter. Kemudian lima angka yang dimasukkan akan diperlakukan sebagai kode ASCII dan diubah menjadi karakter ASCII yang sesuai. Misalnya, angka 65 diubah menjadi karakter 'A'. Lalu Tiga huruf yang dimasukkan oleh pengguna diubah menjadi karakter ASCII berikutnya (karakter dengan nilai ASCII satu lebih tinggi). Contoh, jika 'a' dimasukkan, program akan mencetak 'b', karena 'a' memiliki nilai ASCII 97 dan 'b' memiliki nilai 98.

Soal Latihan Modul 2B

1.
Program ini adalah program  yang mengecek apakah input warna yang diberikan oleh pengguna pada setiap percobaan telah sesuai dengan urutan yang ditentukan ("merah", "kuning", "hijau", dan "ungu"). Pengguna diminta memasukkan empat warna pada setiap percobaan, dan program melakukan lima kali percobaan. Jika semua percobaan berhasil memasukkan urutan warna yang benar, maka program akan mencetak "BERHASIL: true". Jika ada salah satu percobaan yang tidak sesuai, maka program akan mencetak "BERHASIL: false".

2.
Program ini menerima input dari pengguna berupa nama-nama bunga dan terus berulang (looping) hingga pengguna memasukkan kata "SELESAI" untuk menghentikan input. Program kemudian akan menampilkan pita yang berisi daftar bunga yang dimasukkan dan dipisahkan oleh tanda " - ", program juga menampilkan jumlah total bunga yang dimasukkan sebelum pengguna mengetik "SELESAI".

3.
Program ini berfungsi untuk mengecek apakah beban pada kedua kantong terpal yang dibawa oleh sepeda motor Pak Andi seimbang atau tidak. Program akan menerima input berupa berat belanjaan di kedua kantong, yaitu kantong terpal kanan dan kantong terpal kiri. Jika selisih berat antara kedua kantong lebih dari atau sama dengan 9 kg, program akan menyatakan bahwa "Sepeda motor pak Andi akan oleng: true". Sebaliknya, jika selisih kurang dari 9 kg, program akan menyatakan "Sepeda motor pak Andi akan oleng: false". Program akan berhenti jika jumlah total berat dari kedua kantong melebihi batas maksimum 150 kg, yang dianggap tidak aman untuk sepeda motor.

4.
Program ini menghitung nilai perkiraan akar kuadrat dari 2 dengan menggunakan deret tak hingga atau aproksimasi. Pengguna diminta untuk memasukkan nilai K, yaitu jumlah iterasi yang digunakan untuk memperbaiki akurasi hasil. Semakin besar nilai K, semakin akurat hasil perkiraan nilai akar kuadrat dari 2.
Program menggunakan formula deret untuk menghitung akar 2 dengan pola perhitungan yang berulang-ulang (iteratif), di mana setiap iterasi menambahkan faktor pengali pada nilai awal perkiraan akar 2.

Soal Latihan Modul 2C

1.
Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan beratnya dalam gram. Program akan menerima input berupa berat parsel dalam gram, kemudian program menghitung biaya pengiriman berdasarkan berat kilogram (kg) dan sisa gram di bawah 1 kg. Biaya pengiriman dihitung dengan aturan tarif tertentu tergantung pada berat keseluruhan dan sisa gram. Aturan penghitungan biaya adalah sebagai berikut: 
Biaya dasar: Rp 10.000 per kilogram. 
Jika berat lebih dari atau sama dengan 10 kg, tidak ada biaya tambahan untuk sisa gram. 
Jika berat kurang dari 10 kg: Jika sisa gram lebih dari atau sama dengan 500 gram, maka biaya tambahan adalah Rp 5 per gram. 
Jika sisa gram kurang dari 500 gram, maka biaya tambahan adalah Rp 15 per gram. 
Setelah melakukan perhitungan, program menampilkan rincian berat dalam kilogram dan gram, serta rincian biaya dan total biaya pengiriman parsel. 

2.
Program ini bertujuan untuk mengkonversi nilai angka dari mata kuliah menjadi nilai huruf berdasarkan skala tertentu.
Pertanyaan:

a. Jika nam diberikan nilai 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
jawab: Program ini tidak dirancang untuk menangani nilai float sebagai input, dan oleh karena itu, program tidak akan dieksekusi sesuai spesifikasi soal.

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program yang seharusnya!
jawab: 1. Variabel nam dideklarasikan sebagai float64, tetapi dalam blok if, nam diubah menjadi string. Ini akan menyebabkan error tipe data.
       2. Nilai huruf seharusnya disimpan dalam variabel nmk, bukan nam.
       3. Setiap pernyataan if berdiri sendiri, sehingga nilai yang lebih kecil terus menimpa nilai sebelumnya.

c. Perbaiki Program!
jawab: ''' go
      package main

      import "fmt"

      func main() {
      	var nam float64
      	var nmk string
      
      	fmt.Print("Nilai akhir mata kuliah: ")
      	fmt.Scanln(&nam)
      
      	if nam > 80 {
      		nmk = "A"
      	} else if nam > 72.5 {
      		nmk = "AB"
      	} else if nam > 65 {
      		nmk = "B"
      	} else if nam > 57.5 {
      		nmk = "BC"
      	} else if nam > 50 {
      		nmk = "C"
      	} else if nam > 40 {
      		nmk = "D"
      	} else {
      		nmk = "E"
      	}
      
      	fmt.Println("Nilai mata kuliah: ", nmk)
      }
    '''

3. 
Program ini melakukan dua operasi, operasi pertama menentukan faktor-faktor dari bilangan bulat positif. Program akan meminta pengguna untuk memasukkan bilangan bulat positif lebih dari 1, kemudian mencetak semua faktor dari bilangan tersebut. Operasi kedua program memeriksa apakah bilangan tersebut adalah bilangan prima atau bukan. Sebuah bilangan dikatakan prima jika hanya memiliki dua faktor, yaitu 1 dan dirinya sendiri.







