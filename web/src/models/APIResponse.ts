export type APIResponse<T> = {
  data: T;
  success: boolean;
}

export type Ping = {
  message: string;
}
