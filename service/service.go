package service

import (
	"fmt"
)

type Producer interface { // описывает метод со входными данными
	Produce() ([]string, error) //принимает срез строк  и возвращает ошибку
}

type Presenter interface { // описывает метод с перезаписью данных
	Present([]string) error //перезаписывает и возвращает ошибку
}
type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {

	return &Service{prod: prod, pres: pres}

}

func (s *Service) Mask(messege string) string { //функция

	sliceWeb := []byte(messege)    //перевод ссылки в байты
	sliceHttp := []byte("http://") //перевод  в байты http://

	for i := 0; i < len(messege); {
		if i <= len(messege)-len(sliceHttp) && string(sliceWeb[i:i+len(sliceHttp)]) == "http://" { //проверяем элементы на схожесть с http://
			i += len(sliceHttp)
			for i < len(messege) && sliceWeb[i] != ' ' { // маскировка ссылки до пробела
				sliceWeb[i] = '*'
				i++
			}
		} else {
			i++
		}
	}
	return string(sliceWeb)
}

func (s *Service) Run() error {

	newText, err := s.prod.Produce() //чтение файла через продюсера
	if err != nil {
		return fmt.Errorf("ERROR FILE %v", err) //возврат ошибки, если не прочитался файл
	}
	//fmt.Println("Сам текст", newText)
	for ind, str := range newText {
		newText[ind] = s.Mask(str) //маскировка каждой строки
	}

	if err := s.pres.Present(newText); err != nil {
		return fmt.Errorf("Ошибка данных: %v", err) //возвращает ошибку и выводит на экран(если не удалось записать)
	}
	return nil //если успешно все проходит, то возвращает nil
}
