package servicestest

import (
	"os"
	"time"

	"bou.ke/monkey"
	"github.com/go-redis/redis"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SuiteConfig struct {
	*suite.Suite

	relativeModels []models.Models
	withTimeStub   bool
	clearData      bool
	db             *gorm.DB
	repo           repositories.Repositories
	rd             *redis.Client
	timeStub       time.Time
	tearDownTime   func()
}

type SuiteParameters struct {
	RelativeModels []models.Models
	ClearData      bool
	WithTimeStub   bool
}

func NewSuiteConfig(params SuiteParameters) *SuiteConfig {
	return &SuiteConfig{
		Suite:          &suite.Suite{},
		relativeModels: params.RelativeModels,
		withTimeStub:   params.WithTimeStub,
		clearData:      params.ClearData,
	}
}

func (s SuiteConfig) SetupSuite() {
	dm, db, rd, err := InitializeDB()
	if err != nil {
		os.Exit(1)
	}
	s.db = db
	s.repo = dm
	s.rd = rd

	s.tearDownTime = func() {}
}

func (s *SuiteConfig) ClearData() error {
	for _, m := range s.relativeModels {
		if err := s.db.Where("1=1").Delete(m).Error; err != nil {
			return err
		}
	}
	err := s.rd.FlushDB().Err()
	if err != nil {
		return err
	}
	return nil
}

func (s SuiteConfig) SetupTest() {
	if s.clearData {
		s.ClearData()
	}
	if s.withTimeStub {
		s.timeStub = time.Now()
		s.tearDownTime = TimeSetup(s.timeStub)
	}
}

func TimeSetup(t time.Time) (tearDown func()) {
	p := monkey.Patch(time.Now, func() time.Time {
		return t
	})
	return p.Unpatch
}

func (s SuiteConfig) TearDownTest() {
	s.tearDownTime()
}

func (s SuiteConfig) TearDownSuite() {
	s.repo.Close()
}

func (s *SuiteConfig) GetDm() repositories.Repositories {
	return s.repo
}
