package main

type mobil struct {
	totalRoda       int
	kecepatanPerJam int
}

func (mobil *mobil) tambahKecepatan(kecepatanBaru int) {
	mobil.kecepatanPerJam += kecepatanBaru
}

func (mobil *mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func main() {
	mobilCepat := mobil{}
	mobilLamban := mobil{}

	mobilCepat.berjalan()
	mobilLamban.berjalan()
}
