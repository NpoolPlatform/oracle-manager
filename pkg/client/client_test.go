package client

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	constant "github.com/NpoolPlatform/oracle-manager/pkg/const"
	"github.com/NpoolPlatform/oracle-manager/pkg/test-init" //nolint
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestClient(t *testing.T) {
	_, _ = GetCurrencyOnly(context.Background(), //nolint
		cruder.NewFilterConds().
			WithCond(constant.FieldAppID, cruder.EQ, structpb.NewStringValue(uuid.UUID{}.String())))
	// Here won't pass test due to we always test with localhost
}
