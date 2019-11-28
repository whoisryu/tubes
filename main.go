package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

var (
	tanggal                                                 time.Time
	order                                                   []string
	total                                                   float64
	i, id, jml, no, all                                     int
	action, nama, alamat, lagi, pesanFav, value, layoutDate string
	costumer                                                []org
	history                                                 []riwayat
	listMenu                                                [7]makanan
	cek                                                     bool
	angka                                                   []int
)

type riwayat struct {
	nama  string
	date  string
	order []string
	total int
}
type org struct {
	id      int
	nama    string
	menuFav int
	alamat  string
}
type makanan struct {
	no    int
	nama  string
	harga int
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	costumer = append(
		costumer,
		org{
			nama:    "aryuska",
			alamat:  "jatinangor",
			menuFav: 1})

	scanner.Scan()
	action = scanner.Text()
	if action == "turn on" {
		fmt.Println("SELAMAT DATANG DI APLIKASI PESAN MAKANAN")
		fmt.Println("----------------------------------------")
	}
	for {
		fmt.Print("Ketik Aksi: ")
		scanner.Scan()
		action = scanner.Text()
		for !validateValidAction(action) {
			if !validateValidAction(action) {
				fmt.Println("AKSI TIDAK VALID")
				fmt.Print("Ketik Aksi: ")
				scanner.Scan()
				action = scanner.Text()
			}
		}
		fmt.Println()
		if action == "pesan" {
			fmt.Print("Nama Pemesan : ")
			scanner.Scan()
			nama = scanner.Text()
			for !validateEmptyDataNama(nama) {
				if !validateEmptyDataNama(nama) {
					fmt.Println("NAMA TIDAK BOLEH KOSONG!!!")
					fmt.Print("Nama: ")
					scanner.Scan()
					nama = scanner.Text()
				}
			}
			cek, id = search(nama)
			if nama == "done" {
				break
			}
			if cek == true {
				pesan(cek, id)
			} else {
				fmt.Println("ANDA MEMIIKI PELANGGAN BARU")
				create()
				pesan(cek, id)
			}
		} else if action == "menu1" {
			t := table.NewWriter()
			tTemp := table.Table{}
			tTemp.Render()
			t.AppendHeader(table.Row{"NO", "NAMA"})
			for i := 0; i < len(costumer); i++ {
				if costumer[i].menuFav == 1 {
					t.AppendRow([]interface{}{i + 1, costumer[i].nama})
				}
			}
			fmt.Println(t.Render())
		} else if action == "history" {
			sortDate(history)
			showHistory()
		}
		if action == "mati" {
			break
		}
	}

}

//history
func showHistory() {
	if !validateEmptyHistory(history) {
		fmt.Println("ANDA BELUM MELAKUKAN TRANSAKSI")
	} else {
		t := table.NewWriter()
		tTemp := table.Table{}
		tTemp.Render()
		t.AppendHeader(table.Row{"TANGGAL", "NAMA", "ORDER", "TOTAL"})
		for i := 0; i < len(history); i++ {
			t.AppendRow([]interface{}{history[i].date, history[i].nama, history[i].order, history[i].total})
		}
		fmt.Println(t.Render())
	}
}

// sortDate
func sortDate(arr []riwayat) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j].date < arr[min].date {
				min = j
			}
		}
		arr[min].date, arr[i].date = arr[i].date, arr[min].date
	}
}

func tambahOrder(no int) {
	if no == 1 {
		order = append(order, "Nasi Uduk")
	} else if no == 2 {
		order = append(order, "Pecel Lele")
	} else if no == 3 {
		order = append(order, "Es Teh Manis")
	} else if no == 4 {
		order = append(order, "Ayam Bakar")
	} else if no == 5 {
		order = append(order, "Ayam Geprek")
	} else if no == 6 {
		order = append(order, "Es Jeruk")
	} else if no == 7 {
		order = append(order, "Nasi Putih")
	}

}

