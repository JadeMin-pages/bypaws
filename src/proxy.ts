export default async function proxyFetch(
	request: Request
): Promise<Response> {
	// GET 요청만 허용
	// if (request.method !== "GET") {
	// 	return new Response("Access Denied", {
	// 		status: 405
	// 	});
	// }

	const requestUrl = new URL(request.url);
	const { pathname, search } = requestUrl;

	// 게시글 페이지 또는 미디어 파일 요청인지 확인
	const isPostDetail = /\/posts\/\d+/i.test(pathname);
	const isMedia = /\.(jpg|jpeg|png|gif|webp|webm|mp4)$/i.test(pathname);

	const cacheFetch = async (useCache: boolean): Promise<Response> => {
		const targetUrl = `https://e621.net${pathname}${search}`;

		const response = await fetch(targetUrl, {
			method: "GET",
			headers: {
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
			},
			cf: {
				cacheEverything: true,
				cacheTtl: useCache ? 2592000 : 0,
			}
		});

		const resHeaders = new Headers(response.headers);
		if (useCache) {
			resHeaders.set("Cache-Control", "public, max-age=31536000, immutable");
		}

		return new Response(response.body, {
			headers: resHeaders
		});
	};


	if (isPostDetail || isMedia) {
		return cacheFetch(true);
	} else {
		return cacheFetch(false);
	}
};