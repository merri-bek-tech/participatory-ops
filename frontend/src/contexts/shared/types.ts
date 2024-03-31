export interface ComponentDetails {
  hostName: string
}

export interface ComponentStatus {
  status: "online" | "offline"
  uuid: string
  details?: ComponentDetails | null
}

export interface Scheme {
  id: string
  name: string
  hostnames: string[]
}