//pesan
func pesan(cek bool, id int) {
	fmt.Println()
	fmt.Println("Masukkan Format Tanggal Seperti Berikut: dd/mm/yyyy")
	fmt.Print("Tanggal : ")
	scanner.Scan()
	value = scanner.Text()
	layoutDate = "2/1/2006"
	tanggal, _ = time.Parse(layoutDate, value)
	if cek == true {
		menu(id, cek)
		fmt.Print("Ingin memesan ini?: ")
		scanner.Scan()
		pesanFav = scanner.Text()
		for !validateValidPesanFav(pesanFav) {
			if !validateValidPesanFav(pesanFav) {
				fmt.Println("JAWABAN TIDAK VALID")
				fmt.Print("Ingin memesan ini?: ")
				scanner.Scan()
				pesanFav = scanner.Text()
			}
		}

		if pesanFav == "ya, itu saja" {
			fmt.Print("Jumlah : ")
			fmt.Scanln(&jml)
			no = costumer[id].menuFav
			tambahOrder(no)

			all = all + showTotal(no, jml)
			fmt.Println("Total : Rp.", all)
			//all = 0
		} else if pesanFav == "ya, tapi tampilkan menu lain juga" {
			no = costumer[id].menuFav

			tambahOrder(no)
			fmt.Print("Jumlah : ")
			fmt.Scanln(&jml)
			all = all + showTotal(no, jml)
			menu(id, false)
			for {
				fmt.Print("No : ")
				fmt.Scanln(&no)

				tambahOrder(no)

				fmt.Print("Jumlah : ")
				fmt.Scanln(&jml)
				all = all + showTotal(no, jml)
				fmt.Print("Tambah lagi??")
				fmt.Scanln(&lagi)
				if lagi == "tidak" {
					fmt.Println("Total : Rp.", all)
					//all = 0
					break
				}
			}
		} else if pesanFav == "tidak, tampilkan menu lain" {
			menu(id, false)
			for {
				fmt.Print("No : ")
				fmt.Scanln(&no)

				tambahOrder(no)

				fmt.Print("Jumlah : ")
				fmt.Scanln(&jml)
				all = 0
				all = all + showTotal(no, jml)
				fmt.Print("Tambah lagi??")
				fmt.Scanln(&lagi)
				if lagi == "tidak" {
					fmt.Println("Total : Rp.", all)
					//all = 0
					break
				}
			}
		}
	} else {
		menu(id, cek)
		var fav = -1
		for {
			fmt.Print("No : ")
			fmt.Scanln(&no)

			tambahOrder(no)
			fmt.Print("Jumlah : ")
			fmt.Scanln(&jml)
			if fav < jml {
				fav = jml
				costumer[id].menuFav = no
			}
			all = all + showTotal(no, jml)
			fmt.Print("Tambah lagi??")
			scanner.Scan()
			lagi = scanner.Text()
			if lagi == "tidak" {

				fmt.Println("Total : Rp.", all)
				//all = 0
				break
			}
		}
	}

	history = append(history, riwayat{
		nama:  costumer[id].nama,
		date:  tanggal.String(),
		order: order,
		total: all})
	all = 0
	order = nil

}

//create
func create() {
	i = 1
	fmt.Print("Nama: ")
	scanner.Scan()
	nama = scanner.Text()

	for !validateEmptyDataNama(nama) {
		if !validateEmptyDataNama(nama) {
			fmt.Println("NAMA TIDAK BOLEH KOSONG!!!")
			fmt.Print("Nama: ")
			scanner.Scan()
			nama = scanner.Text()
		}
	}

	fmt.Print("Alamat: ")
	scanner.Scan()
	alamat = scanner.Text()

	for !validateEmptyDataNama(alamat) {
		if !validateEmptyDataNama(alamat) {
			fmt.Println("ALAMAT TIDAK BOLEH KOSONG!!!")
			fmt.Print("Alamat: ")
			scanner.Scan()
			alamat = scanner.Text()
		}
	}

	costumer = append(
		costumer,
		org{
			id:      i,
			nama:    nama,
			alamat:  alamat,
			menuFav: 0})
	i++
}

// search
func search(nama string) (bool, int) {
	for i = 0; i < len(costumer); i++ {
		if costumer[i].nama == nama {
			cek = true
			break
		} else {
			cek = false
		}
	}
	return cek, i
}

