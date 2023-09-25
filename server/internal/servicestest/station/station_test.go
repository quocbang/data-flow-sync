package station

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bou.ke/monkey"

	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	repoErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	m "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/setupmock"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
	"github.com/quocbang/data-flow-sync/server/utils/diff"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
	"github.com/quocbang/data-flow-sync/server/utils/types"
	"github.com/stretchr/testify/mock"
)

func (s *Suite) TestCreateStationMergeRequest() {
	assertion := s.Assert()

	goodParams := func(files []string) station.CreateStationMergeRequestParams {
		return station.CreateStationMergeRequestParams{
			HTTPRequest: s.HttpTestRequest("POST", "/api/station/merge-request", nil),
			Body: station.CreateStationMergeRequestBody{
				Files: files,
			},
		}
	}
	year, month, day := time.Now().Local().Date()
	s.Context = goodParams([]string{}).HTTPRequest.Context()

	old := repositories.Station{
		ID:           "TEST_ID",
		SubCompany:   1,
		Factory:      "TEST_FACTORY",
		DepartmentID: "TEST_DEP",
		Alias:        "TEST_A",
		SerialNumber: 1,
		Description:  "TESTER",
		Devices:      []int64{1, 2, 3, 4},
	}
	oldByte, err := json.Marshal(old)
	assertion.NoError(err)
	new := repositories.Station{
		ID:           "TEST_ID",
		SubCompany:   2,
		Factory:      "TEST_FACTORY",
		DepartmentID: "TEST_DEP",
		Alias:        "TEST_A",
		SerialNumber: 2,
		Description:  "TESTER",
		Devices:      []int64{1, 6, 8, 4},
	}

	createMRRequest := func(old, new []byte, fileId string, creator string, createFileAmount int) repositories.CreateMRRequest {
		createMergeRequest := repositories.CreateMRRequest{
			MergeRequest: m.MergeRequest{
				Type: types.Types_STATION,
				File: make(m.Files, createFileAmount),
				Information: m.MergeRequestInfo{
					HistoryChanged: []string{fmt.Sprintf("Create a merge request at %d-%d-%d, by %s", year, month, day, creator)},
					Status: m.MergeRequestStatus{
						IsApproved: false,
						IsMerged:   false,
						IsOpen:     true,
					},
				},
				CreatedBy: creator,
			},
		}
		for i := 0; i < createFileAmount; i++ {
			added, deleted, err := diff.FindDiff[repositories.Station](old, new)
			assertion.NoError(err)
			createMergeRequest.MergeRequest.File[i] = m.FileMergeRequest{
				FileID: fileId,
				FileCompare: m.FileCompares{
					FileBeforeMerge: old,
					FileAfterMerge:  new,
				},
				FileChanged: m.Changed{
					Added:   added,
					Deleted: deleted,
				},
			}
		}

		return createMergeRequest
	}

	// create merge request success
	{
		// Arrange
		newByte, err := json.Marshal(new)
		assertion.NoError(err)

		var files []string
		files = append(files, string(newByte))
		params := goodParams(files)
		principal := models.Principal{
			Email:             "test_auth@gmail.com",
			ID:                "tester",
			IsUnspecifiedUser: false,
			Role:              int64(roles.Roles_LEADER),
		}
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.NewMergeRequestServices(s.T())
				mergeRequest.EXPECT().CreateMergeRequest(s.Context, createMRRequest(oldByte, newByte, new.ID, "tester", 1)).ReturnArguments =
					mock.Arguments{
						repositories.CreateMRReply{
							MergeRequestID: 1,
						}, nil,
					}
				mergeRequest.EXPECT().GetMergeRequestOpeningByFileID(s.Context, new.ID).ReturnArguments =
					mock.Arguments{
						repositories.GetMergeRequestOpeningByFileIDReply{}, repoErr.ErrDataNotFound, // should return not found
					}
				return mergeRequest
			}(),
		}
		mockRepo.EXPECT().File().ReturnArguments = mock.Arguments{
			func() repositories.FileServices {
				file := mocks.NewFileServices(s.T())
				file.EXPECT().GetFile(s.Context, repositories.GetFileRequest{
					ID:   old.ID,
					Type: types.Types_STATION,
				}).ReturnArguments =
					mock.Arguments{repositories.GetFileReply{
						File: m.File{
							ID:          old.ID,
							Type:        types.Types_STATION,
							FileContent: oldByte,
						},
					}, nil}
				return file
			}(),
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.CreateStationMergeRequest(params, &principal).(*station.CreateStationMergeRequestOK)

		// Assert
		assertion.Equal(int64(1), response.Payload.MergeRequestID)
	}

	// create merge request failed: forbidden
	{
		// Arrange
		principal := &models.Principal{
			Email:             "quocbang@gmail.com",
			ID:                "quocbang",
			IsUnspecifiedUser: true,
			Role:              int64(roles.Roles_UNSPECIFIED),
		}
		mockRepo := s.MockRepository()
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.CreateStationMergeRequest(goodParams([]string{}), principal).(*station.CreateStationMergeRequestDefault)

		// Assert
		expected := &models.ErrorResponse{
			Code:    int64(errors.Code_FORBIDDEN),
			Details: "permission denied",
		}
		assertion.Equal(expected, response.Payload)
	}
	// create merge request failed: missing request body
	{
		// Arrange
		principal := &models.Principal{
			Email:             "quocbang@gmail.com",
			ID:                "quocbang",
			IsUnspecifiedUser: true,
			Role:              int64(roles.Roles_LEADER),
		}
		mockRepo := s.MockRepository()
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.CreateStationMergeRequest(goodParams([]string{}), principal).(*station.CreateStationMergeRequestBadRequest)

		// Assert
		expected := &models.ErrorResponse{
			Details: "missing request body",
		}
		assertion.Equal(expected, response.Payload)
	}
	// create merge request failed: failed to  unmarshal file
	{
		// Arrange
		principal := &models.Principal{
			Email:             "quocbang@gmail.com",
			ID:                "quocbang",
			IsUnspecifiedUser: true,
			Role:              int64(roles.Roles_LEADER),
		}
		mockRepo := s.MockRepository()
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))
		monkey.Patch(json.Unmarshal, func([]byte, any) error {
			return fmt.Errorf("unexpected json format")
		})

		// Act
		response := mockServer.Station.CreateStationMergeRequest(goodParams([]string{"bad format file 01", "bad format file 02"}), principal).(*station.CreateStationMergeRequestBadRequest)

		// Assert
		monkey.UnpatchAll()
		expected := &models.ErrorResponse{
			Details: "failed to unmarshal file [0], error: unexpected json format",
		}
		assertion.Equal(expected, response.Payload)
	}
	// create merge request failed: file was conflicted
	{
		newByte, err := json.Marshal(new)
		assertion.NoError(err)

		var files []string
		files = append(files, string(newByte))
		params := goodParams(files)
		principal := models.Principal{
			Email:             "test_auth@gmail.com",
			ID:                "tester",
			IsUnspecifiedUser: false,
			Role:              int64(roles.Roles_LEADER),
		}
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.NewMergeRequestServices(s.T())
				mergeRequest.EXPECT().GetMergeRequestOpeningByFileID(s.Context, new.ID).ReturnArguments =
					mock.Arguments{
						repositories.GetMergeRequestOpeningByFileIDReply{}, nil, // return nil that mean the transaction is complete without error
					}
				return mergeRequest
			}(),
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.CreateStationMergeRequest(params, &principal).(*station.CreateStationMergeRequestBadRequest)

		// Assert
		expected := &models.ErrorResponse{
			Code:    int64(errors.Code_FILE_CONFLICTED),
			Details: "file [TEST_ID] was conflicted",
		}
		assertion.Equal(expected, response.Payload)
	}
}

