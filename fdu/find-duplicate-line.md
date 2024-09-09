Fungsi `Dup1` dalam kode Go ini melakukan hal berikut:

1. `counts := make(map[string]int)`: Membuat sebuah map di mana kunci adalah string (baris dari input) dan nilai adalah integer (jumlah kemunculan setiap baris).

2. `input := bufio.NewScanner(os.Stdin)`: Membuat sebuah scanner baru yang membaca dari input standar.

3. Loop `for input.Scan()` membaca token berikutnya (dalam hal ini, baris input berikutnya) dari input. Jika token berikutnya tersedia, `Scan` mengembalikan `true` dan loop berlanjut; jika tidak, mengembalikan `false` dan loop berakhir.

4. Di dalam loop, `counts[input.Text()]++` menambahkan jumlah untuk baris saat ini. `input.Text()` mengembalikan token terbaru yang dihasilkan oleh panggilan ke `Scan`, yang merupakan baris input saat ini.

5. Setelah semua baris telah dibaca dan dihitung, loop lain `for line, n := range counts` mengulangi semua entri dalam map `counts`. Untuk setiap entri, jika jumlah `n` lebih besar dari 1, mencetak jumlah dan baris menggunakan `fmt.Printf`.

Fungsi `fmt.Printf` memformat dan mencetak argumennya sesuai dengan spesifikasi format yang diberikan. Dalam hal ini, `%d\t%s\n` adalah spesifikasi format, di mana `%d` adalah placeholder untuk angka desimal (jumlah), `\t` memasukkan tab, `%s` adalah placeholder untuk string (baris), dan `\n` memulai baris baru.

Secara keseluruhan, fungsi ini membaca baris dari input standar, menghitung berapa kali setiap baris muncul, dan mencetak baris yang muncul lebih dari sekali, bersama dengan jumlah kemunculannya.

## Dup2

Fungsi `Dup2()` dalam kode Go ini bertujuan untuk menghitung jumlah baris duplikat dalam satu set file. Berikut adalah penjelasannya dalam Bahasa Indonesia:

Fungsi `Dup2()` dimulai dengan membuat peta kosong `counts` di mana setiap kunci adalah string (mewakili baris teks) dan setiap nilai adalah integer (mewakili jumlah baris tersebut). Kemudian, fungsi ini mengambil daftar nama file dari argumen baris perintah.

Jika tidak ada nama file yang disediakan, fungsi ini akan mencetak pesan meminta nama file. Jika nama file disediakan, fungsi ini akan melakukan iterasi atas setiap nama file, mencoba membuka file, dan jika berhasil, memanggil fungsi `countLines()` untuk memperbarui peta `counts` dengan data dari file tersebut. Jika terjadi kesalahan saat membuka file, fungsi ini akan mencetak kesalahan ke output kesalahan standar dan melanjutkan dengan file berikutnya.

Fungsi `countLines()` menerima file dan peta sebagai argumen. Fungsi ini membuat pemindai baru untuk file, kemudian melakukan iterasi atas setiap baris dalam file. Untuk setiap baris, fungsi ini menambahkan entri yang sesuai di peta.

Setelah semua file diproses, `Dup2()` melakukan iterasi atas peta `counts`. Untuk setiap baris yang muncul lebih dari sekali di semua file, fungsi ini mencetak baris dan jumlahnya.
