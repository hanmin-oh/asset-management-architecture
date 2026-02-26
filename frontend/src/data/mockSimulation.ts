// @/data/mockSimulation.ts

// --- TypeScript Interfaces ---

export interface SimulationParams {
  targetDate: string; // 예측 목표 날짜 (ISO 8601)
  monthlyIncome: number; // 월 예상 수입
  monthlyExpense: number; // 월 예상 지출
  inflationRate: number; // 예상 물가 상승률 (%)
  investmentReturnRate: number; // 예상 투자 수익률 (%)
}

export interface SimulationResult {
  date: string; // 시뮬레이션 시점
  totalAssets: number; // 총 자산 예측액
  liquidCash: number; // 가용 현금 예측액
  netWorth: number; // 순자산 예측액
}

// --- Mock Data ---
// 초기 시뮬레이션 결과 예시

export const initialSimulationResult: SimulationResult = {
  date: new Date().toISOString().split('T')[0], // 오늘 날짜
  totalAssets: 150000000,
  liquidCash: 3500000,
  netWorth: 120000000,
};
