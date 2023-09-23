package user

import "time"

type User struct {
	Id                int64     `bson:"id"`
	Username          string    `bson:"username"`
	Email             string    `bson:"email"`
	RegistrationDate  time.Time `bson:"registrationDate"`
	SubExpirationDate time.Time `bson:"subExpirationDate"`
	IsSubscribed      bool      `bson:"isSubscribed"`
	WarningsCounter   uint8     `bson:"warningsCounter"` // 3 warning = permanent penalty
}

const MaxWarnings = 3
const YearInHours = 8760

func (u *User) ShouldBan() bool {
	return u.WarningsCounter >= MaxWarnings
}

func (u *User) IsSubActive() bool {
	return u.SubExpirationDate.After(time.Now())
}

func (u *User) ShouldHaveBonus() bool {
	sub := time.Now().Sub(u.RegistrationDate)
	return sub.Hours() >= YearInHours
}
