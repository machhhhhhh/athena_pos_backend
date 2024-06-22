package middlewares

// import (
// 	"sync"
// 	"time"

// 	"github.com/machhhhhhh/E-Memo/configs"
// 	"github.com/machhhhhhh/E-Memo/models"
// 	"github.com/machhhhhhh/E-Memo/utils"
// 	"gorm.io/gorm"
// )

// type RepresentativeExpire struct {
// 	models.User
// 	Representative models.User `json:"representative,omitempty" gorm:"foreignKey:ID; references:RepresentID"`
// }

// TODO: ***** This is not a library *****
// TODO: Need to understand more about sync_function with Goroutines for wait_group
// TODO: Relate to Cron-Job

// func FocusUserRepresentExpire() {

// 	// Run a periodic task to check and send emails
// 	ticker := time.NewTicker(30 * time.Minute) // Adjust the interval as needed (30 Minute tricker)
// 	defer ticker.Stop()

// 	// Use a wait group to wait for all Goroutines to finish
// 	var waiting_group sync.WaitGroup

// 	for range ticker.C {

// 		now := time.Now()

// 		// Check if the current time is between 00:00 and 01:00
// 		if now.Hour() == 0 && now.Minute() >= 0 && now.Minute() < 60 {

// 			// Find expire_representative
// 			all_user := FindUserWithExpireRepresentative()

// 			if len(all_user) != 0 {
// 				for i := range all_user {

// 					waiting_group.Add(1) // Increment the wait group counter

// 					go func(user RepresentativeExpire) {

// 						utils.SendEmailToOwnerAndRepresentative(
// 							utils.UserEmail{
// 								Name:  user.Name,
// 								Email: user.Email,
// 							},
// 							utils.UserEmail{
// 								Name:  user.Representative.Name,
// 								Email: user.Representative.Name,
// 							},
// 							utils.EmailActionDELETE,
// 							"",
// 						)

// 						defer waiting_group.Done() // Decrement the wait group counter when done

// 					}(all_user[i])

// 				}
// 			}
// 		}

// 	}
// 	waiting_group.Wait()
// }

// func FindUserWithExpireRepresentative() []RepresentativeExpire {

// 	tx := configs.DB.Begin()

// 	var all_user []RepresentativeExpire
// 	if err := configs.DB.
// 		Where("[user].[represent_expire] < GETDATE()").
// 		Where("[user].[represent_expire] IS NOT NULL").
// 		Preload("Representative", func(db *gorm.DB) *gorm.DB {
// 			return db.Unscoped().Select("user_id, user_number, position, name, email")
// 		}).
// 		Find(&all_user).Error; err != nil {
// 		// rollback
// 		tx.Rollback()
// 		panic(err)
// 	}

// 	if len(all_user) != 0 {

// 		if err := tx.
// 			Model(models.User{}).
// 			Where("[user].[represent_expire] < GETDATE()").
// 			Where("[user].[represent_expire] IS NOT NULL").
// 			Updates(map[string]interface{}{"represent_id": nil, "represent_expire": nil}).
// 			Error; err != nil {
// 			// rollback
// 			tx.Rollback()
// 			panic(err)
// 		}

// 	}

// 	tx.Commit()

// 	return all_user
// }
