'use client'

import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';
import type { ChartData } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend);

type TData = any; // Replace 'any' with the actual type for TData
type TLabel = any; // Replace 'any' with the actual type for TLabel


export default function PieChart({names, amounts}: {names: TLabel[], amounts: TData[]}) {


  const data: ChartData<"pie", TData, TLabel> = {
      labels: names,  // category names
      datasets: [
        {
          label: '$ spent this month',
          data: amounts, // spent amount
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(255, 206, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(255, 159, 64, 0.2)',
          ],
          borderColor: [
            'rgba(255, 99, 132, 1)',
            'rgba(54, 162, 235, 1)',
            'rgba(255, 206, 86, 1)',
            'rgba(75, 192, 192, 1)',
            'rgba(153, 102, 255, 1)',
            'rgba(255, 159, 64, 1)',
          ],
          borderWidth: 1,
        },
      ]
    }

    return (
        <div className={'relative w-96'}>
            <Pie data={data} options={{responsive: true}}/>
        </div>
    )
}