
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { ClerkProvider } from '@clerk/nextjs'
import { twMerge } from 'tailwind-merge'
import './globals.css'
import PageHeader from './page-header'
import NavBar from './ui/navbar'


const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Personal Budget App',
  description: 'Created by Will Lapinel for ITIS 5166',
}

export default function RootLayout({
  children
}: {
  children: React.ReactNode
}) {
  return (
    <ClerkProvider>
      <html lang="en" className='h-full'>
        <body className={twMerge(inter.className, 'h-full')}>
          <div className="min-h-full">
            <NavBar />
            <div className="py-10">
              <PageHeader />
              <main>
                <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">{children}</div>
              </main>
            </div>
          </div>
        </body>
      </html>
    </ClerkProvider>
  )
}
