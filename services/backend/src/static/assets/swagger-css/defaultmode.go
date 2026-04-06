package swaggercss

import "html/template"

func GetSwaggerCustomCSS() template.CSS {
	return `
	:root {
		--custom-color: #388e47;
	}
	body {
		margin: 0;
	} 
	body > #swagger-ui {
		min-height: 100vh;
		background-color: #ebebeb;
		padding-bottom: 30px;
	}
	.swagger-ui .topbar .download-url-wrapper input[type=text] {
		border-color: var(--custom-color); 
	}
	.swagger-ui .topbar .download-url-wrapper .download-url-button {
		background: var(--custom-color);
	}
	.swagger-ui .models {
		margin-bottom: 0;
	}

	.topbar-wrapper img {
		content: url('/api/asset/logo-branding')
	}
`
}
