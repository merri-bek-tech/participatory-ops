import { GeneralStatus } from "../shared/types"

export type PhysicalDimensionName = "power" | "compute" | "storage" | "dataLink"

export interface PhysicalDimensionSummary {
  name: PhysicalDimensionName
  status: GeneralStatus
}

export interface PhysicalDimensionSummaries {
  power: PhysicalDimensionSummary
  compute: PhysicalDimensionSummary
  storage: PhysicalDimensionSummary
  dataLink: PhysicalDimensionSummary
}

export interface RackSummary {
  name: string
  id: string
  detailUrl?: string
  dimensions: PhysicalDimensionSummaries
}

export interface SiteSummary {
  id: string
  name: string
  racks: RackSummary[]
}

export interface RackDetails {
  id: string
  name: string
  dimensions: PhysicalDimensionSummaries
}

export interface PhysicalLayerSummary {
  sites: SiteSummary[]
  racks: Record<string, RackDetails>
}

export interface ComponentDetails {
  hostName: string
  productName: string
  sysVendor: string
}

export interface ComponentStatus {
  status: "online" | "offline"
  uuid: string
  details?: ComponentDetails | null
}
