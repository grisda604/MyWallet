<template>
  <!-- Add Transaction Modal -->
  <Teleport to="body">
    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="CloseTransactionModal"
    >
      <div
        class="bg-white rounded-2xl shadow-2xl max-w-md w-full max-h-[90vh] overflow-y-auto"
      >
        <!-- Modal Header -->
        <div
          class="sticky top-0 bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between rounded-t-2xl"
        >
          <h3 class="text-xl font-bold text-gray-800">➕ Add Transaction</h3>
          <button
            @click="CloseTransactionModal"
            class="text-gray-400 hover:text-gray-600 transition"
          >
            <span class="text-2xl">×</span>
          </button>
        </div>

        <!-- Modal Body -->
        <form @submit.prevent="SubmitTranscation" class="p-6 space-y-4">
          <!-- Transaction Type -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Transaction Type
            </label>
            <div class="grid grid-cols-3 gap-2">
              <button
                type="button"
                @click="TransactionForm.type = 'income'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  TransactionForm.type === 'income'
                    ? 'bg-green-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                💵 Income
              </button>
              <button
                type="button"
                @click="TransactionForm.type = 'expense'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  TransactionForm.type === 'expense'
                    ? 'bg-red-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                💸 Expense
              </button>
              <button
                type="button"
                @click="TransactionForm.type = 'transfer'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  TransactionForm.type === 'transfer'
                    ? 'bg-blue-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                🔄 Transfer
              </button>
            </div>
          </div>

          <!-- Amount -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Amount (฿)
            </label>
            <input
              v-model.number="TransactionForm.amount"
              type="number"
              step="0.01"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="0.00"
            />
          </div>

          <!-- Wallet -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              {{
                TransactionForm.type === "transfer" ? "From Wallet" : "Wallet"
              }}
            </label>
            <select
              v-model="TransactionForm.wallet_id"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option :value="null" disabled>Select wallet</option>
              <option
                v-for="wallet in wallets"
                :key="wallet.ID"
                :value="wallet.ID"
              >
                {{ wallet.name }} (฿{{ wallet.balance.toLocaleString() }})
              </option>
            </select>
          </div>

          <!-- Target Wallet (for Transfer only) -->
          <div v-if="TransactionForm.type === 'transfer'">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              To Wallet
            </label>
            <select
              v-model="TransactionForm.target_wallet_id"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option :value="null" disabled>Select target wallet</option>
              <option
                v-for="wallet in wallets?.filter(
                  (w) => w.ID !== TransactionForm.wallet_id
                )"
                :key="wallet.ID"
                :value="wallet.ID"
              >
                {{ wallet.name }} (฿{{ wallet.balance.toLocaleString() }})
              </option>
            </select>
          </div>

          <!-- Category (not for Transfer) -->
          <div v-if="TransactionForm.type !== 'transfer'">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Category
            </label>
            <select
              v-model="TransactionForm.category_id"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option :value="null">No category</option>
              <option
                v-for="category in categories?.filter(
                  (c) => c.type === TransactionForm.type
                )"
                :key="category.ID"
                :value="category.ID"
              >
                {{ category.icon }} {{ category.name }}
              </option>
            </select>
          </div>

          <!-- Date -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Date
            </label>
            <input
              v-model="TransactionForm.date"
              type="date"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>

          <!-- Note -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Note (Optional)
            </label>
            <textarea
              v-model="TransactionForm.note"
              rows="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
              placeholder="Add a note..."
            ></textarea>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-3 pt-4">
            <button
              type="button"
              @click="CloseTransactionModal"
              class="flex-1 px-4 py-2 bg-gray-100 text-gray-700 font-medium rounded-lg hover:bg-gray-200 transition"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-500 text-white font-medium rounded-lg hover:bg-blue-600 transition"
            >
              Add Transaction
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { Wallet, Category } from "~/types";

// Props
const props = defineProps<{
  modelValue: boolean; // สำหรับ v-model (เปิด/ปิด modal)
  wallets: Wallet[]; // รายการ wallets
  categories: Category[]; // รายการ categories
}>();

// Emits
const emit = defineEmits<{
  "update:modelValue": [value: boolean]; // สำหรับ v-model two-way binding
  "transaction-created": []; // เมื่อสร้าง transaction สำเร็จ
}>();

// Computed สำหรับ v-model
const isModalOpen = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

// Form state
const TransactionForm = ref({
  type: "income" as "income" | "expense" | "transfer",
  amount: 0,
  date: new Date().toISOString().split("T")[0], // แก้ไขจาก Date().toLocaleString()
  note: "",
  category_id: null as number | null,
  wallet_id: null as number | null,
  target_wallet_id: null as number | null,
});

// ฟังก์ชันปิด modal
const CloseTransactionModal = () => {
  isModalOpen.value = false;

  // Reset form
  TransactionForm.value = {
    type: "income",
    amount: 0,
    date: new Date().toISOString().split("T")[0],
    note: "",
    category_id: null,
    wallet_id: null,
    target_wallet_id: null,
  };
};

// ฟังก์ชัน submit form
const SubmitTranscation = async () => {
  try {
    // เตรียมข้อมูลก่อนส่ง - แปลงให้ถูก format
    const payload: any = {
      type: TransactionForm.value.type,
      amount: Number(TransactionForm.value.amount),
      date: TransactionForm.value.date
        ? new Date(TransactionForm.value.date).toISOString()
        : new Date().toISOString(), // แปลงเป็น ISO 8601
      note: TransactionForm.value.note || "",
      wallet_id: Number(TransactionForm.value.wallet_id),
    };

    // ส่ง category_id เฉพาะเมื่อมีค่า (ไม่ใช่ null)
    if (TransactionForm.value.category_id) {
      payload.category_id = Number(TransactionForm.value.category_id);
    }

    // ส่ง target_wallet_id เฉพาะเมื่อเป็น transfer และมีค่า
    if (
      TransactionForm.value.type === "transfer" &&
      TransactionForm.value.target_wallet_id
    ) {
      payload.target_wallet_id = Number(TransactionForm.value.target_wallet_id);
    }

    const response = await $fetch("http://localhost:8081/api/transactions", {
      method: "POST",
      body: payload,
    });

    // Emit event เพื่อบอก parent ว่าสร้าง transaction สำเร็จแล้ว
    emit("transaction-created");

    CloseTransactionModal();
    alert("เพิ่ม Transaction สำเร็จ!");
  } catch (error) {
    console.error("Error adding transaction:", error);
    alert("เกิดข้อผิดพลาดในการเพิ่ม Transaction");
  }
};
</script>
