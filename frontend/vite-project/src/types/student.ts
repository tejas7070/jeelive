export interface Student {
  _id?: string
  id: string
  name: string
  percentile: number
  preferences: string[]
  allotedBranch?: string
  status?: string
  createdOn?: number
  modifiedOn?: number
}
