<template>
  <Teleport to="body">
    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="closeModal"
    >
      <div
        class="bg-white rounded-2xl shadow-2xl max-w-4xl w-full max-h-[90vh] overflow-y-auto"
      >
        <!-- Modal Header -->
        <div
          class="sticky top-0 bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between rounded-t-2xl"
        >
          <h3 class="text-xl font-bold text-gray-800">
            {{ wallet?.icon || "💰" }} {{ wallet?.name || "Wallet Details" }}
          </h3>
          <button
            @click="closeModal"
            class="text-gray-400 hover:text-gray-600 transition"
          >
            <span class="text-2xl">×</span>
          </button>
        </div>

        <!-- Modal Body -->
        <div v-if="wallet" class="p-6">
          <!-- Wallet Info -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
            <div class="bg-gray-50 rounded-lg p-4">
              <p class="text-sm text-gray-600 mb-1">Type</p>
              <p class="text-lg font-semibold capitalize">{{ wallet.type }}</p>
            </div>
            <div class="bg-gray-50 rounded-lg p-4">
              <p class="text-sm text-gray-600 mb-1">Currency</p>
              <p class="text-lg font-semibold">
                {{ wallet.currency || "THB" }}
              </p>
            </div>
            <div class="bg-gray-50 rounded-lg p-4">
              <p class="text-sm text-gray-600 mb-1">Current Balance</p>
              <p
                class="text-lg font-semibold"
                :class="{
                  'text-red-500': wallet.balance < 0,
                  'text-green-600': wallet.balance >= 0,
                }"
              >
                ฿{{ wallet.balance.toLocaleString() }}
              </p>
            </div>
          </div>

          <!-- Transaction Statistics -->
          <div v-if="transactions.length > 0" class="mb-6">
            <h4 class="text-lg font-bold text-gray-800 mb-4">📊 Statistics</h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div class="bg-green-50 rounded-lg p-4 border border-green-200">
                <p class="text-sm text-green-700 mb-1">Total Income</p>
                <p class="text-2xl font-bold text-green-600">
                  ฿{{ totalIncome.toLocaleString() }}
                </p>
                <p class="text-xs text-green-600 mt-1">
                  {{ incomeCount }} transactions
                </p>
              </div>
              <div class="bg-red-50 rounded-lg p-4 border border-red-200">
                <p class="text-sm text-red-700 mb-1">Total Expense</p>
                <p class="text-2xl font-bold text-red-600">
                  ฿{{ totalExpense.toLocaleString() }}
                </p>
                <p class="text-xs text-red-600 mt-1">
                  {{ expenseCount }} transactions
                </p>
              </div>
              <div class="bg-blue-50 rounded-lg p-4 border border-blue-200">
                <p class="text-sm text-blue-700 mb-1">Net Flow</p>
                <p
                  class="text-2xl font-bold"
                  :class="{
                    'text-green-600': netFlow >= 0,
                    'text-red-600': netFlow < 0,
                  }"
                >
                  {{ netFlow >= 0 ? "+" : "" }}฿{{ netFlow.toLocaleString() }}
                </p>
                <p class="text-xs text-blue-600 mt-1">
                  {{ transactions.length }} total
                </p>
              </div>
            </div>
          </div>

          <!-- Recent Transactions -->
          <div>
            <h4 class="text-lg font-bold text-gray-800 mb-4">
              📝 Recent Transactions
            </h4>
            <div
              v-if="transactions.length === 0"
              class="text-gray-500 text-center py-8"
            >
              No transactions yet
            </div>
            <div v-else class="bg-gray-50 rounded-lg overflow-hidden">
              <div
                v-for="transaction in recentTransactions"
                :key="transaction.ID"
                class="flex items-center justify-between p-4 border-b border-gray-200 last:border-b-0 hover:bg-white transition"
              >
                <!-- Left: Icon + Info -->
                <div class="flex items-center gap-4 flex-1">
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
                  <div class="flex-1 min-w-0">
                    <div class="flex items-center gap-2 mb-1">
                      <span class="font-medium text-gray-800">
                        {{
                          transaction.note ||
                          transaction.category?.name ||
                          "No description"
                        }}
                      </span>
                      <span
                        v-if="transaction.category"
                        class="text-xs px-2 py-0.5 rounded-full border flex-shrink-0"
                        :style="{
                          backgroundColor: transaction.category.color + '20',
                          borderColor: transaction.category.color,
                          color: transaction.category.color,
                        }"
                      >
                        {{ transaction.category.icon }}
                        {{ transaction.category.name }}
                      </span>
                    </div>
                    <div class="text-sm text-gray-500">
                      {{
                        new Date(
                          transaction.date || transaction.CreatedAt
                        ).toLocaleDateString("th-TH")
                      }}
                    </div>
                  </div>
                </div>
                <!-- Right: Amount -->
                <div
                  class="font-semibold text-lg"
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
              </div>
            </div>
            <div v-if="transactions.length > 10" class="text-center mt-4">
              <p class="text-sm text-gray-500">
                Showing 10 of {{ transactions.length }} transactions
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import type { Wallet, Transaction } from "~/types";

interface Props {
  modelValue: boolean;
  wallet: Wallet | null;
  transactions: Transaction[];
}

const props = defineProps<Props>();

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
}>();

const isModalOpen = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

const closeModal = () => {
  isModalOpen.value = false;
};

// Transaction statistics
const totalIncome = computed(() => {
  return props.transactions
    .filter((t) => t.type === "income")
    .reduce((sum, t) => sum + t.amount, 0);
});

const totalExpense = computed(() => {
  return props.transactions
    .filter((t) => t.type === "expense")
    .reduce((sum, t) => sum + Math.abs(t.amount), 0);
});

const incomeCount = computed(() => {
  return props.transactions.filter((t) => t.type === "income").length;
});

const expenseCount = computed(() => {
  return props.transactions.filter((t) => t.type === "expense").length;
});

const netFlow = computed(() => {
  return totalIncome.value - totalExpense.value;
});

// Recent transactions (last 10)
const recentTransactions = computed(() => {
  return [...props.transactions]
    .sort((a, b) => {
      const dateA = new Date(a.date || a.CreatedAt).getTime();
      const dateB = new Date(b.date || b.CreatedAt).getTime();
      return dateB - dateA;
    })
    .slice(0, 10);
});
</script>