//total
func showTotal(no int, jml int) int {
	total := 0

	if no == 1 {
		total = (listMenu[0].harga * jml)
	} else if no == 2 {
		total = listMenu[1].harga * jml
	} else if no == 3 {
		total = (listMenu[2].harga * jml)
	} else if no == 4 {
		total = (listMenu[3].harga * jml)
	} else if no == 5 {
		total = (listMenu[4].harga * jml)
	} else if no == 6 {
		total = (listMenu[5].harga * jml)
	} else if no == 7 {
		total = (listMenu[6].harga * jml)
	}

	if no == 2 && jml == 2 {
		total = total - ((listMenu[1].harga * jml * 25) / 100)
	} else if no == 1 && jml == 5 {
		total = total - ((listMenu[0].harga * jml * 20) / 100)
	} else if no == 5 {
		total = total - ((listMenu[4].harga * jml * 20) / 100)
	} else if no == 3 && jml == 2 {
		total = total - ((listMenu[2].harga * jml * 25) / 100)
	}

	return total
}

// MENU
func menu(id int, cek bool) {
	listMenu[0].no = 1
	listMenu[0].nama = "Nasi Uduk"
	listMenu[0].harga = 5000

	listMenu[1].no = 2
	listMenu[1].nama = "Pecel Lele"
	listMenu[1].harga = 12000

	listMenu[2].no = 3
	listMenu[2].nama = "Es Teh Manis"
	listMenu[2].harga = 3000

	listMenu[3].no = 4
	listMenu[3].nama = "Ayam Bakar"
	listMenu[3].harga = 15000

	listMenu[4].no = 5
	listMenu[4].nama = "Ayam Geprek"
	listMenu[4].harga = 13000

	listMenu[5].no = 6
	listMenu[5].nama = "Es Jeruk"
	listMenu[5].harga = 4000

	listMenu[6].no = 7
	listMenu[6].nama = "Nasi Putih"
	listMenu[6].harga = 3000

	t := table.NewWriter()
	tTemp := table.Table{}
	tTemp.Render()
	if cek == true {
		if costumer[id].menuFav == 1 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"1", listMenu[0].nama, listMenu[0].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 2 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"2", listMenu[1].nama, listMenu[1].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 3 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"3", listMenu[2].nama, listMenu[2].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 4 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"4", listMenu[3].nama, listMenu[3].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 5 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"5", listMenu[4].nama, listMenu[4].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 6 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"6", listMenu[5].nama, listMenu[5].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == 7 {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"7", listMenu[6].nama, listMenu[6].harga})
			fmt.Println(t.Render())
		}
	} else if cek == false {
		t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
		t.AppendRow([]interface{}{"1", listMenu[0].nama, listMenu[0].harga})
		t.AppendRow([]interface{}{"2", listMenu[1].nama, listMenu[1].harga})
		t.AppendRow([]interface{}{"3", listMenu[2].nama, listMenu[2].harga})
		t.AppendRow([]interface{}{"4", listMenu[3].nama, listMenu[3].harga})
		t.AppendRow([]interface{}{"5", listMenu[4].nama, listMenu[4].harga})
		t.AppendRow([]interface{}{"6", listMenu[5].nama, listMenu[5].harga})
		t.AppendRow([]interface{}{"7", listMenu[6].nama, listMenu[6].harga})
		fmt.Println(t.Render())
	}
}

func validateEmptyDataNama(nama string) bool {
	if len(nama) == 0 {
		return false
	}
	return true
}

func validateEmptyDataAlamat(nama string) bool {
	if len(nama) == 0 {
		return false
	}
	return true
}

func validateValidAction(action string) bool {
	if action == "pesan" || action == "menu1" || action == "history" || action == "mati" {
		return true
	}
	return false
}

func validateEmptyHistory(history []riwayat) bool {
	if len(history) == 0 {
		return false
	}
	return true
}

func validateValidPesanFav(pesan string) bool {
	if pesan == "ya, itu saja" || pesan == "ya, tampilkan menu yang lain juga" || pesan == "tidak, tampilkan menu lain" {
		return true
	}
	return false
}
