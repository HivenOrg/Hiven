import axios, { type Axios } from "axios";
import { cookies } from "next/headers";
import { Auth } from "./auth";

class Api {
  private readonly _axios: Axios;
  auth: Auth;

  constructor() {
    this._axios = this.createAxios();
    this.auth = new Auth(this._axios);
  }

  private createAxios() {
    const ax = axios.create({
      baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:5000",
    });

    ax.interceptors.request.use(async (config) => {
      if (typeof window === "undefined") {
        const cookieStore = await cookies();
        const token = cookieStore.get("token");
        if (token) {
          config.headers.Authorization = `${token}`;
        }
        return config;
      }
      const token = localStorage.getItem("token");
      if (token) {
        config.headers.Authorization = `${token}`;
      }
      return config;
    });

    return ax;
  }

  getAxios() {
    return this._axios;
  }
}

const api = new Api();
export default api;
