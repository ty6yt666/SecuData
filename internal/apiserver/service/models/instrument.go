package models

import "time"

//dsn := "host=localhost port=5432 instrument=root password= dbname=instrument"
//db := gorm.Open(postgres.Open(dsn), &gorm.Config{})

type Instrument struct {
	InstrumentId       int64     `json:"instrument_id" gorm:"column:instrument_id;primary_key"`
	InstrumentMnemonic string    `json:"instrument_mnemonic" gorm:"column:instrument_mnemonic"`
	InstrumentSubtype  string    `json:"instrument_subtype" grom:"column:instrument_subtype"`
	InfoCode           string    `json:"info_code" gorm:"column:info_code"`
	InfoName           string    `json:"info_name" gorm:"column:info_name"`
	RicCode            string    `json:"ric_code" gorm:"column:ric_code"`
	ChoiceCode         string    `json:"choice_code" gorm:"column:choice_code"`
	BloombergCode      string    `json:"bloomberg_code" gorm:"column:bloomberg_code"`
	WindCode           string    `json:"wind_code" gorm:"column:wind_code"`
	ListExchange       string    `json:"list_exchange" gorm:"column:list_exchange"`
	FundId             string    `json:"fund_id" gorm:"column:fund_id"`
	CreatedAt          time.Time `json:"created_at,omitempty" gorm:"column:createdAt"`
	UpdateAt           time.Time `json:"update_at,omitempty" gorm:"column:updateAt"`
}

type InstrumentList struct {
	Items []*Instrument `json:"items"`
}

func (instrument *Instrument) TableName() string {
	return "instrument_basic"
}
