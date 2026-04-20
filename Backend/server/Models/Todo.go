//Models/User.go

package Models

import (
	"fmt"
	"todo-apps/backend/server/Config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//GetAllTodos Fetch all todo data

func GetAllTodos(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//CreateTodo ... Insert New data

func CreateTodo(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//GetTodoByID ... Fetch only one todo by Id

func GetTodoByID(todo *Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//UpdateTodo ... Update todo

func UpdateTodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

//DeleteTodo ... Delete todo

func DeleteTodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
