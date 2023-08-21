package mockgenerator

// use mockery to generate mock file
// install mock toll
//   - go install github.com/vektra/mockery/v2@latest

//go:generate mockery --name Repositories --with-expecter --filename mock_repo.go --dir ../../internal/repositories --output ../../internal/mocks
//go:generate mockery --name Services --with-expecter --filename mock_service.go --dir ../../internal/repositories --output ../../internal/mocks
//go:generate mockery --name AccountServices --with-expecter --filename mock_account.go --dir ../../internal/repositories --output ../../internal/mocks
