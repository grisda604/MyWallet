<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
    <h3 class="text-lg font-bold text-gray-800 mb-4">💼 Wallet Distribution</h3>
    <div
      v-if="!wallets || wallets.length === 0"
      class="text-gray-500 text-center py-8"
    >
      No wallets available
    </div>
    <apexchart
      v-else
      type="donut"
      :options="chartOptions"
      :series="series"
      height="300"
    />
  </div>
</template>

<script setup lang="ts">
import type { Wallet } from "~/types";

interface Props {
  wallets: Wallet[];
}

const props = defineProps<Props>();

const series = computed(() => {
  return props.wallets.map((w) => Math.max(0, w.balance));
});

const chartOptions = computed(() => ({
  chart: {
    type: "donut",
    fontFamily: "Inter, sans-serif",
    animations: {
      enabled: true,
      easing: "easeinout",
      speed: 800,
      animateGradually: {
        enabled: true,
        delay: 150,
      },
      dynamicAnimation: {
        enabled: true,
        speed: 350,
      },
    },
  },
  labels: props.wallets.map((w) => w.name),
  colors: ["#3B82F6", "#10B981", "#F59E0B", "#EF4444", "#8B5CF6", "#EC4899"],
  legend: {
    position: "bottom",
    fontSize: "14px",
    fontWeight: 500,
    markers: {
      width: 12,
      height: 12,
      radius: 3,
    },
  },
  plotOptions: {
    pie: {
      donut: {
        size: "70%",
        labels: {
          show: true,
          name: {
            show: true,
            fontSize: "16px",
            fontWeight: 600,
            color: "#374151",
          },
          value: {
            show: true,
            fontSize: "24px",
            fontWeight: 700,
            color: "#111827",
            formatter: (val: number) => `฿${Number(val).toLocaleString()}`,
          },
          total: {
            show: true,
            label: "Total Balance",
            fontSize: "14px",
            fontWeight: 600,
            color: "#6B7280",
            formatter: (w: any) => {
              const total = w.globals.seriesTotals.reduce(
                (a: number, b: number) => a + b,
                0
              );
              return `฿${total.toLocaleString()}`;
            },
          },
        },
      },
    },
  },
  dataLabels: {
    enabled: false,
  },
  tooltip: {
    enabled: true,
    y: {
      formatter: (val: number) => `฿${val.toLocaleString()}`,
    },
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
