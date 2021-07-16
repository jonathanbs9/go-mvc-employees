package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/jonathanbs9/go-mvc-employees/src/webapp/controllers"
)

func BootApplication() {
	fmt.Println("Boot APP init")
	http.HandleFunc("/employees", controllers.GetEmployee)
	http.ListenAndServe(":8080", nil)
}
