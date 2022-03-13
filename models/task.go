package models

import (
	"todorestapi/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	//Id          string `json:"id"`
	Title        string `gorm:""json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
}


func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (b *Task) CreateTask() *Task {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)

	return Tasks
}

func GetTaskById(Id int64) (*Task , *gorm.DB){
	var getTask Task
	db:=db.Where("ID = ?", Id).Find(&getTask)

	return &getTask, db
}


func DeleteTask(ID int64) Task {
	var task Task
	db.Where("ID = ?", ID).Delete(task)
	
	return task
}