package workersgrpc

import (
	"context"

	"github.com/andrei1998Front/grpc_workers_server/internal/domain/models"
	wrkrs "github.com/andrei1998Front/grpc_workers_server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	wrkrs.UnimplementedWorkersServiceServer
	employeeService EmployeeService
}

func New(
	employeeService EmployeeService,
) *serverAPI {
	return &serverAPI{
		employeeService: employeeService,
	}
}

type EmployeeService interface {
	GetInfoByName(ctx context.Context, title string) ([]*models.EmployeeInfo, error)
}

func Register(
	gRPCServer *grpc.Server,
	employeeService EmployeeService,
) {
	sApi := New(employeeService)
	wrkrs.RegisterWorkersServiceServer(gRPCServer, sApi)
}

func (s *serverAPI) GetEployeeFullInfo(ctx context.Context, req *wrkrs.EmployeeJobGetByEmployeeName) (*wrkrs.EmployeeFullInfoList, error) {
	var employeeInfoResponce wrkrs.EmployeeFullInfoList

	if req.EmployeeTitle == "" {
		return &wrkrs.EmployeeFullInfoList{}, status.Error(codes.InvalidArgument, "title employee is empty")
	}

	employeeInfoList, err := s.employeeService.GetInfoByName(ctx, req.EmployeeTitle)

	if err != nil {
		return &wrkrs.EmployeeFullInfoList{}, status.Error(codes.Internal, "internal error")
	}

	for _, employee := range employeeInfoList {
		infoItem := wrkrs.EmployeeFullInfo{
			JobTitle:        employee.JobTitle,
			EmployeeTitle:   employee.EmployeeTitle,
			Hoursrate:       employee.Hoursrate,
			Annualsalary:    employee.Annualsalary,
			Yy:              employee.YY,
			Typefp:          employee.Typefp,
			Typicalhours:    employee.Typicalhours,
			DepartmentTitle: employee.DepartmentTitle,
		}
		employeeInfoResponce.InfoList = append(employeeInfoResponce.InfoList, &infoItem)
	}

	return &employeeInfoResponce, nil
}
