package task

import (
	"taskStorage/pkg/domain/user"
	"testing"
	"time"
)

func TestAppendComment(t *testing.T) {
	testUser := &user.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	testTask := &Task{
		Id:               1,
		Title:            "Test Task",
		Description:      "This is a test task",
		Priority:         1,
		Status:           New,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		Deadline:         nil,
		AssignedToUserId: testUser.Id,
		Categories:       []string{"testing", "example"},
		Comments:         []Comment{},
	}

	commentText := "This is a test comment."

	testTask.AppendComment(commentText, testUser)

	if len(testTask.Comments) != 1 {
		t.Errorf("expected 1 comment, got %d", len(testTask.Comments))
	}

	lastComment := testTask.Comments[0]
	if lastComment.UserId != testUser.Id || lastComment.Text != commentText {
		t.Errorf("appended comment does not match, got %+v", lastComment)
	}

	if lastComment.Timestamp.After(time.Now()) {
		t.Errorf("timestamp of the comment is in the future")
	}
}

func TestChangeStatus_ValidTransitions(t *testing.T) {
	tests := []struct {
		initialStatus TaskStatus
		newStatus     TaskStatus
		expectedError bool
	}{
		{New, InProgress, false},
		{InProgress, Completed, false},
	}

	for _, tt := range tests {
		task := Task{
			Id:        1,
			Title:     "Sample Task",
			Status:    tt.initialStatus,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := task.ChangeStatus(tt.newStatus)

		if (err != nil) != tt.expectedError {
			t.Errorf("ChangeStatus(%v): expected error: %v, got: %v", tt.newStatus, tt.expectedError, err)
		}

		if err == nil && task.Status != tt.newStatus {
			t.Errorf("ChangeStatus(%v): expected status: %v, got: %v", tt.newStatus, tt.newStatus, task.Status)
		}
	}
}

func TestChangeStatus_InvalidTransitions(t *testing.T) {
	tests := []struct {
		initialStatus TaskStatus
		newStatus     TaskStatus
	}{
		{New, Completed},
		{InProgress, New},
		{Completed, InProgress},
		{Completed, New},
	}

	for _, tt := range tests {
		task := Task{
			Id:        1,
			Title:     "Sample Task",
			Status:    tt.initialStatus,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := task.ChangeStatus(tt.newStatus)

		if err == nil {
			t.Errorf("ChangeStatus(%v) from %v: expected an error, got nil", tt.newStatus, tt.initialStatus)
		}
	}
}
