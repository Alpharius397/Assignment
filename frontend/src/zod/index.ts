import z from "zod";

export const DefaultError = z.object({
  message: z.string().nonempty(),
});

export type DefaultErrorType = z.infer<typeof DefaultError>;
