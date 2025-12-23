import { useMutation } from "@tanstack/react-query";
import type { data, HookParams } from "@/hooks";
import Axios, { URL } from "@/axios";
import { RegisterInput, RegisterSuccess } from "@/zod/register";
import { AxiosError } from "axios";
import { DefaultError } from "@/zod";

async function register(data: data) {
  const validateData = await RegisterInput.parseAsync(data);

  const response = await Axios.post(URL.Register, validateData);

  await RegisterSuccess.parseAsync(response.data);

  return true;
}

export default function useRegister({
  onSuccess,
  onError,
  onFailure,
}: HookParams) {
  const { mutate, isPending, isError, isSuccess } = useMutation({
    mutationFn: register,
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
