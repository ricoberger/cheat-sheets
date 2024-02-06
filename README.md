# Cheat Sheets

My personal cheat sheets and cheat sheets generator, which generates cheat sheets from JSON files. Inspired by [vim.rtorr.com](https://vim.rtorr.com).

```sh
go build -o ./generator .

./generator data/gh-dash.json > out/gh-dash.html
./generator data/vim.json > out/vim.html
./generator data/vim-plugins.json > out/vim-plugins.html
```
