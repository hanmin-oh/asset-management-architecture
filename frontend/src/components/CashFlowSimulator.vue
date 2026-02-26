<template>
  <div class="simulator-container p-6 bg-white rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-4 text-gray-800">현금 흐름 시뮬레이터</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- 입력 폼 -->
      <div class="input-section space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">목표 날짜</label>
          <input
            type="date"
            v-model="store.simulationParams.targetDate"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">월 예상 수입 (원)</label>
          <input
            type="number"
            v-model.number="store.simulationParams.monthlyIncome"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">월 예상 지출 (원)</label>
          <input
            type="number"
            v-model.number="store.simulationParams.monthlyExpense"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">예상 투자 수익률 (연 %)</label>
          <input
            type="number"
            v-model.number="store.simulationParams.investmentReturnRate"
            step="0.1"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <button
          @click="store.runSimulation"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          시뮬레이션 실행
        </button>
      </div>

      <!-- 결과 표시 -->
      <div class="result-section bg-gray-50 p-4 rounded-md border border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-3">예측 결과</h3>

        <div v-if="store.simulationResult" class="space-y-3">
          <div class="flex justify-between">
            <span class="text-gray-600">목표 시점:</span>
            <span class="font-medium">{{ store.simulationResult.targetDate }}</span>
          </div>

          <div class="flex justify-between">
            <span class="text-gray-600">예상 순자산:</span>
            <span class="font-bold text-indigo-600 text-xl">
              {{ formatCurrency(store.simulationResult.predictedNetWorth) }}
            </span>
          </div>

          <div class="border-t border-gray-200 my-2 pt-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">순수 현금 증가분:</span>
              <span class="text-gray-900">+{{ formatCurrency(store.simulationResult.cashFlow) }}</span>
            </div>
            <div class="flex justify-between text-sm mt-1">
              <span class="text-gray-500">투자 수익 증가분:</span>
              <span class="text-gray-900">+{{ formatCurrency(store.simulationResult.investmentGrowth) }}</span>
            </div>
          </div>
        </div>

        <div v-else class="text-center text-gray-500 py-10">
          왼쪽의 정보를 입력하고<br/>'시뮬레이션 실행' 버튼을 눌러주세요.
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAssetStore } from '../store/useAssetStore';

// Pinia 스토어 사용
const store = useAssetStore();

// 통화 포맷팅 함수 (유틸리티로 분리 가능하지만, 현재는 컴포넌트 내에 둠)
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ko-KR', {
    style: 'currency',
    currency: 'KRW',
  }).format(amount);
};
</script>

<style scoped>
/* Tailwind CSS를 사용하므로 별도의 CSS는 최소화합니다. */
</style>
