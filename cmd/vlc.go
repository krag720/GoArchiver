package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Файл использует variable-lenght code",
	Run:   pack,
}

const PACKEDEXTENSION = "vlc"

// пользователь указывает местоположение до файла
func pack(_ *cobra.Command, args []string) {

	filePlace := args[0]

	//обработка ошибки
	r, err := os.Open(filePlace)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//cчитываем в переменную
	data, err := ioutil.ReadAll(r)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//packed :=Encode(data)
	packed := "" + string(data)
	//записываем результат в отдельный файл, содержимое файла является слайсом из байт
	//права позволяют 0644 позволяют пользователю записывать и читать файл. остальные могут лишь читать
	err = ioutil.WriteFile(packedFileName(filePlace), []byte(packed), 0644)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// даем название выходящему файлу исходя из его пути
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	//получаем расширение файла
	ext := filepath.Ext(fileName)
	//вырезаем расширение из имени файла
	baseName := strings.TrimSuffix(fileName, ext)
	//устнавливаем свое  расширение
	return baseName + "." + PACKEDEXTENSION
}

func init() {
	packCmd.AddCommand(vlcCmd)

}
