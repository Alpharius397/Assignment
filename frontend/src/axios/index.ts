import { RefreshSuccess } from "@/zod/refresh";
import axios, { AxiosError, type AxiosRequestConfig } from "axios";

const API = "http://localhost:8081";

// all possible api paths
export const URL = {
  Login: `${API}/login`,
  Register: `${API}/register`,
  Refresh: `${API}/refresh`,
  GetData: `${API}/get-data`,
  Swagger: `${API}/swagger/swagger.json`,
  Profile: `${API}/profile`,
} as const;

const Axios = axios.create({
  baseURL: API,
  headers: {
    "Request-Origin": API,
    "Content-Type": "application/json",
  },
});

export const AccessToken = "access";
export const RefreshToken = "refresh";

interface Config extends AxiosRequestConfig {
  retry?: boolean;
}

Axios.interceptors.request.use(async (config) => {
  const token = localStorage.getItem(AccessToken);

  if (token !== null && config.headers) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }

  return config;
});

Axios.interceptors.response.use(
  (response) => response,

  async (error: AxiosError) => {
    const config = error.config as Config;

    if (!config) return Promise.reject(error);

    if (config.retry) {
      return Promise.reject(error);
    }

    const skipRefresh = [URL.Login, URL.Refresh, URL.Register].some((path) =>
      config.url?.includes(path)
    );

    if (skipRefresh) {
      return Promise.reject(error);
    }

    if (error.response && error.response.status >= 500) {
      return Promise.reject(error);
    }

    const token = localStorage.getItem(RefreshToken);

    if (token === null) {
      return Promise.reject(new Error("Refresh Token not found"));
    }

    config.retry = true;

    try {
      const response = await axios.post(
        URL.Refresh,
        {},
        {
          headers: {
            "Request-Origin": API,
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const { access } = await RefreshSuccess.parseAsync(response.data);

      localStorage.setItem(AccessToken, access);

      return Axios(config);
    } catch (err) {
      localStorage.removeItem(AccessToken);
      localStorage.removeItem(RefreshToken);
      return Promise.reject(err);
    }
  }
);

export default Axios;
