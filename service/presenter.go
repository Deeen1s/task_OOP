package service

import (
	"os"
)

//реализуем перезапись файла

type FileWrite struct {
	inFile string
}

func NewFileWrite(inFile string) *FileWrite { //конструктор
	return &FileWrite{inFile: inFile}
}

func (s FileWrite) Present(newText []string) error {

	file, err := os.Create(s.inFile) //открываем файл

	if err != nil { //возвращаем ошибку, если не удалось создать файл
		return err
	}
	defer file.Close() //закрываем файл в конце программы

	for _, str := range newText { //возвращает ошибку, если не записалась строка
		if _, err := file.WriteString(str + "\n"); err != nil {
			return err
		}
	}

	return nil

	//почти то же самое, только преобразовываем файл
	//и записываем либо в тот же файл, либо в новый
	//принимает сообщения(messege) и декодирует их
	//тоже проверяем на nil, будем возвращать nil
	//тоже возвращать
}
