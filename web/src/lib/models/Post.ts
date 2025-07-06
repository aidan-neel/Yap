export type Timestamp =
	`${number}-${number}-${number} ${number}:${number}:${number}.${number}+${number}${number}`; // e.g., 2025-07-05 21:22:13.897577+00

export type Post = {
	ID: number;
	content: string;
	CreatedAt: Timestamp;
	UpdatedAt: Timestamp;
	DeletedAt: Timestamp | null;
	upvotes: number;
	downvotes: number;
	upvoted: boolean;
	downvoted: boolean;
};
