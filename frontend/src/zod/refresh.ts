import z from "zod";
import { DefaultError } from "@/zod";

export const RefreshSuccess = z.object({
  message: z.enum(["ok"]),
  access: z.string().nonempty(),
});

export const RefreshValidator = z.union([RefreshSuccess, DefaultError]);

export type RefreshValidatorType = z.infer<typeof RefreshValidator>;
