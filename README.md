# Movie Picker

Movie Picker menampilkan pilihan film secara random dengan meminta masukan dari pengguna berupa Rating, Tahun, dan Genre

# Installation
  - Install Golang 1.9.2
  - Install Beego
  - Install mysql for Beego
  - Import `cinema_management_system.sql` 
  - Fork https://github.com/ariefsptn/Movie-Picker
 
```sh
$ cd $GOPATH/src
# Clone Forked Repo
$ git clone https://github.com/yourname/Movie-Picker
$ cd $GOPATH/src/github.com/yourname/Movie-Picker
$ bee run
```

# Usage
```sh
# Menu Utama -- Input
localhost:9009/

## Pencarian kembali dengan masukan sama
localhost:9009/searcha

## Cetak Json
localhost:9009/json
```


