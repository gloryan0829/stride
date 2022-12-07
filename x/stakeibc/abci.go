package stakeibc

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/Stride-Labs/stride/v4/x/stakeibc/keeper"
	"github.com/Stride-Labs/stride/v4/x/stakeibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/**
 * @Analysis
 * Begin Block 마다 stakeibc 모듈에서는 redemption rate 를 90% 미만 또는 150% 이상 넘는지 각 HostZone 에서 체크하는데 이때
 * Redemption Rate 라는 것은 stToken 과 실제 nativeToken 간에 비율임.. stToken 은 share 라고 보면 되고 nativeToken 은 결국 reward 가
 * 계속 누적된 총 treasury pool 이라고 보면 될거같음
 */
// BeginBlocker of stakeibc module
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all host zones and verify redemption rate
	for _, hz := range k.GetAllHostZone(ctx) {
		rrSafe, err := k.IsRedemptionRateWithinSafetyBounds(ctx, hz)
		if !rrSafe {
			panic(fmt.Sprintf("[INVARIANT BROKEN!!!] %s's RR is %s. ERR: %v", hz.GetChainId(), hz.RedemptionRate.String(), err.Error()))
		}
	}
}
