package main
import (
	"fmt"
	"strings"
	"strconv"
)
func main() {
	var fungsi, value, data1, data2 string
	var index, status, batasinput, statusdelete, indexdelete int
	var nomorloker [99]int
	var tipeidentitas [99]string
	var nomoridentitas [99]string
	var find int
	var statusfind string
	
	batasinput = 0
	statusdelete = 0
	statusfind = "null"
	
	for {
		fmt.Scanln(&value, &data1, &data2)
		
		fungsi = strings.ToLower(value)
		
		if fungsi == "init" {
			if data1 != "" {
				fmt.Println("Berhasil membuat loker dengan jumlah", data1)
				status = 1
				
				i, err := strconv.Atoi(data1)
				if err == nil {
				
				}
				
				index = i
			} else {
				fmt.Println("Gagal membuat loker, mohon isi jumlah loker yang inginkan.")
			}
		} else if fungsi == "input" {
			if status == 1 {
				if batasinput != index {
					nomorloker[batasinput] = batasinput + 1
					tipeidentitas[batasinput] = strings.ToUpper(data1)
					nomoridentitas[batasinput] = data2
					
					fmt.Println("Kartu identitas tersimpan di loker nomor", batasinput + 1)
					
					batasinput = batasinput + 1
				} else if statusdelete == 1 {
					tipeidentitas[indexdelete] = strings.ToUpper(data1)
					nomoridentitas[indexdelete] = data2
					
					fmt.Println("Kartu identitas tersimpan di loker nomor", indexdelete + 1)
					
					statusdelete = 0
					indexdelete = 0
				} else {
					fmt.Println("Maaf loker sudah penuh")
				}
			} else {
				fmt.Println("Silahkan membuat loker terlebih dahulu dengan perintah 'init'.")
			}
		} else if fungsi == "status" {
			fmt.Println("No Loker\t Tipe Identitas\t No Identitas")
			
			for i := 0; i < batasinput; i++ {
				fmt.Println(nomorloker[i], "\t\t", tipeidentitas[i], "\t\t", nomoridentitas[i])
			}
		} else if fungsi == "leave" {
			i, err := strconv.Atoi(data1)
			if err == nil {
			
			}
			
			indexdelete = i - 1
			if tipeidentitas[i-1] != "" {
				tipeidentitas[i-1] = "kosong"
				nomoridentitas[i-1] = "kosong"
				
				fmt.Println("Loker nomor", i, "berhasil dikosongkan.")
				
				statusdelete = 1
			} else {
				fmt.Println("Nomor loker tidak tersedia.")
			}
		} else if fungsi == "find" {
			for i := 0; i < batasinput; i++ {
				if nomoridentitas[i] == strings.ToUpper(data1) {
					find = i
					statusfind = "masuk"
				}
			}

			if statusfind == "masuk" {
				fmt.Println("Kartu identitas tersimpan di loker nomor", find + 1)
				statusfind = "null"
			} else if statusfind == "null" {
				fmt.Println("Kartu identitas yang anda cari tidak ditemukan.")
			}
		} else if fungsi == "search" {
			var search [99]string
			indexsearch := 0
			for i := 0; i < batasinput; i++ {
				if tipeidentitas[i] == strings.ToUpper(data1) {
					search[indexsearch] = nomoridentitas[i]
					indexsearch = indexsearch + 1
				}
			}
			
			for i := 0; i < indexsearch; i++ {
				fmt.Printf(search[i])
				
				if i != indexsearch - 1 {
					fmt.Printf(",")
				} else {
					fmt.Println("")
				}
			}

			if indexsearch == 0 {
				fmt.Println("Tipe Identitas yang ada cari tidak ditemukan.")
			}
		} else if fungsi == "exit" {
			break
		} else {
			fmt.Println("Perintah tidak ditemukan.")
		}
	}
} 
