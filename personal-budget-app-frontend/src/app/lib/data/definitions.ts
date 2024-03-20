import { Dispatch } from "react";
import { SetStateAction } from "react";

export type SessionContextType = {
    user: User | null;
    setUser: Dispatch<SetStateAction<User | null>>;
    signOut: () => void;
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
    email: string // user email (foreign key)
}

export type Goal = {
    id: string,
    name: string,
    amount: number,
    targetDate: Date,
    category: Category
}

export type Transaction = {
    id: string,
    accountID: Account['id'],
    date: Date,
    payee: string,
    amount: number,
    memo: string,
    categoryID: Category['id'] | Transaction[] | null // for split transactions?  Not sure
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
    id: number,
    email: string, // user email (foreign key)
    categoryID: number,
    month: string,
    year: number,
    amount: number
}

export type AccountType = 'checking' | 'savings' | 'credit' | 'loan' | 'investment' | 'other'