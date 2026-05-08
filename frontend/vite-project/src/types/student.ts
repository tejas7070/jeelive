export interface Student {
    _id?: string
    id: string
    name: string
    percentile: number
    preferences: string[]
    allotted?: string
    createdOn?: number
    modifiedOn?: number
}