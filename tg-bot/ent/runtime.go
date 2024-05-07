// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Simplewallethq/tg-bot/ent/rewardsdata"
	"github.com/Simplewallethq/tg-bot/ent/schema"
	"github.com/Simplewallethq/tg-bot/ent/task"
	"github.com/Simplewallethq/tg-bot/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	rewardsdataFields := schema.RewardsData{}.Fields()
	_ = rewardsdataFields
	// rewardsdataDescFirstEra is the schema descriptor for first_era field.
	rewardsdataDescFirstEra := rewardsdataFields[3].Descriptor()
	// rewardsdata.DefaultFirstEra holds the default value on creation for the first_era field.
	rewardsdata.DefaultFirstEra = rewardsdataDescFirstEra.Default.(int64)
	// rewardsdataDescLastEra is the schema descriptor for last_era field.
	rewardsdataDescLastEra := rewardsdataFields[4].Descriptor()
	// rewardsdata.DefaultLastEra holds the default value on creation for the last_era field.
	rewardsdata.DefaultLastEra = rewardsdataDescLastEra.Default.(int64)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[2].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLoggedIn is the schema descriptor for logged_in field.
	userDescLoggedIn := userFields[3].Descriptor()
	// user.DefaultLoggedIn holds the default value on creation for the logged_in field.
	user.DefaultLoggedIn = userDescLoggedIn.Default.(bool)
	// userDescNotify is the schema descriptor for notify field.
	userDescNotify := userFields[7].Descriptor()
	// user.DefaultNotify holds the default value on creation for the notify field.
	user.DefaultNotify = userDescNotify.Default.(bool)
	// userDescNotifyTime is the schema descriptor for notify_time field.
	userDescNotifyTime := userFields[8].Descriptor()
	// user.DefaultNotifyTime holds the default value on creation for the notify_time field.
	user.DefaultNotifyTime = userDescNotifyTime.Default.(int8)
	// userDescNotifyLastTime is the schema descriptor for notify_last_time field.
	userDescNotifyLastTime := userFields[9].Descriptor()
	// user.DefaultNotifyLastTime holds the default value on creation for the notify_last_time field.
	user.DefaultNotifyLastTime = userDescNotifyLastTime.Default.(time.Time)
	// userDescStorePrivatKey is the schema descriptor for store_privat_key field.
	userDescStorePrivatKey := userFields[10].Descriptor()
	// user.DefaultStorePrivatKey holds the default value on creation for the store_privat_key field.
	user.DefaultStorePrivatKey = userDescStorePrivatKey.Default.(bool)
	// userDescEnableLogging is the schema descriptor for enable_logging field.
	userDescEnableLogging := userFields[11].Descriptor()
	// user.DefaultEnableLogging holds the default value on creation for the enable_logging field.
	user.DefaultEnableLogging = userDescEnableLogging.Default.(bool)
	// userDescRegistered is the schema descriptor for registered field.
	userDescRegistered := userFields[12].Descriptor()
	// user.DefaultRegistered holds the default value on creation for the registered field.
	user.DefaultRegistered = userDescRegistered.Default.(bool)
}