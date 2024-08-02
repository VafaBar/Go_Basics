package main

// Шаг 1.
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Шаг 2.
type Hello struct {
	Greeting string `json:"greeting"`
}

// Шаг 3.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	h := Hello{
		Greeting: "Hello World!",
	}
	helloJSON, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(helloJSON)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)

	var decodedHello Hello
	err = json.Unmarshal(helloJSON, &decodedHello)
	if err != nil {
		// Если задать пустую структуру Hello и переменной задать значение h:=0, то
		// ошибка сработает. Не знаю насколько это все не бред, но я так вижу))
		// Я выбрала код 500.
		log.Fatalf("Error happened in JSON unmarshal: %d", http.StatusInternalServerError)
		return
	}
	fmt.Println("Десериализованная структура:", decodedHello)
}

// Шаг 4.
func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 3. Обработка ошибок: Добавьте обработку ошибок в ваш сервер. Например,
// если обработчик не может декодировать JSON, он должен возвращать
// соответствующий код ошибки и сообщение. Подумайте, какой код лучше
// всего подходит под эту ситуацию.
