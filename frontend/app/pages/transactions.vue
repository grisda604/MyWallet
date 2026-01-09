<template>
  <div class="p-8 bg-gray-50 min-h-screen">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-3xl font-bold text-gray-800">📝 Transactions</h1>
        <p class="text-gray-600 mt-1">Manage all your transactions</p>
      </div>
      <button
        @click="openTransactionModal"
        class="flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-blue-500 to-blue-600 text-white font-semibold rounded-lg shadow-md hover:shadow-lg hover:from-blue-600 hover:to-blue-700 transition-all duration-200 transform hover:scale-105"
      >
        <span class="text-xl">+</span>
        <span>เพิ่ม Transaction</span>
      </button>
    </div>

    <!-- Filters Section -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Search</label
          >
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search transactions..."
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Type</label
          >
          <select
            v-model="filterType"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">All Types</option>
            <option value="income">Income</option>
            <option value="expense">Expense</option>
            <option value="transfer">Transfer</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Wallet</label
          >
          <select
            v-model="filterWallet"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">All Wallets</option>
            <option
              v-for="wallet in wallets"
              :key="wallet.ID"
              :value="wallet.ID"
            >
              {{ wallet.name }}
            </option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Category</label
          >
          <select
            v-model="filterCategory"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">All Categories</option>
            <option
              v-for="category in categories"
              :key="category.ID"
              :value="category.ID"
            >
              {{ category.icon }} {{ category.name }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- Transaction List -->
    <TransactionList
      :transactions="filteredTransactions"
      :loading="transactionsPending"
      :error="transactionsError"
    />

    <!-- Transaction Modal -->
    <TransactionModal
      v-model="isTransactionModalOpen"
      :wallets="wallets || []"
      :categories="categories || []"
      @transaction-created="handleTransactionCreated"
    />
  </div>
</template>

<script setup lang="ts">
import TransactionList from "~/components/transactions/TransactionList.vue";
import TransactionModal from "~/components/transactions/TransactionModal.vue";
import type { Wallet, Transaction, Category } from "~/types";

// Response type for transactions API
interface TransactionsResponse {
  transactions: Transaction[];
}

// Fetch data
const {
  data: wallets,
  pending: walletsPending,
  error: walletsError,
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

// Filter states
const searchQuery = ref("");
const filterType = ref("");
const filterWallet = ref("");
const filterCategory = ref("");

// Filtered transactions
const filteredTransactions = computed(() => {
  let filtered = transactions.value;

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(
      (t) =>
        t.note?.toLowerCase().includes(query) ||
        t.category?.name.toLowerCase().includes(query) ||
        t.wallet?.name.toLowerCase().includes(query)
    );
  }

  // Type filter
  if (filterType.value) {
    filtered = filtered.filter((t) => t.type === filterType.value);
  }

  // Wallet filter
  if (filterWallet.value) {
    filtered = filtered.filter(
      (t) => t.wallet_id === Number(filterWallet.value)
    );
  }

  // Category filter
  if (filterCategory.value) {
    filtered = filtered.filter(
      (t) => t.category_id === Number(filterCategory.value)
    );
  }

  return filtered;
});

// Modal state
const isTransactionModalOpen = ref(false);

// Functions
const openTransactionModal = () => {
  isTransactionModalOpen.value = true;
};

const handleTransactionCreated = async () => {
  await refreshNuxtData();
};
</script>