func (s *Suite) TestGetStationMergeRequest() {
	assertion := s.Assertions
	goodParams := func(ID int64) station.GetStationMergeRequestParams {
		return station.GetStationMergeRequestParams{
			HTTPRequest: s.HttpTestRequest(http.MethodGet, fmt.Sprintf("/api/station/merge-request"), nil),
			ID:          ID,
		}
	}
	s.Context = goodParams(1).HTTPRequest.Context()
	getMergeRequestReply := func(ID int64, added []byte, deleted []byte) repositories.GetMRReply {
		return repositories.GetMRReply{
			MergeRequest: m.MergeRequest{
				ID:   uint64(ID),
				Type: types.Types_STATION,
				File: m.Files{
					m.FileMergeRequest{
						FileID: "TEST_ID",
						FileCompare: m.FileCompares{
							FileBeforeMerge: []byte(`{
								"ID": "TEST_ID",
								"subCompany": 1,
								"factory": "TEST_FAC",
								"departmentID": "TEST_DEP",
								"alias": "TEST_ALIAS",
								"serialNumber": 2,
								"description": "TEST_DE",
								"devices": [1,2,3]}`),
							FileAfterMerge: []byte(`{
								"ID": "TEST_ID",
								"subCompany": 1,
								"factory": "TEST_FAC",
								"departmentID": "TEST_DEP",
								"alias": "TEST_ALIAS",
								"serialNumber": 2,
								"description": "TEST_DE",
								"devices": [1,2,3,4,5]}`),
						},
						FileChanged: m.Changed{
							Added:   added,
							Deleted: deleted,
						},
					},
				},
				Information: m.MergeRequestInfo{
					HistoryChanged: []string{"create"},
					Status: m.MergeRequestStatus{
						IsOpen: true,
					},
				},
			},
		}
	}
	principal := &models.Principal{
		Email:             "tester@gmail.com",
		ID:                "tester",
		IsUnspecifiedUser: true,
		Role:              int64(roles.Roles_UNSPECIFIED),
	}

	// get station merge request successfully.
	{
		// Arrange
		added := []byte(`{"devices": [4,5]}`)
		deleted := []byte(`{}`)
		params := goodParams(1)
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.MergeRequestServices{}
				mergeRequest.EXPECT().GetMergeRequest(s.Context, repositories.GetMRRequest{
					MergeRequestID: 1}).ReturnArguments = mock.Arguments{
					getMergeRequestReply(1, added, deleted), nil,
				}
				return &mergeRequest
			},
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.GetStationMergeRequest(params, principal).(*station.GetStationMergeRequestOK)

		// Assert
		isOpening := true
		isApprove := false
		isMerged := false
		expected := &station.GetStationMergeRequestOKBody{
			Data: []*models.StationMergeRequest{
				{
					Added: &models.Station{
						Devices: []int64{4, 5},
					},
					Deleted: &models.Station{},
					File: &models.Station{
						ID:           "TEST_ID",
						Alias:        "TEST_ALIAS",
						DepartmentID: "TEST_DEP",
						Description:  "TEST_DE",
						Devices:      []int64{1, 2, 3, 4, 5},
						Factory:      "TEST_FAC",
						SerialNumber: 2,
						SubCompany:   1,
					},
				},
			},
			MergeRequestInfo: &station.GetStationMergeRequestOKBodyMergeRequestInfo{
				HistoryChanged: []string{"create"},
			},
			MergeRequestStatus: &station.GetStationMergeRequestOKBodyMergeRequestStatus{
				IsOpening:  &isOpening,
				IsApproved: &isApprove,
				IsMerged:   &isMerged,
			},
		}
		assertion.Equal(expected, response.Payload)
	}

	// get station merge request failed: missing merge request id
	{
		// Arrange
		params := goodParams(0)
		mockRepo := s.MockRepository()
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.GetStationMergeRequest(params, principal).(*station.GetStationMergeRequestBadRequest)

		// Assert
		expected := &models.ErrorResponse{
			Details: "missing merge request id",
		}
		assertion.Equal(expected, response.Payload)
	}
	// get station merge request failed: failed to unmarshal added, error: invalid character ':' after top-level value
	{
		// Arrange
		added := []byte(`"devices": [4,5]}`)
		deleted := []byte(`{}`)
		params := goodParams(1)
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.MergeRequestServices{}
				mergeRequest.EXPECT().GetMergeRequest(s.Context, repositories.GetMRRequest{
					MergeRequestID: 1}).ReturnArguments = mock.Arguments{
					getMergeRequestReply(1, added, deleted), nil,
				}
				return &mergeRequest
			},
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.GetStationMergeRequest(params, principal).(*station.GetStationMergeRequestInternalServerError)

		// Assert
		expected := &models.ErrorResponse{
			Details: "failed to unmarshal added, error: invalid character ':' after top-level value",
		}
		assertion.Equal(expected, response.Payload)
	}
	// get station merge request failed: failed to unmarshal deleted, error: invalid character '}' looking for beginning of value
	{
		// Arrange
		added := []byte(`{"devices": [4,5]}`)
		deleted := []byte(`}`)
		params := goodParams(1)
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.MergeRequestServices{}
				mergeRequest.EXPECT().GetMergeRequest(s.Context, repositories.GetMRRequest{
					MergeRequestID: 1}).ReturnArguments = mock.Arguments{
					getMergeRequestReply(1, added, deleted), nil,
				}
				return &mergeRequest
			},
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.GetStationMergeRequest(params, principal).(*station.GetStationMergeRequestInternalServerError)

		// Assert
		expected := &models.ErrorResponse{
			Details: "failed to unmarshal deleted, error: invalid character '}' looking for beginning of value",
		}
		assertion.Equal(expected, response.Payload)
	}
	// get station merge request failed: failed to unmarshal file after merge, error: unexpected end of JSON input
	{
		// Arrange
		added := []byte(`{"devices": [4,5]}`)
		deleted := []byte(`{}`)
		params := goodParams(1)
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().MergeRequest().ReturnArguments = mock.Arguments{
			func() repositories.MergeRequestServices {
				mergeRequest := mocks.MergeRequestServices{}
				reply := getMergeRequestReply(1, added, deleted)
				reply.MergeRequest.File[0].FileCompare.FileAfterMerge = []byte(`{`)
				mergeRequest.EXPECT().GetMergeRequest(s.Context, repositories.GetMRRequest{
					MergeRequestID: 1}).ReturnArguments = mock.Arguments{
					reply, nil,
				}
				return &mergeRequest
			},
		}
		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Station.GetStationMergeRequest(params, principal).(*station.GetStationMergeRequestInternalServerError)

		// Assert
		expected := &models.ErrorResponse{
			Details: "failed to unmarshal file after merge, error: unexpected end of JSON input",
		}
		assertion.Equal(expected, response.Payload)
	}
}
