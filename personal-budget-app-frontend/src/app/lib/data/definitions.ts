

export type User = {
    id: string,
    password: string,
    firstName: string,
    lastName: string,
    email: string,
}

export type BudgetCategory = {
    id?: string,  // nanoid
    name: string, // e.g. "Groceries"
}

export type Goal = {
    id: string,
    name: string,
    amount: number,
    targetDate: Date,
    category: BudgetCategory
}

export type Transaction = {
    id: string,
    account: Account['id'],
    date: Date,
    payee: string,
    amount: number,
    memo: string,
    category: BudgetCategory['id'] | Transaction[] // for split transactions?  Not sure
}

export type Account = {
    id: string,
    name: string,
    type: AccountType,
    bankName: string,
    balance: number,
}

export type AccountType = 'checking' | 'savings' | 'credit' | 'loan' | 'investment' | 'other'