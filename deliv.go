package main
import "fmt"
const NMAX = 1024
type delivery struct{
	namaPenerima,alamat,status string
	id int
}
type arrdeliv[NMAX] delivery
func main(){
	var menu bool
	var pilih,nPengiriman int
	var listPengiriman arrdeliv
	nPengiriman = 0
	menu = true
	for menu {
		banner()
		printdata(listPengiriman,nPengiriman)
		fmt.Print("(Menu utama) Masukan angka : ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahPengiriman(&listPengiriman,&nPengiriman)
		}else if pilih == 2{
			editData(&listPengiriman,nPengiriman)
		}else if pilih == 3{
			deleteData(&listPengiriman,&nPengiriman)
		}else if pilih == 4{
			cariStatus(listPengiriman,nPengiriman)
		}else if pilih == 5 {
			sortir(&listPengiriman,nPengiriman)
		}else if pilih == 6 {
			menu = false
		}
	}
}

func banner(){
	fmt.Println("====================================")
	fmt.Println("|                                  |")
	fmt.Println("|        Chambre de la Food        |")
	fmt.Println("|          Delivery Order          |")
	fmt.Println("|                                  |")
	fmt.Println("====================================")
	fmt.Println("1. Tambahkan Pengiriman")
	fmt.Println("2. Edit Status")
	fmt.Println("3. Hapus Pengiriman")
	fmt.Println("4. Cari Status Pengiriman")
	fmt.Println("5. Sortir pengiriman")
	fmt.Println("6. Keluar")
}

func tambahPengiriman(A *arrdeliv,n *int){
	var penerima, alamatPenerima,tambah string
	var statusDeliv,antrian,idx int
	var input bool = true
	for input {
		var delivmenu bool = true
		var antrimenu bool = true
		*n++
		fmt.Print("Masukan nama penerima: ")
		fmt.Scan(&penerima)
		fmt.Print("Masukan alamat : ")
		fmt.Scan(&alamatPenerima)
		fmt.Println("Pilih status delivery : ")
		fmt.Println("1. Diproses")
		fmt.Println("2. Dikirim")
		fmt.Println("3. Diterima")
		for delivmenu {
			fmt.Print("Masukan pilihan status : ")
			fmt.Scan(&statusDeliv)
			if statusDeliv == 1 {
				A[*n].status = "Diproses"
				delivmenu = false
			}else if statusDeliv == 2{
				A[*n].status = "Dikirim"
				delivmenu = false
			}else if statusDeliv == 3{
				A[*n].status = "Diterima"
				delivmenu = false
			}
		}
		for antrimenu {
			fmt.Print("Masukan id : ")
			fmt.Scan(&antrian)
			idx = cariData(*A,*n,antrian)
			if idx == -1 && antrian > 0{
				antrimenu = false
			}else if idx <= 0 || antrian == 0{
				fmt.Println("Id harus lebih dari 0")
			}else{
				fmt.Println("Id sudah ada")
			}
		}
		A[*n].id = antrian
		A[*n].namaPenerima = penerima
		A[*n].alamat = alamatPenerima
		fmt.Print("Apakah masih ingin menambah Pengiriman (Y/N) : ")
		fmt.Scan(&tambah)
		if tambah == "N" || tambah == "n" {
			input = false
		}
	}
}

func printdata(A arrdeliv,n int){
	fmt.Println("Data pengiriman : ")
	for i := 1; i <= n; i++ {
		fmt.Printf("Id : %d | Nama penerima : %s | Alamat : %s | Status pengiriman : %s ",A[i].id,A[i].namaPenerima,A[i].alamat,A[i].status)
		fmt.Println()
	}
}

func editData(A *arrdeliv, n int){
	var editData,delivmenu bool
	var idEdit,idx,statusDeliv int
	editData = true
	for editData{
		delivmenu = true
		fmt.Print("Masukan id yang ingin diedit statusnya : ")
		fmt.Scan(&idEdit)
		idx = cariData(*A,n,idEdit)
		if idx != -1{
			editData = false
			fmt.Println("Pilih status delivery : ")
			fmt.Println("1. Diproses")
			fmt.Println("2. Dikirim")
			fmt.Println("3. Diterima")
			for delivmenu {
				fmt.Print("Masukan pilihan status : ")
				fmt.Scan(&statusDeliv)
				if statusDeliv == 1 {
					A[idx].status = "Diproses"
					delivmenu = false
				}else if statusDeliv == 2{
					A[idx].status = "Dikirim"
					delivmenu = false
				}else if statusDeliv == 3{
					A[idx].status = "Diterima"
					delivmenu = false
				}
			}
		}else{
			fmt.Println("Id tidak ditemukan")
		}
	}

}

func cariData(A arrdeliv,n,x int)int{
	var idx int
	idx = -1
	for i := 1; i <= n; i++ {
		if A[i].id == x {
			idx = i
		}
	}
	return idx
}

func deleteData(A *arrdeliv, n *int) {
	var deleteMenu bool
	var idDelete, idx int
	deleteMenu = true
	for deleteMenu {
		fmt.Print("Masukan id yang ingin dihapus : ")
		fmt.Scan(&idDelete)
		idx = cariData(*A, *n, idDelete)
		if idx != -1 {
			for i := idx; i < *n; i++ {
				A[i] = A[i+1]
			}
			*n--
			deleteMenu = false
		} else {
			fmt.Println("Id tidak ditemukan")
		}
	}
}

func cariStatus(A arrdeliv, n int){
	var menu bool
	var statusDicari int
	menu = true
	for menu{
		fmt.Println("Pilih status delivery yang dicari : ")
		fmt.Println("1. Diproses")
		fmt.Println("2. Dikirim")
		fmt.Println("3. Diterima")
		fmt.Print("Masukan pilihan status : ")
		fmt.Scan(&statusDicari)
		if statusDicari == 1 {
			printStatusDicari(A,n,"Diproses")
			menu = false
		}else if statusDicari == 2{
			printStatusDicari(A,n,"Dikirim")
			menu = false
		}else if statusDicari == 3{
			printStatusDicari(A,n,"Diterima")
			menu = false
		}
	}	
}

func printStatusDicari(A arrdeliv,n int,x string){
	var menu bool
	var pilih string
	menu = true
	for menu {
		for i := 1; i <=n; i++ {
			if A[i].status == x {
				fmt.Printf("Id : %d | Nama penerima : %s | Alamat : %s | Status pengiriman : %s ",A[i].id,A[i].namaPenerima,A[i].alamat,A[i].status)
				fmt.Println()
			}
		}
		fmt.Print("Tekan X untuk keluar : ")
		fmt.Scan(&pilih)
		if pilih == "X" || pilih == "x" {
			menu = false
		}
	}
}

func sortir(A *arrdeliv, n int){
	var pass,idx,i int
	var temp delivery
	for pass = 2; pass <= n; pass++ {
		idx = pass - 1
		for i = pass; i <= n; i++ {
			if A[idx].id > A[i].id {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass - 1] = A[idx]
		A[idx] = temp
	}
}