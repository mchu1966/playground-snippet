import {
	PUBLIC_SUPABASE_URL,
	PUBLIC_SUPABASE_ANON_KEY,
	PUBLIC_LOCAL_SUPABASE_URL,
	PUBLIC_LOCAL_SUPABASE_ANON_KEY
} from '$env/static/public';
import type { Database } from '$lib/types/database.types';
import { createSupabaseServerClient } from '@supabase/auth-helpers-sveltekit';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const isDevEnv: boolean = import.meta.env.DEV;
	event.locals.supabase = createSupabaseServerClient<Database>({
		supabaseUrl: isDevEnv ? PUBLIC_LOCAL_SUPABASE_URL : PUBLIC_SUPABASE_URL,
		supabaseKey: isDevEnv ? PUBLIC_LOCAL_SUPABASE_ANON_KEY : PUBLIC_SUPABASE_ANON_KEY,
		event
	});

	/**
	 * A convenience helper so we can just call await getSession() instead const { data: { session } } = await supabase.auth.getSession()
	 */
	event.locals.getSession = async () => {
		const {
			data: { session }
		} = await event.locals.supabase.auth.getSession();
		return session;
	};

	return resolve(event, {
		/**
		 * There´s an issue with `filterSerializedResponseHeaders` not working when using `sequence`
		 *
		 * https://github.com/sveltejs/kit/issues/8061
		 */
		filterSerializedResponseHeaders(name) {
			return name === 'content-range';
		}
	});
};
