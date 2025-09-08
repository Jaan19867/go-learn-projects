package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
) 
type Todo struct {
	// hi

	Title string `json:"id"`

	Completed bool  `json:"description"`

	CreatedAt time.Time  `json: "created_at"`
	
	CompletedAt *time.Time `json: "completed_at ,omitempty" `
}
// variable lenght unlike array 

type TaskManager struct{
	Tasks []Todo `json:"tasks"`
	FilePath string `json:"-`
}
type Todos []Todo



var todos  Todos


var (
	taskFile string 
	taskManager *TaskManager
)


func(todos *Todos) ValidateIndex(index int ) error {


 if(index<0 || index>len(*todos)){
	err:=errors.New("Invalid index")
	fmt.Println(err)
	return err 
 }
 return nil 


}

// To add Todos 

// func(todos *Todos )  add(title string){
// 	todo :=Todo{
// 		Title : title , 
// 		Completed : false , 
// 		CompletedAt : nil , 
// 		CreatedAt : time.Now() , 
// 	}
// 	*todos=append(*todos , todo) 


// }
func loadTasks(cmd *cobra.Command , args []string){
	file :=getTaskFile() 
	taskManager=&TaskManager{
		Tasks : make([]Todo , 0) , 
		FilePath: file , 
	}
	if data , err :=os.ReadFile(file) ; err==nil {
		json.Unmarshal(data , taskManager)
	}
}

func saveTasks(cmd *cobra.Command , args []string){
	data , err :=json.MarshalIndent(taskManager , "" , " ")
	if err!=nil {
		fmt.Printf("Error saving tasks : %v\n " , err)
		return 
	}

	os.WriteFile(taskManager.FilePath , data , 0644)
}

func AddTodo(cmd* cobra.Command , args []string){

	

	if len(args)==0{
		fmt.Println("Please provide a todo title ")
		return 


	}
	

	title :=strings.Join(args , "")
	todo:=Todo{
		Title : title , 
		Completed: false,
		CreatedAt: time.Now(),

	}
	taskManager.Tasks=append(taskManager.Tasks , todo )

	fmt.Println("Added todo: ", todo)
	fmt.Println(taskManager.FilePath)
	fmt.Println("full taks is ", taskManager.Tasks)
}


// To delete Todo's 


func(todos *Todos) delete(index int ) error {
	t:=*todos 
	if err :=t.ValidateIndex(index); err !=nil {
		return err 
	}
   // ... is important spread operator this will append 
   // element one by one , as second arg should be element 
	*todos= append(t[:index] , t[index+1:]...)
	return nil 

}


// Toggle , agar complete ho rakha he to incomplete karna he 
// aur complete nhi he to complete kar dena he 

func(todos *Todos ) toggle(index int ) error {

	t:=*todos 
	todo :=&(*todos)[index] 
	// so todo is a pointer 
    if err := t.ValidateIndex(index); err != nil {
    return err
    }


	// ab dekhna he ki complete he ki nahi 
	if todo.Completed {
		// completed he isko not complete karna he 
        todo.Completed=false 
		todo.CompletedAt=nil 
	}else{
		now :=time.Now()
		todo.CompletedAt = &now
		todo.Completed=true 
	}
	return nil 
}