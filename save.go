package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type q struct {
	A string
	Z string
	S string
}

func main() {
	obj := q{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//пишем чтобы делать запросы из браузера
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//получаем данные с фронта
		markCar := r.URL.Query().Get("mark")
		madelCar := r.URL.Query().Get("model")
		photoCar := r.URL.Query().Get("photo")

		//fmt.Println(json1)

		obj = q{markCar, madelCar, photoCar}
		fmt.Println(obj)

		if obj.A != "" {
			dataFile, _ := ioutil.ReadFile("data.json")
			mas := []q{}
			mas = append(mas, obj)

			fmt.Println(mas)
			json.Unmarshal(dataFile, &mas)
			mas = append(mas, obj)
			jsonData, _ := json.Marshal(mas)
			jsonDataString := string(jsonData)

			//открываем файл для работы с ним
			file, _ := os.Create("data.json")
			defer file.Close()
			//записываем в файл данные
			file.WriteString(jsonDataString)

			fmt.Fprintf(w, jsonDataString)

			//fmt.Println(jsonDataString)
			//fmt.Println(jsonData)
		}
	})

	http.ListenAndServe(":8081", nil)

}
