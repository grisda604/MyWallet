<template>
  <div class="p-8 bg-gray-50 min-h-screen">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-3xl font-bold text-gray-800">💼 Wallets</h1>
        <p class="text-gray-600 mt-1">Manage your wallets and accounts</p>
      </div>
      <button
        @click="openWalletModal"
        class="flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-blue-500 to-blue-600 text-white font-semibold rounded-lg shadow-md hover:shadow-lg hover:from-blue-600 hover:to-blue-700 transition-all duration-200 transform hover:scale-105"
      >
        <span class="text-xl">+</span>
        <span>Add Wallet</span>
      </button>
    </div>

    <!-- Search and Filter -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Search</label
          >
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search wallets..."
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
            <option value="bank">Bank</option>
            <option value="cash">Cash</option>
            <option value="e-wallet">E-Wallet</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2"
            >Sort By</label
          >
          <select
            v-model="sortBy"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="name">Name</option>
            <option value="balance-high">Balance (High to Low)</option>
            <option value="balance-low">Balance (Low to High)</option>
            <option value="type">Type</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="pending" class="text-gray-500 text-center py-8">
      Loading wallets...
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-red-500 text-center py-8">
      Error loading wallets. Make sure Go Backend is running on port 8081
    </div>

    <!-- Wallets Grid -->
    <div
      v-else-if="filteredWallets.length === 0"
      class="text-gray-500 text-center py-8"
    >
      No wallets found. Create your first wallet!
    </div>

    <div
      v-else
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6"
    >
      <div
        v-for="wallet in filteredWallets"
        :key="wallet.ID"
        class="bg-white rounded-xl shadow-sm border border-gray-100 hover:shadow-lg transition-all duration-300 cursor-pointer group"
        @click="viewWalletDetails(wallet)"
      >
        <!-- Wallet Card Header -->
        <div class="p-6 border-b border-gray-100">
          <div class="flex items-start justify-between mb-4">
            <div
              class="w-12 h-12 rounded-lg flex items-center justify-center text-2xl"
              :style="{ backgroundColor: wallet.color + '20' }"
            >
              {{ wallet.icon || "💰" }}
            </div>
            <span
              class="text-xs px-3 py-1 rounded-full font-medium uppercase"
              :class="{
                'bg-blue-100 text-blue-700': wallet.type === 'bank',
                'bg-green-100 text-green-700': wallet.type === 'cash',
                'bg-purple-100 text-purple-700': wallet.type === 'e-wallet',
              }"
            >
              {{ wallet.type }}
            </span>
          </div>
          <h3 class="text-lg font-bold text-gray-800 mb-1">
            {{ wallet.name }}
          </h3>
          <p class="text-xs text-gray-500">{{ wallet.currency || "THB" }}</p>
        </div>

        <!-- Wallet Balance -->
        <div class="p-6">
          <div
            class="text-3xl font-bold mb-4"
            :class="{
              'text-red-500': wallet.balance < 0,
              'text-green-600': wallet.balance >= 0,
            }"
          >
            ฿{{ wallet.balance.toLocaleString() }}
          </div>

          <!-- Quick Actions -->
          <div class="flex gap-2">
            <button
              @click.stop="editWallet(wallet)"
              class="flex-1 px-3 py-2 bg-blue-50 text-blue-600 text-sm font-medium rounded-lg hover:bg-blue-100 transition"
            >
              Edit
            </button>
            <button
              @click.stop="deleteWallet(wallet)"
              class="flex-1 px-3 py-2 bg-red-50 text-red-600 text-sm font-medium rounded-lg hover:bg-red-100 transition"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Wallet Modal -->
    <WalletModal
      v-model="isWalletModalOpen"
      :wallet="selectedWallet"
      @wallet-created="handleWalletCreated"
      @wallet-updated="handleWalletUpdated"
    />

    <!-- Wallet Detail Modal -->
    <WalletDetailModal
      v-model="isDetailModalOpen"
      :wallet="selectedWallet"
      :transactions="walletTransactions"
    />
  </div>
</template>

<script setup lang="ts">
import WalletModal from "~/components/wallets/WalletModal.vue";
import WalletDetailModal from "~/components/wallets/WalletDetailModal.vue";
import type { Wallet, Transaction } from "~/types";

// Response type for transactions API
interface TransactionsResponse {
  transactions: Transaction[];
}

// Fetch wallets
const {
  data: wallets,
  pending,
  error,
} = await useFetch<Wallet[]>("http://localhost:8081/api/wallets");

// Fetch transactions for wallet details
const { data: transactionsData } = await useFetch<TransactionsResponse>(
  "http://localhost:8081/api/transactions"
);

const transactions = computed(() => transactionsData.value?.transactions || []);

// Filter and search states
const searchQuery = ref("");
const filterType = ref("");
const sortBy = ref("name");

// Modal states
const isWalletModalOpen = ref(false);
const isDetailModalOpen = ref(false);
const selectedWallet = ref<Wallet | null>(null);

// Filtered wallets
const filteredWallets = computed(() => {
  let filtered = wallets.value || [];

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(
      (w) =>
        w.name.toLowerCase().includes(query) ||
        w.type.toLowerCase().includes(query)
    );
  }

  // Type filter
  if (filterType.value) {
    filtered = filtered.filter((w) => w.type === filterType.value);
  }

  // Sort
  const sorted = [...filtered];
  switch (sortBy.value) {
    case "name":
      sorted.sort((a, b) => a.name.localeCompare(b.name));
      break;
    case "balance-high":
      sorted.sort((a, b) => b.balance - a.balance);
      break;
    case "balance-low":
      sorted.sort((a, b) => a.balance - b.balance);
      break;
    case "type":
      sorted.sort((a, b) => a.type.localeCompare(b.type));
      break;
  }

  return sorted;
});

// Wallet transactions
const walletTransactions = computed(() => {
  if (!selectedWallet.value) return [];
  return transactions.value.filter(
    (t) => t.wallet_id === selectedWallet.value!.ID
  );
});

// Functions
const openWalletModal = () => {
  selectedWallet.value = null;
  isWalletModalOpen.value = true;
};

const editWallet = (wallet: Wallet) => {
  selectedWallet.value = wallet;
  isWalletModalOpen.value = true;
};

const viewWalletDetails = (wallet: Wallet) => {
  selectedWallet.value = wallet;
  isDetailModalOpen.value = true;
};

const deleteWallet = async (wallet: Wallet) => {
  if (!confirm(`Are you sure you want to delete "${wallet.name}"?`)) {
    return;
  }

  try {
    await $fetch(`http://localhost:8081/api/wallets/${wallet.ID}`, {
      method: "DELETE",
    });
    await refreshNuxtData();
    alert("Wallet deleted successfully!");
  } catch (error) {
    console.error("Error deleting wallet:", error);
    alert("Failed to delete wallet");
  }
};

const handleWalletCreated = async () => {
  await refreshNuxtData();
};

const handleWalletUpdated = async () => {
  await refreshNuxtData();
};
</script>
