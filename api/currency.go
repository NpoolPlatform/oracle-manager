// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/oracle-manager/pkg/const"
	crud "github.com/NpoolPlatform/oracle-manager/pkg/crud/currency"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/oraclemgr"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCurrency(ctx context.Context, in *npool.CreateCurrencyRequest) (*npool.CreateCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorf("invalid request coin type id: %v", err)
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.GetInfo().GetPriceVSUSDT() <= 0 {
		logger.Sugar().Errorf("invalid coin usdt price")
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency parameter")
	}

	if in.GetInfo().GetAppPriceVSUSDT() <= 0 {
		logger.Sugar().Errorf("invalid app coin usdt price")
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency parameter")
	}

	switch in.GetInfo().GetCurrencyMethod() {
	case constant.CurrencyOverPercent:
	case constant.CurrencyFixAmount:
	default:
		logger.Sugar().Errorf("invalid currency method")
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency method")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create currency: %v", err)
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppCurrency(ctx context.Context, in *npool.CreateAppCurrencyRequest) (*npool.CreateAppCurrencyResponse, error) {
	in.GetInfo().AppID = in.GetTargetAppID()
	resp, err := s.CreateCurrency(ctx, &npool.CreateCurrencyRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return &npool.CreateAppCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppCurrencyResponse{
		Info: resp.Info,
	}, nil

}

func (s *Server) UpdateCurrency(ctx context.Context, in *npool.UpdateCurrencyRequest) (*npool.UpdateCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorf("invalid request coin type id: %v", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("invalid currency id: %v", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.GetInfo().GetPriceVSUSDT() <= 0 {
		logger.Sugar().Errorf("invalid coin usdt price")
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency parameter")
	}

	if in.GetInfo().GetAppPriceVSUSDT() <= 0 {
		logger.Sugar().Errorf("invalid app coin usdt price")
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency parameter")
	}

	switch in.GetInfo().GetCurrencyMethod() {
	case constant.CurrencyOverPercent:
	case constant.CurrencyFixAmount:
	default:
		logger.Sugar().Errorf("invalid currency method")
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, "invalid currency method")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update currency: %v", err)
		return &npool.UpdateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCurrencyResponse{
		Info: info,
	}, nil
}

func currencyCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
	newConds := cruder.NewConds()

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}

		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.FieldAppID:
			fallthrough //nolint
		case constant.FieldCoinTypeID:
			fallthrough //nolint
		case constant.CurrencyFieldCurrencyMethod:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.CurrencyFieldOverPercent:
			fallthrough //nolint
		case constant.CurrencyFieldPriceVSUSDT:
			fallthrough //nolint
		case constant.CurrencyFieldAppPriceVSUSDT:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetNumberValue())
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}

	return newConds, nil
}

func (s *Server) GetCurrency(ctx context.Context, in *npool.GetCurrencyRequest) (*npool.GetCurrencyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetCurrencyResponse{}, fmt.Errorf("invalid currency id: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get currency: %v", err)
		return &npool.GetCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCurrencyOnly(ctx context.Context, in *npool.GetCurrencyOnlyRequest) (*npool.GetCurrencyOnlyResponse, error) {
	conds, err := currencyCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid currency fields: %v", err)
		return &npool.GetCurrencyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCurrencyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail get currencies: %v", err)
		return &npool.GetCurrencyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	conds, err := currencyCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid currency fields: %v", err)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	conds = conds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := schema.Rows(ctx, conds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get currencies: %v", err)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrenciesResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetAppCurrencies(ctx context.Context, in *npool.GetAppCurrenciesRequest) (*npool.GetAppCurrenciesResponse, error) {
	conds, err := currencyCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid currency fields: %v", err)
		return &npool.GetAppCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	conds = conds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := schema.Rows(ctx, conds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get currencies: %v", err)
		return &npool.GetAppCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCurrenciesResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) DeleteCurrency(ctx context.Context, in *npool.DeleteCurrencyRequest) (*npool.DeleteCurrencyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteCurrencyResponse{}, fmt.Errorf("invalid currency id: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.DeleteCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete currency: %v", err)
		return &npool.DeleteCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) Currency(ctx context.Context, in *npool.CurrencyRequest) (*npool.CurrencyResponse, error) {
	conds := cruder.NewConds().
		WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID()).
		WithCond(constant.FieldCoinTypeID, cruder.EQ, in.GetCoinTypeID())

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CurrencyResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	resp := &npool.CurrencyResponse{
		Info: &npool.CurrencyAmount{
			CoinTypeID: in.GetCoinTypeID(),
		},
	}

	info, err := schema.RowOnly(ctx, conds)
	if err != nil || info == nil {
		logger.Sugar().Errorf("fail get coin currency: %v", err)
		return &npool.CurrencyResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	resp.Info.Amount = info.AppPriceVSUSDT
	return resp, nil
}
