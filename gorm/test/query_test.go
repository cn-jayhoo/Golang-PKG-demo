package test

import (
	"demo-gorm/model"
	"fmt"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestQuery(t *testing.T) {

}

type QuerySuite struct {
	db *gorm.DB
	suite.Suite
}

func (s *QuerySuite) SetupTest() {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"demo",
		"demo",
		fmt.Sprintf("%s:%d", "172.18.1.194", 3306),
		"demo_schema",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		s.Failf("Init fail", "failed to connect database >> %s", err.Error())
		return
	}
	s.db = db
}

func (s *QuerySuite) TearDownTest() {

}

func (s *QuerySuite) TestQueryOne() {
	priceModel := &model.TPrice{}
	err := s.db.Model(&model.TPrice{}).Where("uid = ?", 1).Scan(&priceModel).Error
	if err != nil {
		s.Failf("Query fail", err.Error())
		return
	}
	s.T().Logf("查询结果值为 %s", priceModel.Price.String())
}

func TestQuerySuite(t *testing.T) {
	suite.Run(t, new(QuerySuite))
}
