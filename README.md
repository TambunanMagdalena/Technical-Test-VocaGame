# Golang_Template
# Parking Lot CLI

Aplikasi ini adalah simulasi sistem parkir sederhana berbasis CLI menggunakan Golang.

## Cara Menjalankan

Jalankan perintah berikut dari folder `cmd`:

```sh
go run main.go <command> [args...]
```

## Perintah CLI

### 1. Membuat Parking Lot

```sh
go run main.go create_parking_lot <jumlah_slot>
```

Contoh:

```
go run main.go create_parking_lot 6
```

Output:

```
Created a parking lot with 6 slots
```

### 2. Parkir Mobil

```sh
go run main.go park <nomor_polisi> <warna>
```

Contoh:

```
go run main.go park B1234XYZ White
```

Output:

```
Allocated slot number: 1
```

Jika parking lot belum dibuat:

```
parking lot not created
```

### 3. Melihat Status Parking Lot

```sh
go run main.go status
```

Output:

```
Slot No.        Registration No.        Colour
1               B1234XYZ                White
```

### 4. Mengeluarkan Mobil

```sh
go run main.go leave <nomor_polisi> <jam_parkir>
```

Contoh:

```
go run main.go leave B1234XYZ 4
```

Output:

```
Registration number B1234XYZ with Slot Number 1 is free with Charge 30
```

Jika mobil tidak ditemukan:

```
Registration number DL-12-AA-9999 not found
```

## Catatan

- Pastikan menjalankan perintah dari folder `cmd`.
- Jika terjadi error, periksa apakah parking lot sudah dibuat sebelum parkir mobil.

## Contoh Penggunaan

```
go run main.go create_parking_lot 6
go run main.go park B1234XYZ White
go run main.go status
go run main.go leave B1234XYZ 4
go run main.go status
```
