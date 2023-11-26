package task

import (
	"taskStorage/pkg/domain/user"
	"time"
)

type TaskStatus string
type PriorityLevel int

const (
	New        TaskStatus = "New"
	InProgress TaskStatus = "InProgress"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	Id               int32
	Title            string
	Description      string
	Priority         PriorityLevel
	Status           TaskStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Deadline         *time.Time
	AssignedToUserId int32
	Categories       []string
	Comments         []Comment
}

type Comment struct {
	ID        int32
	UserId    int32
	Text      string
	Timestamp time.Time
}

func (task *Task) AppendComment(comment string, user *user.User) {
	task.Comments = append(task.Comments, Comment{UserId: user.Id, Text: comment, Timestamp: time.Now()})
}
