# Skilltest RSP 15 April 2021
## RENTAL BUKU
---

### Pengerjaan project alokasi 3 hari 
<hr><br>

**Berikut ini ERD yang kami sediakan : `https://dbdiagram.io/d/6077bcccb6aeb3052d902f8e`**

#### Deskripsi MVP project

1. Terdapat 2 role pengguna yaitu :
    * **Admin**
       * Admin dapat melakukan CRUD pada semua data 
       * Berlaku sebagai superadmin
    * **Member**
       *  Member dapat melihat daftar buku
       *  Member dapat melihat data transaksi dirinya

2. Untuk user member akan memiliki fitur register, login, reset password, verifikasi email dan logout

3. Implementasi Auth menggunakan JWT (Json Web Token)

4. Pada Apps terdapat menu :
    * **Catalog** : Menampilkan semua daftar buku
    * **Category** : Menampilkan daftar buku berdasarkan category
    * **Newest** : Menampilkan daftar 20 buku berdasarkan tahun terbit dalam kurun waktu 2 tahun terakhir

5. Aturan peminjaman:
   * User yang mengembalikan buku melebihi `due date` maka akan dikenakan denda sebanyak 5000 per hari nya.
   * User hanya bisa meminjam buku maksimal 3 dalam satu tanggal yang sama.

### Technology
1. Gunakan Framework, nodejs - express , golang - gin / echo
2. Gunakan ORM, nodejs - sequelize, golang - GORM.
   - Gunakan fitur migrations dan seeder.
