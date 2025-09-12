package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo/backend/controller"
)

func SetupRouter(todoController *controller.TodoController) *mux.Router {

	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// 待办事项路由
	todoRoutes := r.PathPrefix("/api/todos").Subrouter()
	todoRoutes.HandleFunc("", todoController.GetAllTodos).Methods("GET")
	todoRoutes.HandleFunc("/{id:[0-9]+}", todoController.GetTodoByID).Methods("GET")
	todoRoutes.HandleFunc("", todoController.CreateTodo).Methods("POST")
	todoRoutes.HandleFunc("/{id:[0-9]+}", todoController.UpdateTodo).Methods("PUT")
	todoRoutes.HandleFunc("/{id:[0-9]+}/toggle", todoController.ToggleTodoCompletion).Methods("PATCH")
	todoRoutes.HandleFunc("/{id:[0-9]+}", todoController.DeleteTodo).Methods("DELETE")

	return r
}
