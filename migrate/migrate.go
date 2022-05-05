package migrate

import (
	"crypto/rand"
	"strconv"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/mock"
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
	// token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())
	lastModified := strconv.FormatInt(currTime.Unix(), 10)
	hashPassword, salt := generalHashPassword(mock.UserPassword, conf.Secure)
	users := [2]domain.Auths{
		mock.GenerateAuth(),
		{
			UserId:       "rikki",
			UserPassword: string(hashPassword),
			Salt:         string(salt),
			Teacher: domain.Teachers{
				Domain:    "rikki",
				Email:     "rikki@mail.tw",
				NameZh:    "陳瑞奇",
				NameUs:    "Jui-Chi Chen(Rikki)",
				Office:    "HB13",
				Call:      "(04)1234-5678 ext. 1234",
				Education: "國立中興大學資訊科學博士",
				Infos: []domain.Infos{
					{
						LastModified: lastModified,
						BulletinBoards: []domain.InfoBulletinBoards{
							{Content: "(第五週)10月11日起恢復實體授課，請配帶口罩，並按教室座位表入座，謝謝。"},
							{Content: "請益時間(Office Hours): 資訊大樓地下室HB13研究室 TUE(二) 14:10 – 17:00    THU(四) 10:10 – 17:00    Fri.(五) 10:10 – 12:00"},
							{Content: "9/13(一)開學日、開始上課，9/13~9/23全校加退選。"},
							{Content: "即日起進入本校各辦公處所、教室、實驗室、工坊等室內空間時，請一律佩戴口罩。(防疫小組規定)"},
							{Content: "2/22(一)開學日、開始上課，2/23~3/4全校加退選"},
							{Content: "11月4~10日期中考(We have midterm exams on 4-10, Nov.)"},
							{Content: "4月4,5日(星期四、五)清明節放假。"},
						},
					},
				},
				Courses: []domain.Courses{
					{
						LastModified: lastModified,
						NameZh:       "計算機網路",
						NameUs:       "computer network",
						BulletinBoard: []domain.BulletinBoards{
							{Content: "資工2A(第五週)10月14日起恢復實體授課，上課地點:I526教室，請配帶口罩，並按座位表入座，謝謝"},
							{Content: "資工2計算機網路概論--上課地點在I526教室，但前四週(9/13-10/10)請於創課平台TronClass線上同步遠距學習，後續週數再按學校防疫規定另行公布。"},
							{Content: `參考教科書(Textbook): ※請遵守智慧財產權規定，不可非法影印教科書。<br>
						英文版：Andrew S. Tanenbaum, David J. Wetherall, “Computer Networks 5/e” 5th ed., Pearson Education (東華), 2011, ISBN: 9780132553179.<br>
						中文版：邵喻美、潘育群譯, “電腦網路(第5版),” 東華書局, 2012, ISBN：9789862800973.<br>
						參考資料(Reference)：<br>
						James F. Kurose, Keith W. Ross, “Computer Networking - A Top-Down Approach,” 7th ed., Global Edition, Pearson, 2016, ISBN: 9781292153599.`},
						},
						Slide:    []domain.Slides{},
						Homework: []domain.Homeworks{},
					},
					{
						LastModified: lastModified,
						NameZh:       "無線網路",
						NameUs:       "wireless network",
						BulletinBoard: []domain.BulletinBoards{
							{Content: "學期結束。(The semester is over!) "},
							{Content: "資工2A在6/23(三)9:30-11:00無線網路期末考，改成：線上考試或學習報告，範圍:第7章到第14章，謝謝。"},
							{Content: "資工2A在4/21(三)8:30-10:00無線網路期中考，地點改在：L005(行政大樓地下室)，範圍:第1章到第6章，筆試，開書考，本次不得使用計算機，謝謝。"},
							{Content: "Office Hours(請益時間):Tue.(二) 9:10-12:00    Wed.(三) 13:10-15:00    Thu.(四) 13:10-16:00"},
							{Content: "資工CSIE 2A無線網路概論(Wireless Networks)--上課時間(Class & Location): 每週三(WED)9:10-12:00在Room H103教室。"},
							{Content: `※請遵守智慧財產權規定，不可非法影印教科書。參考資料(Reference)：<br>
						1. D.P. Agrawal and Q.-A. Zeng, “Introduction to Wireless and Mobile Systems,” 4th Ed. (International), Cengage Learning (東華書局代理), 2015, ISBN: 9781305259621.<br>
						2. 曾恕銘編譯, “無線通訊系統概論：行動通訊與網路 4/e,” 東華書局, 2016, ISBN-13：9789865632786。<br>
						3. 陳裕賢、張志勇、陳宗禧、石貴平、吳世琳、廖文華、許智舜、林勻蔚, “無線網路與行動計算,” 全華書局, 2013/2, ISBN：9789572188637`},
						},
						Slide:    []domain.Slides{},
						Homework: []domain.Homeworks{},
					},
				},
			},
		},
	}
	db.Create(&users)
}
func generalHashPassword(password string, conf *config.Secure) (string, string) {
	var salt = make([]byte, conf.SaltSize)
	rand.Read(salt[:])
	saltPassword := append([]byte(password), salt...)
	hashPassword, _ := bcrypt.GenerateFromPassword(saltPassword, conf.HashCost)
	return string(hashPassword), string(salt)
}
