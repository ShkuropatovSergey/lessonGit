package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//разрешаем запросы с любого адреса - пишем нужные заголовки для ответа
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// main, err := ioutil.ReadFile("data.json")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println(main)
		// fmt.Println(string(main))
		// fmt.Fprintf(w, string(main))
	})
	http.ListenAndServe(":8080", nil)
}
