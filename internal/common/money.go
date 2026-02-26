package common

import (
	"fmt"
	"math"
)

// Money는 화폐 단위를 안전하게 다루기 위한 값 객체(Value Object)입니다.
// Google Style: 불변성(Immutability)을 지향하며, 연산 메서드는 새로운 Money 객체를 반환합니다.
type Money struct {
	Amount   int64  // 최소 단위 (예: 원, 센트)
	Currency string // ISO 4217 통화 코드 (예: KRW, USD)
}

// NewMoney는 새로운 Money 객체를 생성합니다.
func NewMoney(amount int64, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

// Add는 두 Money 객체를 더합니다. 통화가 다르면 에러를 반환합니다.
func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("cannot add different currencies: %s and %s", m.Currency, other.Currency)
	}
	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

// Subtract는 두 Money 객체를 뺍니다.
func (m Money) Subtract(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("cannot subtract different currencies: %s and %s", m.Currency, other.Currency)
	}
	return Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

// Multiply는 Money 객체에 스칼라 값을 곱합니다.
// 예: 이자 계산 시 사용 (원금 * 이자율)
func (m Money) Multiply(factor float64) Money {
	// 반올림 처리 (Round Half Up)
	amount := int64(math.Round(float64(m.Amount) * factor))
	return Money{
		Amount:   amount,
		Currency: m.Currency,
	}
}

// String은 Money 객체의 문자열 표현을 반환합니다.
func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}
