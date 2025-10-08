import type { Axios } from "axios";
import {
    authResponseSchema,
    type LoginUserInput,
    type RegisterUserInput,
} from "@/lib/validators/auth.schema";

export class Auth {
  axios: Axios;
  constructor(axios: Axios) {
    this.axios = axios;
  }

  async register(data: RegisterUserInput) {
    const response = await this.axios.post("/auth/register", data);
    return authResponseSchema.parse(response.data);
  }

  async login(data: LoginUserInput) {
    const response = await this.axios.post("/auth/login", data);
    return authResponseSchema.parse(response.data);
  }

  async logout() {}
}
