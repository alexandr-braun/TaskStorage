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
