package task

import "fmt"

type Task struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t Task) String() string {
	var isDoneEmoji string
	if t.Done {
		isDoneEmoji = "✔️"
	} else {
		isDoneEmoji = "❌"
	}

	return fmt.Sprintf("%-4d | %-30s | %s", t.ID, t.Title, isDoneEmoji)
}
