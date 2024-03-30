export interface ComponentDetails {
  hostName: string
}

export interface ComponentStatus {
  status: "online" | "offline"
  uuid: string
  details?: ComponentDetails | null
}
