// @/data/mockAssets.ts

// --- TypeScript Interfaces ---
// 백엔드(Go)의 데이터 모델과 거의 1:1로 매칭되는 구조입니다.
// 프론트엔드와 백엔드 간의 데이터 계약(Data Contract) 역할을 합니다.

export interface InterestInfo {
  rate: number; // 연 이자율 (예: 3.5)
  type: 'FIXED' | 'VARIABLE'; // 고정/변동 금리
  nextReset?: string; // 다음 금리 갱신일 (ISO 8601 형식)
}

export interface Asset {
  id: string; // UUID 또는 DB의 PK
  name: string; // 자산 이름 (예: "국민은행 월급통장")
  type: 'Bank' | 'Stock' | 'RealEstate' | 'Loan'; // 자산 대분류
  subCategory: 'Checking' | 'Savings' | 'Stock' | 'Mortgage' | 'Credit'; // 소분류
  
  // 중요: 금액은 float 대신 정수(원 단위)로 다루어 오차를 방지합니다.
  balance: number; 
  currency: 'KRW' | 'USD';

  interest?: InterestInfo; // 이자 정보 (Optional)

  // 유동성 관련 정보
  isLocked: boolean; // 출금/매도 제한 여부
  lockedUntil?: string; // 제한 해제일 (ISO 8601 형식)
  
  // 화면 표시에 필요한 추가 정보
  icon: string; // UI에 표시할 아이콘 이름 (예: 'bank', 'chart-line')
}


// --- Mock Data ---
// 이 데이터를 기반으로 Vue 컴포넌트를 개발합니다.

export const mockAssets: Asset[] = [
  {
    id: 'asset-001',
    name: 'KB국민은행 급여통장',
    type: 'Bank',
    subCategory: 'Checking',
    balance: 3500000,
    currency: 'KRW',
    isLocked: false,
    icon: 'bank',
  },
  {
    id: 'asset-002',
    name: '카카오뱅크 세이프박스',
    type: 'Bank',
    subCategory: 'Savings',
    balance: 15000000,
    currency: 'KRW',
    interest: {
      rate: 2.1,
      type: 'VARIABLE',
    },
    isLocked: false,
    icon: 'safe',
  },
  {
    id: 'asset-003',
    name: '삼성전자',
    type: 'Stock',
    subCategory: 'Stock',
    balance: 8200000, // 현재 평가액
    currency: 'KRW',
    isLocked: false,
    icon: 'chart-line',
  },
  {
    id: 'asset-004',
    name: 'Tesla',
    type: 'Stock',
    subCategory: 'Stock',
    balance: 12000 * 1250, // $12,000, 환율 1250원 가정
    currency: 'USD',
    isLocked: false,
    icon: 'car',
  },
  {
    id: 'asset-005',
    name: '주택담보대출',
    type: 'Loan',
    subCategory: 'Mortgage',
    balance: -150000000, // 부채는 음수로 표현
    currency: 'KRW',
    interest: {
      rate: 3.8,
      type: 'FIXED',
    },
    isLocked: true, // 대출은 상환 외에 다른 거래 불가
    icon: 'home-loan',
  },
];
