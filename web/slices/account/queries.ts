import { useMutation } from "@tanstack/vue-query";

import { pocketbase } from "~/api";
import { type LoginSchema } from "~/pages/auth/login/schema";

export const useLoginMutation = () => {
  return useMutation({
    mutationFn: async (args: LoginSchema) => {
      return pocketbase.collection("accounts").authWithPassword(args.email, args.password);
    },
  });
};
