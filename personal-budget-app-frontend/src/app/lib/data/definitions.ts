import { Dispatch } from "react";
import { SetStateAction } from "react";

export type SessionContextType = {
    user: User | null;
    setUser: Dispatch<SetStateAction<User | null>>;
    signOut: () => void;
}

export type IncomeAndExpenses = [
    {
        month: number,
        income: number,
        expenses: number,
    }
]

export type TimeInfo = {
    today: Date
    viewedMonth: number
    viewedYear: number
    viewedDate: Date
    nextMonth: number
    prevMonth: number
    nextYear: number
    prevYear: number
    monthString: string
}

export type BudgetPageData = {
    currentFunds: number
    projIncome: number
    totalAvailable: number
    totalUnassigned: number
    categoryRows: CategoryData[]
}

export type CategoryData = {
    categoryID: number
    categoryName: string
    assigned: number
    spent: number
    available: number
    goalsSum: number
}

export type MonthlyBudget = {
    id?: number
    email: string
    categoryID: number
    month: number
    year: number
    assigned: number
    spent?: number
    balance?: number
}

export type User = {
    password: string,
    firstName: string,
    lastName: string,
    email: string,
    expiration?: number | null
} | null

export type Category = {
    id: number,
    name: string, // e.g. "Groceries"
    email: string, // user email (foreign key)
}

export type Goal = {
    id: string,
    name: string,
    amount: number,
    targetDate: Date,
    categoryID: number,
    email: string,
    periodicity: Periodicity,
}

export type Periodicity = 'monthly' | 'onetime'

export type Transaction = {
    id: string,
    accountID: Account['id'],
    date: Date,
    payee: string,
    amount: number,
    memo: string,
    categoryID: Category['id'] | Transaction[] | null // for split transactions?  Not sure
    categoryName?: string
    email: string // user email (foreign key)
}

export type Account = {
    email: string, // user email (foreign key)
    id: number, // primary key
    name: string,
    type: AccountType,
    bankName: string,
    startingBalance: number,
    balance: number // not stored in DB, but calculated from transactions
}

export type Assignment = {
    id?: number,
    email: string, // user email (foreign key)
    categoryID: number,
    month: number,
    year: number,
    amount: number
}

export type AccountType = 'checking' | 'savings' | 'credit' | 'loan' | 'investment' | 'other'