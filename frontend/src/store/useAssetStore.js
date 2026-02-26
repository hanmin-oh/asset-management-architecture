import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { mockAssets } from '../data/mockAssets'; // 이전에 정의한 Mock 데이터를 가져옵니다.

// defineStore의 첫 번째 인자는 스토어의 고유 ID입니다.
export const useAssetStore = defineStore('assetStore', () => {
  // --- State ---
  // ref()는 state 속성이 됩니다.
  const assets = ref(mockAssets);
  const simulationParams = ref({
    targetDate: new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0], // 기본값: 1년 후
    monthlyIncome: 5000000,
    monthlyExpense: 3000000,
    investmentReturnRate: 5, // 연 5%
  });
  const simulationResult = ref(null); // 시뮬레이션 결과 저장 (초기값 null)

  // --- Getters ---
  // computed()는 getters가 됩니다.
  const totalAssets = computed(() =>
    assets.value
      .filter(a => a.balance > 0)
      .reduce((sum, asset) => sum + asset.balance, 0)
  );

  const totalLiabilities = computed(() =>
    assets.value
      .filter(a => a.balance < 0)
      .reduce((sum, asset) => sum + asset.balance, 0)
  );

  const netWorth = computed(() => totalAssets.value + totalLiabilities.value);

  const liquidCash = computed(() =>
    assets.value
      .filter(a => a.type === 'Bank' && a.subCategory === 'Checking' && !a.isLocked)
      .reduce((sum, asset) => sum + asset.balance, 0)
  );

  // --- Actions ---
  // function()은 actions가 됩니다.
  function runSimulation() {
    // 지금은 실제 계산 대신 가짜 로직을 사용합니다.
    // 목표: 입력된 파라미터를 기반으로 미래의 자산을 예측하는 로직 구현
    const monthlyNetIncome = simulationParams.value.monthlyIncome - simulationParams.value.monthlyExpense;
    const months = 12; // 단순화를 위해 1년(12개월)으로 고정

    // 1년 후 순수 현금 증가액
    const futureCashIncrease = monthlyNetIncome * months;

    // 1년 후 투자 자산 증가액 (단순 연 복리 계산)
    const currentInvestments = assets.value
      .filter(a => a.type === 'Stock')
      .reduce((sum, asset) => sum + asset.balance, 0);
    
    const futureInvestmentValue = currentInvestments * (1 + simulationParams.value.investmentReturnRate / 100);

    // 최종 예측 결과
    const futureNetWorth = netWorth.value + futureCashIncrease + (futureInvestmentValue - currentInvestments);

    // simulationResult state를 업데이트합니다.
    simulationResult.value = {
      targetDate: simulationParams.value.targetDate,
      // 예측된 순자산, 총자산 등을 여기에 담습니다.
      // 지금은 간단히 예측된 순자산만 보여줍니다.
      predictedNetWorth: Math.round(futureNetWorth),
      cashFlow: futureCashIncrease,
      investmentGrowth: Math.round(futureInvestmentValue - currentInvestments),
    };
    
    console.log('Simulation Executed:', simulationResult.value);
  }

  // 스토어에서 외부로 노출할 state, getters, actions를 반환합니다.
  return {
    // State
    assets,
    simulationParams,
    simulationResult,
    // Getters
    totalAssets,
    totalLiabilities,
    netWorth,
    liquidCash,
    // Actions
    runSimulation,
  };
});
