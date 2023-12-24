package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"io"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld(t *testing.T) {
	// Создаем новый экземпляр веб-сервера с помощью Gin
	router := gin.Default()

	// Определяем маршрут для тестирования
	router.GET(os.Getenv("IP"), func(c *gin.Context) {
		c.JSON(200, "Hello, World!")
	})

	// Запускаем веб-сервер на порту 8000
	router.Run(":8000")

	// Проверяем, что веб-сервер работает правильно
	resp, err := http.Get("http://localhost:8000/driver/get")
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Ошибка в чтении: %s", err)
	}
	fmt.Println(result[len(result)-1])

	// Проверяем, что веб-сервер не выдает никаких ошибок
	err = resp.Request.Context().Err()
	if err != nil {
		t.Errorf("Ошибка в зпросе: %s", err)
	}
}
