import { type InferType, object, string } from "yup";

export const loginSchema = object({
  email: string().label("Email").email().required(),
  password: string().label("Password").required(),
});

export type LoginSchema = InferType<typeof loginSchema>;
