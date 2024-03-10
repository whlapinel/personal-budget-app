
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { twMerge } from 'tailwind-merge'
import './globals.css'
import NavBar from './ui/navbar'
import { SessionProvider } from './session-context'
import UserInactive from './user-inactive'


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

    <html lang="en" className='h-full'>
      <body className={twMerge(inter.className, 'h-full')}>
        <SessionProvider>
          <div className="min-h-full">
            <NavBar />
            <main className={'py-2 px-2'}>
              {children}
            </main>
          </div>
        </SessionProvider>
      </body>
    </html >
  )
}
