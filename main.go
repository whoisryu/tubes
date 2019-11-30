package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

var (
	tanggal                                                             time.Time
	order                                                               []string
	total                                                               float64
	i, id, jml, all                                                     int
	action, nama, alamat, lagi, pesanFav, value, layoutDate, jumlah, no string
	costumer                                                            []org
	history                                                             []riwayat
	listMenu                                                            [7]makanan
	cek                                                                 bool
	angka                                                               []int
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
	menuFav string
	alamat  string
}
type makanan struct {
	no    int
	nama  string
	harga int
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	action = scanner.Text()
	if action == "turn on" {
		fmt.Println("SELAMAT DATANG DI APLIKASI PESAN MAKANAN")
		fmt.Println("----------------------------------------")
	}
	for {
		fmt.Println()
		fmt.Println("1. Pesan")
		fmt.Println("2. Lihat Kostumer Favorit Menu 1")
		fmt.Println("3. Lihat Riwayat Transaksi")
		fmt.Println("4. Keluar Program")
		fmt.Print("Pilih Menu: ")
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
		if action == "1" {
			fmt.Print("Nama Pemesan : ")
			scanner.Scan()
			nama = scanner.Text()
			for !validateEmptyData(nama) {
				if !validateEmptyData(nama) {
					fmt.Println("NAMA TIDAK BOLEH KOSONG!!!")
					fmt.Print("Nama: ")
					scanner.Scan()
					nama = scanner.Text()
				}
			}
			cek, id = search(nama)
			// if nama == "done" {
			// 	break
			// }
			if cek == true {
				pesan(cek, id)
			} else {
				fmt.Println("ANDA MEMIIKI PELANGGAN BARU")
				create(nama)
				pesan(cek, id)
			}
		} else if action == "2" {
			j := 1
			t := table.NewWriter()
			tTemp := table.Table{}
			tTemp.Render()
			t.AppendHeader(table.Row{"NO", "ID", "NAMA"})
			for i := 0; i < len(costumer); i++ {
				if costumer[i].menuFav == "1" {
					t.AppendRow([]interface{}{j, costumer[i].nama})
					j++
				}
			}
			fmt.Println(t.Render())
		} else if action == "3" {
			sortDate(history)
			showHistory()
		}
		if action == "4" {
			fmt.Println("yakin keluar ? (y/n)")
			fmt.Scan(&action)
			if action == "y" {
				break
			} else if action == "n" {
				main()
			}

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
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func tambahOrder(no string) {
	if no == "1" {
		order = append(order, "Nasi Uduk")
	} else if no == "2" {
		order = append(order, "Pecel Lele")
	} else if no == "3" {
		order = append(order, "Es Teh Manis")
	} else if no == "4" {
		order = append(order, "Ayam Bakar")
	} else if no == "5" {
		order = append(order, "Ayam Geprek")
	} else if no == "6" {
		order = append(order, "Es Jeruk")
	} else if no == "7" {
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
	for !validateEmptyData(value) {
		if !validateEmptyData(value) {
			fmt.Println("Tanggal Tidak Boleh Kosong")
			fmt.Println("Masukkan Format Tanggal Seperti Berikut: dd/mm/yyyy")
			fmt.Print("Tanggal : ")
			scanner.Scan()
			value = scanner.Text()
		}
	}
	layoutDate = "2/1/2006"
	tanggal, _ = time.Parse(layoutDate, value)
	if cek == true {
		menu(id, cek)
		fmt.Println("1. ya, itu saja")
		fmt.Println("2. ya, tapi tampilkan menu lain juga")
		fmt.Println("3. tidak, tampilkan menu lain")
		fmt.Print("Ingin memesan ini?: ")
		scanner.Scan()
		pesanFav = scanner.Text()
		_, err := strconv.Atoi(pesanFav)
		for err != nil {
			if _, err = strconv.Atoi(pesanFav); err != nil {
				fmt.Println(" No Menu Harus Dimasukkan dengan angka")
			}
			fmt.Print("No : ")
			fmt.Scanln(&pesanFav)
			_, err = strconv.Atoi(pesanFav)
		}

		if pesanFav == "1" {
			fmt.Print("Jumlah : ")
			scanner.Scan()
			jumlah = scanner.Text()
			jml, err := strconv.Atoi(jumlah)
			for err != nil {
				if _, err = strconv.Atoi(jumlah); err != nil {
					fmt.Println("Jumlah Harus Dimasukkan dengan angka")
				}
				fmt.Print("Jumlah : ")
				scanner.Scan()
				jumlah = scanner.Text()
				jml, err = strconv.Atoi(jumlah)
			}

			no = costumer[id].menuFav
			tambahOrder(no)

			all = all + showTotal(no, jml)
			fmt.Println("Total : Rp.", all)
			//all = 0
		} else if pesanFav == "2" {
			no = costumer[id].menuFav

			tambahOrder(no)
			fmt.Print("Jumlah : ")
			scanner.Scan()
			jumlah = scanner.Text()
			jml, err := strconv.Atoi(jumlah)
			for err != nil {
				if _, err = strconv.Atoi(jumlah); err != nil {
					fmt.Println("Jumlah Harus Dimasukkan dengan angka")
				}
				fmt.Print("Jumlah : ")
				scanner.Scan()
				jumlah = scanner.Text()
				jml, err = strconv.Atoi(jumlah)
			}
			all = all + showTotal(no, jml)
			menu(id, false)
			for {
				fmt.Print("No : ")
				fmt.Scanln(&no)
				_, err = strconv.Atoi(no)
				for err != nil {
					if _, err = strconv.Atoi(no); err != nil {
						fmt.Println(" No Menu Harus Dimasukkan dengan angka")
					}
					fmt.Print("No : ")
					fmt.Scanln(&no)
					_, err = strconv.Atoi(no)
				}
				tambahOrder(no)

				fmt.Print("Jumlah : ")
				scanner.Scan()
				jumlah = scanner.Text()
				jml, err := strconv.Atoi(jumlah)
				for err != nil {
					if _, err = strconv.Atoi(jumlah); err != nil {
						fmt.Println("Jumlah Harus Dimasukkan dengan angka")
					}
					fmt.Print("Jumlah : ")
					scanner.Scan()
					jumlah = scanner.Text()
					jml, err = strconv.Atoi(jumlah)
				}
				all = all + showTotal(no, jml)
				fmt.Print("Tambah lagi??")
				fmt.Scanln(&lagi)
				if lagi == "tidak" {
					fmt.Println("Total : Rp.", all)
					//all = 0
					break
				}
			}
		} else if pesanFav == "3" {
			menu(id, false)
			for {
				fmt.Print("No : ")
				fmt.Scanln(&no)
				_, err := strconv.Atoi(no)
				for err != nil {
					if _, err = strconv.Atoi(no); err != nil {
						fmt.Println(" No Menu Harus Dimasukkan dengan angka")
					}
					fmt.Print("No : ")
					fmt.Scanln(&no)
					_, err = strconv.Atoi(no)
				}

				tambahOrder(no)

				fmt.Print("Jumlah : ")
				scanner.Scan()
				jumlah = scanner.Text()
				jml, err := strconv.Atoi(jumlah)
				for err != nil {
					if _, err = strconv.Atoi(jumlah); err != nil {
						fmt.Println("Jumlah Harus Dimasukkan dengan angka")
					}
					fmt.Print("Jumlah : ")
					scanner.Scan()
					jumlah = scanner.Text()
					jml, err = strconv.Atoi(jumlah)
				}
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
			_, err := strconv.Atoi(no)
			for err != nil {
				if _, err = strconv.Atoi(no); err != nil {
					fmt.Println(" No Menu Harus Dimasukkan dengan angka")
				}
				fmt.Print("No : ")
				fmt.Scanln(&no)
				_, err = strconv.Atoi(no)
			}

			tambahOrder(no)
			fmt.Print("Jumlah : ")
			scanner.Scan()
			jumlah = scanner.Text()
			jml, err := strconv.Atoi(jumlah)
			for err != nil {
				if _, err = strconv.Atoi(jumlah); err != nil {
					fmt.Println("Jumlah Harus Dimasukkan dengan angka")
				}
				fmt.Print("Jumlah : ")
				scanner.Scan()
				jumlah = scanner.Text()
				jml, err = strconv.Atoi(jumlah)
			}
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
func create(nama string) {
	i = 1
	fmt.Println("Nama: ", nama)
	fmt.Print("Alamat: ")
	scanner.Scan()
	alamat = scanner.Text()

	for !validateEmptyData(alamat) {
		if !validateEmptyData(alamat) {
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
			menuFav: ""})
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
func showTotal(no string, jml int) int {
	total := 0

	if no == "1" {
		total = (listMenu[0].harga * jml)
	} else if no == "2" {
		total = listMenu[1].harga * jml
	} else if no == "3" {
		total = (listMenu[2].harga * jml)
	} else if no == "4" {
		total = (listMenu[3].harga * jml)
	} else if no == "5" {
		total = (listMenu[4].harga * jml)
	} else if no == "6" {
		total = (listMenu[5].harga * jml)
	} else if no == "7" {
		total = (listMenu[6].harga * jml)
	}

	if no == "2" && jml == 2 {
		total = total - ((listMenu[1].harga * jml * 25) / 100)
	} else if no == "1" && jml == 5 {
		total = total - ((listMenu[0].harga * jml * 20) / 100)
	} else if no == "5" {
		total = total - ((listMenu[4].harga * jml * 20) / 100)
	} else if no == "3" && jml == 2 {
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
		if costumer[id].menuFav == "1" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"1", listMenu[0].nama, listMenu[0].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "2" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"2", listMenu[1].nama, listMenu[1].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "3" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"3", listMenu[2].nama, listMenu[2].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "4" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"4", listMenu[3].nama, listMenu[3].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "5" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"5", listMenu[4].nama, listMenu[4].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "6" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"6", listMenu[5].nama, listMenu[5].harga})
			fmt.Println(t.Render())
		} else if costumer[id].menuFav == "7" {
			t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
			t.AppendRow([]interface{}{"7", listMenu[6].nama, listMenu[6].harga})
			fmt.Println(t.Render())
		}
	} else if cek == false {
		t.AppendHeader(table.Row{"NO", "NAMA", "HARGA"})
		for i := 0; i < 7; i++ {
			t.AppendRow([]interface{}{i + 1, listMenu[i].nama, listMenu[i].harga})
		}
		fmt.Println(t.Render())
	}
}

func validateEmptyData(isi string) bool {
	if len(isi) == 0 {
		return false
	}
	return true
}

func validateValidAction(action string) bool {
	if action == "1" || action == "2" || action == "3" || action == "4" {
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
	if pesan == "ya, itu saja" || pesan == "ya, tapi tampilkan menu lain juga" || pesan == "tidak, tampilkan menu lain" {
		return true
	}
	return false
}

func validateMenu(no string) bool {
	if no == "1" || no == "2" || no == "3" || no == "4" || no == "5" || no == "6" || no == "7" {
		return true
	}
	return false
}
