package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FernandoMendoza12/godesde0/ejercisios"
)

var fileName string = "./files/txt/tabla.txt"

func GuardarTabla() {
	var texto string = ejercisios.CapturarDatos()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo txt" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercisios.CapturarDatos()
	if !Append(fileName, texto) {
		fmt.Println("Erro al concatenar el contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro al concatenar el archivo" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteStringAppend" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("El archivo no se puede leer" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()

}
