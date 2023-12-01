package task

import (
	"fmt"
	"taskStorage/pkg/domain/user"
	"time"
)

type Status string
type PriorityLevel int

const (
	New        Status = "New"
	InProgress Status = "InProgress"
	Completed  Status = "Completed"
)

type Task struct {
	Id               int32
	Title            string
	Description      string
	Priority         PriorityLevel
	Status           Status
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

func (task *Task) ChangeStatus(newStatus Status) error {
	switch task.Status {
	case New:
		if newStatus == InProgress {
			task.Status = newStatus
		} else {
			return fmt.Errorf("invalid status transition from %s to %s", task.Status, newStatus)
		}
	case InProgress:
		if newStatus == Completed {
			task.Status = newStatus
		} else {
			return fmt.Errorf("invalid status transition from %s to %s", task.Status, newStatus)
		}
	default:
		return fmt.Errorf("invalid status transition from %s", task.Status)
	}
	return nil
}
