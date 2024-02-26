# Cheat Sheets

My personal cheat sheets and cheat sheets generator, which generates cheat sheets from JSON files. Inspired by [vim.rtorr.com](https://vim.rtorr.com).

```sh
go build -o ./generator .

./generator data/index.json > out/index.html

./generator data/flux.json > out/flux.html
./generator data/gh-dash.json > out/gh-dash.html
./generator data/vim.json > out/vim.html
./generator data/vim-plugins.json > out/vim-plugins.html
```

### Formatting

- **Code:** To add a code section a command can be prefixed using `||| `. The code will be rendered after the commands description.
- **Image:** To add an image a command can be prefixed using `!!! ` followed by the image url, e.g. `!!! assets/flux-gitops-toolkit.png`
