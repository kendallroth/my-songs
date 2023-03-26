import { acceptHMRUpdate, defineStore } from "pinia";

import { type Account } from "~/slices/account/types";

interface AccountState {
  account: Account | null;
}

export const useAccountStore = defineStore("account", {
  state: (): AccountState => ({
    account: null,
  }),
  getters: {
    authenticated: (state) => Boolean(state.account),
  },
  actions: {
    /** Clear authenticated user (logout, etc) */
    clearAccount() {
      this.$reset();
    },
    /** Set authenticated account information */
    setAccount(payload: Account | null) {
      this.account = payload;
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAccountStore, import.meta.hot));
}
