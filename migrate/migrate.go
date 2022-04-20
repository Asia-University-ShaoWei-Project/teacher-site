package migrate

import (
	"crypto/rand"
	"strconv"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/mock"
	"teacher-site/pkg/util"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(
		&domain.Auths{},
		&domain.Teachers{},
		&domain.Infos{},
		&domain.InfoBulletinBoards{},
		&domain.Courses{},
		&domain.BulletinBoards{},
		&domain.Slides{},
		&domain.Homeworks{},
	)
	createAll(db)
}
func createAll(db *gorm.DB) {
	currTime := time.Now()
	conf := config.New()
	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
	lastModified := strconv.FormatInt(currTime.Unix(), 10)
	hashPassword, salt := generalHashPassword(mock.UserPassword, conf.Secure)
	data := domain.Auths{
		UserID:       mock.UserID,
		UserPassword: string(hashPassword),
		Salt:         string(salt),
		Token:        token,
		Teacher: domain.Teachers{
			Domain:    "rikki",
			Email:     "rikki@asia.edu.tw",
			NameZH:    "陳瑞奇",
			NameUS:    "Jui-Chi Chen(Rikki)",
			Office:    "HB13",
			Call:      "(04)2332-3456 ext. 20013",
			Education: "國立中興大學資訊科學博士",
			Infos: []domain.Infos{
				{
					LastModified: lastModified,
					BulletinBoards: []domain.InfoBulletinBoards{
						{
							Content: "(第五週)10月11日起恢復實體授課，請配帶口罩，並按教室座位表入座，謝謝。",
						},
						{
							Content: "請益時間(Office Hours): 資訊大樓地下室HB13研究室 TUE(二) 14:10 – 17:00    THU(四) 10:10 – 17:00    Fri.(五) 10:10 – 12:00",
						},
						{
							Content: "9/13(一)開學日、開始上課，9/13~9/23全校加退選。",
						},
					},
				},
			},
			Courses: []domain.Courses{
				{
					LastModified: lastModified,
					NameZH:       "計算機網路",
					NameUS:       "computer network",
					BulletinBoard: []domain.BulletinBoards{
						{Content: "資工2A(第五週)10月14日起恢復實體授課，上課地點:I526教室，請配帶口罩，並按座位表入座，謝謝"},
						{Content: "資工2計算機網路概論--上課地點在I526教室，但前四週(9/13-10/10)請於創課平台TronClass線上同步遠距學習，後續週數再按學校防疫規定另行公布。"},
					},
					Slide: []domain.Slides{
						{Chapter: "00", File: domain.File{Title: "Syllabus and Introduction(課程介紹)", Type: "pdf"}},
						{Chapter: "01", File: domain.File{Title: "ISO OSI 7-layer Model (網路架構與OSI七層參考模式)", Type: "pdf"}},
						{Chapter: "02", File: domain.File{Title: "Physical layer (實體層: Wired/Wireless, Hub)", Type: "pdf"}},
						{Chapter: "03", File: domain.File{Title: "Data link layer (資料鏈結層: Protocol principles, PPP)", Type: "pdf"}},
						{Chapter: "04-1", File: domain.File{Title: "MAC Sublayer (媒體存取控制層: Ethernet, IEEE 802.3, Bridge)", Type: "pdf"}},
						{Chapter: "04-2", File: domain.File{Title: "MAC Sublayer (媒體存取控制層: Switch, IEEE 802.1D, VLAN)", Type: "pdf"}},
					},
					Homework: []domain.Homeworks{
						{Number: "1", File: domain.File{Title: "第1-2章", Type: "pdf"}},
						{Number: "2", File: domain.File{Title: "第3-4章", Type: "pdf"}},
					},
				},
				{
					LastModified: lastModified,
					NameZH:       "無線網路",
					NameUS:       "wireless network",
					BulletinBoard: []domain.BulletinBoards{
						{Content: "資工2A在6/23(三)9:30-11:00無線網路期末考，改成：線上考試或學習報告，範圍:第7章到第14章，謝謝。"},
						{Content: "資工2A在4/21(三)8:30-10:00無線網路期中考，地點改在：L005(行政大樓地下室)，範圍:第1章到第6章，筆試，開書考，本次不得使用計算機，謝謝。"},
					},
					Slide: []domain.Slides{
						{Chapter: "00", File: domain.File{Title: "Syllabus (課程大綱)", Type: "pdf"}},
						{Chapter: "01", File: domain.File{Title: "Introduction (序論)", Type: "pdf"}},
						{Chapter: "03", File: domain.File{Title: "Mobile Radio Propagation (行動無線電傳播)", Type: "pdf"}},
						{Chapter: "04", File: domain.File{Title: "Channel Coding and Error Control (通道編碼與錯誤控制)", Type: "pdf"}},
					},
					Homework: []domain.Homeworks{
						{Number: "1", File: domain.File{Title: "作業一", Type: "pdf"}},
						{Number: "2", File: domain.File{Title: "作業二", Type: "pdf"}},
					},
				},
			},
		},
	}
	db.Create(&data)
}
func generalHashPassword(password string, conf *config.Secure) (string, string) {
	var salt = make([]byte, conf.SaltSize)
	rand.Read(salt[:])
	saltPassword := append([]byte(password), salt...)
	hashPassword, _ := bcrypt.GenerateFromPassword(saltPassword, conf.HashCost)
	return string(hashPassword), string(salt)
}
