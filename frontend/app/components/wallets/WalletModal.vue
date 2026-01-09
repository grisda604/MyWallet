<template>
  <!-- Add Wallet Modal -->
  <Teleport to="body">
    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="closeModal"
    >
      <div
        class="bg-white rounded-2xl shadow-2xl max-w-md w-full max-h-[90vh] overflow-y-auto"
      >
        <!-- Modal Header -->
        <div
          class="sticky top-0 bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between rounded-t-2xl"
        >
          <h3 class="text-xl font-bold text-gray-800">💰 เพิ่ม Wallet</h3>
          <button
            @click="closeModal"
            class="text-gray-400 hover:text-gray-600 transition"
          >
            <span class="text-2xl">×</span>
          </button>
        </div>

        <!-- Modal Body -->
        <form @submit.prevent="submitWallet" class="p-6 space-y-4">
          <!-- Wallet Name -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ชื่อ Wallet <span class="text-red-500">*</span>
            </label>
            <input
              v-model="walletForm.name"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="เช่น KBANK, PromptPay, เงินสด"
            />
          </div>

          <!-- Wallet Type -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ประเภท <span class="text-red-500">*</span>
            </label>
            <div class="grid grid-cols-3 gap-2">
              <button
                type="button"
                @click="walletForm.type = 'bank'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  walletForm.type === 'bank'
                    ? 'bg-blue-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                🏦 ธนาคาร
              </button>
              <button
                type="button"
                @click="walletForm.type = 'cash'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  walletForm.type === 'cash'
                    ? 'bg-green-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                💵 เงินสด
              </button>
              <button
                type="button"
                @click="walletForm.type = 'e-wallet'"
                :class="[
                  'px-4 py-2 rounded-lg font-medium transition',
                  walletForm.type === 'e-wallet'
                    ? 'bg-purple-500 text-white'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200',
                ]"
              >
                📱 E-Wallet
              </button>
            </div>
          </div>

          <!-- Initial Balance -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ยอดเงินเริ่มต้น (฿)
            </label>
            <input
              v-model.number="walletForm.balance"
              type="number"
              step="0.01"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="0.00"
            />
          </div>

          <!-- Color Picker -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              สี
            </label>
            <div class="grid grid-cols-8 gap-2">
              <button
                v-for="color in colorOptions"
                :key="color"
                type="button"
                @click="walletForm.color = color"
                :class="[
                  'w-10 h-10 rounded-lg transition',
                  walletForm.color === color
                    ? 'ring-2 ring-offset-2 ring-blue-500'
                    : 'hover:scale-110',
                ]"
                :style="{ backgroundColor: color }"
              ></button>
            </div>
          </div>

          <!-- Icon Picker -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ไอคอน
            </label>
            <div class="grid grid-cols-8 gap-2">
              <button
                v-for="icon in iconOptions"
                :key="icon"
                type="button"
                @click="walletForm.icon = icon"
                :class="[
                  'w-10 h-10 rounded-lg text-2xl transition flex items-center justify-center',
                  walletForm.icon === icon
                    ? 'bg-blue-100 ring-2 ring-blue-500'
                    : 'bg-gray-100 hover:bg-gray-200',
                ]"
              >
                {{ icon }}
              </button>
            </div>
          </div>

          <!-- Currency -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              สกุลเงิน
            </label>
            <select
              v-model="walletForm.currency"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="THB">THB (บาท)</option>
              <option value="USD">USD (ดอลลาร์)</option>
              <option value="EUR">EUR (ยูโร)</option>
              <option value="JPY">JPY (เยน)</option>
            </select>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-3 pt-4">
            <button
              type="button"
              @click="closeModal"
              class="flex-1 px-4 py-2 bg-gray-100 text-gray-700 font-medium rounded-lg hover:bg-gray-200 transition"
            >
              ยกเลิก
            </button>
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-500 text-white font-medium rounded-lg hover:bg-blue-600 transition"
            >
              เพิ่ม Wallet
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";

// Props
const props = defineProps<{
  modelValue: boolean; // สำหรับ v-model (เปิด/ปิด modal)
}>();

// Emits
const emit = defineEmits<{
  "update:modelValue": [value: boolean];
  "wallet-created": [];
}>();

// Computed สำหรับ v-model
const isModalOpen = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

// Color options
const colorOptions = [
  "#3B82F6", // blue
  "#10B981", // green
  "#F59E0B", // amber
  "#EF4444", // red
  "#8B5CF6", // purple
  "#EC4899", // pink
  "#6366F1", // indigo
  "#14B8A6", // teal
];

// Icon options
const iconOptions = ["💰", "🏦", "💳", "💵", "💸", "🪙", "💴", "💶"];

// Form state
const walletForm = ref({
  name: "",
  type: "bank" as "bank" | "cash" | "e-wallet",
  balance: 0,
  color: "#3B82F6",
  icon: "💰",
  currency: "THB",
});

// ฟังก์ชันปิด modal
const closeModal = () => {
  isModalOpen.value = false;

  // Reset form
  walletForm.value = {
    name: "",
    type: "bank",
    balance: 0,
    color: "#3B82F6",
    icon: "💰",
    currency: "THB",
  };
};

// ฟังก์ชัน submit form
const submitWallet = async () => {
  try {
    const payload = {
      name: walletForm.value.name,
      type: walletForm.value.type,
      balance: Number(walletForm.value.balance),
      color: walletForm.value.color,
      icon: walletForm.value.icon,
      currency: walletForm.value.currency,
    };

    const response = await $fetch("http://localhost:8081/api/wallets", {
      method: "POST",
      body: payload,
    });

    // Emit event เพื่อบอก parent ว่าสร้าง wallet สำเร็จแล้ว
    emit("wallet-created");

    closeModal();
    alert("เพิ่ม Wallet สำเร็จ!");
  } catch (error) {
    console.error("Error adding wallet:", error);
    alert("เกิดข้อผิดพลาดในการเพิ่ม Wallet");
  }
};
</script>
