package motor

import (
	"fmt"

	"github.com/bayuiqballl/package-interface01/contract"
)

type Motor struct {
	Ban, Gigi int
	Warna     string
	Menit     float32
}

func (m Motor) Cepat() float32 {

	waktu := m.Menit * 5 * 2.0

	return waktu
}

func (m Motor) Lambat() float32 {

	waktu := m.Menit * 5 * 0.5

	return waktu
}

func Hitung(b contract.Maju) {
	fmt.Println("cepat :", b.Cepat())
	fmt.Println("lambat :", b.Lambat())
}

func Motors() {
	Motors := [...]Motor{
		Motor{2, 2, "Merah", 20},
		Motor{2, 4, "Hijau", 40},
		Motor{2, 4, "Biru", 60},
		Motor{2, 2, "Pink", 60},
		Motor{2, 2, "Kuning", 20},
	}

	for index, item := range Motors {
		fmt.Println("--------------")
		fmt.Println("Motor ke : ", index+1)
		fmt.Println("Jumlah Gigi: ", item.Gigi)
		fmt.Println("Jumlah Ban: ", item.Ban)
		fmt.Println("Jenis Warna: ", item.Warna)
		fmt.Println("Waktu Tempuh: ", item.Menit)
		Hitung(item)
	}
}
