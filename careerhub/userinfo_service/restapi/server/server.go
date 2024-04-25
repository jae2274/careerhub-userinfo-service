package server

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/service"
)

type RestApiGrpcServer struct {
	conditionService service.ConditionService
	restapi_grpc.UnimplementedRestApiGrpcServer
}

func NewRestApiGrpcServer(conditionService service.ConditionService) restapi_grpc.RestApiGrpcServer {
	return &RestApiGrpcServer{
		conditionService: conditionService,
	}
}

func (r *RestApiGrpcServer) FindConditions(ctx context.Context, req *restapi_grpc.FindConditionsRequest) (*restapi_grpc.Conditions, error) {
	desiredCondition, err := r.conditionService.FindByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	response := make([]*restapi_grpc.Condition, len(desiredCondition.Conditions))
	for i, condition := range desiredCondition.Conditions {
		response[i] = convertConditionToGrpc(condition)
	}

	return &restapi_grpc.Conditions{
		Conditions:  response,
		AgreeToMail: desiredCondition.AgreeToMail,
	}, nil
}
func (r *RestApiGrpcServer) FindCondition(ctx context.Context, req *restapi_grpc.FindConditionRequest) (*restapi_grpc.Condition, error) {
	condition, err := r.conditionService.FindByUserIdAndUUID(ctx, req.UserId, req.ConditionId)
	if err != nil {
		return nil, err
	}

	return convertConditionToGrpc(condition), nil
}
func (r *RestApiGrpcServer) AddCondition(ctx context.Context, req *restapi_grpc.AddConditionRequest) (*restapi_grpc.IsSuccess, error) {
	newCondition := convertConditionToDomain("", req.Condition.ConditionName, req.Condition.Query)
	success, err := r.conditionService.InsertCondition(ctx, req.UserId, uint(req.LimitCount), newCondition)

	return &restapi_grpc.IsSuccess{
		Value: success,
	}, err
}
func (r *RestApiGrpcServer) UpdateCondition(ctx context.Context, req *restapi_grpc.UpdateConditionRequest) (*restapi_grpc.IsSuccess, error) {
	updateCondition := convertConditionToDomain(req.Condition.ConditionId, req.Condition.ConditionName, req.Condition.Query)
	success, err := r.conditionService.UpdateCondition(ctx, req.UserId, updateCondition)

	return &restapi_grpc.IsSuccess{
		Value: success,
	}, err
}
func (r *RestApiGrpcServer) DeleteCondition(ctx context.Context, req *restapi_grpc.DeleteConditionRequest) (*restapi_grpc.IsSuccess, error) {
	success, err := r.conditionService.DeleteCondition(ctx, req.UserId, req.ConditionId)

	return &restapi_grpc.IsSuccess{
		Value: success,
	}, err
}

func convertConditionToGrpc(domainValue *condition.Condition) *restapi_grpc.Condition {
	categories := make([]*restapi_grpc.Category, len(domainValue.Query.Categories))
	for i, category := range domainValue.Query.Categories {
		categories[i] = &restapi_grpc.Category{
			Site:         category.Site,
			CategoryName: category.CategoryName,
		}
	}
	skillNames := make([]*restapi_grpc.Skill, len(domainValue.Query.SkillNames))
	for i, skill := range domainValue.Query.SkillNames {
		skillNames[i] = &restapi_grpc.Skill{
			Or: skill,
		}
	}

	return &restapi_grpc.Condition{
		ConditionId:   domainValue.ConditionId,
		ConditionName: domainValue.ConditionName,
		Query: &restapi_grpc.Query{
			Categories: categories,
			SkillNames: skillNames,
			MinCareer:  domainValue.Query.MinCareer,
			MaxCareer:  domainValue.Query.MaxCareer,
		},
	}
}

func convertConditionToDomain(conditionId, conditionName string, grpcValue *restapi_grpc.Query) *condition.Condition {
	return &condition.Condition{
		ConditionId:   conditionId,
		ConditionName: conditionName,
		Query:         convertQueryToDomain(grpcValue),
	}
}

func convertQueryToDomain(grpcValue *restapi_grpc.Query) *condition.Query {
	categories := make([]*condition.CategoryQuery, len(grpcValue.Categories))
	for i, category := range grpcValue.Categories {
		categories[i] = &condition.CategoryQuery{
			Site:         category.Site,
			CategoryName: category.CategoryName,
		}
	}

	skillNames := make([][]string, len(grpcValue.SkillNames))
	for i, skill := range grpcValue.SkillNames {
		skillNames[i] = skill.Or
	}

	return &condition.Query{
		Categories: categories,
		SkillNames: skillNames,
		MinCareer:  grpcValue.MinCareer,
		MaxCareer:  grpcValue.MaxCareer,
	}
}
