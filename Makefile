run:
	@bun run tailwindcss --config tailwind.config.js -i input.css -o static/css/styles.css
	@templ generate
	@go build -o ./tmp/main ./cmd