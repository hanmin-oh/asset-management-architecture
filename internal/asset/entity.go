package asset

import (
	"time"
)

// AssetType은 자산의 대분류를 정의합니다.
// Google Style: 열거형은 type alias와 iota를 사용하여 정의합니다.
type AssetType int

const (
	AssetTypeUnknown AssetType = iota
	AssetTypeBank              // 은행 (예적금)
	AssetTypeStock             // 주식/증권
	AssetTypeRealEstate        // 부동산
	AssetTypeLoan              // 대출 (부채)
)

// SubCategory는 자산의 구체적인 성격을 정의합니다.
// 예: 입출금통장, 예금, 적금, 주택담보대출 등
type SubCategory string

const (
	SubCategoryChecking SubCategory = "CHECKING" // 수시입출금
	SubCategorySavings  SubCategory = "SAVINGS"  // 예적금
	SubCategoryStock    SubCategory = "STOCK"    // 주식 현물
	SubCategoryMortgage SubCategory = "MORTGAGE" // 주택담보대출
	SubCategoryCredit   SubCategory = "CREDIT"   // 신용대출
)

// InterestType은 금리 유형(고정/변동)을 정의합니다.
type InterestType string

const (
	InterestTypeFixed    InterestType = "FIXED"
	InterestTypeVariable InterestType = "VARIABLE"
)

// InterestInfo는 자산(주로 예적금 및 대출)의 금리 정보를 담습니다.
type InterestInfo struct {
	Rate      float64      // 연 이자율 (예: 3.5 = 3.5%)
	Type      InterestType // 고정/변동 여부
	NextReset *time.Time   // 변동금리일 경우 다음 금리 갱신일 (Optional)
}

// Asset은 시스템의 핵심 도메인 엔티티입니다.
// ID는 DB의 Primary Key와 매핑되며, 외부 노출 시 UUID 사용을 권장합니다.
type Asset struct {
	ID   int64  // 내부 식별자
	Name string // 자산 이름 (예: "국민은행 월급통장")

	Type        AssetType
	SubCategory SubCategory

	// Balance는 자산의 현재 잔액을 나타냅니다.
	// 중요: 금융 계산의 정확성을 위해 float64 대신 int64(원 단위)를 사용합니다.
	Balance int64

	// Currency는 통화 코드를 나타냅니다 (ISO 4217, 예: "KRW", "USD").
	Currency string

	// Interest는 이자 정보가 있는 경우에만 설정됩니다 (Nullable).
	Interest *InterestInfo

	// Liquidity 정보
	IsLocked    bool       // 출금/매도 제한 여부 (예: 예금 담보 설정 등)
	LockedUntil *time.Time // 언제까지 묶여있는지 (예: 적금 만기일)

	CreatedAt time.Time
	UpdatedAt time.Time
}

// IsLiquid는 현재 시점에서 즉시 현금화가 가능한지 판단합니다.
// 도메인 로직은 엔티티 내부에 위치하여 응집도를 높입니다.
func (a *Asset) IsLiquid() bool {
	if a.IsLocked {
		if a.LockedUntil != nil && a.LockedUntil.After(time.Now()) {
			return false
		}
	}
	// 부동산 등은 즉시 현금화 불가 자산으로 분류할 수 있음
	if a.Type == AssetTypeRealEstate {
		return false
	}
	return true
}
