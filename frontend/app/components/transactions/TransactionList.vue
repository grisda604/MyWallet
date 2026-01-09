<template>
  <!-- Loading State -->
  <div v-if="loading" class="text-gray-500">กำลังโหลด transactions...</div>

  <!-- Error State -->
  <div v-else-if="error" class="text-red-500">
    เกิดข้อผิดพลาดในการโหลด transactions: {{ error.message }}
  </div>

  <!-- Empty State -->
  <div
    v-else-if="!transactions || transactions.length === 0"
    class="text-gray-500"
  >
    ยังไม่มี transaction เริ่มเพิ่มกันเลย!
  </div>

  <!-- Transaction List -->
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
                "ไม่มีคำอธิบาย"
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
            {{ transaction.wallet?.name || "ไม่ทราบกระเป๋าเงิน" }}
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
</template>

<script setup lang="ts">
import type { Transaction } from "~/types";

// Props
defineProps<{
  transactions: Transaction[];
  loading?: boolean;
  error?: Error | null;
}>();
</script>
