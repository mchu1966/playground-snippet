const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			lineHeight: {
				'extra-loose': '2.5',
				12: '3rem'
			},
			darkMode: 'class'
		}
	},
	darkMode: 'class'
};

module.exports = config;
