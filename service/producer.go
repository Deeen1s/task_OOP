package service

import (
	"bufio"
	"os"
)

//реализуем чтение файла

type FileRead struct {
	inFile string
}

func NewFileRead(inFile string) *FileRead { //конструктор
	return &FileRead{inFile: inFile}
}

func (f *FileRead) Produce() ([]string, error) { //метод реализации чтения данных из файла

	file, err := os.Open(f.inFile) //открываем файл

	if err != nil { // смотрим нет ли ошибок, если не удалось открыть файл, возвращаем ошибку
		return nil, err
	}
	defer file.Close() // закрываем файл в конце программы

	var text []string
	scan := bufio.NewScanner(file) //чтение из файла

	if err := scan.Err(); err != nil { //возврат ошибки, в случае, если текст какой-то не такой
		return nil, err
	}

	for scan.Scan() { //добавляем маскированную строку с помощью методов Scan() и Text() в наш срез данных
		text = append(text, scan.Text())
	}

	return text, nil //возврат текста, нулевого значения ошибки//
}
