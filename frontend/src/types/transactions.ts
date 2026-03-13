export type TransactionType = "income" | "export";

export type Transaction = {
    id: number
    amount: number
    date: string
    type: TransactionType
    description: string
    category: string
}

export type TransactionInput = {
    amount: number
    type: TransactionType
    description: string
    category: string
    date?: string
}