

export type User = {
    id: string,
    firstName: string,
    lastName: string,
    email: string,
    accounts: Account[],
    budgets: BudgetCategory[],
    transactions: Transaction[]
}

export type BudgetCategory = {
    id: string,  // nanoid
    name: string, // e.g. "Groceries"
    needed: number, // calculated
    assigned: number, // calculated
    spent: number, // calculated
    available: number, // assigned - spent
    goals: Goal[], // determines needed amount
    transactions: Transaction[]
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