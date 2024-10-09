package utils

import (
	"log/slog"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func StartCronJobs() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Error("Failed to load location")
		return
	}

	ns, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		slog.Error("Failed to create scheduler")
		return
	}

	ns.Start()

	var user model.User

	if err := db.DB.First(&user, "user_id = ?", "your-user-id").Error; err != nil {
		slog.Error("Failed to retrieve user from database:")
		return
	}

	date := time.Now().Format("2006-01-02")
	accounts, err := GetAccounts(user.UserID)
	if err != nil {
		slog.Error("Failed to retrieve accounts from database")
		return
	}
	_, err = ns.NewJob(
		gocron.CronJob("0 0 0 * * *", false),
		gocron.NewTask(notes.GenerateNote(user.UserID, date, *accounts)),
	)
	if err != nil {
		slog.Error("Failed to create job")
		return
	}
}
