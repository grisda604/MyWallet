<template>
  <div class="p-8 bg-gray-50 min-h-screen">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-800">📊 Dashboard</h1>
      <p class="text-gray-600 mt-1">Overview of your financial status</p>
    </div>

    <!-- Loading State -->
    <div v-if="pending || transactionsPending" class="text-gray-500">
      Loading dashboard...
    </div>

    <!-- Error State -->
    <div v-else-if="error || transactionsError" class="text-red-500">
      Error loading data. Make sure Go Backend is running on port 8081
    </div>

    <!-- Dashboard Content -->
    <div v-else>
      <!-- Summary Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <StatCard
          title="Total Balance"
          :value="totalBalance"
          icon="💰"
          color="blue"
        />
        <StatCard
          title="Monthly Income"
          :value="monthlyIncome"
          icon="📈"
          color="green"
        />
        <StatCard
          title="Monthly Expense"
          :value="monthlyExpense"
          icon="📉"
          color="red"
        />
        <StatCard
          title="Transactions"
          :value="transactions.length"
          icon="📝"
          color="purple"
          prefix=""
        />
      </div>

      <!-- Charts Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <!-- Wallet Distribution Chart -->
        <WalletDistributionChart :wallets="wallets || []" />

        <!-- Income vs Expense Chart -->
        <IncomeExpenseChart :transactions="transactions" />
      </div>

      <!-- Category Spending Chart - Full Width -->
      <div class="mb-8">
        <CategorySpendingChart :transactions="transactions" />
      </div>

      <!-- Wallets Section -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-2xl font-bold text-gray-800">💼 My Wallets</h2>
          <button
            @click="openWalletModal"
            class="flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-blue-500 to-blue-600 text-white font-semibold rounded-lg shadow-md hover:shadow-lg hover:from-blue-600 hover:to-blue-700 transition-all duration-200 transform hover:scale-105"
          >
            <span class="text-xl">+</span>
            <span>Add Wallet</span>
          </button>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4">
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
        </div>
      </div>

      <!-- Recent Transactions -->
      <div>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-2xl font-bold text-gray-800">
            📝 Recent Transactions
          </h2>
          <NuxtLink
            to="/transactions"
            class="text-blue-600 hover:text-blue-700 font-semibold flex items-center gap-1 transition"
          >
            View All
            <span>→</span>
          </NuxtLink>
        </div>
        <TransactionList
          :transactions="recentTransactions"
          :loading="false"
          :error="null"
        />
      </div>
    </div>

    <!-- Wallet Modal -->
    <WalletModal
      v-model="isWalletModalOpen"
      @wallet-created="handleWalletCreated"
    />
  </div>
</template>

<script setup lang="ts">
import StatCard from "~/components/dashboard/StatCard.vue";
import WalletDistributionChart from "~/components/dashboard/WalletDistributionChart.vue";
import IncomeExpenseChart from "~/components/dashboard/IncomeExpenseChart.vue";
import CategorySpendingChart from "~/components/dashboard/CategorySpendingChart.vue";
import TransactionList from "~/components/transactions/TransactionList.vue";
import WalletModal from "~/components/wallets/WalletModal.vue";
import type { Wallet, Transaction, Category } from "~/types";

// Response type for transactions API
interface TransactionsResponse {
  transactions: Transaction[];
}

// Fetch data
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

// Computed values for stats
const totalBalance = computed(() => {
  return wallets.value?.reduce((sum, wallet) => sum + wallet.balance, 0) || 0;
});

const monthlyIncome = computed(() => {
  const now = new Date();
  const currentMonth = now.getMonth();
  const currentYear = now.getFullYear();

  return transactions.value
    .filter((t) => {
      const date = new Date(t.date || t.CreatedAt);
      return (
        t.type === "income" &&
        date.getMonth() === currentMonth &&
        date.getFullYear() === currentYear
      );
    })
    .reduce((sum, t) => sum + t.amount, 0);
});

const monthlyExpense = computed(() => {
  const now = new Date();
  const currentMonth = now.getMonth();
  const currentYear = now.getFullYear();

  return transactions.value
    .filter((t) => {
      const date = new Date(t.date || t.CreatedAt);
      return (
        t.type === "expense" &&
        date.getMonth() === currentMonth &&
        date.getFullYear() === currentYear
      );
    })
    .reduce((sum, t) => sum + Math.abs(t.amount), 0);
});

// Recent transactions (last 5)
const recentTransactions = computed(() => {
  return [...transactions.value]
    .sort((a, b) => {
      const dateA = new Date(a.date || a.CreatedAt).getTime();
      const dateB = new Date(b.date || b.CreatedAt).getTime();
      return dateB - dateA;
    })
    .slice(0, 5);
});

// Modal state
const isWalletModalOpen = ref(false);

// Functions
const openWalletModal = () => {
  isWalletModalOpen.value = true;
};

const handleWalletCreated = async () => {
  await refreshNuxtData();
};
</script>
