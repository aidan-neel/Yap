import type { Post } from '$lib/models/Post';
import { userId } from '$lib/user';
import { API_URL } from '$lib/utils';

async function addDownvote(post_id: number) {
	const res = await fetch(`${API_URL}/api/posts/vote`, {
		method: 'POST',
		body: JSON.stringify({
			post_id: post_id,
			user_id: userId.current,
			is_up: false
		}),
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = (await res.json()) as {
		error: any | null;
		ok: boolean;
	};

	if (!data.ok || !res.ok || data.error !== null) {
		return;
	}
}

async function removeDownvote(post_id: number) {
	const res = await fetch(`${API_URL}/api/posts/vote`, {
		method: 'DELETE',
		body: JSON.stringify({
			post_id: post_id,
			user_id: userId.current
		}),
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = (await res.json()) as {
		error: any | null;
		ok: boolean;
	};

	if (!data.ok || !res.ok || data.error !== null) {
		return;
	}
}

async function addUpvote(post_id: number) {
	const res = await fetch(`${API_URL}/api/posts/vote`, {
		method: 'POST',
		body: JSON.stringify({
			post_id: post_id,
			user_id: userId.current,
			is_up: true
		}),
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = (await res.json()) as {
		error: any | null;
		ok: boolean;
	};

	if (!data.ok || !res.ok || data.error !== null) {
		return;
	}
}

async function removeUpvote(post_id: number) {
	const res = await fetch(`${API_URL}/api/posts/vote`, {
		method: 'DELETE',
		body: JSON.stringify({
			post_id: post_id,
			user_id: userId.current
		}),
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = (await res.json()) as {
		error: any | null;
		ok: boolean;
	};

	if (!data.ok || !res.ok || data.error !== null) {
		return;
	}
}

async function getAllPosts(): Promise<Post[]> {
	const res = await fetch(`${API_URL}/api/posts?user_id=${userId.current}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = await res.json();
	const posts = data.posts as Post[];
	return posts;
}

async function createPost(content: string) {
	const res = await fetch(`${API_URL}/api/posts`, {
		method: 'POST',
		body: JSON.stringify({
			user_id: userId.current,
			content: content
		}),
		headers: {
			'Content-Type': 'application/json'
		}
	});

	const data = (await res.json()) as {
		error: any | null;
		ok: boolean;
		post: Post;
	};

	if (!data.ok || !res.ok || data.error !== null) {
		return;
	}

	return data.post;
}

export const postService = {
	addDownvote,
	removeDownvote,
	addUpvote,
	removeUpvote,

	getAllPosts,

	createPost
};
