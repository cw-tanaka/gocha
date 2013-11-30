package gocha

type Task struct {
    id     int
    to_ids []int
    body   string
    room   *Room
}
