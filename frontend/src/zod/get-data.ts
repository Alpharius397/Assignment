import z from "zod";
import { DefaultError } from "@/zod";

export const UserData = z.object({
  id: z.coerce.number(),
  user_name: z.string().nonempty(),
  email: z.email().nonempty(),
  aadhar: z.string().transform((aadhar) => {
    const regexp = new RegExp(/^\d{12}$/); // check whether we received an  encrypted aadhar or decrypted one

    if (regexp.exec(aadhar) !== null) {
      return (
        aadhar.slice(0, 4) + "-" + aadhar.slice(4, 8) + "-" + aadhar.slice(8)
      );
    } else {
      return aadhar;
    }
  }),
});

export type UserDataType = z.infer<typeof UserData>;

export const GetDataSuccess = z.object({
  message: z.enum(["ok"]),
  data: z.array(UserData),
  total: z.number(),
});

export type GetDataSuccessType = z.infer<typeof GetDataSuccess>;

export const GetDataValidator = z.union([GetDataSuccess, DefaultError]);

export type GetDataValidatorType = z.infer<typeof GetDataValidator>;
