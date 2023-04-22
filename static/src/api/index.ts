/* eslint-disable camelcase, @typescript-eslint/no-explicit-any, @typescript-eslint/no-non-null-assertion */
export type GamesDataActionV1 = {
  data: number[]
  user: number
}

export type GamesDataGameUserV1 = {
  score: number
  new_score: number
  user: GamesDataUserV1
}

export type GamesDataGameV1 = {
  ts: string
  participants: GamesDataGameUserV1[]
  winner: number
  options: number[]
  ticks: GamesDataTickV1[]
}

export type GamesDataReqV1 = {
  id: string
}

export type GamesDataTickV1 = {
  tick: number
  state: number[]
  actions: GamesDataActionV1[]
}

export type GamesDataUserV1 = {
  id: number
  gh_login: string
  name: string
  avatar_url: string
}

export type GamesListGameV1 = {
  id: string
  debug: boolean
  ts: string
  finished?: string
  participants: GamesListUserV1[]
  winner: number
  timeout: number
}

export type GamesListReqV1 = Record<string, never>

export type GamesListUserV1 = {
  name: string
}

export class ApiError extends Error {
  private readonly _code: string
  private readonly _message: string
  private readonly _data: unknown

  constructor(code: string, message: string, data: unknown) {
    super(message)
    this._code = code
    this._message = message
    this._data = data
  }

  get code(): string {
    return this._code
  }

  get message(): string {
    return this._message
  }

  get data(): unknown {
    return this._data
  }
}

export default class API {
  static url = '/api'
  static customHeaders: () => Promise<Record<string, string>> | undefined

  private static requestToFormData(request: any): FormData {
    const form = new FormData()
    const json_data: any = {}
    for (const name in request) {
      if (request[name] instanceof Blob) {
        form.append(name, request[name])
        continue
      }
      json_data[name] = request[name]
    }
    if (Object.keys(json_data).length !== 0) form.append('json_data', JSON.stringify(json_data))
    return form
  }

  private static async post(method: string, request: unknown, contentType: string): Promise<unknown> {
    return fetch(
      this.url + method,
      {
        method: 'post',
        headers: Object.assign(this.customHeaders ? await this.customHeaders()! : {},
          contentType === 'application/json' ? {'Content-Type': contentType} : {}
        ),
        body: contentType === 'application/json' ? JSON.stringify(request) : this.requestToFormData(request)
      }
    )
      .then(response => {
        return new Promise<Response>((resolve, reject) => {
          switch (response.status) {
            case 200:
              resolve(response)
              break
            case 400:
              response.json().then(err => {
                reject(new ApiError(err.code, err.message, err.data))
              })
              break
            default:
              response.text().then(text => {
                reject(new Error(text || response.statusText))
              })
          }
        })
      })
      .then((response) => response.json())
  }

  // Returns the game data
  public static GamesDataV1(request: GamesDataReqV1): Promise<GamesDataGameV1> {
    return this.post('/games/data/v1', request, 'application/json') as Promise<GamesDataGameV1>
  }

  // Returns games list
  public static GamesListV1(request: GamesListReqV1): Promise<GamesListGameV1[]> {
    return this.post('/games/list/v1', request, 'application/json') as Promise<GamesListGameV1[]>
  }
}
