'use client'

import React from 'react';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ChartData
} from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

export async function BarChart({months, income, expenses}:{months: string[], income: number[], expenses: number[]}) {
    const options = {
        responsive: true,
        plugins: {
          legend: {
            position: 'top' as const,
          },
          title: {
            display: true,
            text: 'Chart.js Bar Chart',
          },
        },
      };
      
      const labels = months;
      
      const data = {
        labels,
        datasets: [
          {
            label: 'Income',
            data: income,
            backgroundColor: 'rgba(255, 99, 132, 0.5)',
          },
          {
            label: 'Expenses',
            data: expenses,
            backgroundColor: 'rgba(53, 162, 235, 0.5)',
          },
        ],
      };
      

return (
    <div className={'relative w-96'}>
        <Bar data={data} options={options}/>
    </div>
)

}