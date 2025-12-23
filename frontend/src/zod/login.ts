import z from "zod";
import { DefaultError } from "@/zod";

export const LoginSuccess = z.object({
  message: z.enum(["ok"]),
  access: z.string().nonempty(),
  refresh: z.string().nonempty(),
});

export const LoginValidator = z.union([LoginSuccess, DefaultError]);

export type LoginValidatorType = z.infer<typeof LoginValidator>;

export const LoginInput = z.object({
  user_name: z.string().nonempty({ error: "Username cannot be empty" }),
  password: z.string().nonempty({ error: "Password cannot be empty" }),
});

export type LoginInputType = z.infer<typeof LoginInput>;
