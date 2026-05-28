import proxyFetch from "./proxy.ts";



export default {
	async fetch(request, env, ctx): Promise<Response> {
		try {
			return await proxyFetch(request);
		} catch (e: any) {
			return new Response(e.stack || e.message, { status: 500 });
		}
	},
} satisfies ExportedHandler<Env>;