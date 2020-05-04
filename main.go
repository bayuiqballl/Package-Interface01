package main

import (
	"fmt"

	"github.com/bayuiqballl/package-interface01/motor"
	"github.com/bayuiqballl/package-interface01/sepeda"
)

func main() {

	SepedaOntel := [...]sepeda.Sepeda{
		sepeda.Sepeda{2, 2, "Merah", 20},
		sepeda.Sepeda{2, 4, "Hijau", 40},
		sepeda.Sepeda{2, 4, "Biru", 60},
		sepeda.Sepeda{2, 2, "Pink", 60},
		sepeda.Sepeda{2, 2, "Kuning", 20},
	}

	for index, item := range SepedaOntel {
		fmt.Println("Sepeda ke : ", index+1)
		fmt.Println("Jumlah Gigi: ", item.Gigi)
		fmt.Println("Jumlah Ban: ", item.Ban)
		fmt.Println("Jenis Warna: ", item.Warna)
		fmt.Println("Waktu Tempuh: ", item.Menit)
		sepeda.Hitung(item)
		fmt.Println("--------------")

	}

	fmt.Println("")

	Motors := [...]motor.Motor{
		motor.Motor{2, 2, "Merah", 20},
		motor.Motor{2, 4, "Hijau", 40},
		motor.Motor{2, 4, "Biru", 60},
		motor.Motor{2, 2, "Pink", 60},
		motor.Motor{2, 2, "Kuning", 20},
	}

	for index, item := range Motors {
		fmt.Println("Motor ke : ", index+1)
		fmt.Println("Jumlah Gigi: ", item.Gigi)
		fmt.Println("Jumlah Ban: ", item.Ban)
		fmt.Println("Jenis Warna: ", item.Warna)
		fmt.Println("Waktu Tempuh: ", item.Menit)
		sepeda.Hitung(item)
		fmt.Println("--------------")

	}

}
