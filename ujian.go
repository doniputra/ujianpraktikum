package main
import (
	"fmt"
)
type barang struct {
	NamaBarang  string
	HargaBarang int
}
type belanja struct {
	JumlahBarang int
}
func main() {
	brng1 := barang{
		NamaBarang:  "Sabun Cuci Piring",
		HargaBarang: 50000,
	}
	brng2 := barang{
		NamaBarang:  "Sabun Mandi",
		HargaBarang: 25000,
	}
	brng3 := barang{
		NamaBarang:  "Shampoo",
		HargaBarang: 35000,
	}
	brng4 := barang{
		NamaBarang:  "Pasta Gigi",
		HargaBarang: 10000,
	}
	brng5 := barang{
		NamaBarang:  "Tissue",
		HargaBarang: 7500,
	}
	blja1 := belanja{
		JumlahBarang: 7,
	}
	blja2 := belanja{
		JumlahBarang: 9,
	}
	blja3 := belanja{
		JumlahBarang: 10,
	}
	blja4 := belanja{
		JumlahBarang: 11,
	}
	blja5 := belanja{
		JumlahBarang: 20,
	}

	var total int
	total = (brng1.HargaBarang * blja1.JumlahBarang) + (brng2.HargaBarang * blja2.JumlahBarang) +
		(brng3.HargaBarang * blja3.JumlahBarang) + (brng4.HargaBarang * blja4.JumlahBarang) +
		(brng5.HargaBarang * blja5.JumlahBarang)

	fmt.Println("---List Barang Belanja---")
	fmt.Println(brng1.NamaBarang, "=", blja1.JumlahBarang, "buah")
	fmt.Println(brng2.NamaBarang, "=", blja2.JumlahBarang, "buah")
	fmt.Println(brng3.NamaBarang, "=", blja3.JumlahBarang, "buah")
	fmt.Println(brng4.NamaBarang, "=", blja4.JumlahBarang, "buah")
	fmt.Println(brng5.NamaBarang, "=", blja5.JumlahBarang, "buah")
	fmt.Println("Total Biaya yang harus dibayar:", total)
}