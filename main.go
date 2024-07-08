package main

import (
	"fmt"
	"os"
	"stepic-go-basic/OOP/service"
)

func main() {
	if len(os.Args) < 2 { //проверка наличия командной строки на количество введенной в ней аргументов
		fmt.Printf("Используйте go run main.go *название входного файла* *название выходного файла*")
		return
	}
	inFile := os.Args[1]    //путь к файлу, которые введены в командной строке
	outFile := "output.txt" //наименование выходного файла
	if len(os.Args) >= 3 {
		outFile = os.Args[2]
	}
	//передал данные в наш сервис
	prod := service.NewFileProducer(inFile)
	pres := service.NewFilePresenter(outFile)

	serv := service.NewService(prod, pres) //создание сервиса с маскировкой в файл

	if err := serv.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err) //вывод ошибки
		os.Exit(1)                                  //завершение с кодом ошибки
	}

	fmt.Println("Выполнено!\nВыходной файл:", outFile)

}
