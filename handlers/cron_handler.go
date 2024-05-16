package handlers

import (
	"academ_be/services"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CronJobHander(cron *gocron.Scheduler) {

	// 10 am every days
	cron.CronWithSeconds("0 0 10  * * *").Do(func() {

		projects, err := services.GetAllProjectCron()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, mp := range projects {

			if !mp.IsArchive && time.Until(*mp.ProjectEndDate) <= 72*time.Hour && time.Until(*mp.ProjectEndDate) > 48*time.Hour {

				var ownerId primitive.ObjectID = primitive.NilObjectID
				for _, r := range mp.Roles {
					if r.RoleName == ROLE_DEFAULT_OWNER {
						ownerId = r.RoleId
					}
				}

				for _, m := range mp.Members {
					if m.RoleId == ownerId && ownerId != primitive.NilObjectID {

						sendNotificationByUserId(m.UserId,
							&mp.ProjectProfile,
							NOTI_HEADER_PROJECT_DEADLINE,
							NOTI_BODY_PROJECT_DEADLINE)

						invitationToken := generateInvitationToken()
						sendInvite(m.Emaill, mp.ProjectProfile.ProjectName, invitationToken)

					}
				}

			}
		}

	})

	cron.StartAsync()
}
