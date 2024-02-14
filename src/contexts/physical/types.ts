export type PhysicalDimensionStatus = "planned" | "active";

export interface PhysicalDimensionSummary {
  name: string;
  status: PhysicalDimensionStatus;
}

export interface PhysicalDimensionSummaries {
  power: PhysicalDimensionSummary;
  compute: PhysicalDimensionSummary;
  storage: PhysicalDimensionSummary;
  dataLink: PhysicalDimensionSummary;
}

export interface RackSummary {
  name: string;
  id: string;
  dimensions: PhysicalDimensionSummaries;
}

export interface SiteSummary {
  name: string;
  racks: RackSummary[];
}
