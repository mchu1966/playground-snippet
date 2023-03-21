const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			lineHeight: {
				'extra-loose': '2.5',
				12: '3rem'
			},
			darkMode: 'class'
		},
		container: {
			center: true,
		},
	},
	darkMode: 'class'
};

module.exports = config;
