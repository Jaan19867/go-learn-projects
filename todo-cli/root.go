package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)



var rootCmd=&cobra.Command{

	Use:"taskman" , 
	Short : "A personal task manager" , 
	PersistentPreRun : loadTasks , 
	PersistentPostRun: saveTasks,
	Long : ` taskman is a CLI task manager that helps you organize your work.

Store tasks locally with priorities, mark them complete, and keep
track of your productivity over time.`,
    Run : func(cmd *cobra.Command , args []string){
		fmt.Println("Welcome to Taskman ! greated code ")
	} , 
}

func getTaskFile() string {
	if(taskFile!=""){
		return taskFile
	}
	return viper.GetString("file")
}



var greetCmd= &cobra.Command{
	Use:"hi" , 
	Short :"greet hte terminal " , 
	Run : func(cmd *cobra.Command , args []string){
		fmt.Println("Hello World to terminal ")
	},
}

var addCmd =&cobra.Command{
	Use:"add" , 
	Short : "Add a todo in the list" , 
	Run : AddTodo,
}



func init(){
	rootCmd.PersistentFlags().StringVar(&taskFile , "file", "" , "task file defautl is something else")
	rootCmd.AddCommand(greetCmd , addCmd)
	setupConfig()
}

func setupConfig(){
	viper.SetConfigName("taskman")
	viper.SetConfigType("json")
	viper.AddConfigPath("/home/riz/Desktop/Golang-Projects/todo-cli/")
	viper.SetDefault("file", filepath.Join(os.Getenv("HOME") , ".taskman.json"))
	viper.ReadInConfig()
}


func Execute(){
	if err :=rootCmd.Execute() ; err!=nil {
		fmt.Println(err)

	}
}


