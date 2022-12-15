# Microservices Product

## Deskripsi

```markdown
Program sederhana ini merupakan sample REST Api menggunakan bahasa pemrograman Golang
```

### Fitur

> - Authentication
>   - Register User
>     - Input data user untuk digunakan di fungsi login
>   - Login User
>     - Saat login mendapatkan informasi **Token**
> - Products
>   - Create Product
>     - Membuat data product terdiri dari (id_product, name, stock, price)
>   - Update Product
>     - Memperbarui data product terdiri dari (name, stock, price)
>   - Delete Product
>     - Menghapus data product berdasarkan id product
>   - Get Product All
>     - Menampilkan keseluruhan data product
>   - Get Product By ID Product
>     - Hanya menampilkan product yang di cari
> - Orders
>   - Order Cart
>     - Membuat order product dengan status masih dalam **cart** dan **Stock belum berkurang jika masih dalam cart**
>     - Update **Qty** product, jika ada product yang sudah pernah di order
>     - Membuat order product baru, jika ada product yang sudah pernah di order dalam status **cancel** dan **checkout**
>   - Order Checkout
>     - Update status order menjadi **checkout** dan **Stock** berkurang atau diperbarui di tabel *products*
>     - Tidak bisa melakukan **checkout** jika status order **cancel**
>   - Order Cancel
>     - Update **Stock** pada tabel product atau rollback **Stock** dan update status order menjadi **cancel**
>     - Dibuat kondisi supaya jika di lakukan ***Order Cancel*** tidak dapat melakukan ***Hit Request***

### Running Aplikasi

1. Buka **folder project micro_product** lalu menuju terminal atau pada direktori ***`./micro_product`***
2. Masukkan perintah berikut pada terminal

```markdown
go run main.go
```

### Test Menggunakan Postman

> - Menuju [**Link Berikut**](./micro_product.postman_collection.json) untuk melakukan unduh file atau pada area **code** anda dapat melakukan ***CTRL + S*** atau klik simpan sebagai **file** dengan **nama_file.json**
> - Buka aplikasi postman lalu import ***file json*** pada aplikasi Postman
> - Terdapat **collection request** dan lakukan tes pada setiap **Endpoint**
> - ***Alternatif*** membuka **collection request** menggunakan **Postman Public Api Network** [**Link Berikut**](https://postman.com/arsetsoft/workspace/go-microservices/documentation/18056562-47a0fed6-ecb7-4241-925e-ce6e49b881ce)

### Test Menggunakan Swagger

> - Buka [**Link Berikut**](http://localhost:8080/swagger/index.html) pada browser ***Chrome atau Mozila***
> - Halaman **swagger** akan terbuka dan Lakukan tes pada setiap Endpoint
> - ***Catatan*** sebelum melakukan test menggunakan swagger, pastikan aplikasi sudah **Running**

### Tentang Saya

> | Nama          | Kota              |
> | ------------- | ----------------- |
> | Arif Setiawan | Bandar Lampung    |
