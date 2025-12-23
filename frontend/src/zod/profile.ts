import z from "zod";
import { DefaultError } from "@/zod";

export const ProfileData = z.object({
  id: z.coerce.number(),
  user_name: z.string().nonempty(),
  email: z.email().nonempty(),
  aadhar: z.string().regex(/^\d{12}$/),
});

export type ProfileDataType = z.infer<typeof ProfileData>;

export const ProfileValidator = z.union([ProfileData, DefaultError]);

export type ProfileValidatorType = z.infer<typeof ProfileValidator>;
