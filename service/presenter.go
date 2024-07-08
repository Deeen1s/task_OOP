package service

import (
	"os"
)

//реализуем перезапись файла

type FilePresenter struct {
	inFile string
}

func NewFilePresenter(inFile string) *FilePresenter { //конструктор
	return &FilePresenter{inFile: inFile}
}

func (s *FilePresenter) Present(newText []string) error { //перезапись файла

	file, err := os.Create(s.inFile) //открываем файл

	if err != nil { //возвращаем ошибку, если не удалось создать файл
		return err
	}

	defer file.Close() // закрываем файл в конце программы

	for _, str := range newText { //возвращает ошибку, если не записалась строка
		if _, err := file.WriteString(str + "\n"); err != nil {
			return err
		}
	}
	return nil // возврат нулевого значения ошибки
}
