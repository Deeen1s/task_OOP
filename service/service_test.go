package service_test

import (
	"stepic-go-basic/OOP/service"
	"stepic-go-basic/OOP/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) { //тест без ошибок

	mockProducer := new(mocks.Producer)   // Создаем экземпляр mockProducer
	mockPresenter := new(mocks.Presenter) // Создаем экземпляр mockPresenter

	// Определяем, что метод  должны возвращать
	mockProducer.On("Produce").Return([]string{"http://mail.ru", "www.wolf.com"}, nil)
	mockPresenter.On("Present", []string{"http://*******", "www.wolf.com"}).Return(nil)

	mockServ := service.NewService(mockProducer, mockPresenter)
	mockRun := mockServ.Run()

	//Используем assert для проверки значений

	assert.NotNil(t, mockServ)
	assert.Nil(t, mockRun)
	assert.Empty(t, mockRun)
	assert.NoError(t, mockRun, mockServ)
	assert.Equal(t, mockRun, nil)

}
