import { UnknownComponent } from "./types"

const defaultHeaders = {
  "Content-Type": "application/json",
}

const default_host = "http://localhost:1323/api"
function getApiHost(path: string | undefined = ""): string {
  const base_url: string = import.meta.env.VITE_API_HOST || default_host
  return base_url + path
}

class Api {
  base_url: string

  constructor(base_url?: string) {
    this.base_url = base_url || getApiHost()
  }

  inbox(): Promise<UnknownComponent[]> {
    return fetch(`${this.base_url}/inbox`, {
      method: "GET",
      headers: defaultHeaders,
    })
      .then((response: Response) => response.json())
      .then((json: any) => json as UnknownComponent[])
  }
}

export default Api
