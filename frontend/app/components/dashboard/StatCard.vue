<template>
  <div
    class="bg-gradient-to-br rounded-xl shadow-sm border border-gray-100 p-6 hover:shadow-md transition-all duration-300 transform hover:scale-105"
    :class="gradientClass"
  >
    <div class="flex items-center justify-between mb-2">
      <span class="text-sm font-medium text-gray-600">{{ title }}</span>
      <span class="text-2xl">{{ icon }}</span>
    </div>
    <div class="flex items-end justify-between">
      <div>
        <div class="text-3xl font-bold text-gray-800">
          {{ formattedValue }}
        </div>
        <div v-if="trend" class="flex items-center gap-1 mt-2 text-sm">
          <span :class="trendColor">
            {{ trend > 0 ? "↑" : "↓" }} {{ Math.abs(trend) }}%
          </span>
          <span class="text-gray-500">vs last month</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string;
  value: number;
  icon: string;
  trend?: number;
  color?: "blue" | "green" | "red" | "purple";
  prefix?: string;
}

const props = withDefaults(defineProps<Props>(), {
  color: "blue",
  prefix: "฿",
});

const formattedValue = computed(() => {
  if (props.prefix === "฿") {
    return `฿${props.value.toLocaleString()}`;
  }
  return props.value.toLocaleString();
});

const gradientClass = computed(() => {
  const gradients = {
    blue: "from-blue-50 to-blue-100",
    green: "from-green-50 to-green-100",
    red: "from-red-50 to-red-100",
    purple: "from-purple-50 to-purple-100",
  };
  return gradients[props.color];
});

const trendColor = computed(() => {
  if (!props.trend) return "";
  return props.trend > 0
    ? "text-green-600 font-semibold"
    : "text-red-600 font-semibold";
});
</script>
