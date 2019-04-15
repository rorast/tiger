package models

type StarInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主鍵ID') INT(11)" form:"id"`
	NameZh       string `xorm:"not null default '' comment('中文名') VARCHAR(50)" form:"name_zh"`
	NameEn       string `xorm:"not null default '' comment('英文名') VARCHAR(50)" form:"name_en"`
	Avatar       string `xorm:"comment('頭像') VARCHAR(255)" form:"avatar"`
	Birthday     string `xorm:"comment('生日') VARCHAR(50)" form:"birthday"`
	Height       int    `xorm:"comment('身高') INT(10)" form:"height"`
	Weight       int    `xorm:"comment('體重') INT(10)" form:"weight"`
	Club         string `xorm:"comment('俱樂部') VARCHAR(50)" form:"club"`
	Jersy        string `xorm:"comment('球衣號碼以及主隊') VARCHAR(50)" form:"jersy"`
	Country      string `xorm:"comment('國籍') VARCHAR(50)" form:"country"`
	Birthaddress string `xorm:"comment('出生地') VARCHAR(255)" form:"birthaddress"`
	Feature      string `xorm:"comment('個人特點') VARCHAR(255)" form:"feature"`
	Moreinfo     string `xorm:"comment('更多介紹') TEXT" form:"moreinfo"`
	SysStatus    int    `xorm:"default 0 comment('狀態，默認值0') TINYINT(3)" form:"-"`
	SysCreated   int    `xorm:"default 0 INT(10)" form:"-"`
	SysUpdated   int    `xorm:"default 0 INT(10)" form:""-`
}
