package transaction

import (
	"time"
)

// TransactionType은 거래 유형을 정의합니다.
type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "DEPOSIT"    // 입금
	TransactionTypeWithdrawal TransactionType = "WITHDRAWAL" // 출금
	TransactionTypeTransfer   TransactionType = "TRANSFER"   // 이체
	TransactionTypeInterest   TransactionType = "INTEREST"   // 이자 지급
	TransactionTypeFee        TransactionType = "FEE"        // 수수료
)

// TransactionStatus는 거래 상태를 정의합니다.
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"   // 대기 중
	TransactionStatusCompleted TransactionStatus = "COMPLETED" // 완료됨
	TransactionStatusFailed    TransactionStatus = "FAILED"    // 실패함
	TransactionStatusCancelled TransactionStatus = "CANCELLED" // 취소됨
)

// Transaction은 자산의 변동 내역을 기록하는 엔티티입니다.
// Google Style: 불변성을 지향하며, 생성 후에는 상태 변경을 최소화합니다.
type Transaction struct {
	ID            int64
	AssetID       int64             // 연관된 자산 ID
	Type          TransactionType   // 거래 유형
	Amount        int64             // 거래 금액 (양수: 입금, 음수: 출금)
	Currency      string            // 통화 코드
	Description   string            // 거래 설명
	Status        TransactionStatus // 거래 상태
	TransactionAt time.Time         // 실제 거래 발생 시각
	CreatedAt     time.Time         // 시스템 기록 시각
	UpdatedAt     time.Time         // 최종 수정 시각
}

// IsValid는 거래 데이터의 유효성을 검사합니다.
// 도메인 규칙: 금액은 0이 될 수 없으며, 통화 코드는 필수입니다.
func (t *Transaction) IsValid() bool {
	if t.Amount == 0 {
		return false
	}
	if t.Currency == "" {
		return false
	}
	return true
}
