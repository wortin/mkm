package fyin

type Shm string

const (
	ShmB  Shm = "b"
	ShmP  Shm = "p"
	ShmM  Shm = "m"
	ShmF  Shm = "f"
	ShmD  Shm = "d"
	ShmT  Shm = "t"
	ShmN  Shm = "n"
	ShmL  Shm = "l"
	ShmG  Shm = "g"
	ShmK  Shm = "k"
	ShmH  Shm = "h"
	ShmJ  Shm = "j"
	ShmQ  Shm = "q"
	ShmX  Shm = "x"
	ShmZh Shm = "zh"
	ShmCh Shm = "ch"
	ShmSh Shm = "sh"
	ShmR  Shm = "r"
	ShmZ  Shm = "z"
	ShmC  Shm = "c"
	ShmS  Shm = "s"
	ShmY  Shm = "y"
	SHmW  Shm = "w"
)

var ShmChar = [23]Shm{ShmB, ShmP, ShmM, ShmF, ShmD, ShmT, ShmN, ShmL, ShmG, ShmK, ShmH, ShmJ, ShmQ, ShmX, ShmZh, ShmCh, ShmSh, ShmR, ShmZ, ShmC, ShmS, ShmY, SHmW}

func (s Shm) IsJi(o Shm) bool {
	return false
}
