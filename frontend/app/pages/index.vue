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
        class="border-2 border-dashed border-gray-300 rounded-xl p-6 flex items-center justify-center text-gray-400 hover:border-blue-500 hover:text-blue-500 transition"
      >
        + Add New Wallet
      </button>
    </div>

    <!-- Transactions Section -->
    <h2 class="text-2xl font-bold mt-10 mb-4">📝 Recent Transactions</h2>

    <div v-if="transactionsPending" class="text-gray-500">
      Loading transactions...
    </div>

    <div v-else-if="transactionsError" class="text-red-500">
      Error loading transactions: {{ transactionsError.message }}
    </div>

    <div
      v-else-if="!transactions || transactions.length === 0"
      class="text-gray-500"
    >
      No transactions yet. Start adding some!
    </div>

    <div
      v-else
      class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden"
    >
      <div
        v-for="transaction in transactions"
        :key="transaction.ID"
        class="flex items-center justify-between p-4 border-b border-gray-100 last:border-b-0 hover:bg-gray-50 transition"
      >
        <!-- ซ้าย: Icon + ข้อมูล -->
        <div class="flex items-center gap-4 flex-1">
          <!-- Icon based on type -->
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center text-lg flex-shrink-0"
            :class="{
              'bg-green-100': transaction.type === 'income',
              'bg-red-100': transaction.type === 'expense',
              'bg-blue-100': transaction.type === 'transfer',
            }"
          >
            <span v-if="transaction.type === 'income'">💵</span>
            <span v-else-if="transaction.type === 'expense'">💸</span>
            <span v-else>🔄</span>
          </div>
          <!-- ข้อมูล Transaction -->
          <div class="flex-1 min-w-0">
            <!-- บรรทัดแรก: ชื่อ + Category Badge -->
            <div class="flex items-center gap-2 mb-1">
              <span class="font-medium text-gray-800">
                {{
                  transaction.note ||
                  transaction.category?.name ||
                  "No description"
                }}
              </span>

              <!-- Category Badge -->
              <span
                v-if="transaction.category"
                class="text-xs px-2 py-0.5 rounded-full border flex-shrink-0"
                :style="{
                  backgroundColor: transaction.category.color + '20',
                  borderColor: transaction.category.color,
                  color: transaction.category.color,
                }"
              >
                {{ transaction.category.icon }} {{ transaction.category.name }}
              </span>
            </div>
            <!-- บรรทัดสอง: Wallet info -->
            <div class="text-sm text-gray-500">
              {{ transaction.wallet?.name || "Unknown Wallet" }}
              <span
                v-if="
                  transaction.type === 'transfer' && transaction.target_wallet
                "
              >
                → {{ transaction.target_wallet.name }}
              </span>
            </div>
          </div>
        </div>
        <!-- ขวา: ยอดเงิน + วันที่ -->
        <div class="text-right flex-shrink-0 ml-4">
          <div
            class="font-semibold"
            :class="{
              'text-green-600': transaction.type === 'income',
              'text-red-500': transaction.type === 'expense',
              'text-blue-600': transaction.type === 'transfer',
            }"
          >
            <span v-if="transaction.type === 'income'">+</span>
            <span v-else-if="transaction.type === 'expense'">-</span>
            ฿{{ Math.abs(transaction.amount).toLocaleString() }}
          </div>
          <div class="text-xs text-gray-400">
            {{
              new Date(
                transaction.date || transaction.CreatedAt
              ).toLocaleDateString("th-TH")
            }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
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
</script>
