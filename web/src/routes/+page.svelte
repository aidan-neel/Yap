<script lang="ts">
    import { Textarea } from "$lib/components/ui/textarea";
    import { Button } from "$lib/components/ui/button";
	import PostComponent from "$lib/components/Post.svelte";
    import type { Post } from "$lib/models/Post";
	import { onMount } from "svelte";
	import { postService } from "$lib/services/postService";
    import ArrowUp from "@lucide/svelte/icons/arrow-up";
    import LoaderCircle from "@lucide/svelte/icons/loader-circle";

    let posts = $state<Post[]>([]);
    let posting = $state<boolean>(false);
    let postContent = $state<string>('');

    async function post() {
        if (postContent.length < 3) { return }

        posting = true;
        let post = await postService.createPost(postContent);
        if(post) {
            posts = [...posts, post]
        }
        posts.sort((a, b) => new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime());
        postContent = '';
        posting = false;
    }
    
    onMount(async() => {
        posts = await postService.getAllPosts();
        posts.sort((a, b) => new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime());
    })
</script>

<main class="h-screen w-full md:w-2/3 lg:w-[40rem] p-3 gap-3 overflow-hidden flex flex-col relative mx-auto">
    <div class="flex flex-col w-full hide-scrollbar h-full overflow-y-auto">
        {#each posts as post}
            <PostComponent {post} />
        {/each}
    </div>
    <div class="flex bg-background w-full flex-shrink-0 relative flex-col gap-3">
        <Textarea bind:value={postContent} onkeydown={async(e) => {
            if (e.key === 'Enter') {
                await post()
            }
        }} disabled={posting} placeholder="What's on your mind?" class="resize-none pb-6 min-h-20" />
        <Button onclick={post} disabled={posting} class="w-fit size-7 ml-auto absolute bottom-2 right-2">
            {#if posting}
                <LoaderCircle size={15} class="animate-spin" />
            {:else}
                <ArrowUp size={15} />
            {/if}
        </Button>
    </div>
</main>