<script lang="ts">
	import type { Post } from "$lib/models/Post";
    import { Button } from "$lib/components/ui/button";
    import ArrowUp from "@lucide/svelte/icons/arrow-up";
    import ArrowDown from "@lucide/svelte/icons/arrow-down";
	import { timeAgo } from "$lib/utils";
	import { postService } from "$lib/services/postService";

    type Props = {
        post: Post
    }

    const { post }: Props = $props();
    
    let upvoted = $derived(post.upvoted);
    let downvoted = $derived(post.downvoted);
    
    async function upvote() {
        if (upvoted) {
            upvoted = false;
            post.upvotes -= 1;
            await postService.removeUpvote(post.ID);
        } else {
            if (downvoted) {
                downvoted = false;
                post.downvotes -= 1;
                await postService.removeDownvote(post.ID);
            }
            upvoted = true;
            post.upvotes += 1;
            await postService.addUpvote(post.ID);
        }
    }

    async function downvote() {
        if (downvoted) {
            downvoted = false;
            post.downvotes -= 1;
            await postService.removeDownvote(post.ID);
        } else {
            if (upvoted) {
                upvoted = false;
                post.upvotes -= 1;
                await postService.removeUpvote(post.ID);
            }
            downvoted = true;
            post.downvotes += 1;
            await postService.addDownvote(post.ID);
        }
    }
</script>

<div class="flex py-6 border-b flex-col gap-2">
    <header class="flex flex-row items-start gap-3">
        <img class="size-10 rounded-full border" src="https://api.dicebear.com/9.x/notionists-neutral/svg?seed=Brian" />
        <div class="flex flex-col gap-1 w-full">
            <div class="flex flex-row w-full justify-between">
                <h1 class="font-medium text-sm">Anonymous</h1>
                <p class="text-muted-foreground text-xs">
                    {timeAgo(post.CreatedAt)}
                </p>
            </div>
            <p class="">
                {post.content}
            </p>

            <div class="flex flex-row mt-1 gap-2">
                <button onclick={upvote} class="p-1 bg-card border {upvoted ? 'bg-green-950 hover:bg-green-900' : 'hover:bg-secondary'} duration-150 rounded-md">
                    <ArrowUp class="{upvoted ? 'text-green-500' : ''} duration-150" size={16} />
                </button>
                <button onclick={downvote} class="p-1 bg-card border {downvoted ? 'bg-red-950 hover:bg-red-900' : 'hover:bg-secondary'} duration-150 rounded-md">
                    <ArrowDown class="{downvoted ? 'text-red-500' : ''} duration-150" size={16} />
                </button>
            </div>

            <div>
                <p class="text-sm text-muted-foreground mt-1">
                    {post.upvotes - post.downvotes} votes
                </p>
            </div>
        </div>
    </header>
</div>