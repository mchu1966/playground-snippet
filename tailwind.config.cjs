const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			lineHeight: {
				'extra-loose': '2.5',
				12: '3rem'
			},
			darkMode: 'class',
			boxShadow: {
				'bt-xl': '0 35px 60px -60px rgba(0, 0, 0, 0.3)'
			}
		},
		container: {
			center: true
		}
	},
	darkMode: 'class'
};

module.exports = config;
