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
