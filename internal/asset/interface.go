package asset

import (
	"context"
	"time"
)

// LiquidityEvaluator는 자산의 실질 가용 현금(Real Liquidity)을 계산하는 인터페이스입니다.
// 단순 잔고(Balance)가 아닌, 출금 가능 여부, 만기일, 결제일(T+2) 등을 고려합니다.
//
// Google Style: 인터페이스 이름은 '-er'로 끝나거나, 행위를 명확히 나타내야 합니다.
// 여기서는 'Evaluator'를 사용하여 평가 행위를 강조합니다.
type LiquidityEvaluator interface {
	// CalculateLiquidCash는 특정 시점(asOf) 기준으로 즉시 사용 가능한 현금을 계산합니다.
	// context를 통해 타임아웃 및 취소 처리를 지원해야 합니다.
	CalculateLiquidCash(ctx context.Context, assets []Asset, asOf time.Time) (int64, error)

	// EstimateFutureLiquidity는 미래 특정 시점의 예상 가용 현금을 추정합니다.
	// 예: 적금 만기일 이후의 현금 흐름 예측
	EstimateFutureLiquidity(ctx context.Context, assets []Asset, targetDate time.Time) (int64, error)
}

// AssetRepository는 자산 데이터에 대한 영속성 계층(Persistence Layer) 인터페이스입니다.
// 도메인 로직은 Repository 구현체(Postgres 등)에 의존하지 않고 인터페이스에 의존합니다 (DIP).
type AssetRepository interface {
	// FindByID는 ID로 자산을 조회합니다.
	FindByID(ctx context.Context, id int64) (*Asset, error)

	// FindAllByUserID는 특정 사용자의 모든 자산을 조회합니다.
	FindAllByUserID(ctx context.Context, userID int64) ([]Asset, error)

	// Save는 자산을 생성하거나 업데이트합니다.
	Save(ctx context.Context, asset *Asset) error

	// Delete는 자산을 삭제합니다 (Soft Delete 권장).
	Delete(ctx context.Context, id int64) error
}

// AssetService는 자산 도메인의 유스케이스(Use Case)를 정의합니다.
// 외부(API 핸들러 등)에서는 이 인터페이스를 통해 도메인 로직을 실행합니다.
type AssetService interface {
	// GetPortfolioSummary는 사용자의 전체 자산 포트폴리오 요약을 반환합니다.
	GetPortfolioSummary(ctx context.Context, userID int64) (*PortfolioSummary, error)

	// RegisterAsset은 새로운 자산을 등록합니다.
	RegisterAsset(ctx context.Context, cmd RegisterAssetCommand) (*Asset, error)
}

// PortfolioSummary는 포트폴리오 요약 정보를 담는 DTO(Data Transfer Object)입니다.
type PortfolioSummary struct {
	TotalAssets      int64 // 총 자산
	TotalLiabilities int64 // 총 부채
	NetWorth         int64 // 순자산 (자산 - 부채)
	LiquidCash       int64 // 실질 가용 현금
}

// RegisterAssetCommand는 자산 등록에 필요한 입력 데이터입니다.
type RegisterAssetCommand struct {
	UserID      int64
	Name        string
	Type        AssetType
	SubCategory SubCategory
	Balance     int64
	Currency    string
}
