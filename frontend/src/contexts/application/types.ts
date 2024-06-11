import { GeneralStatus } from "../shared/types"

export type AppDimensionName = "availability"

export interface AppDimensionSummary {
  name: AppDimensionName
  status: GeneralStatus
}

export interface AppSummary {
  id: string
  name: string
}
