package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Duration  string `json:"time"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Duration: "10 min", Completed: false},
	{ID: "2", Item: "Brush Teeth", Duration: "10 min", Completed: false},
	{ID: "3", Item: "Go For a Walk", Duration: "15 min", Completed: false},
	{ID: "4", Item: "Breakfast", Duration: "20 min", Completed: false},
	{ID: "5", Item: "Study", Duration: "1 hour", Completed: false},
	{ID: "6", Item: "Attend Meeting", Duration: "30 min", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IntentedJSON(http.StatusOK, todos) //convert structure to json
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")

}
