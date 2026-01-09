<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
    <h3 class="text-lg font-bold text-gray-800 mb-4">📈 Income vs Expense</h3>
    <div
      v-if="!transactions || transactions.length === 0"
      class="text-gray-500 text-center py-8"
    >
      No transactions available
    </div>
    <apexchart
      v-else
      type="area"
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

// Group transactions by month
const monthlyData = computed(() => {
  const data: { [key: string]: { income: number; expense: number } } = {};

  props.transactions.forEach((t) => {
    const date = new Date(t.date || t.CreatedAt);
    const monthKey = `${date.getFullYear()}-${String(
      date.getMonth() + 1
    ).padStart(2, "0")}`;

    if (!data[monthKey]) {
      data[monthKey] = { income: 0, expense: 0 };
    }

    if (t.type === "income") {
      data[monthKey].income += t.amount;
    } else if (t.type === "expense") {
      data[monthKey].expense += Math.abs(t.amount);
    }
  });

  // Sort by month and get last 6 months
  const sorted = Object.entries(data)
    .sort(([a], [b]) => a.localeCompare(b))
    .slice(-6);

  return sorted;
});

const series = computed(() => [
  {
    name: "Income",
    data: monthlyData.value.map(([_, data]) => data.income),
  },
  {
    name: "Expense",
    data: monthlyData.value.map(([_, data]) => data.expense),
  },
]);

const categories = computed(() => {
  return monthlyData.value.map(([month]) => {
    const [year, m] = month.split("-");
    const date = new Date(parseInt(year), parseInt(m) - 1);
    return date.toLocaleDateString("th-TH", {
      month: "short",
      year: "2-digit",
    });
  });
});

const chartOptions = computed(() => ({
  chart: {
    type: "area",
    fontFamily: "Inter, sans-serif",
    toolbar: {
      show: false,
    },
    zoom: {
      enabled: false,
    },
    animations: {
      enabled: true,
      easing: "easeinout",
      speed: 800,
    },
  },
  colors: ["#10B981", "#EF4444"],
  dataLabels: {
    enabled: false,
  },
  stroke: {
    curve: "smooth",
    width: 3,
  },
  fill: {
    type: "gradient",
    gradient: {
      shadeIntensity: 1,
      opacityFrom: 0.6,
      opacityTo: 0.1,
      stops: [0, 90, 100],
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
    },
  },
  yaxis: {
    labels: {
      style: {
        fontSize: "12px",
        fontWeight: 500,
        colors: "#6B7280",
      },
      formatter: (val: number) => `฿${val.toLocaleString()}`,
    },
  },
  legend: {
    position: "top",
    horizontalAlign: "right",
    fontSize: "14px",
    fontWeight: 500,
    markers: {
      width: 12,
      height: 12,
      radius: 3,
    },
  },
  tooltip: {
    shared: true,
    intersect: false,
    y: {
      formatter: (val: number) => `฿${val.toLocaleString()}`,
    },
  },
  grid: {
    borderColor: "#F3F4F6",
    strokeDashArray: 4,
  },
  responsive: [
    {
      breakpoint: 480,
      options: {
        chart: {
          height: 250,
        },
        legend: {
          position: "bottom",
        },
      },
    },
  ],
}));
</script>
