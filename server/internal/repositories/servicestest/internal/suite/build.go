package suite

import (
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/utils/random"
)

type SuiteConfig struct {
	*suite.Suite
	PG   *gorm.DB
	RD   *redis.Client
	Repo repositories.Repositories

	relativeModels       []models.Models
	clearDataForEachTest bool
}

type SuiteParameters struct {
	// just do test or clear data in those models also data in database.
	RelativeModels []models.Models
	// clear all data in RelativeModels.
	ClearDataForEachTest bool
}

func NewSuiteTest(params SuiteParameters) *SuiteConfig {
	return &SuiteConfig{
		relativeModels:       params.RelativeModels,
		Suite:                &suite.Suite{},
		clearDataForEachTest: params.ClearDataForEachTest,
	}
}

var logger *zap.Logger

func (s *SuiteConfig) SetupSuite() {
	field := []zap.Field{zap.String("random_seed", random.RandomString(30))}
	logger.Info("seed", field...)

	// initialize database.
	repo, pg, rd, err := InitializeDB()
	if err != nil {
		os.Exit(1)
	}
	s.Repo, s.PG, s.RD = repo, pg, rd
}

func (suite *SuiteConfig) TearDownSuite() {
	err := suite.Repo.Close()
	if err != nil {
		os.Exit(1)
	}
}

func (s *SuiteConfig) ClearPostgresData() error {
	for _, m := range s.relativeModels {
		if err := s.PG.Where("1=1").Delete(m).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *SuiteConfig) ClearRedisData() error {
	return s.RD.FlushDB().Err()
}

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}
