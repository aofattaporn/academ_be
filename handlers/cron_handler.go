package handlers

import (
	"academ_be/services"
	"fmt"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CronJobHander(cron *gocron.Scheduler) {

	// every 10 for test
	cron.CronWithSeconds("*/10 * * * * *").Do(func() {
		fmt.Println("Test CronJob")
	})

	// "0 0 10  * * *"
	// 10 am every days
	cron.CronWithSeconds("*/10 * * * * *").Do(func() {

		projects, err := services.GetAllProjectCron()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("enter CronWithSeconds")

		fmt.Println(projects)

		// time.Until(*mp.ProjectEndDate) < 24*time.Hour
		for _, mp := range projects {

			fmt.Println(mp.IsArchive)

			if mp.IsArchive == false {

				var ownerId primitive.ObjectID = primitive.NilObjectID
				for _, r := range mp.Roles {
					fmt.Println("found r !!! ")

					if r.RoleName == "Owner" {
						ownerId = r.RoleId
					}
				}

				for _, m := range mp.Members {
					if m.RoleId == ownerId && ownerId != primitive.NilObjectID {

						fmt.Println("found ownerId !!! ")

						// Sent Notification on browser
						sendNotificationByUserIdCron(m.UserId, &mp.ProjectProfile)

						// Sent Notification on email
						// invitationToken := generateInvitationToken()
						// err = sendInvite(m.Emaill, mp.ProjectProfile.ProjectName, invitationToken)
						// if err != nil {
						// 	handleTechnicalError(c, err.Error())
						// 	return
						// }

					}
				}

			}
		}

	})

	cron.StartAsync()
}
