import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import type { Timestamp } from './models/Post';

export const API_URL = 'https://yap-api.aidan-neel.com';

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChild<T> = T extends { child?: any } ? Omit<T, 'child'> : T;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChildren<T> = T extends { children?: any } ? Omit<T, 'children'> : T;
export type WithoutChildrenOrChild<T> = WithoutChildren<WithoutChild<T>>;
export type WithElementRef<T, U extends HTMLElement = HTMLElement> = T & { ref?: U | null };

export function timeAgo(dateStr?: string | null): string {
	console.log(dateStr);
	if (!dateStr) return 'invalid';
	const date = new Date(dateStr.replace(' ', 'T'));
	if (isNaN(date.getTime())) return 'invalid';

	const now = new Date();
	const diff = (now.getTime() - date.getTime()) / 1000;

	const minute = 60;
	const hour = 3600;
	const day = 86400;
	const week = 604800;
	const month = 2628000;
	const year = 31536000;

	if (diff < minute) return 'just now';
	if (diff < hour) return `${Math.floor(diff / minute)}m`;
	if (diff < day) return `${Math.floor(diff / hour)}h`;
	if (diff < week) return `${Math.floor(diff / day)}d`;
	if (diff < month) return `${Math.floor(diff / week)}w`;
	if (diff < year) return `${Math.floor(diff / month)}mo`;
	return `${Math.floor(diff / year)}y`;
}
