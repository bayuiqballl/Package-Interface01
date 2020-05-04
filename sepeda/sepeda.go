package sepeda

import (
	"fmt"

	"github.com/bayuiqballl/package-interface01/contract"
)

type Sepeda struct {
	Ban, Gigi int
	Warna     string
	Menit     float32
}

func (s Sepeda) Cepat() float32 {

	waktu := s.Menit * 2.5 * 2.0

	return waktu
}

func (s Sepeda) Lambat() float32 {

	waktu := s.Menit * 2.5 * 0.5

	return waktu
}

func Hitung(b contract.Maju) {
	fmt.Println("cepat :", b.Cepat())
	fmt.Println("lambat :", b.Lambat())
}

func Ontel() {
	SepedaOntel := [...]Sepeda{
		Sepeda{2, 2, "Merah", 20},
		Sepeda{2, 4, "Hijau", 40},
		Sepeda{2, 4, "Biru", 60},
		Sepeda{2, 2, "Pink", 60},
		Sepeda{2, 2, "Kuning", 20},
	}

	for index, item := range SepedaOntel {
		fmt.Println("--------------")
		fmt.Println("Sepeda ke : ", index+1)
		fmt.Println("Jumlah Gigi: ", item.Gigi)
		fmt.Println("Jumlah Ban: ", item.Ban)
		fmt.Println("Jenis Warna: ", item.Warna)
		fmt.Println("Waktu Tempuh: ", item.Menit)
		Hitung(item)

	}
}
