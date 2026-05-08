export type RawBranch = Record<string, any>

export type NormalizedBranch = {
  id: string
  name: string
  seats: number
  allocatedSeats: number
  remainingSeats: number
  cutoff: number
  status: "Available" | "Full"
}

const pickString = (branch: RawBranch, keys: string[], fallback = "") => {
  for (const key of keys) {
    const value = branch[key]
    if (typeof value === "string" && value.trim()) {
      return value
    }
  }
  return fallback
}

const pickNumber = (branch: RawBranch, keys: string[], fallback = 0) => {
  for (const key of keys) {
    const value = branch[key]
    if (typeof value === "number" && Number.isFinite(value)) {
      return value
    }

    if (typeof value === "string" && value.trim() && !Number.isNaN(Number(value))) {
      return Number(value)
    }
  }
  return fallback
}

export const normalizeBranch = (branch: RawBranch, index = 0): NormalizedBranch => {
  const name = pickString(branch, ["name", "branchName", "title", "branch"], `Branch ${index + 1}`)
  const seats = pickNumber(branch, ["seats", "totalSeats", "capacity", "strength"], 0)
  const allocatedSeats = pickNumber(
    branch,
    ["allotedSeats", "allottedSeats", "allocatedSeats", "filledSeats", "filled"],
    0
  )
  const cutoff = pickNumber(branch, ["cutoff", "cutoffPercentile", "percentileCutoff"], 0)
  const id = pickString(branch, ["id", "_id"], String(index))
  const remainingSeats = Math.max(seats - allocatedSeats, 0)

  return {
    id,
    name,
    seats,
    allocatedSeats,
    remainingSeats,
    cutoff,
    status: remainingSeats > 0 ? "Available" : "Full"
  }
}

export const getBranchDisplayValue = (branch: RawBranch, keys: string[], fallback = "—") => {
  for (const key of keys) {
    const value = branch[key]
    if (typeof value === "string" && value.trim()) return value
    if (typeof value === "number" && Number.isFinite(value)) return String(value)
  }
  return fallback
}
