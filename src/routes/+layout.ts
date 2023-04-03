// import { invalidate } from '$app/navigation';
import {
	PUBLIC_SUPABASE_ANON_KEY,
	PUBLIC_SUPABASE_URL,
	PUBLIC_LOCAL_SUPABASE_URL,
	PUBLIC_LOCAL_SUPABASE_ANON_KEY
} from '$env/static/public';
import { createSupabaseLoadClient } from '@supabase/auth-helpers-sveltekit';
import type { LayoutLoad } from './$types';
import type { snippet } from '$lib/types/snippet';

export const load: LayoutLoad = async ({ fetch, data, depends }) => {
	console.log('layout load.');

	depends('supabase:auth');

	const isDevEnv: boolean = import.meta.env.DEV;
	const supabase = createSupabaseLoadClient({
		supabaseUrl: isDevEnv ? PUBLIC_LOCAL_SUPABASE_URL : PUBLIC_SUPABASE_URL,
		supabaseKey: isDevEnv ? PUBLIC_LOCAL_SUPABASE_ANON_KEY : PUBLIC_SUPABASE_ANON_KEY,
		event: { fetch },
		serverSession: data.session
	});

	const {
		data: { session }
	} = await supabase.auth.getSession();

	const mockSnippets: snippet[] = [
		{
			name: 'Hello World',
			code: `// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
`
		}
	];

	return { supabase, session, mockSnippets };
};
