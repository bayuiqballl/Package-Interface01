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
