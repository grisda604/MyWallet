<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
    <h3 class="text-lg font-bold text-gray-800 mb-4">
      🏷️ Spending by Category
    </h3>
    <div
      v-if="!transactions || transactions.length === 0"
      class="text-gray-500 text-center py-8"
    >
      No expense transactions available
    </div>
    <apexchart
      v-else
      type="bar"
      :options="chartOptions"
      :series="series"
      height="300"
    />
  </div>
</template>

<script setup lang="ts">
import type { Transaction } from "~/types";

interface Props {
  transactions: Transaction[];
}

const props = defineProps<Props>();

// Group expenses by category
const categoryData = computed(() => {
  const data: {
    [key: string]: { amount: number; color: string; icon: string };
  } = {};

  props.transactions
    .filter((t) => t.type === "expense" && t.category)
    .forEach((t) => {
      const categoryName = t.category!.name;
      if (!data[categoryName]) {
        data[categoryName] = {
          amount: 0,
          color: t.category!.color || "#6B7280",
          icon: t.category!.icon || "📌",
        };
      }
      data[categoryName].amount += Math.abs(t.amount);
    });

  // Sort by amount (highest first) and take top 8
  const sorted = Object.entries(data)
    .sort(([, a], [, b]) => b.amount - a.amount)
    .slice(0, 8);

  return sorted;
});

const series = computed(() => [
  {
    name: "Spending",
    data: categoryData.value.map(([_, data]) => data.amount),
  },
]);

const categories = computed(() => {
  return categoryData.value.map(([name, data]) => `${data.icon} ${name}`);
});

const colors = computed(() => {
  return categoryData.value.map(([_, data]) => data.color);
});

const chartOptions = computed(() => ({
  chart: {
    type: "bar",
    fontFamily: "Inter, sans-serif",
    toolbar: {
      show: false,
    },
    animations: {
      enabled: true,
      easing: "easeinout",
      speed: 800,
    },
  },
  colors: colors.value,
  plotOptions: {
    bar: {
      horizontal: true,
      borderRadius: 8,
      distributed: true,
      dataLabels: {
        position: "top",
      },
    },
  },
  dataLabels: {
    enabled: true,
    formatter: (val: number) => `฿${val.toLocaleString()}`,
    offsetX: 30,
    style: {
      fontSize: "12px",
      fontWeight: 600,
      colors: ["#374151"],
    },
  },
  xaxis: {
    categories: categories.value,
    labels: {
      style: {
        fontSize: "12px",
        fontWeight: 500,
        colors: "#6B7280",
      },
      formatter: (val: number) => `฿${val.toLocaleString()}`,
    },
  },
  yaxis: {
    labels: {
      style: {
        fontSize: "13px",
        fontWeight: 500,
        colors: colors.value,
      },
    },
  },
  legend: {
    show: false,
  },
  tooltip: {
    enabled: true,
    y: {
      formatter: (val: number) => `฿${val.toLocaleString()}`,
    },
  },
  grid: {
    borderColor: "#F3F4F6",
    strokeDashArray: 4,
    xaxis: {
      lines: {
        show: true,
      },
    },
    yaxis: {
      lines: {
        show: false,
      },
    },
  },
  responsive: [
    {
      breakpoint: 480,
      options: {
        chart: {
          height: 250,
        },
        plotOptions: {
          bar: {
            horizontal: false,
          },
        },
      },
    },
  ],
}));
</script>
