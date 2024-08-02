package main

// Шаг 1.
import (
	"encoding/json"
	"log"
	"net/http"
)

// Шаг 2.
type Hello struct {
	Greeting string `json:"greeting"`
}
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
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
}

// "Назови хоть jopa, все равно будет работать" (с) Edgar Sipki
func jopa(w http.ResponseWriter, r *http.Request) {
	u := User{Name: "Jane", Age: 35}
	userJSON, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(userJSON)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	p := Product{Name: "Doll", Price: 40}
	productJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(productJSON)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}

// Шаг 4.
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/jopa", jopa)
	// это путь /user, замененный шутки ради.
	http.HandleFunc("/product", productHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 1. Создание простого HTTP сервера: Создайте HTTP сервер, который будет
// слушать на порту 8080 и возвращать приветственное сообщение в
// формате JSON при обращении по пути /hello.

// 2. Обработчики маршрутов: Создайте несколько обработчиков маршрутов.
// Например, /user , который будет возвращать информацию о пользователе,
// и /product , который будет возвращать информацию о продукте.
