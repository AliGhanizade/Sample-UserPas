package model

import config "userpas/config"

type User struct {
	Id       string `json:"id,omitempty" gorm:"primary_key" `
	Username string `json:"username,omitempty" gorm:"index" `
	Rulse    int16  `json:"rulse,omitempty" `
	Password string `json:"password,omitempty"`
}

func (m *User) TableName() string {
	return "User_table"
}

type Users []User

func (m *User) Get() (err error) {
	if err = config.DB.Find(m).Error; err != nil {
		return err
	}
	return nil
}
func (m *Users) GetAdmin(Rulse int16) (err error) {
	if err = config.DB.Model(&m).Where("rulse = ?", Rulse).Find(m).Error; err != nil {
		return err
	}
	return nil
}
