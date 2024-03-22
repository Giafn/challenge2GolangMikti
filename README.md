
# Tugas Challenge 2 Mikti Golang

Ini adalah repositori untuk tugas Golang yang saya kerjakan. Repositori ini berisi kode sumber untuk penyelesaian soal dari Challenge 2 Mikti golang yang diberikan.


## Deskripsi Tugas

### Soal 1
Ada dua tim senam, Tim Lumba-lumba dan Tim Koala. Mereka bertanding satu sama lain sebanyak 3 kali. Pemenang dengan skor rata-rata tertinggi memenangkan trofi!

**Soal:**
Hitung skor rata-rata untuk setiap tim, menggunakan data uji di bawah ini.
Bandingkan skor rata-rata tim untuk menentukan pemenang kompetisi, dan cetak ke konsol. Jangan lupa bahwa bisa ada hasil seri, jadi uji juga untuk itu (seri berarti mereka memiliki skor rata-rata yang sama).

**Bonus 1:** Sertakan persyaratan untuk skor minimum 100. Dengan aturan ini, sebuah tim hanya menang jika memiliki skor lebih tinggi dari tim lain, dan pada saat yang sama skor minimal 100 poin.

**Bonus 2:** Skor minimum juga berlaku untuk seri! Jadi seri hanya terjadi ketika kedua tim memiliki skor yang sama dan keduanya memiliki skor lebih besar atau sama dengan 100 poin. Jika tidak, tidak ada tim yang memenangkan trofi.

**Data Uji:**
- Data 1: 
  - Lumba-lumba mendapat skor 96, 108, dan 89. 
  - Koala mendapat skor 88, 91, dan 110.
- Data Bonus 1: 
  - Lumba-lumba mendapat skor 97, 112, dan 101. 
  - Koala mendapat skor 109, 95, dan 123.
- Data Bonus 2: 
  - Lumba-lumba mendapat skor 97, 112, dan 101. 
  - Koala mendapat skor 109, 95, dan 106

### Soal 2
Mark dan John sedang mencoba membandingkan BMI (Body Mass Index) mereka, yang dihitung menggunakan rumus: BMI = massa / tinggi^2 = massa / (tinggi * tinggi) (massa dalam kg dan tinggi dalam meter).

**Soal:**
Simpan massa dan tinggi badan Mark dan John dalam variabel.
Hitung kedua BMI mereka menggunakan rumus (kamu bahkan dapat menerapkan kedua versi).
Buat variabel Boolean 'markHigherBMI' yang berisi informasi tentang apakah Mark memiliki BMI lebih tinggi dari John.

**Data Uji:**
- Data 1: 
  - Berat Mark 78 kg dan tinggi 1.69 m. 
  - Berat John 92 kg dan tinggi 1.95 m.
- Data 2: 
  - Berat Mark 95 kg dan tinggi 1.88 m. 
  - Berat John 85 kg dan tinggi 1.76 m.
## Struktur Proyek

Proyek ini terdiri dari beberapa file dan folder, dengan struktur berikut:

- `main.go`: File utama dari proyek yang berisi fungsi `main()` untuk menjalankan program.
- `go.mod`: File konfigurasi golang module.
- `soal/`: Folder yang berisi file-file jawaban yang saya buat
  - `soal_satu.go`: File yang berisi penyelesaian untuk soal satu.
  - `soal_dua.go`: File yang berisi penyelesaian untuk soal dua.
## Cara Penggunaan

Anda dapat menggunakan proyek ini dengan mengikuti langkah-langkah berikut:

##### 1. Pastikan Anda memiliki Golang yang terinstal di komputer Anda.
    
##### 2. Clone repositori ini ke komputer Anda dengan.
    >`git clone https://github.com/Giafn/challenge2GolangMikti.git`
##### 3. Jalankan perintah `go run main.go` di dalam direktori proyek untuk menjalankan program.

## Demo

![](https://github.com/Giafn/challenge2GolangMikti/blob/main/gif/run_code.gif)
