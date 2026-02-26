<template>
  <div class="asset-list-container">
    <h2>내 자산 목록</h2>
    <div class="asset-grid">
      <div v-for="asset in assets" :key="asset.id" class="asset-card">
        <div class="asset-header">
          <span class="asset-icon">{{ getIcon(asset.icon) }}</span>
          <span class="asset-name">{{ asset.name }}</span>
        </div>
        <div class="asset-body">
          <div class="asset-balance">
            {{ formatCurrency(asset.balance, asset.currency) }}
          </div>
          <div class="asset-details">
            <span class="asset-type">{{ asset.type }}</span>
            <span v-if="asset.interest" class="asset-interest">
              {{ asset.interest.rate }}% ({{ asset.interest.type }})
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { mockAssets, Asset } from '../data/mockAssets';

// 반응형 상태로 자산 목록 관리
const assets = ref<Asset[]>(mockAssets);

// 통화 포맷팅 함수 (Intl.NumberFormat 사용)
const formatCurrency = (amount: number, currency: string) => {
  return new Intl.NumberFormat('ko-KR', {
    style: 'currency',
    currency: currency,
  }).format(amount);
};

// 아이콘 매핑 함수 (실제 아이콘 라이브러리 사용 전 임시)
const getIcon = (iconName: string) => {
  const icons: Record<string, string> = {
    'bank': '🏦',
    'safe': '🔒',
    'chart-line': '📈',
    'car': '🚗',
    'home-loan': '🏠',
  };
  return icons[iconName] || '💰';
};
</script>

<style scoped>
.asset-list-container {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

.asset-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.asset-card {
  background-color: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s;
}

.asset-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.asset-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.asset-icon {
  font-size: 1.5em;
  margin-right: 10px;
}

.asset-name {
  font-weight: bold;
  font-size: 1.1em;
  color: #333;
}

.asset-balance {
  font-size: 1.4em;
  font-weight: bold;
  color: #2c3e50;
  margin-bottom: 5px;
}

.asset-details {
  display: flex;
  justify-content: space-between;
  font-size: 0.9em;
  color: #777;
}

.asset-interest {
  color: #e67e22; /* 오렌지색 강조 */
}
</style>
