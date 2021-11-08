package sql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID    int64
	UID int64
	Count int64
}

func (User) TableName() string {
	return "users"
}

type Envelope struct {
	ID         int64 `json:"envelope_id"`
	UID        int64 `json:"uid"`
	Opened     bool  `json:"opened"`
	Value      int64 `json:"value"`
	SnatchTime int64 `json:"snatch_time"`
}

func (Envelope) TableName() string {
	return "envelopes"
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:3306@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.AutoMigrate(&User{}, &Envelope{})
		return db, err
	}
	return nil, err
}

//User
func GetUser(uid int64) (user User) {
	DB.FirstOrCreate(&user, User{UID: uid})
	return
}

func UpdateCount(user *User) {
	user.Count++
	DB.Model(&user).Update("count", user.Count)
}

//Envelope
func GetAllEnvelopesByUID(uid int64) ([]*Envelope, error) {

	var envelopes []*Envelope
	conditions := map[string]interface{}{
		"uid": uid,
	}
	if err := DB.Table(Envelope{}.TableName()).Where(conditions).Find(&envelopes).Error; err != nil {
		return nil, err
	}
	return envelopes, nil
}

func GetEnvelopeByEnvelopeID(envelope_id int64) (envelope Envelope) {
	DB.Where("id = ?", envelope_id).First(&envelope)
	return
}

func CreateEnvelope(user User) (envelope Envelope) {

	snatchTime := time.Now().UnixNano()
	var a int64 = 10
	envelope = Envelope{UID: user.UID, Opened: false, Value: a, SnatchTime: snatchTime}
	DB.Create(&envelope)
	return envelope
}

func UpdateState(envelopeId int64) (envelope Envelope) {

	//查询条件
	envelope.ID = envelopeId
	DB.Model(&envelope).Update("opened", true)
	return
}