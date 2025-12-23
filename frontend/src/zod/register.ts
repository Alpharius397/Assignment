import z from "zod";
import { DefaultError } from "@/zod";

export const RegisterInput = z
  .object({
    user_name: z.string().nonempty({ error: "Username cannot be empty" }),
    email: z
      .email({ error: "Invalid email address" })
      .nonempty({ error: "Email cannot be empty" }),
    password: z.string().nonempty({ error: "Password cannot be empty" }),
    confirm_password: z.string().nonempty(),
    aadhar: z.string().regex(/^\d{12}$/, { error: "Invalid Aadhar ID" }),
  })
  .refine((data) => data.confirm_password == data.password, {
    error: "Password must match",
    path: ["confirm_password"],
  });

export type RegisterInputType = z.infer<typeof RegisterInput>;

export const RegisterSuccess = z.object({
  message: z.enum(["ok"]),
});

export const RegisterValidator = z.union([RegisterSuccess, DefaultError]);
