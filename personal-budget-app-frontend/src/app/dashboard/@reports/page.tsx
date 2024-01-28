'use client';

import Card from "@/app/ui/card"
import PieChart from '@/app/dashboard/@reports/test-visualization'


export default function ReportsPage() {
  return (
    <Card className='bg-white-200'>
        <PieChart />
    </Card>
  )
}