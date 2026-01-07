<template>
  <div
    class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden"
  >
    <div class="p-4 border-b border-gray-100 flex justify-between items-center">
      <h2 class="text-lg font-bold text-gray-800">Recent Transactions</h2>
      <NuxtLink to="/transactions" class="text-sm text-blue-500 hover:underline"
        >View All</NuxtLink
      >
    </div>

    <div v-if="transactions.length === 0" class="p-8 text-center text-gray-400">
      No transactions found.
    </div>

    <ul v-else class="divide-y divide-gray-50">
      <li
        v-for="tx in transactions"
        :key="tx.ID"
        class="p-4 hover:bg-gray-50 transition flex items-center justify-between"
      >
        <div class="flex items-center gap-4">
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center text-xl"
            :class="{
              'bg-green-100 text-green-600': tx.type === 'income',
              'bg-red-100 text-red-600': tx.type === 'expense',
              'bg-blue-100 text-blue-600': tx.type === 'transfer',
            }"
          >
            <span v-if="tx.type === 'income'">↓</span>
            <span v-else-if="tx.type === 'expense'">↑</span>
            <span v-else>⇄</span>
          </div>

          <div>
            <p class="font-medium text-gray-800">
              {{ tx.note || tx.category?.name || "No Description" }}
            </p>
            <p class="text-xs text-gray-500">
              {{ formatDate(tx.date) }} • {{ tx.wallet?.name }}
              <span v-if="tx.type === 'transfer'">
                ➝ {{ tx.target_wallet?.name }}
              </span>
            </p>
          </div>
        </div>

        <div
          class="font-bold text-right"
          :class="{
            'text-green-600': tx.type === 'income',
            'text-red-500': tx.type === 'expense',
            'text-blue-500': tx.type === 'transfer',
          }"
        >
          {{ tx.type === "expense" ? "-" : "+" }}
          ฿{{ tx.amount.toLocaleString() }}
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import type { Transaction } from "~/types";

// รับ Props เป็น List ของ Transaction
defineProps<{
  transactions: Transaction[];
}>();

// Helper แปลงวันที่
const formatDate = (dateString: string | undefined | null) => {
  if (!dateString) return "-";

  const date = new Date(dateString);
  if (isNaN(date.getTime())) return "-";

  return new Intl.DateTimeFormat("th-TH", {
    day: "numeric",
    month: "short",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
};
</script>
