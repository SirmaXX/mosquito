package lib

import (
	"fmt"
	"os"
	
)

func CreateFile(p string) *os.File {

	fmt.Println("dosya  olusturuluyor")
	output, err := os.Create(p)
	if err != nil {
		fmt.Println("bu dosya olusturulamıyor", p, "dosya adı")
	}
	return output
}
func CreateFolder(p string) *os.File {

	fmt.Println("dosya  olusturuluyor")
	output, err := os.Mkdir(p)
	if err != nil {
		fmt.Println("bu dosya olusturulamıyor", p, "dosya adı")
	}
	return output
}

func DeleteFile(path string) string {
	fmt.Println("dosya",path, "siliniyor")
	err := os.Remove(path)
	if err != nil {
		fmt.Println("dosya silinemedi")
	}
	return "dosya silindi"
}
