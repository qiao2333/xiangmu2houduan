package models

//LightStatistics 灯饰统计表
type LightStatistics struct {
	LightCode  int `xorm:"not null pk comment('灯饰Code') INT(11)"`
	StoreCount int `xorm:"default 0 comment('库存量') INT(11)"`
	SaleCount  int `xorm:"default 0 comment('销售数量') INT(11)"`
}

//LightStatisticsInterface 灯饰统计表接口
type LightStatisticsInterface interface {
	New() error
	Save(conditions LightStatistics) error
}
