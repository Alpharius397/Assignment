import { useMutation } from "@tanstack/react-query";
import type { data, HookParams } from "@/hooks";
import Axios, { AccessToken, RefreshToken, URL } from "@/axios";
import { LoginInput, LoginSuccess } from "@/zod/login";
import { AxiosError } from "axios";
import { DefaultError } from "@/zod";

async function login(data: data) {
  const validateData = await LoginInput.parseAsync(data);

  const response = await Axios.post(URL.Login, validateData);

  const responseValid = await LoginSuccess.parseAsync(response.data);

  localStorage.setItem(AccessToken, responseValid.access);
  localStorage.setItem(RefreshToken, responseValid.refresh);

  return responseValid.message;
}

export default function useLogin({
  onSuccess,
  onError,
  onFailure,
}: HookParams) {
  const { mutate, isPending, isError, isSuccess } = useMutation({
    mutationFn: login,
    onSuccess: onSuccess,
    onError: (err) => {
      if (err instanceof AxiosError) {
        const response = DefaultError.safeParse(err.response?.data);

        if (response.success) {
          return onError(response.data.message);
        }
      }
      onFailure();
    },
  });

  return { mutate, isPending, isError, isSuccess };
}
