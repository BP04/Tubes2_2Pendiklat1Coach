# Tugas Besar 2 IF2211 - Strategi Algoritma

By: 2Pendiklat1Coach

![image](https://github.com/user-attachments/assets/57a59147-c904-4441-8c5a-d7c1d4c15d8a)

Ini adalah repositori yang dibuat untuk pemenuhan tugas besar IF2211 Strategi Algoritma. Tugas besar ini menuntut pembuatan website untuk melakukan visualisasi pohon pembangun elemen dalam permainan Little Alchemy 2.

## Algorithm

Pengguna memasukkan nama elemen yang dicari beserta algoritma pencarian yang diinginkan. Berdasarkan masukan ini, pencarian pada graf akan dijalankan oleh program, baik dengan DFS atau BFS sambil mengkalkulasikan durasi dari pencarian dan jumlah simpul yang dikunjungi. Dua parameter ini, beserta jalur recipes yang ditemukan akan disimpan dalam suatu struktur data. Algoritma DFS dan BFS yang diimplementasikan dapat mencari recipe paling pendek (yang memerlukan penggabungan paling sedikit) ataupun mencari beberapa recipes berbeda. Untuk melakukan pencarian, pertama dibangun semua pohon recipe yang mungkin, lalu ditentukan pohon yang akan diambil. Dilakukan juga optimasi menggunakan caching.

## Built With

* [![Go][Go-img]][Go-url]
* [![React][React-img]][React-url]
* [![Gorilla WebSocket][Gorilla-img]][Gorilla-url]
* [![react-d3-tree][D3Tree-img]][D3Tree-url]

[Go-img]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/

[React-img]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/

[Gorilla-img]: https://img.shields.io/badge/Gorilla%20WebSocket-009688?style=for-the-badge&logo=websocket&logoColor=white
[Gorilla-url]: https://github.com/gorilla/websocket

[D3Tree-img]: https://img.shields.io/badge/react--d3--tree-282c34?style=for-the-badge&logo=react&logoColor=61DAFB
[D3Tree-url]: https://github.com/bkrem/react-d3-tree

## Prerequisites

### Backend (Go)

> Make sure you have Go installed. [Install Go](https://golang.org/doc/install)

```bash
# Create project directory
mkdir backend && cd backend

# Initialize a Go module
go mod init your-module-name

# Install Gorilla WebSocket
go get github.com/gorilla/websocket
```

### Frontend (React.js)

```bash
# Create a new React app
npx create-react-app frontend

# Move into the project folder
cd frontend

# Install react-d3-tree
npm install react-d3-tree
```

## How To Run

#### Using Docker
Run on root
```
docker compose up --build
```

#### Local
On ```src/fe``` run
```
npm run dev
```

On ```src/be``` run
```
go run main.go
```

## Contributors

- Front End Engineer - Faqih Muhammad Syuhada - [LinkedIn](https://www.linkedin.com/in/faqihmuhammadsyuhada/) - 13523057@std.stei.itb.ac.id
- Back End Engineer - Benedict Presley - [LinkedIn](https://www.linkedin.com/in/benedict-presley-b834ab29a/) - 13523067@std.stei.itb.ac.id
- Production Engineer - Zulfaqqar Nayaka Athadiansyah - [LinkedIn](https://www.linkedin.com/in/nayaka-zna/) - 13523094@std.stei.itb.ac.id
