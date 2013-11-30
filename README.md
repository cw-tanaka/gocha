gocha
=====

Chatwork client with go lang

It is currently in development, not working now.

現在開発中につき、まだ動きません


### Install

```bash
go get github.com/tan-yuki/gocha
```

### Sample

```go
package main


import (
    "github.com/tan-yuki/gocha"
)

func main() {
    client := gocha.NewClient(map[string]string {
        "token": "YOUR API TOKEN",
    })

    // GET /rooms
    client.GetRooms()

    // GET /rooms/{id}
    client.GetRoom(room_id)

    // GET /rooms/{id}/tasks
    client.GetTasks(room_id)

    // GET /rooms/{id}/task/{id}
    client.GetTask(room_id, task_id)

    // GET /rooms/{id}/files
    client.GetFiles(room_id)

    // GET /rooms/{id}/files/{id}
    client.GetFile(room_id, file_id)

    // POST /rooms/{id}/tasks
    client.SendTask(new(gocha.Task))

    // POST /rooms/{id}/messages
    client.SendMessage(new(gocha.Message))
}
```
