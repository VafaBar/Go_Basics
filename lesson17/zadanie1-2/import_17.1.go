package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.
		Request) {
		w.Write([]byte("Hello, guys!"))
	}).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}

/*
1. Создание файла go.mod : Создайте новый проект и инициализируйте файл
go.mod с помощью команды go mod init <your-project-name> . Убедитесь, что
файл go.mod был создан и содержит информацию о вашем проекте.

2. Импорт внешнего пакета: Установите пакет github.com/gorilla/mux с
помощью команды go get github.com/gorilla/mux . Создайте HTTP сервер с
использованием этого пакета, который будет возвращать приветственное
сообщение при обращении к корневому URL.*/
