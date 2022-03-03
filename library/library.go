package library
import "fmt"

type Mahasiswa struct{
  nama string
  umur uint8
}
func Apaankek(){
  mahasiswa := Mahasiswa{
    nama :"Faldy",
    umur : 24,
  }
  mahasiswa1 := mahasiswa
  fmt.Println(mahasiswa1)
}
