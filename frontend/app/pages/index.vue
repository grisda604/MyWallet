<template>
  <div class="p-8 bg-gray-50 min-h-screen">
    <h1 class="text-3xl font-bold mb-6">💰 My Wallets</h1>

    <div v-if="pending" class="text-gray-500">Loading wallets...</div>

    <div v-else-if="error" class="text-red-500">
      Error loading wallets: {{ error.message }}
      <p class="text-sm mt-2">Make sure Go Backend is running on port 8080</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div
        v-for="wallet in wallets"
        :key="wallet.ID"
        class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition"
      >
        <div class="flex items-center justify-between mb-4">
          <span class="text-lg text-gray-800 font-semibold">{{
            wallet.name
          }}</span>
          <span class="text-xs px-2 py-1 rounded bg-gray-100 uppercase">{{
            wallet.type
          }}</span>
        </div>

        <div
          class="text-2xl font-bold"
          :class="{
            'text-red-500': wallet.balance < 0,
            'text-green-600': wallet.balance >= 0,
          }"
        >
          ฿{{ wallet.balance.toLocaleString() }}
        </div>
      </div>
      <button
        @click="openWalletModal"
        class="border-2 border-dashed border-gray-300 rounded-xl p-6 flex items-center justify-center text-gray-400 hover:border-blue-500 hover:text-blue-500 transition"
      >
        + Add New Wallet
      </button>
    </div>

    <!-- Transactions Section -->
    <div class="flex items-center justify-between mt-10 mb-4 space-x-4">
      <h2 class="text-lg font-bold">📝 Recent Transactions</h2>
      <button
        @click="openTransactionModal"
        class="flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-blue-500 to-blue-600 text-white font-semibold rounded-lg shadow-md hover:shadow-lg hover:from-blue-600 hover:to-blue-700 transition-all duration-200 transform hover:scale-105"
      >
        <span class="text-xl">+</span>
        <span>เพิ่ม Transaction</span>
      </button>
    </div>

    <!-- ใช้ TransactionList Component -->
    <TransactionList
      :transactions="transactions"
      :loading="transactionsPending"
      :error="transactionsError"
    />

    <!-- ใช้ TransactionModal Component -->
    <TransactionModal
      v-model="isTransactionModalOpen"
      :wallets="wallets || []"
      :categories="categories || []"
      @transaction-created="handleTransactionCreated"
    />

    <!-- ใช้ WalletModal Component -->
    <WalletModal
      v-model="isWalletModalOpen"
      @wallet-created="handleWalletCreated"
    />
  </div>
</template>

<script setup lang="ts">
import TransactionList from "~/components/transactions/TransactionList.vue";
import TransactionModal from "~/components/transactions/TransactionModal.vue";
import WalletModal from "~/components/wallets/WalletModal.vue";
import type { Wallet, Transaction, Category } from "~/types";

// Response type for transactions API
interface TransactionsResponse {
  transactions: Transaction[];
}

// เรียก API ไปที่ Go Backend
const {
  data: wallets,
  pending,
  error,
} = await useFetch<Wallet[]>("http://localhost:8081/api/wallets");

const {
  data: transactionsData,
  pending: transactionsPending,
  error: transactionsError,
} = await useFetch<TransactionsResponse>(
  "http://localhost:8081/api/transactions"
);

const {
  data: categories,
  pending: categoriesPending,
  error: categoriesError,
} = await useFetch<Category[]>("http://localhost:8081/api/categories");

// Extract transactions array from response
const transactions = computed(() => transactionsData.value?.transactions || []);

// Modal states
const isTransactionModalOpen = ref(false);
const isWalletModalOpen = ref(false);

// ฟังก์ชันเปิด transaction modal
const openTransactionModal = () => {
  isTransactionModalOpen.value = true;
};

// ฟังก์ชันเปิด wallet modal
const openWalletModal = () => {
  isWalletModalOpen.value = true;
};

// ฟังก์ชัน handle เมื่อสร้าง transaction สำเร็จ
const handleTransactionCreated = async () => {
  // Refresh ข้อมูล wallets และ transactions
  await refreshNuxtData();
};

// ฟังก์ชัน handle เมื่อสร้าง wallet สำเร็จ
const handleWalletCreated = async () => {
  // Refresh ข้อมูล wallets
  await refreshNuxtData();
};
</script>
