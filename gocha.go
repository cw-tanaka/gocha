package gocha

import (
    "fmt"
)

type ChatworkClient struct {
    token string
}

func (client ChatworkClient) GetRooms() []*Room{
    fmt.Println("called GetRooms")
    return nil
}

func (client ChatworkClient) GetRoom(room_id int) *Room {
    fmt.Println("called GetRoom")
    return new(Room)
}

func (client ChatworkClient) GetTasks(room_id int) []*Task {
    fmt.Println("called GetTasks")
    return nil
}


func (client ChatworkClient) GetTask(room_id int, task_id int) *Task {
    fmt.Println("called GetTask")
    return new(Task)
}

func (client ChatworkClient) GetFiles(room_id int) []*File {
    fmt.Println("called GetFiles")
    return nil
}

func (client ChatworkClient) GetFile(room_id int, file_id int) *File {
    fmt.Println("called GetFile")
    return new(File)
}

func (client ChatworkClient) SendTask(task *Task) {
    fmt.Println("called sendTask")
}

func (client ChatworkClient) SendMessage(message *Message) {
    fmt.Println("called sendMessage")
}

func NewClient(options map[string]string) *ChatworkClient {
    return &ChatworkClient{
        token: options["token"],
    }
}
