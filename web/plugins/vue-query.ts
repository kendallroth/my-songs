import {
  // dehydrate,
  // type DehydratedState,
  // hydrate,
  QueryClient,
  VueQueryPlugin,
  type VueQueryPluginOptions,
} from "@tanstack/vue-query";

export default defineNuxtPlugin((nuxtApp) => {
  // const vueQueryState = useState<DehydratedState | null>("vue-query");

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        // Time until inactive queries (no observers registered) will be removed from cache
        cacheTime: 5 * 60 * 1000,
        refetchOnWindowFocus: true,
        retry: 2,
        // Time until a query transitions from fresh (read from cache) to stale (possible refetch).
        //   Refetched when new instances mount, window is refocused, or a refetch interval.
        staleTime: 1 * 60 * 1000,
      },
    },
  });

  const options: VueQueryPluginOptions = {
    queryClient,
  };

  nuxtApp.vueApp.use(VueQueryPlugin, options);

  // TODO: Figure out how to properly get SSR hydration working
  /*if (process.server) {
    nuxtApp.hooks.hook("app:rendered", () => {
      vueQueryState.value = dehydrate(queryClient);
    });
  }

  if (process.client) {
    nuxtApp.hooks.hook("app:created", () => {
      hydrate(queryClient, vueQueryState.value);
    });
  }*/
});
