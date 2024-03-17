import { Dispatch } from "react";
import { SetStateAction } from "react";



export type SessionContextType = {
    user: User | null;
    setUser: Dispatch<SetStateAction<User | null>>;
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
    account: Account['id'],
    date: Date,
    payee: string,
    amount: number,
    memo: string,
    category: Category['id'] | Transaction[] // for split transactions?  Not sure
}

export type Account = {
    email: string, // user email (foreign key)
    id: number, // primary key
    name: string,
    type: AccountType,
    bankName: string,
    balance: number,
}

export type AccountType = 'checking' | 'savings' | 'credit' | 'loan' | 'investment' | 'other'